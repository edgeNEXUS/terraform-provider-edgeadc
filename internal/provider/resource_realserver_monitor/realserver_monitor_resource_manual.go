package resource_realserver_monitor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
					"Custom monitor script names are also accepted. " +
					"Friendly names are normalized to backend names when sent to the API.",
			},
			"ssl": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "SSL/TLS setting for the monitor. Valid values: auto, secured, unsecured. Defaults to auto if not specified.",
				MarkdownDescription: "SSL/TLS setting for the monitor. Valid values: `auto`, `secured`, `unsecured`. Defaults to `auto` if not specified.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
