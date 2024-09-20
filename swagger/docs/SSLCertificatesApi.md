# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET19Get**](SSLCertificatesApi.md#GET19Get) | **Get** /GET/19 | SSL page all combo options
[**POST19downloadsslexportiAction4nameTestcert2pastest12345preDownloadCheckyesPost**](SSLCertificatesApi.md#POST19downloadsslexportiAction4nameTestcert2pastest12345preDownloadCheckyesPost) | **Post** /POST/19?download&#x3D;sslexport&amp;iAction&#x3D;4&amp;name&#x3D;Testcert2&amp;pas&#x3D;test12345&amp;preDownloadCheck&#x3D;yes | Export single/multiple certificate
[**POST19iAction1iType1Post**](SSLCertificatesApi.md#POST19iAction1iType1Post) | **Post** /POST/19?iAction&#x3D;1&amp;iType&#x3D;1 | Create new SSL
[**POST19iAction1iType2requestcertPost**](SSLCertificatesApi.md#POST19iAction1iType2requestcertPost) | **Post** /POST/19?iAction&#x3D;1&amp;iType&#x3D;2&amp;request&#x3D;cert | Create new SSL cert request
[**POST19iAction2iType1showcertPost**](SSLCertificatesApi.md#POST19iAction2iType1showcertPost) | **Post** /POST/19?iAction&#x3D;2&amp;iType&#x3D;1&amp;show&#x3D;cert | Show details of SSL cert
[**POST19iAction2iType2Post**](SSLCertificatesApi.md#POST19iAction2iType2Post) | **Post** /POST/19?iAction&#x3D;2&amp;iType&#x3D;2 | Install ssl cert
[**POST19iAction2iType3Post**](SSLCertificatesApi.md#POST19iAction2iType3Post) | **Post** /POST/19?iAction&#x3D;2&amp;iType&#x3D;3 | Intermediate install ssl cert
[**POST19iAction2iType4Post**](SSLCertificatesApi.md#POST19iAction2iType4Post) | **Post** /POST/19?iAction&#x3D;2&amp;iType&#x3D;4 | Remove ssl cert
[**POST19iAction2iType5renewcertPost**](SSLCertificatesApi.md#POST19iAction2iType5renewcertPost) | **Post** /POST/19?iAction&#x3D;2&amp;iType&#x3D;5&amp;renew&#x3D;cert | Renew ssl cert
[**POST19iAction2iType6Post**](SSLCertificatesApi.md#POST19iAction2iType6Post) | **Post** /POST/19?iAction&#x3D;2&amp;iType&#x3D;6 | Reorder ssl cert
[**POST19iAction2iType7Post**](SSLCertificatesApi.md#POST19iAction2iType7Post) | **Post** /POST/19?iAction&#x3D;2&amp;iType&#x3D;7 | Update Reorder
[**POST19iAction3iType1sendsslimportPost**](SSLCertificatesApi.md#POST19iAction3iType1sendsslimportPost) | **Post** /POST/19?iAction&#x3D;3&amp;iType&#x3D;1&amp;send&#x3D;sslimport | Import single SSL cert
[**POST19iAction5iType1sendsslimportPost**](SSLCertificatesApi.md#POST19iAction5iType1sendsslimportPost) | **Post** /POST/19?iAction&#x3D;5&amp;iType&#x3D;1&amp;send&#x3D;sslimport | Import multiple SSL cert

# **GET19Get**
> GetCertificate GET19Get(ctx, )
SSL page all combo options

This service is used to FETCH all data for combo boxes used in SSL Certificate page, it uses three API structures .    Adv_CertCreate: Used to get Contry and Cert key length combo items.    Adv_CertMgmt: Used to get list of all SSL Certificates available with status(Self signed, pending etc).    Adv_CertExport: Used to get all SSL Cert available to export.          The nodes of JSON received, contains the given information        CerKeyLenghtCombo: This node contains options of certificate key length.     CountryCombo: This node contains options of countries.     CertificateManageCombo: This node contains options of ssl certificates available to manage(install, delete etc).     ACE_CertificateListString: This node contains options of ssl certificates available to export. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetCertificate**](GetCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19downloadsslexportiAction4nameTestcert2pastest12345preDownloadCheckyesPost**
> POST19downloadsslexportiAction4nameTestcert2pastest12345preDownloadCheckyesPost(ctx, )
Export single/multiple certificate

This service is used to export(download) Single/Multiple SSL Certificate, you have to send certificate_name and password in query string. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19iAction1iType1Post**
> ManageCertificate POST19iAction1iType1Post(ctx, body)
Create new SSL

This service is used to create new SSL certificate, you have to send information of SSL Certificate(Name, Organization, City etc) that needs to be created, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-      SslCertificateNameText: Certificate name.     SslOrganisationText: Organisation name.     SslOrganisationUnitNameText: Organisation unit .     SslCountryText: id(Received as id in GET/12 service, node CountryCombo) of option selected from Country options.     SslCityText: City information.     SslCountyText:      SslDomainText: Domain Name.     SslKeyLengthText: Key length.     SSLPeriod: Period     ButtonFlag:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSsl**](CreateSsl.md)| Object that needs to be sent to the server | 

### Return type

[**ManageCertificate**](ManageCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19iAction1iType2requestcertPost**
> ManageCertificate POST19iAction1iType2requestcertPost(ctx, body)
Create new SSL cert request

This service is used to create new SSL certificate request that'll generate a key, you have to send information of SSL Certificate(Name, Organization, City etc) that needs to be created, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        SslCertificateNameText: Certificate name.     SslOrganisationText: Organisation name.     SslOrganisationUnitNameText: Organisation unit .     SslCountryText: id(Received as id in GET/12 service, node CountryCombo) of option selected from Country options.     SslCityText: City information.     SslCountyText:      SslDomainText: Domain Name.     SslKeyLengthText: Key length.     SSLPeriod: Period     ButtonFlag:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSslCert**](CreateSslCert.md)| Object that needs to be sent to the server | 

### Return type

[**ManageCertificate**](ManageCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19iAction2iType1showcertPost**
> DetailsSsl POST19iAction2iType1showcertPost(ctx, body)
Show details of SSL cert

This service is used to show details of avaialable SSL certificate.    JSON that needs to be sent as request payload, Should contains given information-        CetificateName: id(Received as id in GET/19 service, node CertificateManageCombo) of option selected from Certificate options.     PasteSignature:PasteSignature/n

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DetailsSslReq**](DetailsSslReq.md)| Object that needs to be sent to the server | 

### Return type

[**DetailsSsl**](DetailsSSL.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19iAction2iType2Post**
> ManageCertificate POST19iAction2iType2Post(ctx, body)
Install ssl cert

This service is used to install avaialable SSL certificate, you have to send Cert name with signature key, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-          CetificateName: id(Received as id in GET/19 service, node CertificateManageCombo) of option selected from Certificate options.     PasteSignature: Certificate signature text.     ButtonFlag: '' 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InstallSsl**](InstallSsl.md)| Object that needs to be sent to the server | 

### Return type

[**ManageCertificate**](ManageCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19iAction2iType3Post**
> ManageCertificate POST19iAction2iType3Post(ctx, body)
Intermediate install ssl cert

This service is used to intermediate install avaialable SSL certificate, you have to send Cert name with signature key, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CetificateName: id(Received as id in GET/19 service, node CertificateManageCombo) of option selected from Certificate options.     PasteSignature: Certificate signature text.      ButtonFlag: '' 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IntermediateInstallSsl**](IntermediateInstallSsl.md)| Object that needs to be sent to the server | 

### Return type

[**ManageCertificate**](ManageCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19iAction2iType4Post**
> ManageCertificate POST19iAction2iType4Post(ctx, body)
Remove ssl cert

This service is used to remove avaialable SSL certificate, you have to send Cert name, that needs to be removed, with POST request json    JSON that needs to be sent as request payload, Should contains given information-        CetificateName: id(Received as id in GET/19 service, node CertificateManageCombo) of option selected from Certificate options.     PasteSignature: Certificate signature text.     ButtonFlag: '' 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveSsl**](RemoveSsl.md)| Object that needs to be sent to the server | 

### Return type

[**ManageCertificate**](ManageCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19iAction2iType5renewcertPost**
> RenewSslResponse POST19iAction2iType5renewcertPost(ctx, body)
Renew ssl cert

This service is used to renew avaialable SSL certificate, you have to send Cert name, that needs to be renewed, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CetificateName: id(Received as id in GET/19 service, node CertificateManageCombo) of option selected from Certificate options.     PasteSignature: Certificate signature text.     ButtonFlag: '' 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RenewSsl**](RenewSsl.md)| Object that needs to be sent to the server | 

### Return type

[**RenewSslResponse**](RenewSSL_Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19iAction2iType6Post**
> ReorderSslResponse POST19iAction2iType6Post(ctx, body)
Reorder ssl cert

This service is used to reorder avaialable SSL certificate, you have to send Cert name, that needs to be reordered, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CetificateName: id(Received as id in GET/19 service, node CertificateManageCombo) of option selected from Certificate options. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReorderSsl**](ReorderSsl.md)| Object that needs to be sent to the server | 

### Return type

[**ReorderSslResponse**](ReorderSSL_Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19iAction2iType7Post**
> ManageCertificate POST19iAction2iType7Post(ctx, body)
Update Reorder

This service is used to Update Reorder SSL certificate, you have to send information of SSL Certificate(CetificateName, PasteSignature) that needs to be created, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-      CetificateName: Certificate name.     PasteSignature: PasteSignature. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateReorderSsl**](UpdateReorderSsl.md)| Object that needs to be sent to the server | 

### Return type

[**ManageCertificate**](ManageCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19iAction3iType1sendsslimportPost**
> ManageCertificate POST19iAction3iType1sendsslimportPost(ctx, name, value, file)
Import single SSL cert

This service is used to upload SSL Certificate, you have to send name and password in form field(SslCertificatesImportCertificateNameText,SslCertificatesImportPasswordText) with .pfx file, that needs to be uploaded, by using form submit.    Form fields that needs to be sent-        SslCertificatesImportCertificateNameText: Certificate name     SslCertificatesImportPasswordText: Password text 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **value** | **string**|  | 
  **file** | ***os.File*****os.File**|  | 

### Return type

[**ManageCertificate**](ManageCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST19iAction5iType1sendsslimportPost**
> ManageCertificate POST19iAction5iType1sendsslimportPost(ctx, name, value, file)
Import multiple SSL cert

This service is used to upload Multiple SSL Certificate, you have to send password in form field(SslCertificatesImportPasswordText) with .jnbk file, that needs to be uploaded, by using form submit.    Form fields that needs to be sent-        SslMultiCertificatesImportCertificateNameText: Certificate name     SslMultiCertificatesImportPasswordText: Password text 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **value** | **string**|  | 
  **file** | ***os.File*****os.File**|  | 

### Return type

[**ManageCertificate**](ManageCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

