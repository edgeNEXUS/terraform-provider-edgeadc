package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"terraform-provider-edgeadc/internal/provider/resource_ssl_certificates"
	"terraform-provider-edgeadc/swagger"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = (*sslCertificatesResource)(nil)

func NewSslCertificatesResource() resource.Resource {
	return &sslCertificatesResource{}
}

type sslCertificatesResource struct {
	client *API
}

func (r *sslCertificatesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *sslCertificatesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssl_certificates"
}

func (r *sslCertificatesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_ssl_certificates.CertificatesResourceSchema(ctx)
}

func (r *sslCertificatesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_ssl_certificates.CertificatesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	_, err := CreateCertificate(r.client, data.Id.ValueString(), data.Password.ValueString(), data.FilePath.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Certificate",
			err.Error(),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *sslCertificatesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_ssl_certificates.CertificatesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	_, err := ReadCertificate(r.client, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Certificate",
			err.Error(),
		)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *sslCertificatesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_ssl_certificates.CertificatesResourceModel

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

func (r *sslCertificatesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_ssl_certificates.CertificatesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	err := DeleteCertificate(r.client, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Certificate",
			err.Error(),
		)
		return
	}
}

func ReadCertificate(client *API, name string) (out string, err error) {
	jsonResponse, err := client.GetEdgeADCObject("/GET/19")
	if err != nil {
		return "", err
	}
	certificateData := swagger.GetCertificate{}
	jsonErr := json.Unmarshal([]byte(jsonResponse), &certificateData)
	if jsonErr != nil {
		return out, jsonErr
	}
	cert := CertificateExists(certificateData, name)
	if !cert {
		return name, errors.New("certificate not found")
	}

	return name, nil
}

func CreateCertificate(client *API, name string, password string, filePath string) (string, error) {
	// Always GET before POST to avoid "Another user has made changes" error
	_, _ = client.GetEdgeADCObject("/GET/19")

	fileName := filepath.Base(filePath)
	extraParams := map[string]string{
		"SslCertificatesImportCertificateNameText": name,
		"SslCertificatesImportFileName":            fmt.Sprintf("C:\\fakepath\\%s", fileName),
		"SslCertificatesImportKeyFileName":         "Go",
		"Go":                                       fileName,
		"SslCertificatesImportPasswordText":        password,
	}

	return client.UploadFileEdgeADCApi("/POST/19?send=sslimport&iAction=3&iType=1", extraParams, filePath)
}

func DeleteCertificate(client *API, name string) error {
	certDelete := swagger.RemoveSsl{CetificateName: name, PasteSignature: "", ButtonFlag: "3"}
	jsonBytes, err := json.Marshal(certDelete)
	// Always GET before POST to avoid "Another user has made changes" error
	_, _ = client.GetEdgeADCObject("/GET/19")
	_, err = client.PostEdgeADCApi("/POST/19?iAction=2&iType=4", jsonBytes)
	return err
}

func CertificateExists(certificateData swagger.GetCertificate, certName string) bool {
	for _, option := range certificateData.ACECertificateListString.Options.Option {
		if option.Id == certName {
			return true
		}
	}
	return false
}
