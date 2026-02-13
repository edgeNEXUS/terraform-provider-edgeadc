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

	// Save user's original type value before normalization so we can
	// preserve it in state (the API stores backend names but the user
	// may have used a friendly name like "HTTP Head").
	userType := data.Type

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
	// Restore the user's original type value so state matches the config,
	// preventing perpetual diffs when friendly names are used.
	data.Type = userType

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

	// Preserve the type value from state. The user may have used a
	// friendly name (e.g. "HTTP Head") which was stored in state. The
	// API returns the backend name ("CheckHead"), so we restore the
	// user's original value to prevent perpetual diffs.
	stateType := data.Type

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

	// Restore the user's original type from state. If the backend name
	// maps to the same normalised value as the state type, keep the
	// state value. This handles both "HTTP Head" ↔ "CheckHead" and
	// direct backend name usage.
	if NormalizeMonitorType(stateType.ValueString()) == NormalizeMonitorType(data.Type.ValueString()) {
		data.Type = stateType
	}

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

	// Save user's original type value before normalization
	userType := data.Type

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
	// Restore the user's original type value
	data.Type = userType

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

// monitorTypeFriendlyToBackend maps UI-friendly monitor type names to the
// backend names that the ADC API requires. Users may specify either form
// in their Terraform configuration; the provider normalises to the backend
// name before sending to the API.
var monitorTypeFriendlyToBackend = map[string]string{
	"HTTP 200 OK":   "Check200",
	"HTTP Head":     "CheckHead",
	"HTTP Response": "CheckResponse",
	"TCP Connect":   "Connect",
	"ICMP Ping":     "Ping",
	"None":          "None",
	"DICOM":         "DICOM",
	"HTTP Post":     "CheckPOST",
}

// monitorTypeBackendToFriendly is the reverse of monitorTypeFriendlyToBackend.
var monitorTypeBackendToFriendly map[string]string

func init() {
	monitorTypeBackendToFriendly = make(map[string]string, len(monitorTypeFriendlyToBackend))
	for friendly, backend := range monitorTypeFriendlyToBackend {
		monitorTypeBackendToFriendly[backend] = friendly
	}
}

// NormalizeMonitorType converts a friendly UI monitor type name to its
// backend equivalent. If the value is already a backend name (or unknown)
// it is returned unchanged.
func NormalizeMonitorType(t string) string {
	if backend, ok := monitorTypeFriendlyToBackend[t]; ok {
		return backend
	}
	return t // already a backend name or custom monitor
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
	_, err = client.PostEdgeADCApi("/POST/13?iAction=2&iType=1", jsonBytes)
	if err != nil {
		return swagger.RealConfigMonitoringOpt{}, err
	}
	time.Sleep(1 * time.Second)
	// Do a fresh GET read to get the updated monitor data
	// The POST response may not include the full grid data
	return ReadRealserverMonitor(client, model.Name)
}

func DeleteRealserverMonitor(client *API, name string) (err error) {
	configMonitoring, err := ReadRealserverMonitor(client, name)
	if err != nil {
		// If we can't read the monitor, it may already be gone.
		// Treat as already deleted to allow destroy to succeed after
		// a previously failed apply.
		return nil
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
		// Normalize the type so that users can supply either the friendly UI
		// name (e.g. "HTTP Head") or the backend name (e.g. "CheckHead").
		Type_:     NormalizeMonitorType(data.Type.ValueString()),
		Ssl:       data.Ssl.ValueString(),
		Url:       data.Url.ValueString(),
		Content:   data.Content.ValueString(),
		Username:  data.Username.ValueString(),
		Password:  data.Password.ValueString(),
		Threshold: data.Threshold.ValueString(),
	}
}

func ToUpdateRealRequestInner(realConfigMonitoringOpt swagger.RealConfigMonitoringOpt) swagger.UpdateRealRequestInner {
	return swagger.UpdateRealRequestInner{
		Id:          realConfigMonitoringOpt.Id,
		Name:        realConfigMonitoringOpt.Name,
		Description: realConfigMonitoringOpt.Description,
		Type_:       realConfigMonitoringOpt.Type_,
		Ssl:         realConfigMonitoringOpt.Ssl,
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
		Ssl:         types.StringValue(realConfigMonitoringOpt.Ssl),
		Url:         types.StringValue(realConfigMonitoringOpt.Url),
		Content:     types.StringValue(realConfigMonitoringOpt.Content),
		Username:    types.StringValue(realConfigMonitoringOpt.Username),
		Password:    types.StringValue(realConfigMonitoringOpt.Password),
		Threshold:   types.StringValue(realConfigMonitoringOpt.Threshold),
	}
}

func GetRealConfigMonitoringOptByName(configMonitoringData swagger.ConfigMonitoring, name string) swagger.RealConfigMonitoringOpt {
	if configMonitoringData.ConfigMonitoringGrid == nil {
		return swagger.RealConfigMonitoringOpt{}
	}
	if configMonitoringData.ConfigMonitoringGrid.Dataset == nil {
		return swagger.RealConfigMonitoringOpt{}
	}
	for _, row := range configMonitoringData.ConfigMonitoringGrid.Dataset.Row {
		if row.Name == name {
			return row
		}
	}
	return swagger.RealConfigMonitoringOpt{}
}
