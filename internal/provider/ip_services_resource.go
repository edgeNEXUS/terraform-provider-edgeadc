package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-edgeadc/internal/provider/resource_ip_services"
	"terraform-provider-edgeadc/swagger"
)

var _ resource.Resource = (*ipServicesResource)(nil)

func NewIpServicesResource() resource.Resource {
	return &ipServicesResource{}
}

type ipServicesResource struct {
	client *API
}

func (r *ipServicesResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ipServicesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_services"
}

func (r *ipServicesResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_ip_services.IpServiceResourceSchema(ctx)
}

func (r *ipServicesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_ip_services.IpServicesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	ipAddr := data.IpAddr.ValueString()
	port := data.Port.ValueString()

	model := ToIpService(data)

	// mutex to allow only a single resource to be managed at once
	lockName := fmt.Sprintf("ip_services")
	r.client.mutexKV.Lock(lockName)
	defer r.client.mutexKV.Unlock(lockName)

	// Retry the creation of the IP Service
	// The API has timeouts and other issues creating IP Services
	// Which we work around with retry logic
	_, retryErr := Retry(
		CreateIpServiceWithRetries, []interface{}{r.client, model}, 5, 1*time.Second, 4*time.Second,
	)
	if retryErr != nil {
		resp.Diagnostics.AddError(
			"Unable to Create IP Services",
			retryErr.Error(),
		)
		return
	}

	// After an IP Service is created, the Basic and Advanced tabs need to be updated
	createdIpService, createdErr := ReadIPService(r.client, ipAddr, port)
	if createdErr != nil {
		resp.Diagnostics.AddError(
			"Unable to Retrieve Created IP Service",
			createdErr.Error(),
		)
		return
	}
	basicTab := FromIpServiceToUpdateBasicTab(createdIpService)
	mergedBasicTab := MergeBasicTabs(basicTab, data)
	basicTabErr := UpdateIpServiceBasicTab(r.client, mergedBasicTab)
	if basicTabErr != nil {
		resp.Diagnostics.AddError(
			"Unable to Update IP Service Basic Tab",
			basicTabErr.Error(),
		)
		return
	}
	basicIpService, _ := ReadIPService(r.client, ipAddr, port)
	advanceTab := FromIpServiceToUpdateAdvanceTab(basicIpService)
	mergedAdvanceTab := MergeAdvanceTabs(advanceTab, data)
	advanceTabErr := UpdateIpServiceAdvancedTab(r.client, mergedAdvanceTab)
	if advanceTabErr != nil {
		resp.Diagnostics.AddError(
			"Unable to Update IP Service Advance Tab",
			advanceTabErr.Error(),
		)
		return
	}
	// Read the final result of all the IP Service changes
	finalIpService, _ := ReadIPService(r.client, ipAddr, port)

	// Don't save the state if the IP Service was not created correctly
	if finalIpService.SId == "" {
		resp.Diagnostics.AddError(
			"IP Service could not be created",
			fmt.Sprintf("IP Service could not be created with IP Address: %s and Port: %s", data.IpAddr.ValueString(), data.Port.ValueString()),
		)
		return
	}

	data = FromIpService(finalIpService)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ipServicesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_ip_services.IpServicesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	ipAddr := data.IpAddr.ValueString()
	port := data.Port.ValueString()

	// The API call will return all IP Service data in one response
	// Since this is the read phase and API client pointer is shared across all ip_services
	// We can retrieve the data once and reference it in the other resources
	if r.client.cachedIpServices.Data == nil {
		ipServices, ipServicesErr := ReadIPServices(r.client)
		if ipServicesErr != nil {
			resp.Diagnostics.AddError(
				"Unable to Read IP Services",
				ipServicesErr.Error(),
			)
			return
		}
		r.client.cachedIpServices = ipServices
	}

	if r.client.cachedIpServiceCombo.MonitorPolicyCombo == nil {
		ipServicesCombo, comboErr := ReadIPServicesComboBox(r.client)
		if comboErr != nil {
			resp.Diagnostics.AddError(
				"Unable to Read IP Services Combo",
				comboErr.Error(),
			)
			return
		}
		r.client.cachedIpServiceCombo = ipServicesCombo
	}

	ipService, err := CachedReadIPService(r.client.cachedIpServices, r.client.cachedIpServiceCombo, ipAddr, port)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read IP Services",
			err.Error(),
		)
		return
	}
	if ipService.SId == "" {
		return
	}

	data = FromIpService(ipService)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ipServicesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_ip_services.IpServicesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// mutex to allow only a single resource to be managed at once
	lockName := fmt.Sprintf("ip_services")
	r.client.mutexKV.Lock(lockName)
	defer r.client.mutexKV.Unlock(lockName)

	// Update API call logic
	ipAddr := data.IpAddr.ValueString()
	port := data.Port.ValueString()
	model := ToIpService(data)
	ipService, err := UpdateIpService(r.client, model)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update IP Services",
			err.Error(),
		)
		return
	}
	// After an IP Service is created, the Basic and Advanced tabs need to be updated
	basicTab := FromIpServiceToUpdateBasicTab(ipService)
	mergedBasicTab := MergeBasicTabs(basicTab, data)
	err = UpdateIpServiceBasicTab(r.client, mergedBasicTab)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update IP Service Basic Tab",
			err.Error(),
		)
		return
	}
	basicIpService, _ := ReadIPService(r.client, ipAddr, port)

	advanceTab := FromIpServiceToUpdateAdvanceTab(basicIpService)
	mergedAdvanceTab := MergeAdvanceTabs(advanceTab, data)
	err = UpdateIpServiceAdvancedTab(r.client, mergedAdvanceTab)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update IP Service Advance Tab",
			err.Error(),
		)
		return
	}

	// Read the final result of all the IP Service changes
	finalIpService, _ := ReadIPService(r.client, ipAddr, port)
	data = FromIpService(finalIpService)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ipServicesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_ip_services.IpServicesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// mutex to allow only a single resource to be managed at once
	lockName := fmt.Sprintf("ip_services")
	r.client.mutexKV.Lock(lockName)
	defer r.client.mutexKV.Unlock(lockName)

	// Delete API call logic
	err := DeleteIPService(r.client, data.IpAddr.ValueString(), data.Port.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete IP Services",
			err.Error(),
		)
		return
	}
}

func (r *ipServicesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	ipAddr, port, err := GetAddressAndPortFromId(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to import %s", req.ID),
			err.Error(),
		)
		return
	}
	resp.State.SetAttribute(ctx, path.Root("ip_addr"), ipAddr)
	resp.State.SetAttribute(ctx, path.Root("port"), port)
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func ReadIPServicesComboBox(client *API) (swagger.IpServicescombo, error) {
	model := swagger.IpServicescombo{}
	jsonResponse, err := client.GetEdgeADCObject("/GET/29")
	if err != nil {
		return model, err
	}
	// If there are no ip_services this API call returns an empty string
	// instead of a success response with an empty array
	// ToDo: Remove this once the IP call returns an empty array
	jsonErr := json.Unmarshal([]byte(jsonResponse), &model)
	if jsonErr != nil {
		return model, nil
	}
	return model, nil
}

func GetMonitorPolicyComboByValueFromComboOptions(comboOptions swagger.IpServicescomboServiceTypeComboOptions, name string) (swagger.TypeComboOne, error) {
	for _, option := range comboOptions.Option {
		if option.Value == name {
			return option, nil
		}
	}
	return swagger.TypeComboOne{}, errors.New(fmt.Sprintf("custom monitor not found with name: %s", name))
}

func GetMonitorPolicyComboByIdFromComboOptions(comboOptions swagger.IpServicescomboServiceTypeComboOptions, id string) (swagger.TypeComboOne, error) {
	for _, option := range comboOptions.Option {
		if option.Id == id {
			return option, nil
		}
	}
	return swagger.TypeComboOne{}, errors.New(fmt.Sprintf("custom monitor not found with id: %s", id))
}

func CachedReadIPService(ipServices swagger.IpServices, ipServicesCombo swagger.IpServicescombo, ipAddr string, port string) (swagger.IpService, error) {
	// ToDo: Use GetIpServiceById once the id is no longer mutable
	model, getErr := GetIpServiceByAddressAndPort(ipServices, ipAddr, port)
	if getErr != nil {
		return swagger.IpService{}, getErr
	}

	comboOptions := *ipServicesCombo.MonitorPolicyCombo.Options
	serverMonitoring, _ := ConvertComboOptionsToNames(comboOptions, model.ServerMonitoring)

	model.ServerMonitoring = serverMonitoring
	return model, nil
}

func ReadIPServices(client *API) (swagger.IpServices, error) {
	model := swagger.IpServices{}
	jsonResponse, err := client.GetEdgeADCObject("/GET/9")
	if err != nil {
		return model, err
	}
	// If there are no ip_services this API call returns an empty string
	// instead of a success response with an empty array
	// ToDo: Remove this once the IP call returns an empty array
	jsonErr := json.Unmarshal([]byte(jsonResponse), &model)
	if jsonErr != nil {
		return model, nil
	}

	return model, nil
}

func ReadIPService(client *API, ipAddr string, port string) (swagger.IpService, error) {
	ipServices, err := ReadIPServices(client)
	if err != nil {
		return swagger.IpService{}, err
	}
	// ToDo: Use GetIpServiceById once the id is no longer mutable
	model, getErr := GetIpServiceByAddressAndPort(ipServices, ipAddr, port)
	if getErr != nil {
		return swagger.IpService{}, getErr
	}

	// Convert server monitoring from lookup values to friendly names
	// Need to use lookup table for combo options
	ipServicesCombo, comboErr := ReadIPServicesComboBox(client)
	if comboErr != nil {
		return model, comboErr
	}
	comboOptions := *ipServicesCombo.MonitorPolicyCombo.Options
	serverMonitoring, _ := ConvertComboOptionsToNames(comboOptions, model.ServerMonitoring)

	model.ServerMonitoring = serverMonitoring
	return model, nil
}

func CreateIpServiceTemplate(client *API) error {
	_, err := ReadIPServices(client)
	if err != nil {
		return err
	}
	// Create a blank IP Service
	toCopy := ToCopyIpBlankTemplate()
	jsonBytes, _ := json.Marshal(toCopy)
	_, postErr := client.PostEdgeADCApi("/POST/9?iAction=3&iType=1", jsonBytes)
	if postErr != nil {
		return postErr
	}
	return nil
}

// UpdateIpServiceTemplate updates a templated ip service
// It assumes all the IDs in the object are correct
func UpdateIpServiceTemplate(client *API, toCopy swagger.CopyIp) error {
	_, _ = ReadIPServices(client)
	jsonBytes, _ := json.Marshal(toCopy)
	_, err := client.PostEdgeADCApi("/POST/9?iAction=2&iType=1", jsonBytes)
	if err != nil {
		return err
	}
	return nil
}

// CreateIpServiceWithRetries creates an IP Service and reads the result
// to ensure it was created correctly, which can be used in a retry loop
func CreateIpServiceWithRetries(client *API, model swagger.IpService) error {
	err := CreateIpService(client, model)
	if err != nil {
		return errors.New("unable to Create IP Services")
	}
	time.Sleep(100 * time.Millisecond)

	// After an IP Service is created, the Basic and Advanced tabs need to be updated
	_, createdErr := ReadIPService(client, model.IpAddr, model.Port)
	if createdErr != nil {
		return errors.New("unable to retrieve created IP Services")
	}
	return nil
}

func CreateIpService(client *API, initial swagger.IpService) error {
	// Create a blank IP Service
	err := CreateIpServiceTemplate(client)
	if err != nil {
		return err
	}

	// Get the templated IP Service
	ipServices, templateErr := ReadIPServices(client)
	if templateErr != nil {
		return templateErr
	}
	template, emptyErr := GetEmptyIpService(ipServices)
	if emptyErr != nil {
		return emptyErr
	}

	terraform := FromIpService(initial)
	model := ToCopyIp(terraform)

	// Use the interface and channel from the templated item
	model.EditedInterface = template.InterfaceID
	model.EditedChannel = template.ChannelID

	// Update the templated IP Service
	// Wait a second for it to be fully created
	return UpdateIpServiceTemplate(client, model)
}

func UpdateIpService(client *API, initial swagger.IpService) (swagger.IpService, error) {
	terraform := FromIpService(initial)
	model := ToCopyIp(terraform)
	jsonBytes, _ := json.Marshal(model)
	// Always GET before POST to avoid "Another user has made changes" error
	_, _ = ReadIPServices(client)
	// Do the update
	_, updateErr := client.PostEdgeADCApi("/POST/9?iAction=2&iType=1", jsonBytes)
	if updateErr != nil {
		return swagger.IpService{}, updateErr
	}
	// Read the final result
	return ReadIPService(client, initial.IpAddr, initial.Port)
}

func UpdateIpServiceBasicTab(client *API, model swagger.UpdateBasicTab) error {
	// Convert server monitoring from friendly names to lookup values
	ipServicesCombo, comboErr := ReadIPServicesComboBox(client)
	if comboErr != nil {
		return comboErr
	}
	comboOptions := *ipServicesCombo.MonitorPolicyCombo.Options
	serverMonitoring, _ := ConvertComboOptionsToIds(comboOptions, model.ServerMonitoring)
	model.ServerMonitoring = serverMonitoring

	// Update
	jsonBytes, _ := json.Marshal(model)
	_, err := client.PostEdgeADCApi("/POST/9?iAction=2&iType=3", jsonBytes)
	return err
}

func FromIpServiceToUpdateBasicTab(input swagger.IpService) swagger.UpdateBasicTab {
	terraform := FromIpService(input)
	return ToUpdateBasicTab(terraform)
}

func FromIpServiceToUpdateAdvanceTab(input swagger.IpService) swagger.UpdateAdvanceTab {
	terraform := FromIpService(input)
	return ToUpdateAdvanceTab(terraform)
}

func UpdateIpServiceAdvancedTab(client *API, model swagger.UpdateAdvanceTab) error {
	jsonBytes, _ := json.Marshal(model)
	_, err := client.PostEdgeADCApi("/POST/9?iAction=2&iType=4", jsonBytes)
	return err
}

func DeleteIPService(client *API, ipAddr string, port string) error {
	// Always GET before POST to avoid "Another user has made changes" error
	ipService, getErr := ReadIPService(client, ipAddr, port)
	if getErr != nil {
		return getErr
	}

	removeIp := swagger.RemoveIp{
		EditedInterface: ipService.InterfaceID,
		EditedChannel:   ipService.ChannelID,
	}
	jsonBytes, _ := json.Marshal(removeIp)
	_, err := client.PostEdgeADCApi("/POST/9?iAction=3&iType=4", jsonBytes)
	return err
}

func HandleEmptyResponse(jsonResponse string) swagger.IpServices {
	model := swagger.IpServices{}
	jsonErr := json.Unmarshal([]byte(jsonResponse), &model)
	if jsonErr != nil {
		return model
	}
	return model
}

// GetIpServiceIds gets a list of all the IDs of the ip services
func GetIpServiceIds(ipServices swagger.IpServices) []string {
	ids := make([]string, 0)
	if ipServices.Data.Dataset.IpService == nil {
		return ids
	}
	for _, svc := range ipServices.Data.Dataset.IpService {
		for _, ipService := range svc {
			ids = append(ids, ipService.SId)
		}
	}
	return ids
}

// GetIpServiceByAddressAndPort gets the ip service by the ip address and port
// ToDo: Use GetIpServiceById once the Ids are no longer mutable
func GetIpServiceByAddressAndPort(ipServices swagger.IpServices, ipAddr string, port string) (swagger.IpService, error) {
	for _, svc := range ipServices.Data.Dataset.IpService {
		for _, ipService := range svc {
			if ipService.IpAddr == ipAddr && ipService.Port == port {
				return ipService, nil
			}
		}
	}
	return swagger.IpService{}, errors.New(fmt.Sprintf("ip service not found with address: %s and port: %s", ipAddr, port))
}

// GetEmptyIpService finds an ip service created by the template
func GetEmptyIpService(ipServices swagger.IpServices) (swagger.IpService, error) {
	for _, svc := range ipServices.Data.Dataset.IpService {
		for _, ipService := range svc {
			if ipService.IpAddr == "" {
				return ipService, nil
			}
		}
	}
	return swagger.IpService{}, errors.New(fmt.Sprintf("Templated ip service not found"))
}

func GetIpServiceById(ipServices swagger.IpServices, id string) (swagger.IpService, error) {
	for _, svc := range ipServices.Data.Dataset.IpService {
		for _, ipService := range svc {
			if ipService.SId == id {
				return ipService, nil
			}
		}
	}
	return swagger.IpService{}, errors.New(fmt.Sprintf("ip service not found with id: %s", id))
}

func ToCopyIpBlankTemplate() swagger.CopyIp {
	return swagger.CopyIp{
		EditedInterface:         "",
		EditedChannel:           "",
		CopyVIP:                 "0",
		IpAddr:                  "",
		LocalPortEnabledChecked: "",
		Port:                    "",
		PrimaryChecked:          "",
		ServiceName:             "blank-terraform",
		ServiceType:             "",
		SubnetMask:              "",
	}
}

func ToCopyIp(data resource_ip_services.IpServicesResourceModel) swagger.CopyIp {
	return swagger.CopyIp{
		EditedInterface:         data.InterfaceID.ValueString(),
		EditedChannel:           data.ChannelID.ValueString(),
		CopyVIP:                 "0", //data.?.ValueString(),
		IpAddr:                  data.IpAddr.ValueString(),
		LocalPortEnabledChecked: data.LocalPortEnabledChecked.ValueString(),
		Port:                    data.Port.ValueString(),
		PrimaryChecked:          data.PrimaryChecked.ValueString(),
		ServiceName:             data.ServiceName.ValueString(),
		ServiceType:             data.ServiceType.ValueString(),
		SubnetMask:              data.SubnetMask.ValueString(),
	}
}

func ToRemoveIp(data resource_ip_services.IpServicesResourceModel) swagger.RemoveIp {
	return swagger.RemoveIp{
		EditedInterface: data.InterfaceID.ValueString(),
		EditedChannel:   data.ChannelID.ValueString(),
	}
}

func ToUpdateBasicTab(data resource_ip_services.IpServicesResourceModel) swagger.UpdateBasicTab {
	return swagger.UpdateBasicTab{
		EditedInterface:      data.InterfaceID.ValueString(),
		EditedChannel:        data.ChannelID.ValueString(),
		ServerMonitoring:     data.ServerMonitoring.ValueString(),
		Acceleration:         data.Acceleration.ValueString(),
		LoadBalancingPolicy:  data.LoadBalancingPolicy.ValueString(),
		SslCertificate:       data.SslCertificate.ValueString(),
		SslClientCertificate: data.SslClientCertificate.ValueString(),
		CachingRule:          data.CachingRule.ValueString(),
	}
}

// MergeBasicTabs merges the basic tabs with the initial terraform value
// It overlays values from terraform only when the value is known
// ToDo: Is there a better way to handle this?
func MergeBasicTabs(input swagger.UpdateBasicTab, data resource_ip_services.IpServicesResourceModel) swagger.UpdateBasicTab {
	assignIfKnown := func(target *string, source types.String) {
		if !source.IsUnknown() {
			*target = source.ValueString()
		}
	}
	assignIfKnown(&input.ServerMonitoring, data.ServerMonitoring)
	assignIfKnown(&input.Acceleration, data.Acceleration)
	assignIfKnown(&input.LoadBalancingPolicy, data.LoadBalancingPolicy)
	assignIfKnown(&input.SslCertificate, data.SslCertificate)
	assignIfKnown(&input.SslClientCertificate, data.SslClientCertificate)
	assignIfKnown(&input.CachingRule, data.CachingRule)

	return input
}

// MergeAdvanceTabs merges the advance tabs with the initial terraform value
// It overlays values from terraform only when the value is known
// ToDo: Is there a better way to handle this?
func MergeAdvanceTabs(input swagger.UpdateAdvanceTab, data resource_ip_services.IpServicesResourceModel) swagger.UpdateAdvanceTab {
	assignIfKnown := func(target *string, source types.String) {
		if !source.IsUnknown() {
			*target = source.ValueString()
		}
	}
	assignIfKnown(&input.Connectivity, data.Connectivity)
	assignIfKnown(&input.ConnectionTimeout, data.ConnectionTimeout)
	assignIfKnown(&input.MonitoringInterval, data.MonitoringInterval)
	assignIfKnown(&input.MonitoringTimeout, data.MonitoringTimeout)
	assignIfKnown(&input.MaxConn, data.MaxConn)
	assignIfKnown(&input.InCount, data.InCount)
	assignIfKnown(&input.OutCount, data.OutCount)
	assignIfKnown(&input.InCount, data.InCount)
	assignIfKnown(&input.CipherName, data.CipherName)
	assignIfKnown(&input.SecurityLog, data.SecurityLog)
	assignIfKnown(&input.SslRenegotiation, data.SslRenegotiation)
	assignIfKnown(&input.SNIDefaultCertificateName, data.SNIDefaultCertificateName)
	assignIfKnown(&input.SslResumption, data.SslResumption)
	assignIfKnown(&input.Offlinonfailure, data.Offlinonfailure)

	// ToDo: Enable once the API is updated
	input.ServerProxyProtocolVariant = "none"
	input.ClientProxyProtocolVariant = "none"
	input.RsSourceAddr = ""
	input.PersistenceTime = ""
	input.DrainBehaviour = "RetireSessions"
	input.KrbRemote = "None"

	return input
}

func ToUpdateAdvanceTab(data resource_ip_services.IpServicesResourceModel) swagger.UpdateAdvanceTab {
	return swagger.UpdateAdvanceTab{
		EditedInterface:           data.InterfaceID.ValueString(),
		EditedChannel:             data.ChannelID.ValueString(),
		Connectivity:              data.Connectivity.ValueString(),
		ConnectionTimeout:         data.ConnectionTimeout.ValueString(),
		MonitoringInterval:        data.MonitoringInterval.ValueString(),
		MonitoringTimeout:         data.MonitoringTimeout.ValueString(),
		MaxConn:                   data.MaxConn.ValueString(),
		InCount:                   data.InCount.ValueString(),
		OutCount:                  data.OutCount.ValueString(),
		CipherName:                data.CipherName.ValueString(),
		SecurityLog:               data.SecurityLog.ValueString(),
		SslRenegotiation:          data.SslRenegotiation.ValueString(),
		SNIDefaultCertificateName: data.SNIDefaultCertificateName.ValueString(),
		SslResumption:             data.SslResumption.ValueString(),
		Offlinonfailure:           data.Offlinonfailure.ValueString(),

		// ToDo: Enable once the API is updated
		ServerProxyProtocolVariant: "none",
		ClientProxyProtocolVariant: "none",
		RsSourceAddr:               "",
		PersistenceTime:            "",
		DrainBehaviour:             "RetireSessions",
		KrbRemote:                  "None",
	}
}

func ToIpService(data resource_ip_services.IpServicesResourceModel) swagger.IpService {
	swaggerModel := swagger.IpService{
		SId:                       data.SId.ValueString(),
		InterfaceID:               data.InterfaceID.ValueString(),
		ChannelID:                 data.ChannelID.ValueString(),
		InterfaceKey:              data.InterfaceKey.ValueString(),
		ChannelKey:                data.ChannelKey.ValueString(),
		IpAddr:                    data.IpAddr.ValueString(),
		MaskState:                 data.MaskState.ValueString(),
		SubnetMask:                data.SubnetMask.ValueString(),
		ServiceName:               data.ServiceName.ValueString(),
		LocalPortEnabledChecked:   data.LocalPortEnabledChecked.ValueString(),
		PrimaryChecked:            data.PrimaryChecked.ValueString(),
		ServiceType:               data.ServiceType.ValueString(),
		MaxConn:                   data.MaxConn.ValueString(),
		Port:                      data.Port.ValueString(),
		Offlinonfailure:           data.Offlinonfailure.ValueString(),
		ContentServerGroupName:    data.ContentServerGroupName.ValueString(),
		EnableConnectionPool:      data.EnableConnectionPool.ValueString(),
		ConnectionPoolSize:        data.ConnectionPoolSize.ValueString(),
		ServerMonitoring:          data.ServerMonitoring.ValueString(),
		LoadBalancingPolicy:       data.LoadBalancingPolicy.ValueString(),
		Connectivity:              data.Connectivity.ValueString(),
		Acceleration:              data.Acceleration.ValueString(),
		SslCertificate:            data.SslCertificate.ValueString(),
		SslClientCertificate:      data.SslClientCertificate.ValueString(),
		CachingRule:               data.CachingRule.ValueString(),
		MonitoringInterval:        data.MonitoringInterval.ValueString(),
		MonitoringTimeout:         data.MonitoringTimeout.ValueString(),
		ConnectionTimeout:         data.ConnectionTimeout.ValueString(),
		InCount:                   data.InCount.ValueString(),
		OutCount:                  data.OutCount.ValueString(),
		CipherName:                data.CipherName.ValueString(),
		SNIDefaultCertificateName: data.SNIDefaultCertificateName.ValueString(),
		SslRenegotiation:          data.SslRenegotiation.ValueString(),
		SslResumption:             data.SslResumption.ValueString(),
		SecurityLog:               data.SecurityLog.ValueString(),
	}
	return swaggerModel
}

func ConvertComboOptionsToNames(comboOptions swagger.IpServicescomboServiceTypeComboOptions, ids string) (string, error) {
	if ids == "" {
		return "", nil
	}
	comboIds := strings.Split(ids, ",")
	var results []string
	for _, id := range comboIds {
		combo, err := GetMonitorPolicyComboByIdFromComboOptions(comboOptions, id)
		if err != nil {
			return strings.Join(results, ","), err
		}
		results = append(results, combo.Value)
	}
	return strings.Join(results, ","), nil
}

func ConvertComboOptionsToIds(comboOptions swagger.IpServicescomboServiceTypeComboOptions, names string) (string, error) {
	if names == "" {
		return "", nil
	}
	comboNames := strings.Split(names, ",")
	var results []string
	for _, name := range comboNames {
		combo, err := GetMonitorPolicyComboByValueFromComboOptions(comboOptions, name)
		if err != nil {
			return strings.Join(results, ","), err
		}
		results = append(results, combo.Id)
	}
	return strings.Join(results, ","), nil
}

func FromIpService(swaggerModel swagger.IpService) resource_ip_services.IpServicesResourceModel {
	terraformModel := resource_ip_services.IpServicesResourceModel{
		Id:                        types.StringValue(fmt.Sprintf("%s:%s", swaggerModel.IpAddr, swaggerModel.Port)),
		SId:                       types.StringValue(swaggerModel.SId),
		InterfaceID:               types.StringValue(swaggerModel.InterfaceID),
		ChannelID:                 types.StringValue(swaggerModel.ChannelID),
		InterfaceKey:              types.StringValue(swaggerModel.InterfaceKey),
		ChannelKey:                types.StringValue(swaggerModel.ChannelKey),
		IpAddr:                    types.StringValue(swaggerModel.IpAddr),
		MaskState:                 types.StringValue(swaggerModel.MaskState),
		SubnetMask:                types.StringValue(swaggerModel.SubnetMask),
		ServiceName:               types.StringValue(swaggerModel.ServiceName),
		LocalPortEnabledChecked:   types.StringValue(swaggerModel.LocalPortEnabledChecked),
		PrimaryChecked:            types.StringValue(swaggerModel.PrimaryChecked),
		ServiceType:               types.StringValue(swaggerModel.ServiceType),
		MaxConn:                   types.StringValue(swaggerModel.MaxConn),
		Port:                      types.StringValue(swaggerModel.Port),
		Offlinonfailure:           types.StringValue(swaggerModel.Offlinonfailure),
		ContentServerGroupName:    types.StringValue(swaggerModel.ContentServerGroupName),
		EnableConnectionPool:      types.StringValue(swaggerModel.EnableConnectionPool),
		ConnectionPoolSize:        types.StringValue(swaggerModel.ConnectionPoolSize),
		ServerMonitoring:          types.StringValue(swaggerModel.ServerMonitoring),
		LoadBalancingPolicy:       types.StringValue(swaggerModel.LoadBalancingPolicy),
		Connectivity:              types.StringValue(swaggerModel.Connectivity),
		Acceleration:              types.StringValue(swaggerModel.Acceleration),
		SslCertificate:            types.StringValue(swaggerModel.SslCertificate),
		SslClientCertificate:      types.StringValue(swaggerModel.SslClientCertificate),
		CachingRule:               types.StringValue(swaggerModel.CachingRule),
		MonitoringInterval:        types.StringValue(swaggerModel.MonitoringInterval),
		MonitoringTimeout:         types.StringValue(swaggerModel.MonitoringTimeout),
		ConnectionTimeout:         types.StringValue(swaggerModel.ConnectionTimeout),
		InCount:                   types.StringValue(swaggerModel.InCount),
		OutCount:                  types.StringValue(swaggerModel.OutCount),
		CipherName:                types.StringValue(swaggerModel.CipherName),
		SNIDefaultCertificateName: types.StringValue(swaggerModel.SNIDefaultCertificateName),
		SslRenegotiation:          types.StringValue(swaggerModel.SslRenegotiation),
		SslResumption:             types.StringValue(swaggerModel.SslResumption),
		SecurityLog:               types.StringValue(swaggerModel.SecurityLog),
	}
	return terraformModel
}
