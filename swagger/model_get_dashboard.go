/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GetDashboard struct {
	// true
	Success string `json:"success,omitempty"`
	// images/light-hd_red.gif
	StatusImage string `json:"StatusImage,omitempty"`
	// get
	StatusText string `json:"StatusText,omitempty"`
	ConfiguredWidgetsComboStore *GetDashboardConfiguredWidgetsComboStore `json:"ConfiguredWidgetsComboStore,omitempty"`
	// no_data
	IsEdited string `json:"isEdited,omitempty"`
	Data []DashdataOpt `json:"data,omitempty"`
}
