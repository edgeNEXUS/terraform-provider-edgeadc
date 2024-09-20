# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET39pageAppstoreGet**](AppStoreApi.md#GET39pageAppstoreGet) | **Get** /GET/39?page&#x3D;Appstore | Licence Details
[**GET63Get**](AppStoreApi.md#GET63Get) | **Get** /GET/63 | Licence Details

# **GET39pageAppstoreGet**
> AppData GET39pageAppstoreGet(ctx, )
Licence Details

This service is used to FETCH licence details.    The nodes of JSON received, contains the given information        UserID: This node contains user id.     MachineID: This node contains machine id.      Location: This node contains location of browser.     PublicIP: This node contains public ip.      LicenceID: This node contains licence id.     SoftwareVersion: This node contains software version.     AppStoreUrl: This node contains app store url. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AppData**](App_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET63Get**
> AppData69Get GET63Get(ctx, )
Licence Details

This service is used to FETCH licence details.    The nodes of JSON received, contains the given information        CsmKey: This node contains Csm Key.     TimeofDay: This node contains datetime stamp.  

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AppData69Get**](App_Data_69_Get.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

