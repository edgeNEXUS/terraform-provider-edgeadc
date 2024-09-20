/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateEmailEvent struct {
	CEEAuthorisationRequired string `json:"CEEAuthorisationRequired,omitempty"`
	CEEEMailServerPassword string `json:"CEEEMailServerPassword,omitempty"`
	CEELocalEmailAddress string `json:"CEELocalEmailAddress,omitempty"`
	CEEMailServer string `json:"CEEMailServer,omitempty"`
	CEEMailServerLogon string `json:"CEEMailServerLogon,omitempty"`
	CEEPort string `json:"CEEPort,omitempty"`
	CEERemoteEmailAddress string `json:"CEERemoteEmailAddress,omitempty"`
	CEESecurity string `json:"CEESecurity,omitempty"`
	CEETimeout string `json:"CEETimeout,omitempty"`
	Test string `json:"Test,omitempty"`
}
