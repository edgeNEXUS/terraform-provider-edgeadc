/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DragDropFlight struct {
	// 1
	EditedInterface string `json:"editedInterface,omitempty"`
	// 0
	EditedChannel string `json:"editedChannel,omitempty"`
	// 2
	FlightPathDragId string `json:"flightPathDragId,omitempty"`
	// Close Folders
	FlightPathDragName string `json:"flightPathDragName,omitempty"`
	// 4
	FlightPathDropId string `json:"flightPathDropId,omitempty"`
	// HTML Extension
	FlightPathDropName string `json:"flightPathDropName,omitempty"`
}
