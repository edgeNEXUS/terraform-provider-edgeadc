# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET24Get**](SoftwareApi.md#GET24Get) | **Get** /GET/24 | Current software version
[**GET39pageSoftwareGet**](SoftwareApi.md#GET39pageSoftwareGet) | **Get** /GET/39?page&#x3D;Software | Licence details
[**GET49Get**](SoftwareApi.md#GET49Get) | **Get** /GET/49 | Smart file
[**GET59Get**](SoftwareApi.md#GET59Get) | **Get** /GET/59 | Location details
[**GET60Get**](SoftwareApi.md#GET60Get) | **Get** /GET/60 | Check valid .software.alb file
[**GET64Get**](SoftwareApi.md#GET64Get) | **Get** /GET/64 | Check Service running status
[**POST26iAction11iType1sendsoftwarePost**](SoftwareApi.md#POST26iAction11iType1sendsoftwarePost) | **Post** /POST/26?iAction&#x3D;11&amp;iType&#x3D;1&amp;send&#x3D;software | Upload smart file
[**POST26iAction13iType1Post**](SoftwareApi.md#POST26iAction13iType1Post) | **Post** /POST/26?iAction&#x3D;13&amp;iType&#x3D;1 | Download from cloud
[**POST26iAction2iType1sendsoftwarePost**](SoftwareApi.md#POST26iAction2iType1sendsoftwarePost) | **Post** /POST/26?iAction&#x3D;2&amp;iType&#x3D;1&amp;send&#x3D;software | Apply software
[**POST26iAction3iType1Post**](SoftwareApi.md#POST26iAction3iType1Post) | **Post** /POST/26?iAction&#x3D;3&amp;iType&#x3D;1 | Download status or check file exist
[**POST26iAction4iType1Post**](SoftwareApi.md#POST26iAction4iType1Post) | **Post** /POST/26?iAction&#x3D;4&amp;iType&#x3D;1 | Download status
[**POST26iAction5iType1Post**](SoftwareApi.md#POST26iAction5iType1Post) | **Post** /POST/26?iAction&#x3D;5&amp;iType&#x3D;1 | Remove software
[**POST26iAction7iType1Post**](SoftwareApi.md#POST26iAction7iType1Post) | **Post** /POST/26?iAction&#x3D;7&amp;iType&#x3D;1 | Upload software
[**POST42Post**](SoftwareApi.md#POST42Post) | **Post** /POST/42 | Save location details

# **GET24Get**
> SoftwareVersion GET24Get(ctx, )
Current software version

This service is used to get current runing software version details.    The nodes of JSON received, contains the given information        loadevent:     DevVer: 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareVersion**](Software_Version.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET39pageSoftwareGet**
> SoftwareLicence GET39pageSoftwareGet(ctx, )
Licence details

This service is used to fetch licence details.     The nodes of JSON received, contains the given information        UserID: This node contains user id.     MachineID: This node contains user machine id     Location: This node contains location.     PublicIP: This node contains public id.     LicenceID: This node contains licence id.      SoftwareVersion: This node contains software version .      AppStoreUrl: 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareLicence**](Software_Licence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET49Get**
> SoftwareData GET49Get(ctx, )
Smart file

This service is used to fetch smart file details, it uses one API structure.    The nodes of JSON received, contains the given information        SmartFiles:      File_Name: This node contains File Name.      Version: This node contains Version.     Message_Digest:      Machine_Id: This node contains Machine id.      Product_Id: This node contains Product id.     Order_Id: his node contains order id.      Icon: This node contains icon.     Description:     Name:     Type:     Certified:     Vendor:     ReleaseDate:     DockerImageName: 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareData**](Software_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET59Get**
> SoftwareLocation GET59Get(ctx, )
Location details

This service is used to fetch location details , it uses one API structure.    The nodes of JSON received, contains the given information        as: This node contains network name.      city: This node contains city name.      country: This node contains country.     countryCode: This node contains countryCode.      isp: This node contains isp.      lat:      lon:      org:     query: This node contains ip.     region: This node contains region.     regionName: This node contains regionName.     status:     timezone: This node contains timezone.     zip: This node contains pin code. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareLocation**](Software_Location.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET60Get**
> SoftwareLicenceChecKvalid GET60Get(ctx, )
Check valid .software.alb file

This service is used to Check valid .software.alb file.   

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareLicenceChecKvalid**](Software_LicenceCHECKvalid.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET64Get**
> SoftwareLicenceChecKservice GET64Get(ctx, )
Check Service running status

This service is used to Check Service running status.   

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareLicenceChecKservice**](Software_LicenceCHECKservice.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST26iAction11iType1sendsoftwarePost**
> SoftwareResponse POST26iAction11iType1sendsoftwarePost(ctx, )
Upload smart file

This service is used to upload smart file from local . 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareResponse**](Software_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST26iAction13iType1Post**
> SoftwareDownResp POST26iAction13iType1Post(ctx, body)
Download from cloud

This service is used to download new builds.    JSON that needs to be sent as request payload, Should contains given information-        URL: build url. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SoftwareCloudDownload**](SoftwareCloudDownload.md)| Object that needs to be sent to the server | 

### Return type

[**SoftwareDownResp**](Software_DownResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST26iAction2iType1sendsoftwarePost**
> SoftwareResponse POST26iAction2iType1sendsoftwarePost(ctx, )
Apply software

This service is used to upload and apply software file from local . 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareResponse**](Software_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST26iAction3iType1Post**
> SoftwareDownloadStatusResponse POST26iAction3iType1Post(ctx, body)
Download status or check file exist

This service is used to download status or check file exist.    JSON that needs to be sent as request payload, Should contains given information-        URL: build url.     Name: Name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SoftwareDownloadStatus**](SoftwareDownloadStatus.md)| Object that needs to be sent to the server | 

### Return type

[**SoftwareDownloadStatusResponse**](Software_DownloadStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST26iAction4iType1Post**
> SoftwareDownloadStatusResponse POST26iAction4iType1Post(ctx, body)
Download status

This service is used to download status.    JSON that needs to be sent as request payload, Should contains given information-        URL: build url.     Name: Name.     TotalSize: TotalSize. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SoftwareDownloadStatus**](SoftwareDownloadStatus.md)| Object that needs to be sent to the server | 

### Return type

[**SoftwareDownloadStatusResponse**](Software_DownloadStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST26iAction5iType1Post**
> SoftwareUploadResp POST26iAction5iType1Post(ctx, body)
Remove software

This service is used to remove smart file from local .    JSON that needs to be sent as request payload, Should contains given information-        File_Name: Name of file which is to delete. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SoftwareUpload**](SoftwareUpload.md)| Object that needs to be sent to the server | 

### Return type

[**SoftwareUploadResp**](Software_UploadResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST26iAction7iType1Post**
> SoftwareUploadResp POST26iAction7iType1Post(ctx, body, sequence)
Upload software

This service is used to upload smart file.    JSON that needs to be sent as request payload, Should contains given information-        File_Name: Name of file which is to upload. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SoftwareUpload**](SoftwareUpload.md)| Object that needs to be sent to the server | 
  **sequence** | **string**| request type | [default to time]

### Return type

[**SoftwareUploadResp**](Software_UploadResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST42Post**
> SoftwareSaveLocationResp POST42Post(ctx, body)
Save location details

This service is used to Save location details.    JSON that needs to be sent as request payload, Should contains given information-        Location: Brower location.     PublicIP: public ip. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SoftwareSaveLoc**](SoftwareSaveLoc.md)| Object that needs to be sent to the server | 

### Return type

[**SoftwareSaveLocationResp**](Software_SaveLocationResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

