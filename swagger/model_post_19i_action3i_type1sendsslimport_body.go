/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"os"
)

type Post19iAction3iType1sendsslimportBody struct {
	// SslCertificatesImportCertificateNameText
	Name string `json:"name"`
	// ABC
	Value string `json:"value"`
	// .pfx file
	File **os.File `json:"file"`
}
