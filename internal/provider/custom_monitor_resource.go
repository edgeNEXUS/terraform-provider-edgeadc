package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"path/filepath"

	"terraform-provider-edgeadc/internal/provider/resource_custom_monitor"
	"terraform-provider-edgeadc/swagger"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = (*customMonitorResource)(nil)

func NewCustomMonitorResource() resource.Resource {
	return &customMonitorResource{}
}

type customMonitorResource struct {
	client *API
}

func (r *customMonitorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*API)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *API, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}
	r.client = client
}

func (r *customMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_monitor"
}

func (r *customMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_custom_monitor.CustomMonitorResourceSchema(ctx)
}

func (r *customMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_custom_monitor.CustomMonitorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	_, err := CreateCustomMonitor(r.client, data.Name.ValueString(), data.FilePath.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Custom Monitor",
			err.Error(),
		)
		return
	}

	id, readErr := ReadCustomMonitor(r.client, data.Name.ValueString())
	if readErr != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Created Custom Monitor",
			readErr.Error(),
		)
		return
	}
	// Set the id which was assigned by the server
	data.Id = types.StringValue(id)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *customMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_custom_monitor.CustomMonitorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	id, err := ReadCustomMonitor(r.client, data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Custom Monitor",
			err.Error(),
		)
		return
	}
	// return if the custom monitor is not found
	if id == "" {
		return
	}

	// Set the id which was assigned by the server
	data.Id = types.StringValue(id)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *customMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_custom_monitor.CustomMonitorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	// Update is currently not supported

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *customMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_custom_monitor.CustomMonitorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	err := DeleteCustomMonitor(r.client, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Custom Monitor",
			err.Error(),
		)
		return
	}
}

func ReadCustomMonitor(client *API, name string) (id string, err error) {
	jsonResponse, err := client.GetEdgeADCObject("/GET/13")
	if err != nil {
		return "", err
	}
	configMonitoringData := swagger.ConfigMonitoring{}
	jsonErr := json.Unmarshal([]byte(jsonResponse), &configMonitoringData)
	if jsonErr != nil {
		return "", jsonErr
	}
	customMonitorId := GetCustomMonitorId(configMonitoringData, name)
	if customMonitorId == "" {
		return name, errors.New("custom monitor not found")
	}

	return customMonitorId, nil
}

func CreateCustomMonitor(client *API, name string, filePath string) (string, error) {
	// Always GET before POST to avoid "Another user has made changes" error
	_, _ = client.GetEdgeADCObject("/GET/13")

	fileName := filepath.Base(filePath)
	extraParams := map[string]string{
		"UploadMointorName": name,
		"Monitorupload":     fileName,
	}

	return client.UploadCustomMonitorFileEdgeADCApi("/POST/13?iAction=5&send=csm", extraParams, filePath)
}

func DeleteCustomMonitor(client *API, name string) error {
	csmDelete := swagger.RemoveCustomMonitor{CustomMonitor: name}
	jsonBytes, err := json.Marshal(csmDelete)
	// Always GET before POST to avoid "Another user has made changes" error
	_, _ = client.GetEdgeADCObject("/GET/13")
	_, err = client.PostEdgeADCApi("/POST/13?iAction=4", jsonBytes)
	return err
}

func GetCustomMonitorId(configMonitoringData swagger.ConfigMonitoring, name string) string {
	for _, row := range configMonitoringData.MethodCombo.Options.Option {
		if row.Value == name {
			return row.Id
		}
	}
	return ""
}
