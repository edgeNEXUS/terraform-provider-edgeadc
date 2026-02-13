package provider

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"terraform-provider-edgeadc/swagger"
)

// mockHTTPClient returns canned responses for specific URL paths.
// This lets us exercise the real provider code without an ADC.
type mockHTTPClient struct {
	responses map[string]string // path suffix -> JSON response body
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	path := req.URL.Path
	for suffix, body := range m.responses {
		if len(path) >= len(suffix) && path[len(path)-len(suffix):] == suffix {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewBufferString(body)),
				Header:     make(http.Header),
			}, nil
		}
	}
	// Default: return empty 200
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(`{}`)),
		Header:     make(http.Header),
	}, nil
}

func newTestAPI(responses map[string]string) *API {
	api := NewAPI("http://localhost", "admin", "admin", "8080")
	api.cookieGuid = "test-guid" // skip authentication
	api.client = &mockHTTPClient{responses: responses}
	return api
}

// --------------------------------------------------------------------------
// These tests exercise the REAL DeleteServer / DeleteIPService /
// DeleteRealserverMonitor functions with a mock HTTP backend, proving
// the full code path returns nil (not an error) when the resource is gone.
// --------------------------------------------------------------------------

// TestDeleteServer_NotFound_IntegrationStyle calls the real DeleteServer
// function with a mock API that returns an IP service with no matching server.
// This is exactly what happens when a server has already been deleted.
//
// This test reproduces the customer error:
//   "server not found with df-ad-qa-int-linux0.saab-qa.danf.org.uk:8461
//    for ip_service 10.51.164.2:8461"
func TestDeleteServer_NotFound_IntegrationStyle(t *testing.T) {
	// Mock API returns an IP service at 10.51.164.2:8461 with NO servers
	api := newTestAPI(map[string]string{
		"/GET/9": `{
			"data": {
				"dataset": {
					"ipService": [[{
						"sId": "1",
						"InterfaceID": "0",
						"ChannelID": "0",
						"ipAddr": "10.51.164.2",
						"port": "8461",
						"contentServer": {"cServerId": []}
					}]]
				}
			}
		}`,
	})

	// Try to delete a server that doesn't exist (hostname-based, like the customer's)
	model := swagger.CServerId{
		CSIPAddr: "df-ad-qa-int-linux0.saab-qa.danf.org.uk",
		CSPort:   "8461",
	}

	err := DeleteServer(api, model, "10.51.164.2", "8461")
	if err != nil {
		t.Errorf("DeleteServer should return nil for not-found server, got: %v", err)
	}
}

// TestDeleteServer_ParentIpServiceGone calls DeleteServer when the parent
// IP service itself no longer exists. This happens during terraform destroy
// when the IP service is destroyed before the server.
func TestDeleteServer_ParentIpServiceGone_IntegrationStyle(t *testing.T) {
	// Mock API returns empty IP services (the VIP is already gone)
	api := newTestAPI(map[string]string{
		"/GET/9": `{
			"data": {
				"dataset": {
					"ipService": []
				}
			}
		}`,
	})

	model := swagger.CServerId{
		CSIPAddr: "10.0.0.2",
		CSPort:   "8080",
	}

	err := DeleteServer(api, model, "10.0.0.1", "80")
	if err != nil {
		t.Errorf("DeleteServer should return nil when parent IP service is gone, got: %v", err)
	}
}

// TestDeleteIPService_NotFound_IntegrationStyle calls the real DeleteIPService
// function when the IP service doesn't exist.
func TestDeleteIPService_NotFound_IntegrationStyle(t *testing.T) {
	// Mock API returns empty IP services
	api := newTestAPI(map[string]string{
		"/GET/9": `{
			"data": {
				"dataset": {
					"ipService": []
				}
			}
		}`,
		"/GET/29": `{}`,
	})

	err := DeleteIPService(api, "10.0.0.1", "80")
	if err != nil {
		t.Errorf("DeleteIPService should return nil for not-found service, got: %v", err)
	}
}

// TestDeleteRealserverMonitor_NotFound_IntegrationStyle calls the real
// DeleteRealserverMonitor function when the monitor doesn't exist.
func TestDeleteRealserverMonitor_NotFound_IntegrationStyle(t *testing.T) {
	// Mock API returns monitors but not the one we're looking for
	api := newTestAPI(map[string]string{
		"/GET/13": `{
			"ConfigMonitoringGrid": {
				"dataset": {
					"Row": [
						{"id": "1", "name": "ExistingMonitor"}
					]
				}
			}
		}`,
	})

	err := DeleteRealserverMonitor(api, "NonExistentMonitor")
	if err != nil {
		t.Errorf("DeleteRealserverMonitor should return nil for not-found monitor, got: %v", err)
	}
}

// statefulMockHTTPClient returns different responses for GET after a POST
// has been received, simulating a resource being deleted.
type statefulMockHTTPClient struct {
	beforeDelete map[string]string
	afterDelete  map[string]string
	deleted      bool
}

func (m *statefulMockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	path := req.URL.Path
	// Track when a POST (delete) happens
	if req.Method == http.MethodPost {
		m.deleted = true
	}
	responses := m.beforeDelete
	if m.deleted {
		responses = m.afterDelete
	}
	for suffix, body := range responses {
		if len(path) >= len(suffix) && path[len(path)-len(suffix):] == suffix {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewBufferString(body)),
				Header:     make(http.Header),
			}, nil
		}
	}
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(`{}`)),
		Header:     make(http.Header),
	}, nil
}

// TestDeleteServer_ServerExists_IntegrationStyle verifies that when a server
// DOES exist, DeleteServer proceeds to call the delete API endpoint (doesn't
// short-circuit) and the verification GET confirms deletion.
func TestDeleteServer_ServerExists_IntegrationStyle(t *testing.T) {
	serverExistsJSON := `{
		"data": {
			"dataset": {
				"ipService": [[{
					"sId": "1",
					"InterfaceID": "0",
					"ChannelID": "0",
					"ipAddr": "10.0.0.1",
					"port": "80",
					"contentServer": {"cServerId": [{
						"cId": "0",
						"CSIPAddr": "10.0.0.2",
						"CSPort": "8080",
						"WeightFactor": "100"
					}]}
				}]]
			}
		}
	}`
	serverGoneJSON := `{
		"data": {
			"dataset": {
				"ipService": [[{
					"sId": "1",
					"InterfaceID": "0",
					"ChannelID": "0",
					"ipAddr": "10.0.0.1",
					"port": "80",
					"contentServer": {"cServerId": []}
				}]]
			}
		}
	}`

	mock := &statefulMockHTTPClient{
		beforeDelete: map[string]string{"/GET/9": serverExistsJSON},
		afterDelete:  map[string]string{"/GET/9": serverGoneJSON},
	}
	api := NewAPI("http://localhost", "admin", "admin", "8080")
	api.cookieGuid = "test-guid"
	api.client = mock

	model := swagger.CServerId{
		CSIPAddr: "10.0.0.2",
		CSPort:   "8080",
	}

	// This should NOT return an error - the server exists, gets deleted,
	// and the verification GET confirms it's gone
	err := DeleteServer(api, model, "10.0.0.1", "80")
	if err != nil {
		t.Errorf("DeleteServer for existing server should succeed, got: %v", err)
	}
}
