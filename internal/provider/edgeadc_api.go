package provider

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
	"time"

	"terraform-provider-edgeadc/swagger"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	DoRequestError = "error: do request: %w"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type API struct {
	client         HTTPClient
	baseURL        string
	hostPort       string
	password       string
	username       string
	cookieGuid     string
	loggingContext context.Context
	mutexKV        *MutexKV

	// Cache for Read-Phase ONLY for performance
	cachedIpServices     swagger.IpServices
	cachedIpServiceCombo swagger.IpServicescombo
}

type LoginResponse struct {
	LoginStatus string `json:"LoginStatus"`
	UserName    string `json:"UserName"`
	GUID        string `json:"GUID"`
}

type EdgeResponse struct {
	LoginStatus string `json:"LoginStatus"`
	Success     string `json:"success"`
	StatusImage string `json:"StatusImage"`
	StatusText  string `json:"StatusText"`
}

func NewAPI(baseURL, username, password, hostPort string) *API {
	api := &API{
		baseURL:    baseURL,
		username:   username,
		password:   password,
		hostPort:   hostPort,
		cookieGuid: "",
		// Initialize logging context
		loggingContext: context.Background(),
		// Set a mutexKV to prevent concurrent requests which can be used globally
		mutexKV: NewMutexKV(),
		// Initialize a new cache for the API
		cachedIpServices:     swagger.IpServices{},
		cachedIpServiceCombo: swagger.IpServicescombo{},
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{MinVersion: tls.VersionTLS12},
	}

	api.client = &http.Client{Transport: tr}
	return api
}

// buildURL constructs the full API URL, handling the case where hostPort may be empty
func (api *API) buildURL(path string) string {
	if api.hostPort == "" {
		return fmt.Sprintf("%s%s", api.baseURL, path)
	}
	return fmt.Sprintf("%s:%s%s", api.baseURL, api.hostPort, path)
}

func (api *API) getCookie() (string, error) {
	endpointURL := api.buildURL("/POST/32")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	// Log URL (exclude body as it contains sensitive information)
	tflog.Debug(api.loggingContext, endpointURL)
	postBody, _ := json.Marshal(map[string]string{
		api.username: api.password,
	})
	getCookieRequest, err := http.NewRequestWithContext(ctx, "POST", endpointURL, bytes.NewBuffer(postBody))
	if err != nil {
		tflog.Error(api.loggingContext, err.Error())
		return "", fmt.Errorf("error: creating cookie request: %w", err)
	}
	response, responseErr := api.client.Do(getCookieRequest)
	if responseErr != nil {
		tflog.Error(api.loggingContext, responseErr.Error())
		return "", fmt.Errorf("error: creating cookie response: %w", responseErr)
	}
	loginResponse := &LoginResponse{}
	loginResponseErr := ReadLoginResponse(response, loginResponse)
	if loginResponseErr != nil {
		tflog.Error(api.loggingContext, loginResponseErr.Error())
		return "", fmt.Errorf("error: processing Login response: %w", loginResponseErr)
	}
	return loginResponse.GUID, nil
}

// doAuthenticatedRequest performs an authenticated request to the EdgeADC API
// The primary function being to retrieve and set the GUID cookie if it doesn't exist
func (api *API) doAuthenticatedRequest(r *http.Request) (*http.Response, error) {
	// Get the cookie if it doesn't exist
	if api.cookieGuid == "" {
		cookieGuid, errGuid := api.getCookie()
		if errGuid != nil {
			tflog.Debug(api.loggingContext, errGuid.Error())
			return nil, errGuid
		}
		api.cookieGuid = cookieGuid
	}
	cookie := http.Cookie{
		Name:  "GUID",
		Value: api.cookieGuid,
	}
	// Use the cookie for subsequent authentication
	r.AddCookie(&cookie)
	// Other headers
	r.Header.Set("Accept", "application/json")
	return api.client.Do(r)
}

// GetEdgeADCObject performs a GET request to the EdgeADC API
func (api *API) GetEdgeADCObject(path string) (string, error) {
	return api.getEdgeADCObjectInternal(path, false)
}

func (api *API) getEdgeADCObjectInternal(path string, isRetry bool) (string, error) {
	endpointURL := api.buildURL(path)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	// Log URL and body
	tflog.Debug(api.loggingContext, endpointURL)
	req, errGet := http.NewRequestWithContext(ctx, "GET", endpointURL, http.NoBody)
	if errGet != nil {
		return "", fmt.Errorf("error: creating EdgeADC request: %w", errGet)
	}

	req.Header.Set("Content-Type", "text/plain")
	var resp, err = api.doAuthenticatedRequest(req) //nolint:bodyclose // linters bug
	if err != nil {
		return "", fmt.Errorf(DoRequestError, err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			return
		}
	}()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", fmt.Errorf("EdgeADC authentication failed: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(api.loggingContext, err.Error())
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// EdgeADC Returns success code and status but
	// actually not successful
	edgeResponse := EdgeResponse{}
	_ = json.Unmarshal(b, &edgeResponse)

	// Check for session expiration and re-authenticate once
	if edgeResponse.LoginStatus == "Session Expired" && !isRetry {
		tflog.Debug(api.loggingContext, fmt.Sprintf("[GET] Session expired for %s, clearing cookie and retrying...", path))
		api.cookieGuid = ""
		return api.getEdgeADCObjectInternal(path, true)
	}
	if edgeResponse.LoginStatus == "Session Expired" && isRetry {
		tflog.Error(api.loggingContext, fmt.Sprintf("[GET] Session still expired after re-auth for %s", path))
	}

	if edgeResponse.StatusImage == "jetError" {
		tflog.Error(api.loggingContext, string(b))
		return "", fmt.Errorf("EdgeADC: %s", edgeResponse.StatusText)
	}

	// Return Object
	return string(b), nil
}

// PostEdgeADCApi performs a POST request to the EdgeADC API with the default headers
func (api *API) PostEdgeADCApi(path string, body []byte) (string, error) {
	headers := map[string]string{"Content-Type": "text/plain"}
	return api.PostEdgeADCApiWithHeaders(path, body, headers)
}

// PostEdgeADCApiWithHeaders performs a POST request to the EdgeADC API with headers
// This is primarily used for POST requests that require a specific content type (e.g. multipart/form-data)
func (api *API) PostEdgeADCApiWithHeaders(path string, body []byte, headers map[string]string) (string, error) {
	return api.postEdgeADCApiWithHeadersInternal(path, body, headers, false)
}

func (api *API) postEdgeADCApiWithHeadersInternal(path string, body []byte, headers map[string]string, isRetry bool) (string, error) {
	endpointURL := api.buildURL(path)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	// Log URL and body
	tflog.Debug(api.loggingContext, endpointURL)
	tflog.Debug(api.loggingContext, string(body))
	getRequest, err := http.NewRequestWithContext(ctx, "POST", endpointURL, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error: creating EdgeADC request: %w", err)
	}

	// Set headers
	for k, v := range headers {
		getRequest.Header.Set(k, v)
	}

	resp, errDo := api.doAuthenticatedRequest(getRequest) //nolint:bodyclose // linters bug
	if errDo != nil {
		tflog.Error(api.loggingContext, errDo.Error())
		return "", fmt.Errorf(DoRequestError, errDo)
	}

	b, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		tflog.Error(api.loggingContext, errRead.Error())
		return "", fmt.Errorf("error reading response body: %w", errRead)
	}

	// EdgeADC Returns success code and status but
	// actually not successful
	edgeResponse := EdgeResponse{}
	_ = json.Unmarshal(b, &edgeResponse)

	// Check for session expiration and re-authenticate once
	if edgeResponse.LoginStatus == "Session Expired" && !isRetry {
		tflog.Debug(api.loggingContext, "Session expired on POST, re-authenticating...")
		api.cookieGuid = ""
		return api.postEdgeADCApiWithHeadersInternal(path, body, headers, true)
	}

	if edgeResponse.StatusImage == "jetError" {
		tflog.Error(api.loggingContext, string(b))
		return "", fmt.Errorf("EdgeADC: %s - %s", path, edgeResponse.StatusText)
	}
	if edgeResponse.StatusImage == "jetWarning" {
		tflog.Warn(api.loggingContext, fmt.Sprintf("EdgeADC warning on %s: %s", path, edgeResponse.StatusText))
	}

	// Return Object
	return string(b), nil
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// UploadFileEdgeADCApi uploads a file to the EdgeADC API
// This is primarily used for uploading SSL certificates
func (api *API) UploadFileEdgeADCApi(path string, params map[string]string, filePath string) (string, error) {
	paramName := "SslCertificatesImportUploadText"
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Need to set a custom header otherwise would have used
	// part, err := writer.CreateFormFile(paramName, filepath.Base(filePath))
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(paramName), escapeQuotes(filepath.Base(filePath))))
	h.Set("Content-Type", "application/x-pkcs12")
	part, err := writer.CreatePart(h)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}
	headers := map[string]string{"Content-Type": writer.FormDataContentType()}
	return api.PostEdgeADCApiWithHeaders(path, body.Bytes(), headers)
}

// UploadCustomMonitorFileEdgeADCApi uploads a file to the EdgeADC API
func (api *API) UploadCustomMonitorFileEdgeADCApi(path string, params map[string]string, filePath string) (string, error) {
	paramName := "UploadMointorupload"
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Need to set a custom header otherwise would have used
	// part, err := writer.CreateFormFile(paramName, filepath.Base(filePath))
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(paramName), escapeQuotes(filepath.Base(filePath))))
	h.Set("Content-Type", "application/octet-stream")
	part, err := writer.CreatePart(h)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}
	headers := map[string]string{"Content-Type": writer.FormDataContentType()}
	return api.PostEdgeADCApiWithHeaders(path, body.Bytes(), headers)
}

func ReadLoginResponse(resp *http.Response, object *LoginResponse) error {
	var buf bytes.Buffer
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			return
		}
	}()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("EdgeADC authentication failed: %d %s", resp.StatusCode, buf.String())
	}
	_, err := buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf.Bytes(), object)
}

// GetAddressAndPortFromId gets the ipaddress and port from an ID
// We use the combination of ip:port as the unique identifier for many resources
func GetAddressAndPortFromId(id string) (ipAddr string, port string, err error) {
	split := strings.Split(id, ":")
	if len(split) != 2 {
		return "", "", errors.New(fmt.Sprintf("Unable to parse id: %s", id))
	}
	return split[0], split[1], nil
}
