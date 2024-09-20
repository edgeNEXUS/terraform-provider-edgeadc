/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateRemotelog struct {
	// true
	RemoteLogStorage string `json:"RemoteLogStorage,omitempty"`
	// 192.168.1.23
	RemoteLogIP string `json:"RemoteLogIP,omitempty"`
	// w3c
	RemoteLogShare string `json:"RemoteLogShare,omitempty"`
	// test
	RemoteLogDir string `json:"RemoteLogDir,omitempty"`
	// test sd
	RemoteLogUser string `json:"RemoteLogUser,omitempty"`
	RemoteLogPassword string `json:"RemoteLogPassword,omitempty"`
}
