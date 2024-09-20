package resource_ip_services

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
			},
			"channel_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
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
			},
			"local_port_enabled_checked": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"primary_checked": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"service_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"max_conn": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"port": schema.StringAttribute{
				Required: true,
				Computed: false,
			},
			"offlinonfailure": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"content_server_group_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"enable_connection_pool": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"connection_pool_size": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"server_monitoring": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"load_balancing_policy": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"connectivity": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"acceleration": schema.StringAttribute{
				Optional: true,
				Computed: true,
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
			},
			"caching_rule": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"monitoring_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"monitoring_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"connection_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"in_count": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"out_count": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cipher_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"sni_default_certificate_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"ssl_renegotiation": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"ssl_resumption": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"security_log": schema.StringAttribute{
				Optional: true,
				Computed: true,
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
