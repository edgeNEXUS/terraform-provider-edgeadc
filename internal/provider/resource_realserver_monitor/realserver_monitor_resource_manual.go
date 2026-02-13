package resource_realserver_monitor

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// monitorTypeNormalizer is a plan modifier that converts friendly UI monitor
// type names (e.g. "HTTP Head") to their backend equivalents (e.g. "CheckHead")
// during the plan phase. This ensures the planned value matches what the API
// will store and return, preventing "inconsistent result after apply" errors.
type monitorTypeNormalizer struct{}

func (m monitorTypeNormalizer) Description(_ context.Context) string {
	return "Normalizes friendly monitor type names to backend names"
}

func (m monitorTypeNormalizer) MarkdownDescription(_ context.Context) string {
	return "Normalizes friendly monitor type names to backend names"
}

func (m monitorTypeNormalizer) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}
	// Map of friendly UI names to backend names
	friendly := map[string]string{
		"HTTP 200 OK":   "Check200",
		"HTTP Head":     "CheckHead",
		"HTTP Response": "CheckResponse",
		"TCP Connect":   "Connect",
		"ICMP Ping":     "Ping",
		"None":          "None",
		"DICOM":         "DICOM",
		"HTTP Post":     "CheckPOST",
	}
	if backend, ok := friendly[req.PlanValue.ValueString()]; ok {
		resp.PlanValue = types.StringValue(backend)
	}
}

func RealserverMonitorResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    false,
				Computed:    true,
				Description: "The internal ADC id assigned to the realserver monitor. This value may change when monitors are added or removed.",
				// Note: UseStateForUnknown is deliberately NOT used here.
				// The ADC renumbers monitor IDs when monitors are added or
				// removed, so the ID can change between applies. Leaving it
				// as unknown in the plan allows any returned value to be accepted.
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the realserver monitor",
			},
			"description": schema.StringAttribute{
				Required:    true,
				Description: "A description of what this monitor checks",
			},
			"type": schema.StringAttribute{
				Required: true,
				Description: "The monitor type. Accepts either backend names or friendly UI names. " +
					"Backend names: Check200, CheckHead, CheckResponse, Connect, Ping, None, DICOM, CheckPOST. " +
					"Friendly names: HTTP 200 OK, HTTP Head, HTTP Response, TCP Connect, ICMP Ping, None, DICOM, HTTP Post. " +
					"Custom monitor script names are also accepted.",
				PlanModifiers: []planmodifier.String{
					monitorTypeNormalizer{},
				},
			},
			"ssl": schema.StringAttribute{
				Optional:            true,
				Description:         "SSL/TLS setting for the monitor. Valid values: auto, secured, unsecured",
				MarkdownDescription: "SSL/TLS setting for the monitor. Valid values: `auto`, `secured`, `unsecured`",
			},
			"url": schema.StringAttribute{
				Required:    true,
				Description: "The URL path to monitor (e.g. /health)",
			},
			"content": schema.StringAttribute{
				Required:    true,
				Description: "The expected content in the response",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Username for authenticated monitors (empty string if not needed)",
			},
			"password": schema.StringAttribute{
				Required:    true,
				Description: "Password for authenticated monitors (empty string if not needed)",
			},
			"threshold": schema.StringAttribute{
				Required:    true,
				Description: "Number of consecutive failures before marking server as down",
			},
		},
	}
}

type RealserverMonitorModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Type        types.String `tfsdk:"type"`
	Ssl         types.String `tfsdk:"ssl"`
	Url         types.String `tfsdk:"url"`
	Content     types.String `tfsdk:"content"`
	Username    types.String `tfsdk:"username"`
	Password    types.String `tfsdk:"password"`
	Threshold   types.String `tfsdk:"threshold"`
}
