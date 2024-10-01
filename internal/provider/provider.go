package provider

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = (*edgeadcProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &edgeadcProvider{}
	}
}

type edgeadcProvider struct{}

type edgeadcProviderModel struct {
	Endpoint             types.String `tfsdk:"endpoint"`
	Username             types.String `tfsdk:"username"`
	Password             types.String `tfsdk:"password"`
	SkipCertVerification types.Bool   `tfsdk:"skip_cert_verification"`
}

func (p *edgeadcProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "EdgeADC Endpoint",
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "API Username",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "API Password",
				Optional:            true,
			},
			"skip_cert_verification": schema.BoolAttribute{
				MarkdownDescription: "Whether to skip TLS verification",
				Optional:            true,
			},
		},
	}
}

func (p *edgeadcProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config edgeadcProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Endpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown EdgeADC Endpoint",
			"The provider cannot create the EdgeADC Endpoint API client as there is an unknown configuration value. ",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown EdgeADC Username",
			"The provider cannot create the EdgeADC Endpoint API client as there is an unknown configuration value. ",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown EdgeADC Password",
			"The provider cannot create the EdgeADC Endpoint API client as there is an unknown configuration value. ",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	endpoint := os.Getenv("EDGEADC_ENDPOINT")
	username := os.Getenv("EDGEADC_USERNAME")
	password := os.Getenv("EDGEADC_PASSWORD")

	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown EdgeADC Endpoint",
			"The provider cannot create the EdgeADC Endpoint API client as there is an unknown configuration value. ",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown EdgeADC Username",
			"The provider cannot create the EdgeADC Endpoint API client as there is an unknown configuration value. ",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown EdgeADC Password",
			"The provider cannot create the EdgeADC Endpoint API client as there is an unknown configuration value. ",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new client using the configuration values
	client := NewAPI(endpoint, username, password, "")

	// Disable TLS verification if the configuration value is set
	if !config.SkipCertVerification.IsNull() {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: config.SkipCertVerification.ValueBool()},
		}
		client.client = &http.Client{Transport: tr}
	}

	// Set the context for logging purposes
	client.loggingContext = ctx

	// Make the client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *edgeadcProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "edgeadc"
}

func (p *edgeadcProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *edgeadcProvider) Resources(ctx context.Context) []func() resource.Resource {
	//return []func() resource.Resource{}
	return []func() resource.Resource{
		NewUsersResource,
		NewIpServicesResource,
		NewRealserverMonitorResource,
		NewServerResource,
		NewSslCertificatesResource,
		NewCustomMonitorResource,
	}
}
