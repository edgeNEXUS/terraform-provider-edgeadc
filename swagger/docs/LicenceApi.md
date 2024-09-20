# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET27Get**](LicenceApi.md#GET27Get) | **Get** /GET/27 | Page Footer data
[**GET7Get**](LicenceApi.md#GET7Get) | **Get** /GET/7 | License page data
[**GET8Get**](LicenceApi.md#GET8Get) | **Get** /GET/8 | License support information
[**POST7iAction1sendlicensePost**](LicenceApi.md#POST7iAction1sendlicensePost) | **Post** /POST/7?iAction&#x3D;1&amp;send&#x3D;license | Upload license
[**POST7iAction2iType1Post**](LicenceApi.md#POST7iAction2iType1Post) | **Post** /POST/7?iAction&#x3D;2&amp;iType&#x3D;1 | Upload License(Licence Text)

# **GET27Get**
> LicenceFooterData GET27Get(ctx, )
Page Footer data

This service is used to FETCH footer details.    success:       StatusText:       StatusImage:       PageTitle:  

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LicenceFooterData**](Licence_FooterData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET7Get**
> LicenceData GET7Get(ctx, )
License page data

This service is used to FETCH licence details, it uses two API structure.    Setup_LicenseRows: Used to get License information like license of how many VS and RS and for how much time it'll be valid.       Config_Bonding: Used to get Machine ID.        The nodes of JSON received, contains the given information          success:       IssuedTo: This node contains information, to whom this license is issued.       ContactPerson: This node contains information of contact person.       DateIssued: This node contains date when licence was issued.       ServerRef: This node contains ALB name.       ALB: This node contains information of ALB licecne(Expiry and other information).       SSL: This node contains information of SSL licecne(Expiry and other information).       Cache: This node contains information of Cache licecne(Expiry and other information).       FlightPath: This node contains information of FlightPath licecne(Expiry and other information).       JetSTREAM: This node contains information of JetSTREAM licecne(Expiry and other information).       Interfaces: This node contains information of Interface licecne(Expiry and other information).       RealServers: This node contains information of Real Server licecne(Expiry and other information).       MachineID: This node contains Machine ID of this ALB.       UUID: This node contains UUID.       GLBLicense:       AuthLicense:       FirewallLicense:       Throughput:       B2BLicense:       Layer4BaseLicense:       HttpBaseLicense:       ALB1:       SSL1:       Cache1:       FlightPath1:       JetSTREAM1:       Interfaces1:       RealServers1:       MachineID1:       GLBLicense1:       AuthLicense1:       FirewallLicense1:       Throughput1:       B2BLicense1:       Layer4BaseLicense1:       HttpBaseLicense1:       data:data. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LicenceData**](Licence_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET8Get**
> GET8Get(ctx, )
License support information

This service is used to FETCH licence support information details, it uses one API structure.     Support_Information: Used to get License support information like config file, license conf and system logs.      Config_Bonding: Used to get Machine ID.      We receive a string separated by a unique combination of characters(^#$@) because a huge amount of data(all data from config, license and syslog files).      It contains the given information in indexes when splitted by ^#$@          0: This index contains error information otherwise 'true'.     1: This index contains MAC address.     2: This index contains ALB name.     3: This index contains OS VERSION.     4: This index contains all contents of configuration file.     5: This index contains all contents of licence configuration file.     6: This index contains curent software version.     7: This index contains contents of most recent syslog file. 

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

# **POST7iAction1sendlicensePost**
> UploadLicence POST7iAction1sendlicensePost(ctx, file)
Upload license

This service is used to upload License, you have to send license file, that needs to be uploaded, by using form submit.     

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 

### Return type

[**UploadLicence**](Upload_licence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST7iAction2iType1Post**
> UploadLicence POST7iAction2iType1Post(ctx, body)
Upload License(Licence Text)

This service is used to Upload License, you have to send Licence Text , that needs to be uploaded.    JSON that needs to be sent as request payload, Should contains given information-        sLicenceText: licence text. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PastLicence**](PastLicence.md)| Object that needs to be sent to the server | 

### Return type

[**UploadLicence**](Upload_licence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

