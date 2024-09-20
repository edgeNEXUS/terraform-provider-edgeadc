package resource_ssl_certificates

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func CertificatesResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Computed:    false,
				Description: "The ID (unique name) of the certificate",
			},
			"password": schema.StringAttribute{
				Required:    true,
				Computed:    false,
				Sensitive:   true,
				Description: "The password for the certificate",
			},
			"file_path": schema.StringAttribute{
				Required:    true,
				Computed:    false,
				Sensitive:   false,
				Description: "The path to a PFX certificate",
			},
		},
	}
}

type CertificatesResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Password types.String `tfsdk:"password"`
	FilePath types.String `tfsdk:"file_path"`
}
