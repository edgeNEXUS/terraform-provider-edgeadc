/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ApplianceGridOpt struct {
	// 1
	Id string `json:"id,omitempty"`
	// eth0
	Ethtype string `json:"ethtype,omitempty"`
	// 192.168.1.175
	Address string `json:"address,omitempty"`
	// 255.255.255.0
	Mask string `json:"mask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	// Green Side
	Desc string `json:"desc,omitempty"`
	// 0
	RestEnabled string `json:"RestEnabled,omitempty"`
	// 1
	WebconsoleChecked string `json:"webconsoleChecked,omitempty"`
	// true
	RpFilter string `json:"RpFilter,omitempty"`
	EthVlan string `json:"ethVlan,omitempty"`
}
