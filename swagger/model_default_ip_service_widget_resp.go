/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DefaultIpServiceWidgetResp struct {
	// true
	Success string `json:"success,omitempty"`
	// images/light-hd_red.gif
	StatusImage string `json:"StatusImage,omitempty"`
	// get
	StatusText string `json:"StatusText,omitempty"`
	// IPService4
	PartID string `json:"PartID,omitempty"`
	// System
	IPServiceType string `json:"IPService_Type,omitempty"`
	IPServiceVSRS string `json:"IPService_VSRS,omitempty"`
	// minute
	IPServicePeriod string `json:"IPService_Period,omitempty"`
	// Average Bytes in
	IPServiceColumns string `json:"IPService_Columns,omitempty"`
}
