/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SoftwareDownloadStatusResponse struct {
	Success string `json:"success,omitempty"`
	StatusImage string `json:"StatusImage,omitempty"`
	StatusText string `json:"StatusText,omitempty"`
	TotalSize string `json:"TotalSize,omitempty"`
	DownloadedSize string `json:"DownloadedSize,omitempty"`
}
