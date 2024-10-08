/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LoggingData struct {
	// true
	Success string `json:"success,omitempty"`
	// images/light-hd_red.gif
	StatusImage string `json:"StatusImage,omitempty"`
	// get
	StatusText string `json:"StatusText,omitempty"`
	W3CCombo *LoggingDataW3CCombo `json:"W3CCombo,omitempty"`
	W3CFCombo *LoggingDataW3CfCombo `json:"W3CFCombo,omitempty"`
	SecurityLogCombo *LoggingDataSecurityLogCombo `json:"SecurityLogCombo,omitempty"`
	// On
	SecurityLog string `json:"SecurityLog,omitempty"`
	// 0
	RemoteLogStorage string `json:"RemoteLogStorage,omitempty"`
	RemoteLogIP string `json:"RemoteLogIP,omitempty"`
	// w3c
	RemoteLogShare string `json:"RemoteLogShare,omitempty"`
	RemoteLogDir string `json:"RemoteLogDir,omitempty"`
	// 0
	AdvancedW3C string `json:"AdvancedW3C,omitempty"`
	// 0
	LogMessageLevel string `json:"LogMessageLevel,omitempty"`
	RemoteLogUser string `json:"RemoteLogUser,omitempty"`
	// full
	W3cLoggingModeList string `json:"W3cLoggingModeList,omitempty"`
	XffLoggingList string `json:"XffLoggingList,omitempty"`
	RsysLog1IP string `json:"RsysLog1IP,omitempty"`
	RsysLog2IP string `json:"RsysLog2IP,omitempty"`
	RsysLog1PROTOCOL string `json:"RsysLog1PROTOCOL,omitempty"`
	RsysLog2PROTOCOL string `json:"RsysLog2PROTOCOL,omitempty"`
	// 0
	RsysServerPort1 string `json:"RsysServerPort1,omitempty"`
	// 0
	RsysServerPort2 string `json:"RsysServerPort2,omitempty"`
	// 0
	RsysServer1enabled string `json:"RsysServer1enabled,omitempty"`
	// 0
	RsysServer2enabled string `json:"RsysServer2enabled,omitempty"`
	RsysProtocolCombo *LoggingDataRsysProtocolCombo `json:"RsysProtocolCombo,omitempty"`
	LogClearListString *LoggingDataLogClearListString `json:"LogClearListString,omitempty"`
	LogMessageLevelCombo *LoggingDataLogMessageLevelCombo `json:"LogMessageLevelCombo,omitempty"`
}
