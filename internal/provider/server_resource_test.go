package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-edgeadc/internal/provider/resource_server"
	"terraform-provider-edgeadc/swagger"
)

// TestToUpdateServer_AllFieldsMapped verifies that ToUpdateServer maps every
// user-specified field from the Terraform model to the swagger UpdateServer
// struct. This would have caught the critical bug where UpdateServer created
// an empty struct and only set the IDs, discarding all actual field values.
func TestToUpdateServer_AllFieldsMapped(t *testing.T) {
	r := &serverResource{}

	data := resource_server.ServerResourceModel{
		Id:                types.StringValue("10.0.0.2:8080"),
		IpService:         types.StringValue("10.0.0.1:80"),
		InterfaceId:       types.StringValue("0"),
		ChannelId:         types.StringValue("0"),
		ServerKey:         types.StringValue("1.1.1"),
		CId:               types.StringValue("0"),
		CSActivity:        types.StringValue("1"),
		CSIPAddr:          types.StringValue("10.0.0.2"),
		CSPort:            types.StringValue("8080"),
		CSNotes:           types.StringValue("backend server 1"),
		CSMonitorEndPoint: types.StringValue("10.0.0.3"),
		WeightFactor:      types.StringValue("75"),
	}

	result := r.ToUpdateServer(data)

	checks := map[string]struct {
		got  string
		want string
	}{
		"CSActivity":        {result.CSActivity, "1"},
		"CSIPAddr":          {result.CSIPAddr, "10.0.0.2"},
		"CSPort":            {result.CSPort, "8080"},
		"CSNotes":           {result.CSNotes, "backend server 1"},
		"WeightFactor":      {result.WeightFactor, "75"},
		"CSMonitorEndPoint": {result.CSMonitorEndPoint, "10.0.0.3"},
	}

	for field, c := range checks {
		if c.got != c.want {
			t.Errorf("ToUpdateServer field %s = %q, want %q", field, c.got, c.want)
		}
	}
}

// TestToUpdateServer_EmptyFieldsAreEmpty verifies that when the user doesn't
// set optional fields, they come through as empty strings rather than garbage.
func TestToUpdateServer_EmptyFieldsAreEmpty(t *testing.T) {
	r := &serverResource{}

	data := resource_server.ServerResourceModel{
		InterfaceId:       types.StringValue(""),
		ChannelId:         types.StringValue(""),
		CSActivity:        types.StringValue(""),
		CSIPAddr:          types.StringValue("10.0.0.2"),
		CSPort:            types.StringValue("80"),
		CSNotes:           types.StringValue(""),
		CSMonitorEndPoint: types.StringValue(""),
		WeightFactor:      types.StringValue(""),
	}

	result := r.ToUpdateServer(data)

	if result.CSActivity != "" {
		t.Errorf("CSActivity should be empty, got %q", result.CSActivity)
	}
	if result.CSNotes != "" {
		t.Errorf("CSNotes should be empty, got %q", result.CSNotes)
	}
	if result.WeightFactor != "" {
		t.Errorf("WeightFactor should be empty, got %q", result.WeightFactor)
	}
}

// TestFromCServerId_AllFieldsMapped verifies that the swagger-to-terraform
// conversion for servers includes all fields.
func TestFromCServerId_AllFieldsMapped(t *testing.T) {
	r := &serverResource{}

	input := swagger.CServerId{
		ServerKey:    "1.1.1",
		CId:          "0",
		CSActivity:   "1",
		CSIPAddr:     "10.0.0.2",
		CSPort:       "8080",
		CSNotes:      "test server",
		WeightFactor: "100",
	}

	result := r.FromCServerId(input, "0", "0")

	if result.Id.ValueString() != "10.0.0.2:8080" {
		t.Errorf("Id = %q, want %q", result.Id.ValueString(), "10.0.0.2:8080")
	}
	if result.CSActivity.ValueString() != "1" {
		t.Errorf("CSActivity = %q, want %q", result.CSActivity.ValueString(), "1")
	}
	if result.CSIPAddr.ValueString() != "10.0.0.2" {
		t.Errorf("CSIPAddr = %q, want %q", result.CSIPAddr.ValueString(), "10.0.0.2")
	}
	if result.WeightFactor.ValueString() != "100" {
		t.Errorf("WeightFactor = %q, want %q", result.WeightFactor.ValueString(), "100")
	}
}

// TestToCServerId_AllFieldsMapped verifies that ToCServerId maps all fields
// needed for server lookups.
func TestToCServerId_AllFieldsMapped(t *testing.T) {
	r := &serverResource{}

	data := resource_server.ServerResourceModel{
		ServerKey:    types.StringValue("1.1.1"),
		CId:         types.StringValue("0"),
		CSActivity:  types.StringValue("1"),
		CSIPAddr:    types.StringValue("10.0.0.2"),
		CSPort:      types.StringValue("8080"),
		CSNotes:     types.StringValue("test server"),
		WeightFactor: types.StringValue("100"),
	}

	result := r.ToCServerId(data)

	if result.CSIPAddr != "10.0.0.2" {
		t.Errorf("CSIPAddr = %q, want %q", result.CSIPAddr, "10.0.0.2")
	}
	if result.CSPort != "8080" {
		t.Errorf("CSPort = %q, want %q", result.CSPort, "8080")
	}
	if result.WeightFactor != "100" {
		t.Errorf("WeightFactor = %q, want %q", result.WeightFactor, "100")
	}
}

// TestGetServerByAddressAndPortFromIpServices_NotFound verifies that looking
// up a non-existent server returns an error with the correct prefix for
// graceful delete handling.
func TestGetServerByAddressAndPortFromIpServices_NotFound(t *testing.T) {
	ipServices := swagger.IpServices{
		Data: &swagger.IpServicesData{
			Dataset: &swagger.IpServicesDataDataset{
				IpService: [][]swagger.IpService{
					{
						{
							IpAddr: "10.0.0.1",
							Port:   "80",
							ContentServer: &swagger.IpServiceContentServer{
								CServerId: []swagger.CServerId{
									{CSIPAddr: "10.0.0.2", CSPort: "8080"},
								},
							},
						},
					},
				},
			},
		},
	}

	_, _, err := GetServerByAddressAndPortFromIpServices(ipServices, "10.0.0.1", "80", "10.0.0.99", "9999")
	if err == nil {
		t.Error("expected error for non-existent server, got nil")
	}
	if len(err.Error()) < len(errServerNotFound) || err.Error()[:len(errServerNotFound)] != errServerNotFound {
		t.Errorf("error %q should start with %q", err.Error(), errServerNotFound)
	}
}

// TestGetServerByAddressAndPortFromIpServices_NilData verifies graceful handling
// of nil data without panic.
func TestGetServerByAddressAndPortFromIpServices_NilData(t *testing.T) {
	ipServices := swagger.IpServices{Data: nil}

	_, _, err := GetServerByAddressAndPortFromIpServices(ipServices, "10.0.0.1", "80", "10.0.0.2", "8080")
	if err == nil {
		t.Error("expected error for nil data, got nil")
	}
}
