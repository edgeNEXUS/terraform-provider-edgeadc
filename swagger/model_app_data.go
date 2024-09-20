/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AppData struct {
	// true
	Success string `json:"success,omitempty"`
	// images/light-hd_red.gif.
	StatusImage string `json:"StatusImage,omitempty"`
	// get.
	StatusText string `json:"StatusText,omitempty"`
	// admin.
	UserID string `json:"UserID,omitempty"`
	// 0800-2745-80BF.
	MachineID string `json:"MachineID,omitempty"`
	// Noida, India.
	Location string `json:"Location,omitempty"`
	// 113.193.186.241.
	PublicIP string `json:"PublicIP,omitempty"`
	// E845FF30-5AC9-467C-80AF-1B61731B6615.
	LicenceID string `json:"LicenceID,omitempty"`
	// 4.2.1 (Build 1700) 7g0859.
	SoftwareVersion string `json:"SoftwareVersion,omitempty"`
	// https://appstore.jetnexuscloud.com.
	AppStoreUrl string `json:"AppStoreUrl,omitempty"`
}
