/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ConditionsRow struct {
	// 0
	CId string `json:"cId,omitempty"`
	// Path
	Condition string `json:"condition,omitempty"`
	// Does
	Sense string `json:"sense,omitempty"`
	// Match RegEx
	Check string `json:"check,omitempty"`
	Match string `json:"match,omitempty"`
	// \\\\.htm$
	CondValue string `json:"condValue,omitempty"`
}
