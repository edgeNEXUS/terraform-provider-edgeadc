/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TroubleSupportResp struct {
	// true
	Success string `json:"success,omitempty"`
	// images/2light-hd.gif
	StatusImage string `json:"StatusImage,omitempty"`
	// success
	StatusText string `json:"StatusText,omitempty"`
	// 20438
	ProcessID string `json:"Process_ID,omitempty"`
}
