/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DetailsSsl struct {
	// true
	Success string `json:"success,omitempty"`
	// true
	StatusImage string `json:"StatusImage,omitempty"`
	// Your changes have been applied/or Error text
	StatusText string `json:"StatusText,omitempty"`
	// NewCertSign
	CertificateName string `json:"CertificateName,omitempty"`
	// mit
	Organization string `json:"Organization,omitempty"`
	// 23
	OrganizationalUnit string `json:"OrganizationalUnit,omitempty"`
	// test
	CityLocality string `json:"CityLocality,omitempty"`
	// test
	StateProvince string `json:"StateProvince,omitempty"`
	// com
	Country string `json:"Country,omitempty"`
	// test
	DomainName string `json:"DomainName,omitempty"`
	// 512
	KeyLength string `json:"KeyLength,omitempty"`
	Periods string `json:"Periods,omitempty"`
	Expires string `json:"Expires,omitempty"`
}
