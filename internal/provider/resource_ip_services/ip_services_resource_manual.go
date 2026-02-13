package resource_ip_services

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure types import is used (referenced by IpServicesResourceModel below).
var _ = types.StringType

func IpServiceResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"s_id": schema.StringAttribute{
				Optional: false,
				Computed: true,
				Required: false,
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
			"interface_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"channel_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ip_addr": schema.StringAttribute{
				Required: true,
				Computed: false,
			},
			"mask_state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"subnet_mask": schema.StringAttribute{
				Required: true,
				Computed: false,
			},
			"service_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"local_port_enabled_checked": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// primary_checked controls whether the VIP is Active or Passive.
			// If omitted, the API assigns a default (typically "Active").
			// Users should set "Active" or "Passive" explicitly, or omit
			// the field to accept the server default.
			"primary_checked": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"service_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"max_conn": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"port": schema.StringAttribute{
				Required: true,
				Computed: false,
			},
			"offlinonfailure": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"content_server_group_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enable_connection_pool": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"connection_pool_size": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"server_monitoring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of custom monitor names",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"load_balancing_policy": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"connectivity": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"acceleration": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ssl_certificate": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ssl_client_certificate": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"caching_rule": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"monitoring_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"monitoring_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"connection_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"in_count": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"out_count": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cipher_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"sni_default_certificate_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ssl_renegotiation": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ssl_resumption": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"security_log": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			//"flight_path": schema.ObjectAttribute{
			//	Optional:       true,
			//	Computed:       true,
			//	AttributeTypes: map[string]schema.Attribute{
			//		// Define fields as per IpServiceFlightPath struct
			//	},
			//},
			//"content_server": schema.ObjectAttribute{
			//	Optional:       true,
			//	Computed:       true,
			//	AttributeTypes: map[string]schema.Attribute{
			//		// Define fields as per IpServiceContentServer struct
			//	},
			//},
		},
	}
}

type IpServicesResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	SId                       types.String `tfsdk:"s_id"`
	InterfaceID               types.String `tfsdk:"interface_id"`
	ChannelID                 types.String `tfsdk:"channel_id"`
	InterfaceKey              types.String `tfsdk:"interface_key"`
	ChannelKey                types.String `tfsdk:"channel_key"`
	IpAddr                    types.String `tfsdk:"ip_addr"`
	MaskState                 types.String `tfsdk:"mask_state"`
	SubnetMask                types.String `tfsdk:"subnet_mask"`
	ServiceName               types.String `tfsdk:"service_name"`
	LocalPortEnabledChecked   types.String `tfsdk:"local_port_enabled_checked"`
	PrimaryChecked            types.String `tfsdk:"primary_checked"`
	ServiceType               types.String `tfsdk:"service_type"`
	MaxConn                   types.String `tfsdk:"max_conn"`
	Port                      types.String `tfsdk:"port"`
	Offlinonfailure           types.String `tfsdk:"offlinonfailure"`
	ContentServerGroupName    types.String `tfsdk:"content_server_group_name"`
	EnableConnectionPool      types.String `tfsdk:"enable_connection_pool"`
	ConnectionPoolSize        types.String `tfsdk:"connection_pool_size"`
	ServerMonitoring          types.String `tfsdk:"server_monitoring"`
	LoadBalancingPolicy       types.String `tfsdk:"load_balancing_policy"`
	Connectivity              types.String `tfsdk:"connectivity"`
	Acceleration              types.String `tfsdk:"acceleration"`
	SslCertificate            types.String `tfsdk:"ssl_certificate"`
	SslClientCertificate      types.String `tfsdk:"ssl_client_certificate"`
	CachingRule               types.String `tfsdk:"caching_rule"`
	MonitoringInterval        types.String `tfsdk:"monitoring_interval"`
	MonitoringTimeout         types.String `tfsdk:"monitoring_timeout"`
	ConnectionTimeout         types.String `tfsdk:"connection_timeout"`
	InCount                   types.String `tfsdk:"in_count"`
	OutCount                  types.String `tfsdk:"out_count"`
	CipherName                types.String `tfsdk:"cipher_name"`
	SNIDefaultCertificateName types.String `tfsdk:"sni_default_certificate_name"`
	SslRenegotiation          types.String `tfsdk:"ssl_renegotiation"`
	SslResumption             types.String `tfsdk:"ssl_resumption"`
	SecurityLog               types.String `tfsdk:"security_log"`
}
