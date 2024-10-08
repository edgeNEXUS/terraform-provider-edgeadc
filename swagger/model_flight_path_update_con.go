/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type FlightPathUpdateCon struct {
	// 1
	FId string `json:"fId,omitempty"`
	// 1
	CId string `json:"cId,omitempty"`
	// Host
	Condition string `json:"condition,omitempty"`
	Match string `json:"match,omitempty"`
	// Does
	Sense string `json:"sense,omitempty"`
	// Exist
	Check string `json:"check,omitempty"`
	CondValue string `json:"condValue,omitempty"`
}
