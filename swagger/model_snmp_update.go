/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SnmpUpdate struct {
	// true
	SNMPV1V2Checked string `json:"SNMPV1V2Checked,omitempty"`
	// jetnexus
	SNMPCommunityString string `json:"SNMPCommunityString,omitempty"`
	// true
	SNMPV3Checked string `json:"SNMPV3Checked,omitempty"`
	// 12
	OldPassPhrase string `json:"oldPassPhrase,omitempty"`
	// 23455667
	NewPassPhrase string `json:"newPassPhrase,omitempty"`
	// 12233413
	ConfirmNewPassPhrase string `json:"confirmNewPassPhrase,omitempty"`
}
