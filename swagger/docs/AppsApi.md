# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET39Get**](AppsApi.md#GET39Get) | **Get** /GET/39 | Licence Details
[**GET49smartInfoappsGet**](AppsApi.md#GET49smartInfoappsGet) | **Get** /GET/49?smartInfo&#x3D;apps | Avaliable download
[**POST26iAction10iType1Post**](AppsApi.md#POST26iAction10iType1Post) | **Post** /POST/26?iAction&#x3D;10&amp;iType&#x3D;1 | Deploy
[**POST26iAction6iType1Post**](AppsApi.md#POST26iAction6iType1Post) | **Post** /POST/26?iAction&#x3D;6&amp;iType&#x3D;1 | Remove
[**POST26iAction8iType1Post**](AppsApi.md#POST26iAction8iType1Post) | **Post** /POST/26?iAction&#x3D;8&amp;iType&#x3D;1 | Download
[**POST26iAction9iType1Post**](AppsApi.md#POST26iAction9iType1Post) | **Post** /POST/26?iAction&#x3D;9&amp;iType&#x3D;1 | Download Status

# **GET39Get**
> AppStoreData GET39Get(ctx, )
Licence Details

This service is used to FETCH licence details.    The nodes of JSON received, contains the given information        UserID: This node contains user id.     MachineID: This node contains machine id.     Location: This node contains location.     PublicIP: This node contains public ip.     LicenceID: This node contains licence id.     SoftwareVersion: This node contains software version.     AppStoreUrl:  

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AppStoreData**](AppStore_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET49smartInfoappsGet**
> AppsAvailable GET49smartInfoappsGet(ctx, )
Avaliable download

This service is used to fetch smart file details of avaliable softwares for download.    The nodes of JSON received, contains the given information        SmartFiles:      File_Name: This node contains File Name.      Version: This node contains Version.     Message_Digest:      Machine_Id: This node contains Machine id.      Product_Id: This node contains Product id.     Order_Id: This node contains order id.      Icon: This node contains icon.     Description:     Name:     Type:     Certified:     Vendor:     ReleaseDate:     DockerImageName: 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AppsAvailable**](Apps_Available.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST26iAction10iType1Post**
> AppsAvailable POST26iAction10iType1Post(ctx, body)
Deploy

This service is used to deploy the software.    JSON that needs to be sent as request payload, Should contains given information-        Name: file name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AppsDeploy**](AppsDeploy.md)| Object that needs to be sent to the store. | 

### Return type

[**AppsAvailable**](Apps_Available.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST26iAction6iType1Post**
> AppsAvailable POST26iAction6iType1Post(ctx, body)
Remove

This service is used to remove the downloaded sofware.    JSON that needs to be sent as request payload, Should contains given information-        FileName: file name.     DockerImageName: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AppsRemove**](AppsRemove.md)| Object that needs to be sent to the store. | 

### Return type

[**AppsAvailable**](Apps_Available.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST26iAction8iType1Post**
> AppsDownResponse POST26iAction8iType1Post(ctx, body)
Download

This service is used to download the software.    JSON that needs to be sent as request payload, Should contains given information-        URL: file url.      Name: file name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AppsDownload**](AppsDownload.md)| Object that needs to be sent to the store. | 

### Return type

[**AppsDownResponse**](Apps_DownResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST26iAction9iType1Post**
> AppsDownResponse POST26iAction9iType1Post(ctx, body)
Download Status

This service is used to fetch download status of downloading software.    JSON that needs to be sent as request payload, Should contains given information-        URL: file url.      Name: file name.     TotalSize: file size. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AppsDownloadOne**](AppsDownloadOne.md)| Object that needs to be sent to the store. | 

### Return type

[**AppsDownResponse**](Apps_DownResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

