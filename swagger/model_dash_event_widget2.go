/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DashEventWidget2 struct {
	// Event-event
	SectionName string `json:"Section_Name,omitempty"`
	// Cpu
	CPU string `json:"CPU,omitempty"`
	// Event-Disk
	Disk string `json:"Disk,omitempty"`
	// GraphMemory
	GraphMemory string `json:"GraphMemory,omitempty"`
	// PageTitle
	PageTitle string `json:"PageTitle,omitempty"`
	Headlines *DashEventWidget2Headlines `json:"Headlines,omitempty"`
}
