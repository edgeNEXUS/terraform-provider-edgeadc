/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type HisdatasetOpt struct {
	// 192.168.1.34:67
	IpPort string `json:"ipPort,omitempty"`
	RS []HisRsOpt `json:"RS,omitempty"`
}
