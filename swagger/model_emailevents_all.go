/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmaileventsAll struct {
	// true
	Success string `json:"success,omitempty"`
	// images/light-hd_red.gif
	StatusImage string `json:"StatusImage,omitempty"`
	// get
	StatusText string `json:"StatusText,omitempty"`
	SecurityCombo *EmaileventsAllSecurityCombo `json:"SecurityCombo,omitempty"`
	// a.x@a.com
	CEERemoteEmailAddress string `json:"CEERemoteEmailAddress,omitempty"`
	// a.b@x.com
	CEELocalEmailAddress string `json:"CEELocalEmailAddress,omitempty"`
	// mail.x.com
	CEEMailServer string `json:"CEEMailServer,omitempty"`
	// 25
	CEEPort string `json:"CEEPort,omitempty"`
	// 1
	CEETimeout string `json:"CEETimeout,omitempty"`
	// 1
	CEEAuthorisationRequired string `json:"CEEAuthorisationRequired,omitempty"`
	// 0
	CEESecurity string `json:"CEESecurity,omitempty"`
	// a.x@x.co
	CEEMailServerLogon string `json:"CEEMailServerLogon,omitempty"`
	// 1
	CEEEvIpServiceNoticeEnabled string `json:"CEEEvIpServiceNoticeEnabled,omitempty"`
	// 1
	CEEEvIpServiceNoticeText string `json:"CEEEvIpServiceNoticeText,omitempty"`
	// 1
	CEEEvIpServiceAlertTextEnabled string `json:"CEEEvIpServiceAlertTextEnabled,omitempty"`
	// 10
	CEEEvIpServiceAlertText string `json:"CEEEvIpServiceAlertText,omitempty"`
	// 0
	CEEEvChannelNoticeEnabled string `json:"CEEEvChannelNoticeEnabled,omitempty"`
	// 2
	CEEEvChannelNoticeText string `json:"CEEEvChannelNoticeText,omitempty"`
	// 1
	CEEEvChannelAlertEnabled string `json:"CEEEvChannelAlertEnabled,omitempty"`
	// 11
	CEEEvChannelAlertText string `json:"CEEEvChannelAlertText,omitempty"`
	// 1
	CEEEvContentServerNoticeEnabled string `json:"CEEEvContentServerNoticeEnabled,omitempty"`
	// 3
	CEEEvContentServerNoticeText string `json:"CEEEvContentServerNoticeText,omitempty"`
	// 1
	CEEEvContentServerAlertEnabled string `json:"CEEEvContentServer AlertEnabled,omitempty"`
	// 12
	CEEEvContentServerAlertText string `json:"CEEEvContentServerAlertText,omitempty"`
	// 0
	CEEEvFlightPathEnabled string `json:"CEEEvFlightPathEnabled,omitempty"`
	// 4
	CEEEvFlightPathText string `json:"CEEEvFlightPathText,omitempty"`
	// 1
	CEEGroupEnabled string `json:"CEEGroupEnabled,omitempty"`
	// 5
	CEEGroupText string `json:"CEEGroupText,omitempty"`
	// 6
	CEEGroupTime string `json:"CEEGroupTime,omitempty"`
	// 1
	CEEEvDiskSpaceEnabled string `json:"CEEEvDiskSpaceEnabled,omitempty"`
	// 8
	CEEEvDiskSpaceText string `json:"CEEEvDiskSpaceText,omitempty"`
	// 7
	CEEEvLicenceRenevalText string `json:"CEEEvLicenceRenevalText,omitempty"`
	// 1
	CEEEvLicenceRenevalEnabled string `json:"CEEEvLicenceRenevalEnabled,omitempty"`
	// 9
	CEEEvDiskSpacePercent string `json:"CEEEvDiskSpacePercent,omitempty"`
	CEEEvTestTempFile string `json:"CEEEvTestTempFile,omitempty"`
	CEEEvTestProcessID string `json:"CEEEvTestProcessID,omitempty"`
	CEEEvTestEndTime string `json:"CEEEvTestEndTime,omitempty"`
	CEEEvTestStatus string `json:"CEEEvTestStatus,omitempty"`
}
