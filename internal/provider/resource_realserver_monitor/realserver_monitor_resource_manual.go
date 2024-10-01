package resource_realserver_monitor

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func RealserverMonitorResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    false,
				Computed:    true,
				Description: "The id assigned to the realserver monitor",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "200OK",
			},
			"description": schema.StringAttribute{
				Required:            true,
				Description:         "Check home page for 200 OK",
				MarkdownDescription: "0",
			},
			"type": schema.StringAttribute{
				Required:            true,
				Description:         "Check200",
				MarkdownDescription: "0",
			},
			"url": schema.StringAttribute{
				Required:            true,
				Description:         "/",
				MarkdownDescription: "0",
			},
			"content": schema.StringAttribute{
				Required:            true,
				Description:         "",
				MarkdownDescription: "0",
			},
			"username": schema.StringAttribute{
				Required:            true,
				Description:         "",
				MarkdownDescription: "0",
			},
			"password": schema.StringAttribute{
				Required:            true,
				Description:         "",
				MarkdownDescription: "0",
			},
			"threshold": schema.StringAttribute{
				Required:            true,
				Description:         "",
				MarkdownDescription: "0",
			},
		},
	}
}

type RealserverMonitorModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Type        types.String `tfsdk:"type"`
	Url         types.String `tfsdk:"url"`
	Content     types.String `tfsdk:"content"`
	Username    types.String `tfsdk:"username"`
	Password    types.String `tfsdk:"password"`
	Threshold   types.String `tfsdk:"threshold"`
}
