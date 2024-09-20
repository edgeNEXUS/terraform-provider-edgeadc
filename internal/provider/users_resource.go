package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-edgeadc/internal/provider/resource_users"
	"terraform-provider-edgeadc/swagger"
)

var _ resource.Resource = (*usersResource)(nil)
var _ resource.ResourceWithConfigure = (*usersResource)(nil)

func NewUsersResource() resource.Resource {
	return &usersResource{}
}

type usersResource struct {
	client *API
}

func (r *usersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *usersResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_users"
}

func (r *usersResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_users.UsersResourceSchema(ctx)
}

func (r *usersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_users.UsersModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	model := r.ToUsersAdd(data)
	usersMemberOpt, err := CreateUser(r.client, model)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create User",
			err.Error(),
		)
		return
	}

	data = r.ToTerraformModel(usersMemberOpt)
	// The API won't return the new password, so we need to set it manually
	data.NewPassword = types.StringValue(model.NewPassword)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *usersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_users.UsersModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	model := r.ToUsersMembersOpt(data)
	usersMemberOpt, err := ReadUser(r.client, model)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read User",
			err.Error(),
		)
		return
	}

	data = r.ToTerraformModel(usersMemberOpt)
	// The API won't return the new password, so we need to set it manually
	data.NewPassword = types.StringValue(model.NewPassword)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *usersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_users.UsersModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	model := r.ToUsersUpdate(data)
	usersMemberOpt, err := UpdateUser(r.client, model)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update User",
			err.Error(),
		)
		return
	}

	data = r.ToTerraformModel(usersMemberOpt)
	// The API won't return the new password, so we need to set it manually
	data.NewPassword = types.StringValue(model.NewPassword)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *usersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_users.UsersModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Return if ID is not set
	if data.UserName.ValueString() == "" {
		resp.Diagnostics.AddError(
			"Error deleting user",
			"Could not delete user: UserName is not set",
		)
		return
	}

	// Delete API call logic
	err := DeleteUser(r.client, data.UserName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete User",
			err.Error(),
		)
		return
	}
}

func (r *usersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("user_name"), req, resp)
}

func ReadUser(client *API, model swagger.UsersMembersOpt) (out swagger.UsersMembersOpt, err error) {
	jsonResponse, err := client.GetEdgeADCObject("/GET/33")
	if err != nil {
		return model, err
	}
	usersData := swagger.UsersData{}
	jsonErr := json.Unmarshal([]byte(jsonResponse), &usersData)
	if jsonErr != nil {
		return out, jsonErr
	}
	usersMemberOpt := GetUsersMembersOptByName(usersData, model.UserName)
	return usersMemberOpt, nil
}

func CreateUser(client *API, model swagger.UsersAdd) (out swagger.UsersMembersOpt, err error) {
	jsonBytes, _ := json.Marshal(model)
	// Always GET before POST to avoid "Another user has made changes" error
	_, _ = client.GetEdgeADCObject("/GET/33")
	jsonResponse, err := client.PostEdgeADCApi("/POST/33?iAction=1&iType=1", jsonBytes)
	if err != nil {
		return out, err
	}
	usersData := swagger.UsersData{}
	jsonErr := json.Unmarshal([]byte(jsonResponse), &usersData)
	if jsonErr != nil {
		return out, jsonErr
	}
	usersMemberOpt := GetUsersMembersOptByName(usersData, model.UserName)
	return usersMemberOpt, nil
}

func UpdateUser(client *API, model swagger.UsersUpdate) (out swagger.UsersMembersOpt, err error) {
	jsonBytes, _ := json.Marshal(model)
	// Always GET before POST to avoid "Another user has made changes" error
	_, _ = client.GetEdgeADCObject("/GET/33")
	jsonResponse, err := client.PostEdgeADCApi("/POST/33?iAction=1&iType=2", jsonBytes)
	if err != nil {
		return out, err
	}
	usersData := swagger.UsersData{}
	jsonErr := json.Unmarshal([]byte(jsonResponse), &usersData)
	if jsonErr != nil {
		return out, jsonErr
	}
	usersMemberOpt := GetUsersMembersOptByName(usersData, model.UserName)
	return usersMemberOpt, nil
}

func DeleteUser(client *API, userName string) error {
	usersDelete := swagger.UsersDelete{UserName: userName}
	jsonBytes, err := json.Marshal(usersDelete)
	// Always GET before POST to avoid "Another user has made changes" error
	_, _ = client.GetEdgeADCObject("/GET/33")
	_, err = client.PostEdgeADCApi("/POST/33?iAction=1&iType=3", jsonBytes)
	return err
}

func (r *usersResource) ToUsersMembersOpt(data resource_users.UsersModel) swagger.UsersMembersOpt {
	usersData := swagger.UsersMembersOpt{
		UserName:    data.UserName.ValueString(),
		NewPassword: data.NewPassword.ValueString(),
		OldPassword: "",
		IsAddOn:     data.IsAddOn.ValueString(),
		IsAdmin:     data.IsAdmin.ValueString(),
		IsAPI:       data.IsApi.ValueString(),
		IsGUIR:      data.IsGuir.ValueString(),
		IsGUIW:      data.IsGuiw.ValueString(),
		IsSSH:       data.IsSsh.ValueString(),
	}
	return usersData
}

func (r *usersResource) ToUsersAdd(data resource_users.UsersModel) swagger.UsersAdd {
	return swagger.UsersAdd{
		UserName:    data.UserName.ValueString(),
		NewPassword: data.NewPassword.ValueString(),
		OldPassword: "",
		IsAddOn:     data.IsAddOn.ValueString(),
		IsAdmin:     data.IsAdmin.ValueString(),
		IsAPI:       data.IsApi.ValueString(),
		IsGUIR:      data.IsGuir.ValueString(),
		IsGUIW:      data.IsGuiw.ValueString(),
		IsSSH:       data.IsSsh.ValueString(),
	}
}

func (r *usersResource) ToUsersUpdate(data resource_users.UsersModel) swagger.UsersUpdate {
	return swagger.UsersUpdate{
		UserName:    data.UserName.ValueString(),
		NewPassword: data.NewPassword.ValueString(),
		OldPassword: "",
		IsAddOn:     data.IsAddOn.ValueString(),
		IsAdmin:     data.IsAdmin.ValueString(),
		IsAPI:       data.IsApi.ValueString(),
		IsGUIR:      data.IsGuir.ValueString(),
		IsGUIW:      data.IsGuiw.ValueString(),
		IsSSH:       data.IsSsh.ValueString(),
	}
}

func (r *usersResource) ToTerraformModel(users swagger.UsersMembersOpt) resource_users.UsersModel {
	usersModel := resource_users.UsersModel{
		IsAddOn:     types.StringValue(users.IsAddOn),
		IsAdmin:     types.StringValue(users.IsAdmin),
		IsApi:       types.StringValue(users.IsAPI),
		IsGuir:      types.StringValue(users.IsGUIR),
		IsGuiw:      types.StringValue(users.IsGUIW),
		IsSsh:       types.StringValue(users.IsSSH),
		NewPassword: types.StringValue(users.NewPassword),
		UserName:    types.StringValue(users.UserName),
	}
	return usersModel
}

func GetUsersMembersOptByName(usersData swagger.UsersData, username string) swagger.UsersMembersOpt {
	for _, user := range usersData.Members.Dataset {
		if user.UserName == username {
			return user
		}
	}
	return swagger.UsersMembersOpt{}
}
