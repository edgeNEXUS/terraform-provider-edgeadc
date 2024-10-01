package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-edgeadc/internal/provider/resource_server"
	"terraform-provider-edgeadc/swagger"
)

var _ resource.Resource = (*serverResource)(nil)

func NewServerResource() resource.Resource {
	return &serverResource{}
}

type serverResource struct {
	client *API
}

func (r *serverResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *serverResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

func (r *serverResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_server.ServerResourceSchema(ctx)
}

func (r *serverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_server.ServerResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ipServiceId := data.IpService.ValueString()

	if ipServiceId == "" {
		resp.Diagnostics.AddError(
			"Unable to Create Server",
			"ip_service_id must not be empty",
		)
		return
	}

	// mutex to allow only a single resource to be managed at once
	lockName := fmt.Sprintf("server")
	r.client.mutexKV.Lock(lockName)
	defer r.client.mutexKV.Unlock(lockName)

	// Create API call logic
	ipServiceIpAddr, ipServicePort, _ := GetAddressAndPortFromId(ipServiceId)
	createErr := CreateServerTemplate(r.client, ipServiceIpAddr, ipServicePort)
	if createErr != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Server Template",
			createErr.Error(),
		)
		return
	}
	server := r.ToUpdateServer(data)
	updateErr := UpdateServerTemplate(r.client, server, ipServiceIpAddr, ipServicePort)
	if updateErr != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Server Template",
			updateErr.Error(),
		)
		return
	}

	ipService, cServerId, finalErr := ReadServer(r.client, server.CSIPAddr, server.CSPort, ipServiceIpAddr, ipServicePort)
	if finalErr != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Created Server",
			finalErr.Error(),
		)
		return
	}
	data = r.FromCServerId(cServerId, ipService.InterfaceID, ipService.ChannelID)
	data.IpService = types.StringValue(ipServiceId)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *serverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_server.ServerResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ipServiceId := data.IpService.ValueString()

	// Read API call logic
	model := r.ToCServerId(data)
	ipServiceIpAddr, ipServicePort, _ := GetAddressAndPortFromId(data.IpService.ValueString())
	ipService, cServer, err := ReadServer(r.client, model.CSIPAddr, model.CSPort, ipServiceIpAddr, ipServicePort)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Server",
			err.Error(),
		)
		return
	}

	data = r.FromCServerId(cServer, ipService.InterfaceID, ipService.ChannelID)
	data.IpService = types.StringValue(ipServiceId)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *serverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_server.ServerResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ipServiceId := data.IpService.ValueString()

	// mutex to allow only a single resource to be managed at once
	lockName := fmt.Sprintf("server")
	r.client.mutexKV.Lock(lockName)
	defer r.client.mutexKV.Unlock(lockName)

	// Update API call logic
	model := r.ToCServerId(data)
	ipServiceIpAddr, ipServicePort, _ := GetAddressAndPortFromId(data.IpService.ValueString())
	err := UpdateServer(r.client, model, ipServiceIpAddr, ipServicePort)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Server",
			err.Error(),
		)
		return
	}

	ipService, cServerId, finalErr := ReadServer(r.client, model.CSIPAddr, model.CSPort, ipServiceIpAddr, ipServicePort)
	if finalErr != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Created Server",
			finalErr.Error(),
		)
		return
	}
	data = r.FromCServerId(cServerId, ipService.InterfaceID, ipService.ChannelID)
	data.IpService = types.StringValue(ipServiceId)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *serverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_server.ServerResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// mutex to allow only a single resource to be managed at once
	lockName := fmt.Sprintf("server")
	r.client.mutexKV.Lock(lockName)
	defer r.client.mutexKV.Unlock(lockName)

	// Delete API call logic
	ipServiceIpAddr, ipServicePort, _ := GetAddressAndPortFromId(data.IpService.ValueString())
	model := r.ToCServerId(data)
	err := DeleteServer(r.client, model, ipServiceIpAddr, ipServicePort)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Server",
			err.Error(),
		)
		return
	}
}

func (r *serverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	ipAddr, port, err := GetAddressAndPortFromId(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to import %s", req.ID),
			err.Error(),
		)
		return
	}
	resp.State.SetAttribute(ctx, path.Root("cs_ip_addr"), ipAddr)
	resp.State.SetAttribute(ctx, path.Root("cs_port"), port)
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func ReadServer(client *API, serverIPAddr string, serverPort string, ipServiceIpAddr string, ipServicePort string) (swagger.IpService, swagger.CServerId, error) {
	ipServices, ipServiceErr := ReadIPServices(client)
	if ipServiceErr != nil {
		return swagger.IpService{}, swagger.CServerId{}, ipServiceErr
	}
	return GetServerByAddressAndPortFromIpServices(ipServices, ipServiceIpAddr, ipServicePort, serverIPAddr, serverPort)
}

func CreateServerTemplate(client *API, ipServerIpAddr string, ipServerPort string) error {
	existingIpServices, err := ReadIPServices(client)
	if err != nil {
		return err
	}
	ipService, ipServiceErr := GetIpServiceByAddressAndPort(existingIpServices, ipServerIpAddr, ipServerPort)
	if ipServiceErr != nil {
		return ipServiceErr
	}

	// Create a blank Server
	// Note this is subject to race condition/concurrent operation scenario
	// As the ID could change between when we read the data and when we create the server
	r := serverResource{}
	toCopy := r.ToAddServerBlankTemplate(ipService.InterfaceID, ipService.ChannelID)
	jsonBytes, _ := json.Marshal(toCopy)
	_, postErr := client.PostEdgeADCApi("/POST/9?iAction=3&iType=3", jsonBytes)
	if postErr != nil {
		return postErr
	}
	return nil
}

func UpdateServerTemplate(client *API, updateServer swagger.UpdateServer, ipServiceIpAddr string, ipServicePort string) error {
	// Refresh server details
	ipServices, readErr := ReadIPServices(client)
	if readErr != nil {
		return readErr
	}
	ipService, ipServiceErr := GetIpServiceByAddressAndPort(ipServices, ipServiceIpAddr, ipServicePort)
	if ipServiceErr != nil {
		return ipServiceErr
	}
	emptyServer, emptyErr := GetEmptyServer(client, ipService)
	if emptyErr != nil {
		return emptyErr
	}

	// Update blank Server
	// Note this is subject to race condition/concurrent operation scenario
	// As the ID could change between when we read the data and when we create the server
	updateServer.EditedInterface = ipService.InterfaceID
	updateServer.EditedChannel = ipService.ChannelID
	updateServer.CId = emptyServer.CId

	jsonBytes, _ := json.Marshal(updateServer)
	_, postErr := client.PostEdgeADCApi("/POST/9?iAction=2&iType=2", jsonBytes)
	if postErr != nil {
		return postErr
	}
	return nil
}

func UpdateServer(client *API, model swagger.CServerId, ipServiceIpAddr string, ipServicePort string) error {
	// Refresh server details
	ipService, cServer, readErr := ReadServer(client, model.CSIPAddr, model.CSPort, ipServiceIpAddr, ipServicePort)
	if readErr != nil {
		return readErr
	}
	// Update details
	// Note this is subject to race condition/concurrent operation scenario
	// As the ID could change between when we read the data and when we create the server
	updateServer := swagger.UpdateServer{}
	updateServer.EditedInterface = ipService.InterfaceID
	updateServer.EditedChannel = ipService.ChannelID
	updateServer.CId = cServer.CId

	// Update the server
	jsonBytes, _ := json.Marshal(updateServer)
	_, err := client.PostEdgeADCApi("/POST/9?iAction=2&iType=2", jsonBytes)
	if err != nil {
		return err
	}
	return nil
}

func DeleteServer(client *API, model swagger.CServerId, ipServiceIpAddr string, ipServicePort string) error {
	// Refresh server details
	ipService, cServer, readErr := ReadServer(client, model.CSIPAddr, model.CSPort, ipServiceIpAddr, ipServicePort)
	if readErr != nil {
		return readErr
	}
	// Update details
	// Note this is subject to race condition/concurrent operation scenario
	// As the ID could change between when we read the data and when we create the server
	toDelete := swagger.RemoveServer{}
	toDelete.EditedInterface = ipService.InterfaceID
	toDelete.EditedChannel = ipService.ChannelID
	toDelete.CId = cServer.CId
	jsonBytes, _ := json.Marshal(toDelete)
	_, err := client.PostEdgeADCApi("/POST/9?iAction=3&iType=5", jsonBytes)
	return err
}

// GetEmptyServer finds a server created by the template
// It will only delete one server at a time
// (When you create a new ip_service it creates a blank server with the same port)
func GetEmptyServer(client *API, ipService swagger.IpService) (swagger.CServerId, error) {
	for _, server := range ipService.ContentServer.CServerId {
		if server.CSIPAddr == "" {
			return server, nil
		}
	}
	return swagger.CServerId{}, nil
}

// DeleteEmptyServer deletes a server which has an empty IP address
// It will only delete one server at a time
// (When you create a new ip_service it creates a blank server with the same port)
func DeleteEmptyServer(client *API, ipServices swagger.IpServices, port string) error {
	for _, svc := range ipServices.Data.Dataset.IpService {
		for _, ipService := range svc {
			for _, server := range ipService.ContentServer.CServerId {
				if server.CSIPAddr == "" && server.CSPort == port {
					deleteErr := DeleteServer(client, server, ipService.IpAddr, ipService.Port)
					if deleteErr != nil {
						return deleteErr
					}
				}
			}
		}
	}
	return nil
}

// GetServerIds gets a list of all the IDs of the servers
func GetServerIds(ipServices swagger.IpServices, ipServiceIpAddr string, ipServicePort string) []string {
	ids := make([]string, 0)
	if ipServices.Data.Dataset.IpService == nil {
		return ids
	}
	for _, svc := range ipServices.Data.Dataset.IpService {
		for _, ipService := range svc {
			if ipService.IpAddr == ipServiceIpAddr && ipService.Port == ipServicePort {
				for _, server := range ipService.ContentServer.CServerId {
					ids = append(ids, server.CId)
				}
			}
		}
	}
	return ids
}

func GetServerByAddressAndPortFromIpServices(ipServices swagger.IpServices, ipServiceIpAddr string, ipServicePort string, serverIpAddr string, serverPort string) (swagger.IpService, swagger.CServerId, error) {
	if ipServices.Data.Dataset.IpService == nil {
		return swagger.IpService{}, swagger.CServerId{}, errors.New(fmt.Sprintf("server not found with %s:%s for ip_service %s:%s", serverIpAddr, serverPort, ipServiceIpAddr, ipServicePort))
	}
	for _, svc := range ipServices.Data.Dataset.IpService {
		for _, ipService := range svc {
			if ipService.IpAddr == ipServiceIpAddr && ipService.Port == ipServicePort {
				for _, server := range ipService.ContentServer.CServerId {
					if server.CSIPAddr == serverIpAddr && server.CSPort == serverPort {
						return ipService, server, nil
					}
				}
			}
		}
	}
	return swagger.IpService{}, swagger.CServerId{}, errors.New(fmt.Sprintf("server not found with %s:%s for ip_service %s:%s", serverIpAddr, serverPort, ipServiceIpAddr, ipServicePort))
}

func GetServerByAddressAndPortFromIpService(ipService swagger.IpService, ipServiceIpAddr string, ipServicePort string, serverIpAddr string, serverPort string) (swagger.CServerId, error) {
	if ipService.IpAddr == ipServiceIpAddr && ipService.Port == ipServicePort {
		for _, server := range ipService.ContentServer.CServerId {
			if server.CSIPAddr == serverIpAddr && server.CSPort == serverPort {
				return server, nil
			}
		}
	}
	return swagger.CServerId{}, errors.New(fmt.Sprintf("server not found with %s:%s for ip_service %s:%s", serverIpAddr, serverPort, ipServiceIpAddr, ipServicePort))
}

func (r *serverResource) ToAddServerBlankTemplate(interfaceId string, channelId string) swagger.AddServer {
	return swagger.AddServer{
		EditedInterface: interfaceId,
		EditedChannel:   channelId,
		CSActivity:      "",
		CSIPAddr:        "",
		CSPort:          "",
		CSNotes:         "blank-terraform",
		//		ServerId:               data.ServerId.ValueString(),
		WeightFactor: "",
	}
}

func (r *serverResource) ToAddServer(data resource_server.ServerResourceModel) swagger.AddServer {
	return swagger.AddServer{
		EditedInterface: data.InterfaceId.ValueString(),
		EditedChannel:   data.ChannelId.ValueString(),
		CSActivity:      data.CSActivity.ValueString(),
		CSIPAddr:        data.CSIPAddr.ValueString(),
		CSPort:          data.CSPort.ValueString(),
		CSNotes:         data.CSNotes.ValueString(),
		//		ServerId:               data.ServerId.ValueString(),
		WeightFactor: data.WeightFactor.ValueString(),
		//WeightFactorCalculated: data.WeightFactorCalculated.ValueString(),
	}
}

func (r *serverResource) ToUpdateServer(data resource_server.ServerResourceModel) swagger.UpdateServer {
	return swagger.UpdateServer{
		EditedInterface: data.InterfaceId.ValueString(),
		EditedChannel:   data.ChannelId.ValueString(),
		CSActivity:      data.CSActivity.ValueString(),
		CSIPAddr:        data.CSIPAddr.ValueString(),
		CSPort:          data.CSPort.ValueString(),
		CSNotes:         data.CSNotes.ValueString(),
		//		ServerId:               data.ServerId.ValueString(),
		WeightFactor: data.WeightFactor.ValueString(),
		//WeightFactorCalculated: data.WeightFactorCalculated.ValueString(),
	}
}

func (r *serverResource) ToRemoveServer(data resource_server.ServerResourceModel) swagger.RemoveServer {
	return swagger.RemoveServer{
		EditedInterface: data.InterfaceId.ValueString(),
		EditedChannel:   data.ChannelId.ValueString(),
		CId:             data.CId.ValueString(),
	}
}

func (r *serverResource) ToCServerId(data resource_server.ServerResourceModel) swagger.CServerId {
	return swagger.CServerId{
		ServerKey:  data.ServerKey.ValueString(),
		CId:        data.CId.ValueString(),
		CSActivity: data.CSActivity.ValueString(),
		CSIPAddr:   data.CSIPAddr.ValueString(),
		CSPort:     data.CSPort.ValueString(),
		CSNotes:    data.CSNotes.ValueString(),
		//		ServerId:               data.ServerId.ValueString(),
		WeightFactor: data.WeightFactor.ValueString(),
	}
}

func (r *serverResource) FromCServerId(swaggerModel swagger.CServerId, interfaceId string, channelId string) resource_server.ServerResourceModel {
	return resource_server.ServerResourceModel{
		Id:          types.StringValue(fmt.Sprintf("%s:%s", swaggerModel.CSIPAddr, swaggerModel.CSPort)),
		InterfaceId: types.StringValue(interfaceId),
		ChannelId:   types.StringValue(channelId),
		ServerKey:   types.StringValue(swaggerModel.ServerKey),
		CId:         types.StringValue(swaggerModel.CId),
		CSActivity:  types.StringValue(swaggerModel.CSActivity),
		CSIPAddr:    types.StringValue(swaggerModel.CSIPAddr),
		CSPort:      types.StringValue(swaggerModel.CSPort),
		CSNotes:     types.StringValue(swaggerModel.CSNotes),
		// ServerId:               types.StringValue(swaggerModel.ServerId),
		WeightFactor: types.StringValue(swaggerModel.WeightFactor),
	}
}
