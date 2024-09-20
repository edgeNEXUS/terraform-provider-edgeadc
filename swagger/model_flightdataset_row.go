/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type FlightdatasetRow struct {
	// 1
	FId string `json:"fId,omitempty"`
	// false
	DiagnosticTracing string `json:"DiagnosticTracing,omitempty"`
	// HTML Extension
	FlightPathName string `json:"flightPathName,omitempty"`
	// Fixes all .htm requests to .html
	FlightPathDesc string `json:"flightPathDesc,omitempty"`
	// 192.168.1.34,192.168.1.250,192.168.1.250
	FlightPathInUse string `json:"flightPathInUse,omitempty"`
	Conditions *FlightdatasetRowConditions `json:"conditions,omitempty"`
	Values *FlightdatasetRowValues `json:"values,omitempty"`
	Actions *FlightdatasetRowActions `json:"actions,omitempty"`
}
