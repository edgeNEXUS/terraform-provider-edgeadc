package provider

import (
	"context"
	"testing"

	schemalib "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-edgeadc/internal/provider/resource_ip_services"
	"terraform-provider-edgeadc/swagger"
)

// --------------------------------------------------------------------------
// FromIpService / ToIpService round-trip tests
// --------------------------------------------------------------------------

// TestFromIpService_PrimaryCheckedDefault verifies that FromIpService correctly
// converts the API's "Passive" value for primary_checked. This would have caught
// the original "inconsistent result after apply" bug where the plan had "" but
// the API returned "Passive".
func TestFromIpService_PrimaryCheckedPreserved(t *testing.T) {
	tests := []struct {
		name     string
		apiValue string
		want     string
	}{
		{"Passive from API", "Passive", "Passive"},
		{"Active from API", "Active", "Active"},
		{"Empty from API", "", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := swagger.IpService{
				IpAddr:         "10.0.0.1",
				Port:           "80",
				PrimaryChecked: tc.apiValue,
			}
			result := FromIpService(model)
			if result.PrimaryChecked.ValueString() != tc.want {
				t.Errorf("FromIpService PrimaryChecked = %q, want %q",
					result.PrimaryChecked.ValueString(), tc.want)
			}
		})
	}
}

// TestToIpService_AllFieldsMapped verifies that ToIpService maps every field
// from the Terraform model to the swagger model. Missing fields would cause
// data loss on Create/Update.
func TestToIpService_AllFieldsMapped(t *testing.T) {
	data := resource_ip_services.IpServicesResourceModel{
		SId:                       types.StringValue("1"),
		InterfaceID:               types.StringValue("0"),
		ChannelID:                 types.StringValue("0"),
		InterfaceKey:              types.StringValue("1"),
		ChannelKey:                types.StringValue("1.1"),
		IpAddr:                    types.StringValue("10.0.0.1"),
		MaskState:                 types.StringValue("0"),
		SubnetMask:                types.StringValue("255.255.255.0"),
		ServiceName:               types.StringValue("test-svc"),
		LocalPortEnabledChecked:   types.StringValue("true"),
		PrimaryChecked:            types.StringValue("Active"),
		ServiceType:               types.StringValue("HTTP"),
		MaxConn:                   types.StringValue("20"),
		Port:                      types.StringValue("80"),
		Offlinonfailure:           types.StringValue("1"),
		ContentServerGroupName:    types.StringValue("Server Group"),
		EnableConnectionPool:      types.StringValue("false"),
		ConnectionPoolSize:        types.StringValue("2000"),
		ServerMonitoring:          types.StringValue("Connect"),
		LoadBalancingPolicy:       types.StringValue("RoundRobin"),
		Connectivity:              types.StringValue("managed"),
		Acceleration:              types.StringValue("Compress"),
		SslCertificate:            types.StringValue("No SSL"),
		SslClientCertificate:      types.StringValue("No SSL"),
		CachingRule:               types.StringValue("0"),
		MonitoringInterval:        types.StringValue("1"),
		MonitoringTimeout:         types.StringValue("1"),
		ConnectionTimeout:         types.StringValue("600"),
		InCount:                   types.StringValue("2"),
		OutCount:                  types.StringValue("3"),
		CipherName:                types.StringValue("Defaults"),
		SNIDefaultCertificateName: types.StringValue("None"),
		SslRenegotiation:          types.StringValue("on"),
		SslResumption:             types.StringValue("on"),
		SecurityLog:               types.StringValue("On"),
	}

	result := ToIpService(data)

	checks := map[string]struct {
		got  string
		want string
	}{
		"SId":                       {result.SId, "1"},
		"IpAddr":                    {result.IpAddr, "10.0.0.1"},
		"Port":                      {result.Port, "80"},
		"PrimaryChecked":            {result.PrimaryChecked, "Active"},
		"ServiceType":               {result.ServiceType, "HTTP"},
		"ServiceName":               {result.ServiceName, "test-svc"},
		"ServerMonitoring":          {result.ServerMonitoring, "Connect"},
		"SslCertificate":            {result.SslCertificate, "No SSL"},
		"MaxConn":                   {result.MaxConn, "20"},
		"Connectivity":              {result.Connectivity, "managed"},
		"ConnectionTimeout":         {result.ConnectionTimeout, "600"},
		"MonitoringInterval":        {result.MonitoringInterval, "1"},
		"CipherName":                {result.CipherName, "Defaults"},
		"SNIDefaultCertificateName": {result.SNIDefaultCertificateName, "None"},
		"SslRenegotiation":          {result.SslRenegotiation, "on"},
		"SslResumption":             {result.SslResumption, "on"},
		"SecurityLog":               {result.SecurityLog, "On"},
	}

	for field, c := range checks {
		if c.got != c.want {
			t.Errorf("ToIpService field %s = %q, want %q", field, c.got, c.want)
		}
	}
}

// TestFromIpServiceToIpService_RoundTrip verifies that converting from swagger
// to terraform and back produces the same values.
func TestFromIpServiceToIpService_RoundTrip(t *testing.T) {
	original := swagger.IpService{
		SId:                       "1",
		InterfaceID:               "0",
		ChannelID:                 "0",
		IpAddr:                    "10.0.0.1",
		Port:                      "80",
		SubnetMask:                "255.255.255.0",
		ServiceName:               "test-svc",
		PrimaryChecked:            "Active",
		ServiceType:               "HTTP",
		LocalPortEnabledChecked:   "true",
		ServerMonitoring:          "Connect",
		SslCertificate:            "No SSL",
		SslClientCertificate:      "No SSL",
		LoadBalancingPolicy:       "RoundRobin",
		Acceleration:              "Compress",
		MaxConn:                   "20",
		Connectivity:              "managed",
		ConnectionTimeout:         "600",
		MonitoringInterval:        "1",
		MonitoringTimeout:         "1",
		CipherName:                "Defaults",
		SNIDefaultCertificateName: "None",
		SslRenegotiation:          "on",
		SslResumption:             "on",
		SecurityLog:               "On",
		InCount:                   "2",
		OutCount:                  "3",
	}

	terraform := FromIpService(original)
	roundTripped := ToIpService(terraform)

	fieldsToCheck := map[string]struct {
		got  string
		want string
	}{
		"IpAddr":           {roundTripped.IpAddr, original.IpAddr},
		"Port":             {roundTripped.Port, original.Port},
		"PrimaryChecked":   {roundTripped.PrimaryChecked, original.PrimaryChecked},
		"ServiceType":      {roundTripped.ServiceType, original.ServiceType},
		"ServiceName":      {roundTripped.ServiceName, original.ServiceName},
		"ServerMonitoring": {roundTripped.ServerMonitoring, original.ServerMonitoring},
		"SslCertificate":   {roundTripped.SslCertificate, original.SslCertificate},
		"MaxConn":          {roundTripped.MaxConn, original.MaxConn},
		"CipherName":       {roundTripped.CipherName, original.CipherName},
		"SslRenegotiation": {roundTripped.SslRenegotiation, original.SslRenegotiation},
	}

	for field, c := range fieldsToCheck {
		if c.got != c.want {
			t.Errorf("Round-trip field %s = %q, want %q", field, c.got, c.want)
		}
	}
}

// --------------------------------------------------------------------------
// MergeBasicTabs tests
// --------------------------------------------------------------------------

// TestMergeBasicTabs_OverlaysSetValues verifies that MergeBasicTabs correctly
// overlays user-specified values onto the API-returned defaults.
func TestMergeBasicTabs_OverlaysSetValues(t *testing.T) {
	apiDefaults := swagger.UpdateBasicTab{
		EditedInterface:      "0",
		EditedChannel:        "0",
		ServerMonitoring:     "Connect",
		Acceleration:         "Compress",
		LoadBalancingPolicy:  "RoundRobin",
		SslCertificate:       "No SSL",
		SslClientCertificate: "No SSL",
		CachingRule:          "0",
	}

	userOverrides := resource_ip_services.IpServicesResourceModel{
		ServerMonitoring:     types.StringValue("Check200"),
		Acceleration:         types.StringValue("None"),
		LoadBalancingPolicy:  types.StringValue("LeastConnections"),
		SslCertificate:       types.StringValue("MyCert"),
		SslClientCertificate: types.StringValue("ClientCert"),
		CachingRule:          types.StringValue("1"),
	}

	result := MergeBasicTabs(apiDefaults, userOverrides)

	if result.ServerMonitoring != "Check200" {
		t.Errorf("ServerMonitoring = %q, want %q", result.ServerMonitoring, "Check200")
	}
	if result.SslCertificate != "MyCert" {
		t.Errorf("SslCertificate = %q, want %q", result.SslCertificate, "MyCert")
	}
	if result.LoadBalancingPolicy != "LeastConnections" {
		t.Errorf("LoadBalancingPolicy = %q, want %q", result.LoadBalancingPolicy, "LeastConnections")
	}
}

// TestMergeBasicTabs_SkipsNullValues verifies that null (unset) Terraform values
// do NOT overwrite API defaults with empty strings. This was a bug where null
// config values would send empty strings to the API.
func TestMergeBasicTabs_SkipsNullValues(t *testing.T) {
	apiDefaults := swagger.UpdateBasicTab{
		ServerMonitoring:    "Connect",
		Acceleration:        "Compress",
		LoadBalancingPolicy: "RoundRobin",
		SslCertificate:      "No SSL",
	}

	// Simulate user not setting any basic tab values (all null)
	nullData := resource_ip_services.IpServicesResourceModel{
		ServerMonitoring:     types.StringNull(),
		Acceleration:         types.StringNull(),
		LoadBalancingPolicy:  types.StringNull(),
		SslCertificate:       types.StringNull(),
		SslClientCertificate: types.StringNull(),
		CachingRule:          types.StringNull(),
	}

	result := MergeBasicTabs(apiDefaults, nullData)

	if result.ServerMonitoring != "Connect" {
		t.Errorf("ServerMonitoring was overwritten by null, got %q want %q", result.ServerMonitoring, "Connect")
	}
	if result.SslCertificate != "No SSL" {
		t.Errorf("SslCertificate was overwritten by null, got %q want %q", result.SslCertificate, "No SSL")
	}
}

// TestMergeBasicTabs_SkipsUnknownValues verifies that unknown Terraform values
// do NOT overwrite API defaults.
func TestMergeBasicTabs_SkipsUnknownValues(t *testing.T) {
	apiDefaults := swagger.UpdateBasicTab{
		ServerMonitoring:    "Connect",
		Acceleration:        "Compress",
		SslCertificate:      "No SSL",
	}

	unknownData := resource_ip_services.IpServicesResourceModel{
		ServerMonitoring:     types.StringUnknown(),
		Acceleration:         types.StringUnknown(),
		LoadBalancingPolicy:  types.StringUnknown(),
		SslCertificate:       types.StringUnknown(),
		SslClientCertificate: types.StringUnknown(),
		CachingRule:          types.StringUnknown(),
	}

	result := MergeBasicTabs(apiDefaults, unknownData)

	if result.ServerMonitoring != "Connect" {
		t.Errorf("ServerMonitoring was overwritten by unknown, got %q want %q", result.ServerMonitoring, "Connect")
	}
	if result.SslCertificate != "No SSL" {
		t.Errorf("SslCertificate was overwritten by unknown, got %q want %q", result.SslCertificate, "No SSL")
	}
}

// TestMergeAdvanceTabs_SkipsNullValues mirrors the basic tabs test for the
// advance tab merge.
func TestMergeAdvanceTabs_SkipsNullValues(t *testing.T) {
	apiDefaults := swagger.UpdateAdvanceTab{
		Connectivity:      "managed",
		ConnectionTimeout: "600",
		MaxConn:           "20",
		CipherName:        "Defaults",
	}

	nullData := resource_ip_services.IpServicesResourceModel{
		Connectivity:              types.StringNull(),
		ConnectionTimeout:         types.StringNull(),
		MonitoringInterval:        types.StringNull(),
		MonitoringTimeout:         types.StringNull(),
		MaxConn:                   types.StringNull(),
		InCount:                   types.StringNull(),
		OutCount:                  types.StringNull(),
		CipherName:                types.StringNull(),
		SecurityLog:               types.StringNull(),
		SslRenegotiation:          types.StringNull(),
		SNIDefaultCertificateName: types.StringNull(),
		SslResumption:             types.StringNull(),
		Offlinonfailure:           types.StringNull(),
	}

	result := MergeAdvanceTabs(apiDefaults, nullData)

	if result.Connectivity != "managed" {
		t.Errorf("Connectivity was overwritten by null, got %q want %q", result.Connectivity, "managed")
	}
	if result.ConnectionTimeout != "600" {
		t.Errorf("ConnectionTimeout was overwritten by null, got %q want %q", result.ConnectionTimeout, "600")
	}
	if result.MaxConn != "20" {
		t.Errorf("MaxConn was overwritten by null, got %q want %q", result.MaxConn, "20")
	}
}

// --------------------------------------------------------------------------
// ConvertComboOptions tests
// --------------------------------------------------------------------------

// TestConvertComboOptionsToNames_EmptyInput verifies that empty input returns
// empty output without error.
func TestConvertComboOptionsToNames_EmptyInput(t *testing.T) {
	comboOptions := swagger.IpServicescomboServiceTypeComboOptions{
		Option: []swagger.TypeComboOne{
			{Id: "2", Value: "Connect"},
		},
	}

	result, err := ConvertComboOptionsToNames(comboOptions, "")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

// TestConvertComboOptionsToNames_Success verifies correct ID-to-name conversion.
func TestConvertComboOptionsToNames_Success(t *testing.T) {
	comboOptions := swagger.IpServicescomboServiceTypeComboOptions{
		Option: []swagger.TypeComboOne{
			{Id: "2", Value: "Connect"},
			{Id: "4", Value: "Check200"},
		},
	}

	result, err := ConvertComboOptionsToNames(comboOptions, "2,4")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != "Connect,Check200" {
		t.Errorf("got %q, want %q", result, "Connect,Check200")
	}
}

// TestConvertComboOptionsToNames_UnknownId verifies that an unknown ID returns
// an error instead of silently producing partial results. This was a bug that
// caused server_monitoring state inconsistencies.
func TestConvertComboOptionsToNames_UnknownId(t *testing.T) {
	comboOptions := swagger.IpServicescomboServiceTypeComboOptions{
		Option: []swagger.TypeComboOne{
			{Id: "2", Value: "Connect"},
		},
	}

	_, err := ConvertComboOptionsToNames(comboOptions, "999")
	if err == nil {
		t.Error("expected error for unknown ID, got nil")
	}
}

// TestConvertComboOptionsToIds_Success verifies correct name-to-ID conversion.
func TestConvertComboOptionsToIds_Success(t *testing.T) {
	comboOptions := swagger.IpServicescomboServiceTypeComboOptions{
		Option: []swagger.TypeComboOne{
			{Id: "2", Value: "Connect"},
			{Id: "4", Value: "Check200"},
		},
	}

	result, err := ConvertComboOptionsToIds(comboOptions, "Connect,Check200")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != "2,4" {
		t.Errorf("got %q, want %q", result, "2,4")
	}
}

// TestConvertComboOptionsToIds_UnknownName verifies that an unknown name returns
// an error. Previously this was silently ignored, causing the API to receive
// empty strings and the subsequent read to return different values.
func TestConvertComboOptionsToIds_UnknownName(t *testing.T) {
	comboOptions := swagger.IpServicescomboServiceTypeComboOptions{
		Option: []swagger.TypeComboOne{
			{Id: "2", Value: "Connect"},
		},
	}

	_, err := ConvertComboOptionsToIds(comboOptions, "NonExistentMonitor")
	if err == nil {
		t.Error("expected error for unknown name, got nil")
	}
}

// --------------------------------------------------------------------------
// IP Service lookup tests
// --------------------------------------------------------------------------

// TestGetIpServiceByAddressAndPort_NotFound verifies proper error message when
// IP service doesn't exist. This error prefix is used by delete functions to
// decide whether to treat the error as "already deleted".
func TestGetIpServiceByAddressAndPort_NotFound(t *testing.T) {
	ipServices := swagger.IpServices{
		Data: &swagger.IpServicesData{
			Dataset: &swagger.IpServicesDataDataset{
				IpService: [][]swagger.IpService{},
			},
		},
	}

	_, err := GetIpServiceByAddressAndPort(ipServices, "10.0.0.1", "80")
	if err == nil {
		t.Error("expected error for non-existent IP service, got nil")
	}
	// Verify the error message starts with errServiceNotFound so that
	// delete handling can recognise it
	if len(err.Error()) < len(errServiceNotFound) || err.Error()[:len(errServiceNotFound)] != errServiceNotFound {
		t.Errorf("error %q should start with %q", err.Error(), errServiceNotFound)
	}
}

// TestGetIpServiceByAddressAndPort_NilData verifies graceful handling of nil data.
func TestGetIpServiceByAddressAndPort_NilData(t *testing.T) {
	ipServices := swagger.IpServices{Data: nil}

	_, err := GetIpServiceByAddressAndPort(ipServices, "10.0.0.1", "80")
	if err == nil {
		t.Error("expected error for nil data, got nil")
	}
}

// --------------------------------------------------------------------------
// ToCopyIp tests
// --------------------------------------------------------------------------

// TestToCopyIp_IncludesPrimaryChecked verifies that ToCopyIp includes the
// primary_checked value. A missing or empty primary_checked was the root
// cause of the original "inconsistent result" bug.
func TestToCopyIp_IncludesPrimaryChecked(t *testing.T) {
	data := resource_ip_services.IpServicesResourceModel{
		InterfaceID:             types.StringValue("0"),
		ChannelID:               types.StringValue("0"),
		IpAddr:                  types.StringValue("10.0.0.1"),
		SubnetMask:              types.StringValue("255.255.255.0"),
		ServiceName:             types.StringValue("test"),
		LocalPortEnabledChecked: types.StringValue("true"),
		Port:                    types.StringValue("80"),
		PrimaryChecked:          types.StringValue("Active"),
		ServiceType:             types.StringValue("HTTP"),
	}

	result := ToCopyIp(data)
	if result.PrimaryChecked != "Active" {
		t.Errorf("ToCopyIp PrimaryChecked = %q, want %q", result.PrimaryChecked, "Active")
	}
}

// TestPrimaryChecked_EmptyStringBecomesUnknown verifies that the
// emptyStringToUnknown plan modifier on primary_checked converts ""
// to unknown. This is critical because the ADC API normalizes "" to
// either "Active" or "Passive" depending on context, and a static
// plan value of "" would cause an inconsistency with whatever the
// API returns.
//
// This test would have caught: "primary_checked: was cty.StringVal(""),
// but now cty.StringVal("Active")"
func TestPrimaryChecked_EmptyStringBecomesUnknown(t *testing.T) {
	ctx := context.Background()
	schema := resource_ip_services.IpServiceResourceSchema(ctx)

	pcAttr, ok := schema.Attributes["primary_checked"]
	if !ok {
		t.Fatal("schema missing 'primary_checked' attribute")
	}

	strAttr, ok := pcAttr.(schemalib.StringAttribute)
	if !ok {
		t.Fatal("'primary_checked' is not a StringAttribute")
	}

	// Verify the field is Optional+Computed
	if !strAttr.Optional || !strAttr.Computed {
		t.Error("primary_checked should be Optional+Computed")
	}

	// Verify there's no static default (static defaults don't work because
	// the API can return either "Active" or "Passive")
	if strAttr.Default != nil {
		t.Error("primary_checked should NOT have a static Default -- the API can return either 'Active' or 'Passive'")
	}

	// Verify that one of the plan modifiers handles empty strings
	// by simulating a plan modify with ""
	found := false
	for _, pm := range strAttr.PlanModifiers {
		req := planmodifier.StringRequest{
			PlanValue: types.StringValue(""),
		}
		resp := &planmodifier.StringResponse{
			PlanValue: types.StringValue(""),
		}
		pm.PlanModifyString(ctx, req, resp)
		if resp.PlanValue.IsUnknown() {
			found = true
			break
		}
	}
	if !found {
		t.Error("primary_checked must have a plan modifier that converts '' to unknown, " +
			"otherwise the API returning 'Active' or 'Passive' will cause an inconsistency")
	}
}

// TestPrimaryChecked_NonEmptyPassesThrough verifies that valid non-empty
// values like "Active" and "Passive" pass through the plan modifier unchanged.
func TestPrimaryChecked_NonEmptyPassesThrough(t *testing.T) {
	ctx := context.Background()
	schema := resource_ip_services.IpServiceResourceSchema(ctx)

	pcAttr := schema.Attributes["primary_checked"]
	strAttr := pcAttr.(schemalib.StringAttribute)

	for _, val := range []string{"Active", "Passive"} {
		req := planmodifier.StringRequest{
			PlanValue: types.StringValue(val),
		}
		resp := &planmodifier.StringResponse{
			PlanValue: types.StringValue(val),
		}
		for _, pm := range strAttr.PlanModifiers {
			pm.PlanModifyString(ctx, req, resp)
		}
		if resp.PlanValue.IsUnknown() {
			t.Errorf("plan modifier should not convert %q to unknown", val)
		}
		if resp.PlanValue.ValueString() != val {
			t.Errorf("plan modifier changed %q to %q", val, resp.PlanValue.ValueString())
		}
	}
}

// TestGetAddressAndPortFromId verifies the ID parsing used by all resources.
func TestGetAddressAndPortFromId(t *testing.T) {
	tests := []struct {
		id       string
		wantAddr string
		wantPort string
		wantErr  bool
	}{
		{"10.0.0.1:80", "10.0.0.1", "80", false},
		{"192.168.1.1:443", "192.168.1.1", "443", false},
		{"invalid", "", "", true},
		{"too:many:colons", "", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.id, func(t *testing.T) {
			addr, port, err := GetAddressAndPortFromId(tc.id)
			if tc.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if addr != tc.wantAddr || port != tc.wantPort {
				t.Errorf("got (%q, %q), want (%q, %q)", addr, port, tc.wantAddr, tc.wantPort)
			}
		})
	}
}
