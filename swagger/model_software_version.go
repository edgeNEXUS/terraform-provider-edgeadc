/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SoftwareVersion struct {
	// true
	Success string `json:"success,omitempty"`
	// images/2light-hd.gif
	StatusImage string `json:"StatusImage,omitempty"`
	// get
	StatusText string `json:"StatusText,omitempty"`
	Loadevent string `json:"loadevent,omitempty"`
	DevVer string `json:"DevVer,omitempty"`
}
