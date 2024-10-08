/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AppsDownResponse struct {
	// true
	Success string `json:"success,omitempty"`
	// images/light-hd_red.gif
	StatusImage string `json:"StatusImage,omitempty"`
	// Download Complete
	StatusText string `json:"StatusText,omitempty"`
	TotalSize string `json:"TotalSize,omitempty"`
	DownloadedSize string `json:"DownloadedSize,omitempty"`
}
