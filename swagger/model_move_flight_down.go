/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MoveFlightDown struct {
	// 1
	EditedInterface string `json:"editedInterface,omitempty"`
	// 0
	EditedChannel string `json:"editedChannel,omitempty"`
	// Swap HTTP to HTTPS
	FlightPathName string `json:"flightPathName,omitempty"`
}
