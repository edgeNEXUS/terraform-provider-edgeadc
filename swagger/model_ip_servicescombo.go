/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type IpServicescombo struct {
	// no
	IsNoVIP string `json:"isNoVIP,omitempty"`
	ServiceTypeCombo *IpServicescomboServiceTypeCombo `json:"ServiceTypeCombo,omitempty"`
	SSLCertCombo *IpServicescomboServiceTypeCombo `json:"SSLCertCombo,omitempty"`
	CacheRuleCombo *IpServicescomboServiceTypeCombo `json:"CacheRuleCombo,omitempty"`
	MonitorPolicyCombo *IpServicescomboServiceTypeCombo `json:"MonitorPolicyCombo,omitempty"`
	SSLClientCertCombo *IpServicescomboServiceTypeCombo `json:"SSLClientCertCombo,omitempty"`
	LBPolicyCombo *IpServicescomboServiceTypeCombo `json:"LBPolicyCombo,omitempty"`
	LBPolicyFTPCombo *IpServicescomboServiceTypeCombo `json:"LBPolicyFTPCombo,omitempty"`
	ConnectivityCombo *IpServicescomboServiceTypeCombo `json:"ConnectivityCombo,omitempty"`
	CSActivityCombo *IpServicescomboServiceTypeCombo `json:"CSActivityCombo,omitempty"`
	AccelerateCombo *IpServicescomboServiceTypeCombo `json:"AccelerateCombo,omitempty"`
	FlightPathSelectionList *IpServicescomboServiceTypeCombo `json:"FlightPathSelectionList,omitempty"`
	SSLCipherListString *IpServicescomboServiceTypeCombo `json:"SSLCipherListString,omitempty"`
	SNICertificateListString *IpServicescomboSniCertificateListString `json:"SNICertificateListString,omitempty"`
	SSLRenegotiationString *IpServicescomboSslRenegotiationString `json:"SSLRenegotiationString,omitempty"`
	SecurityLogString *IpServicescomboSecurityLogString `json:"SecurityLogString,omitempty"`
}
