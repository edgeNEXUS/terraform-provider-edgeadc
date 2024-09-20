/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RemoveWidgetResp struct {
	// true
	Success string `json:"success,omitempty"`
	// images/light-hd_red.gif
	StatusImage string `json:"StatusImage,omitempty"`
	// Your changes have been applied
	StatusText string `json:"StatusText,omitempty"`
	DashboardEventName string `json:"DashboardEvent_Name,omitempty"`
	DashboardEventFilter string `json:"DashboardEvent_Filter,omitempty"`
	IPServiceName string `json:"IPService_Name,omitempty"`
	IPServiceType string `json:"IPService_Type,omitempty"`
	IPServiceVSRS string `json:"IPService_VSRS,omitempty"`
	IPServicePeriod string `json:"IPService_Period,omitempty"`
	IPServiceColumns string `json:"IPService_Columns,omitempty"`
	IPStatusLayout string `json:"IPStatus_Layout,omitempty"`
	SystemName string `json:"System_Name,omitempty"`
	SystemType string `json:"System_Type,omitempty"`
}
