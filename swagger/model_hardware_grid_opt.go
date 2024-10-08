/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type HardwareGridOpt struct {
	// 0
	Id string `json:"id,omitempty"`
	// auto
	Duplex string `json:"duplex,omitempty"`
	// Auto
	Speed string `json:"speed,omitempty"`
	// eth0
	EthName string `json:"ethName,omitempty"`
	// none
	EthBond string `json:"ethBond,omitempty"`
	EthStatus string `json:"ethStatus,omitempty"`
}
