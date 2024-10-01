package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"terraform-provider-edgeadc/internal/provider/resource_realserver_monitor"
	"terraform-provider-edgeadc/swagger"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*realserverMonitorResource)(nil)

func NewRealserverMonitorResource() resource.Resource {
	return &realserverMonitorResource{}
}

type realserverMonitorResource struct {
	client *API
}

func (r *realserverMonitorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	//r.mutex = NewMutexKV()
}

func (r *realserverMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_realserver_monitor"
}

func (r *realserverMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_realserver_monitor.RealserverMonitorResourceSchema(ctx)
}

func (r *realserverMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_realserver_monitor.RealserverMonitorModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// mutex to allow only a single resource to be managed at once
	lockName := fmt.Sprintf("realserver_monitor")
	r.client.mutexKV.Lock(lockName)
	defer r.client.mutexKV.Unlock(lockName)

	// Create API call logic
	model := r.ToRealserverMonitorOpt(data)
	realserverMonitorOpt, err := CreateRealserverMonitor(r.client, model)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Realserver Monitor",
			err.Error(),
		)
		return
	}

	data = r.ToTerraformModel(realserverMonitorOpt)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *realserverMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_realserver_monitor.RealserverMonitorModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	model := r.ToRealserverMonitorOpt(data)
	realserverMonitorOpt, err := ReadRealserverMonitor(r.client, model.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Realserver Monitor",
			err.Error(),
		)
		return
	}

	data = r.ToTerraformModel(realserverMonitorOpt)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *realserverMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_realserver_monitor.RealserverMonitorModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// mutex to allow only a single resource to be managed at once
	lockName := fmt.Sprintf("realserver_monitor")
	r.client.mutexKV.Lock(lockName)
	defer r.client.mutexKV.Unlock(lockName)

	// Update API call logic
	model := r.ToRealserverMonitorOpt(data)
	realserverMonitorOpt, err := UpdateRealserverMonitor(r.client, model, model.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Realserver Monitor",
			err.Error(),
		)
		return
	}
	data = r.ToTerraformModel(realserverMonitorOpt)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *realserverMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_realserver_monitor.RealserverMonitorModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// mutex to allow only a single resource to be managed at once
	lockName := fmt.Sprintf("realserver_monitor")
	r.client.mutexKV.Lock(lockName)
	defer r.client.mutexKV.Unlock(lockName)

	// Delete API call logic
	model := r.ToRealserverMonitorOpt(data)
	err := DeleteRealserverMonitor(r.client, model.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Realserver Monitor",
			err.Error(),
		)
		return
	}
}

func (r *realserverMonitorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}

func ReadRealserverMonitor(client *API, name string) (swagger.RealConfigMonitoringOpt, error) {
	jsonResponse, err := client.GetEdgeADCObject("/GET/13")
	if err != nil {
		return swagger.RealConfigMonitoringOpt{}, err
	}
	configMonitoringData := swagger.ConfigMonitoring{}
	jsonErr := json.Unmarshal([]byte(jsonResponse), &configMonitoringData)
	if jsonErr != nil {
		return swagger.RealConfigMonitoringOpt{}, jsonErr
	}
	realConfigMonitoringOpt := GetRealConfigMonitoringOptByName(configMonitoringData, name)
	return realConfigMonitoringOpt, nil
}

func CreateRealserverMonitor(client *API, model swagger.RealConfigMonitoringOpt) (swagger.RealConfigMonitoringOpt, error) {
	// Always GET before POST to avoid "Another user has made changes" error
	_, _ = client.GetEdgeADCObject("/GET/13")

	// Creating a ReadServerMonitor is a two-step process
	// First you post to the API and it creates a ReadServerMonitor named "New Method"
	// Then you call update on the created template to complete the process
	// Problematic is you cannot have two ReadServerMonitors with the same name
	jsonBytes, _ := json.Marshal(model)
	_, createErr := client.PostEdgeADCApi("/POST/13?iAction=1&iType=1", jsonBytes)
	if createErr != nil {
		return swagger.RealConfigMonitoringOpt{}, createErr
	}

	return UpdateRealserverMonitor(client, model, "New Method")
}

// UpdateRealserverMonitor updates a RealserverMonitor
// It takes a "name" paramenter to be used by the "Create" function to update a monitor with an alternate name
func UpdateRealserverMonitor(client *API, model swagger.RealConfigMonitoringOpt, name string) (swagger.RealConfigMonitoringOpt, error) {
	// Get the ID of the ReadServerMonitor - required in order to update
	configMonitoring, err := ReadRealserverMonitor(client, name)
	if err != nil {
		return swagger.RealConfigMonitoringOpt{}, err
	}

	updateRealRequestInner := ToUpdateRealRequestInner(model)
	updateRealRequestInner.Id = configMonitoring.Id
	// Not reflected in docs but this is an array
	// ToDo: Update swagger
	updateRealServerMonitor := []swagger.UpdateRealRequestInner{updateRealRequestInner}
	jsonBytes, _ := json.Marshal(updateRealServerMonitor)
	jsonResponse, err := client.PostEdgeADCApi("/POST/13?iAction=2&iType=1", jsonBytes)
	if err != nil {
		return swagger.RealConfigMonitoringOpt{}, err
	}
	configMonitoringData := swagger.ConfigMonitoring{}
	jsonErr := json.Unmarshal([]byte(jsonResponse), &configMonitoringData)
	if jsonErr != nil {
		return swagger.RealConfigMonitoringOpt{}, jsonErr
	}
	time.Sleep(1 * time.Second)
	realConfigMonitoringOpt := GetRealConfigMonitoringOptByName(configMonitoringData, model.Name)
	return realConfigMonitoringOpt, nil
}

func DeleteRealserverMonitor(client *API, name string) (err error) {
	configMonitoring, err := ReadRealserverMonitor(client, name)
	if err != nil {
		return err
	}

	// If the ReadServerMonitor does not exist, return nil
	if configMonitoring.Id == "" {
		return nil
	}

	removeRealRequest := swagger.RemoveRealRequest{Id: configMonitoring.Id}
	jsonBytes, _ := json.Marshal(removeRealRequest)
	_, deleteErr := client.PostEdgeADCApi("/POST/13?iAction=3&iType=1", jsonBytes)
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}

func (r *realserverMonitorResource) ToRealserverMonitorOpt(data resource_realserver_monitor.RealserverMonitorModel) swagger.RealConfigMonitoringOpt {
	return swagger.RealConfigMonitoringOpt{
		Id:          data.Id.ValueString(),
		Name:        data.Name.ValueString(),
		Description: data.Description.ValueString(),
		Type_:       data.Type.ValueString(),
		Url:         data.Url.ValueString(),
		Content:     data.Content.ValueString(),
		Username:    data.Username.ValueString(),
		Password:    data.Password.ValueString(),
		Threshold:   data.Threshold.ValueString(),
	}
}

func ToUpdateRealRequestInner(realConfigMonitoringOpt swagger.RealConfigMonitoringOpt) swagger.UpdateRealRequestInner {
	return swagger.UpdateRealRequestInner{
		Id:          realConfigMonitoringOpt.Id,
		Name:        realConfigMonitoringOpt.Name,
		Description: realConfigMonitoringOpt.Description,
		Type_:       realConfigMonitoringOpt.Type_,
		Url:         realConfigMonitoringOpt.Url,
		Content:     realConfigMonitoringOpt.Content,
		Username:    realConfigMonitoringOpt.Username,
		Password:    realConfigMonitoringOpt.Password,
		Threshold:   realConfigMonitoringOpt.Threshold,
	}
}

func (r *realserverMonitorResource) ToTerraformModel(realConfigMonitoringOpt swagger.RealConfigMonitoringOpt) resource_realserver_monitor.RealserverMonitorModel {
	return resource_realserver_monitor.RealserverMonitorModel{
		Id:          types.StringValue(realConfigMonitoringOpt.Id),
		Name:        types.StringValue(realConfigMonitoringOpt.Name),
		Description: types.StringValue(realConfigMonitoringOpt.Description),
		Type:        types.StringValue(realConfigMonitoringOpt.Type_),
		Url:         types.StringValue(realConfigMonitoringOpt.Url),
		Content:     types.StringValue(realConfigMonitoringOpt.Content),
		Username:    types.StringValue(realConfigMonitoringOpt.Username),
		Password:    types.StringValue(realConfigMonitoringOpt.Password),
		Threshold:   types.StringValue(realConfigMonitoringOpt.Threshold),
	}
}

func GetRealConfigMonitoringOptByName(configMonitoringData swagger.ConfigMonitoring, name string) swagger.RealConfigMonitoringOpt {
	for _, row := range configMonitoringData.ConfigMonitoringGrid.Dataset.Row {
		if row.Name == name {
			return row
		}
	}
	return swagger.RealConfigMonitoringOpt{}
}
