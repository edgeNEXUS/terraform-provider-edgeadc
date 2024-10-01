package resource_custom_monitor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func CustomMonitorResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    false,
				Computed:    true,
				Description: "The id assigned to the custom monitor",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Computed:    false,
				Description: "The unique name of the custom monitor",
			},
			"file_path": schema.StringAttribute{
				Required:    true,
				Computed:    false,
				Sensitive:   false,
				Description: "The path to a perl based custom monitor file",
			},
		},
	}
}

type CustomMonitorResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	FilePath types.String `tfsdk:"file_path"`
}
