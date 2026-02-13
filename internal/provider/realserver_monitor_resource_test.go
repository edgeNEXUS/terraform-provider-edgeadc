package provider

import (
	"context"
	"testing"

	"terraform-provider-edgeadc/swagger"

	schemalib "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-edgeadc/internal/provider/resource_realserver_monitor"
)

// TestNormalizeMonitorType_FriendlyToBackend verifies that all known friendly
// UI names are correctly mapped to their backend equivalents. This test would
// have caught the bug where using "HTTP Head" in Terraform created a monitor
// that didn't function because the backend needed "CheckHead".
func TestNormalizeMonitorType_FriendlyToBackend(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HTTP 200 OK", "Check200"},
		{"HTTP Head", "CheckHead"},
		{"HTTP Response", "CheckResponse"},
		{"TCP Connect", "Connect"},
		{"ICMP Ping", "Ping"},
		{"None", "None"},
		{"DICOM", "DICOM"},
		{"HTTP Post", "CheckPOST"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result := NormalizeMonitorType(tc.input)
			if result != tc.expected {
				t.Errorf("NormalizeMonitorType(%q) = %q, want %q", tc.input, result, tc.expected)
			}
		})
	}
}

// TestNormalizeMonitorType_BackendPassthrough verifies that backend names pass
// through unchanged. Users who already use backend names should not be affected.
func TestNormalizeMonitorType_BackendPassthrough(t *testing.T) {
	backendNames := []string{
		"Check200", "CheckHead", "CheckResponse", "Connect", "Ping",
		"None", "DICOM", "CheckPOST",
	}

	for _, name := range backendNames {
		t.Run(name, func(t *testing.T) {
			result := NormalizeMonitorType(name)
			if result != name {
				t.Errorf("NormalizeMonitorType(%q) = %q, want %q (backend names must pass through)", name, result, name)
			}
		})
	}
}

// TestNormalizeMonitorType_CustomMonitorPassthrough verifies that custom monitor
// script names (which are neither friendly nor standard backend names) pass through
// unchanged.
func TestNormalizeMonitorType_CustomMonitorPassthrough(t *testing.T) {
	customNames := []string{
		"my_custom_monitor", "CustomHealthCheck", "perl_check_v2",
	}

	for _, name := range customNames {
		t.Run(name, func(t *testing.T) {
			result := NormalizeMonitorType(name)
			if result != name {
				t.Errorf("NormalizeMonitorType(%q) = %q, want %q (custom names must pass through)", name, result, name)
			}
		})
	}
}

// TestToRealserverMonitorOpt_NormalizesType verifies that the Terraform model
// to swagger conversion normalizes the type field. This is the core fix for the
// monitor type bug.
func TestToRealserverMonitorOpt_NormalizesType(t *testing.T) {
	r := &realserverMonitorResource{}

	tests := []struct {
		inputType    string
		expectedType string
	}{
		{"HTTP Head", "CheckHead"},
		{"HTTP Response", "CheckResponse"},
		{"Check200", "Check200"},
		{"TCP Connect", "Connect"},
	}

	for _, tc := range tests {
		t.Run(tc.inputType, func(t *testing.T) {
			data := resource_realserver_monitor.RealserverMonitorModel{
				Id:          types.StringValue("1"),
				Name:        types.StringValue("test"),
				Description: types.StringValue("test desc"),
				Type:        types.StringValue(tc.inputType),
				Ssl:         types.StringValue("auto"),
				Url:         types.StringValue("/"),
				Content:     types.StringValue(""),
				Username:    types.StringValue(""),
				Password:    types.StringValue(""),
				Threshold:   types.StringValue("3"),
			}

			result := r.ToRealserverMonitorOpt(data)
			if result.Type_ != tc.expectedType {
				t.Errorf("ToRealserverMonitorOpt with type=%q produced Type_=%q, want %q",
					tc.inputType, result.Type_, tc.expectedType)
			}
		})
	}
}

// TestToTerraformModel_PreservesAllFields verifies the swagger-to-terraform
// model conversion preserves all fields correctly.
func TestToTerraformModel_PreservesAllFields(t *testing.T) {
	r := &realserverMonitorResource{}

	input := swagger.RealConfigMonitoringOpt{
		Id:          "42",
		Name:        "TestMonitor",
		Description: "A test monitor",
		Type_:       "CheckHead",
		Ssl:         "auto",
		Url:         "/health",
		Content:     "200 OK",
		Username:    "admin",
		Password:    "secret",
		Threshold:   "5",
	}

	result := r.ToTerraformModel(input)

	checks := map[string]struct {
		got  types.String
		want string
	}{
		"Id":          {result.Id, "42"},
		"Name":        {result.Name, "TestMonitor"},
		"Description": {result.Description, "A test monitor"},
		"Type":        {result.Type, "CheckHead"},
		"Ssl":         {result.Ssl, "auto"},
		"Url":         {result.Url, "/health"},
		"Content":     {result.Content, "200 OK"},
		"Username":    {result.Username, "admin"},
		"Password":    {result.Password, "secret"},
		"Threshold":   {result.Threshold, "5"},
	}

	for field, c := range checks {
		if c.got.ValueString() != c.want {
			t.Errorf("field %s = %q, want %q", field, c.got.ValueString(), c.want)
		}
	}
}

// TestGetRealConfigMonitoringOptByName_NotFound verifies that looking up a
// non-existent monitor returns an empty struct (not a panic or error), which
// is needed for graceful delete handling.
func TestGetRealConfigMonitoringOptByName_NotFound(t *testing.T) {
	data := swagger.ConfigMonitoring{
		ConfigMonitoringGrid: &swagger.ConfigMonitoringConfigMonitoringGrid{
			Dataset: &swagger.ConfigMonitoringConfigMonitoringGridDataset{
				Row: []swagger.RealConfigMonitoringOpt{
					{Id: "1", Name: "ExistingMonitor"},
				},
			},
		},
	}

	result := GetRealConfigMonitoringOptByName(data, "NonExistent")
	if result.Id != "" {
		t.Errorf("expected empty struct for non-existent monitor, got Id=%q", result.Id)
	}
}

// TestGetRealConfigMonitoringOptByName_NilGrid verifies graceful handling
// when the API returns nil grid data.
func TestGetRealConfigMonitoringOptByName_NilGrid(t *testing.T) {
	data := swagger.ConfigMonitoring{
		ConfigMonitoringGrid: nil,
	}

	result := GetRealConfigMonitoringOptByName(data, "AnyName")
	if result.Id != "" {
		t.Errorf("expected empty struct for nil grid, got Id=%q", result.Id)
	}
}

// TestGetRealConfigMonitoringOptByName_NilDataset verifies graceful handling
// when the dataset is nil.
func TestGetRealConfigMonitoringOptByName_NilDataset(t *testing.T) {
	data := swagger.ConfigMonitoring{
		ConfigMonitoringGrid: &swagger.ConfigMonitoringConfigMonitoringGrid{
			Dataset: nil,
		},
	}

	result := GetRealConfigMonitoringOptByName(data, "AnyName")
	if result.Id != "" {
		t.Errorf("expected empty struct for nil dataset, got Id=%q", result.Id)
	}
}

// TestRealserverMonitorSchema_IdNotUseStateForUnknown verifies that the
// realserver monitor schema does NOT use UseStateForUnknown on the id field.
// The ADC renumbers monitor IDs when monitors are added/removed, so the ID
// is not stable. Using UseStateForUnknown would cause the plan to preserve
// the old ID, but the API returns a new one -> inconsistency.
//
// TestRealserverMonitorSchema_IdIsComputedWithUseStateForUnknown verifies that
// the id field uses UseStateForUnknown so that plans don't perpetually show
// "(known after apply)". The Update handler separately preserves the planned
// id to avoid "inconsistent result" errors when the ADC renumbers IDs.
func TestRealserverMonitorSchema_IdIsComputedWithUseStateForUnknown(t *testing.T) {
	ctx := context.Background()
	schema := resource_realserver_monitor.RealserverMonitorResourceSchema(ctx)

	idAttr, ok := schema.Attributes["id"]
	if !ok {
		t.Fatal("schema missing 'id' attribute")
	}

	strAttr, ok := idAttr.(schemalib.StringAttribute)
	if !ok {
		t.Fatal("'id' attribute is not a StringAttribute")
	}

	if !strAttr.Computed {
		t.Error("'id' should be Computed")
	}

	// Verify exactly one plan modifier (UseStateForUnknown)
	if len(strAttr.PlanModifiers) != 1 {
		t.Fatalf("'id' should have exactly 1 plan modifier (UseStateForUnknown), has %d",
			len(strAttr.PlanModifiers))
	}
}

// TestMonitorTypeBidirectionalMapping verifies that every friendly name has a
// corresponding backend name and vice versa, ensuring the mapping is complete
// and consistent.
func TestMonitorTypeBidirectionalMapping(t *testing.T) {
	for friendly, backend := range monitorTypeFriendlyToBackend {
		reverseFriendly, ok := monitorTypeBackendToFriendly[backend]
		if !ok {
			t.Errorf("backend name %q (from friendly %q) has no reverse mapping", backend, friendly)
		}
		if reverseFriendly != friendly {
			t.Errorf("reverse mapping for %q = %q, want %q", backend, reverseFriendly, friendly)
		}
	}

	for backend, friendly := range monitorTypeBackendToFriendly {
		reverseBackend, ok := monitorTypeFriendlyToBackend[friendly]
		if !ok {
			t.Errorf("friendly name %q (from backend %q) has no forward mapping", friendly, backend)
		}
		if reverseBackend != backend {
			t.Errorf("forward mapping for %q = %q, want %q", friendly, reverseBackend, backend)
		}
	}
}
