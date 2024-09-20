package resource_server

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ServerResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ip_service": schema.StringAttribute{
				Optional:    false,
				Computed:    false,
				Required:    true,
				Description: "The ip service to associate this server",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_id": schema.StringAttribute{
				Optional: false,
				Computed: true,
				Required: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"channel_id": schema.StringAttribute{
				Optional: false,
				Computed: true,
				Required: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"server_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"c_id": schema.StringAttribute{
				Optional: false,
				Computed: true,
				Required: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cs_activity": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cs_ip_addr": schema.StringAttribute{
				Required: true,
				Computed: false,
			},
			"cs_port": schema.StringAttribute{
				Required: true,
				Computed: false,
			},
			"cs_notes": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"server_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"weight_factor": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}

type ServerResourceModel struct {
	Id           types.String `tfsdk:"id"`
	IpService    types.String `tfsdk:"ip_service"`
	InterfaceId  types.String `tfsdk:"interface_id"`
	ChannelId    types.String `tfsdk:"channel_id"`
	ServerKey    types.String `tfsdk:"server_key"`
	CId          types.String `tfsdk:"c_id"`
	CSActivity   types.String `tfsdk:"cs_activity"`
	CSIPAddr     types.String `tfsdk:"cs_ip_addr"`
	CSPort       types.String `tfsdk:"cs_port"`
	CSNotes      types.String `tfsdk:"cs_notes"`
	ServerId     types.String `tfsdk:"server_id"`
	WeightFactor types.String `tfsdk:"weight_factor"`
}
