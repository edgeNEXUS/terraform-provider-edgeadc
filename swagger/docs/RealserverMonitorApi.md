# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET13Get**](RealserverMonitorApi.md#GET13Get) | **Get** /GET/13 | Real Server Monitor page data
[**GET38Get**](RealserverMonitorApi.md#GET38Get) | **Get** /GET/38 | Custom monitor combo
[**POST13iAction1iType1Post**](RealserverMonitorApi.md#POST13iAction1iType1Post) | **Post** /POST/13?iAction&#x3D;1&amp;iType&#x3D;1 | Add Real Server Monitor
[**POST13iAction2iType1Post**](RealserverMonitorApi.md#POST13iAction2iType1Post) | **Post** /POST/13?iAction&#x3D;2&amp;iType&#x3D;1 | Update Real Server Monitor
[**POST13iAction3iType1Post**](RealserverMonitorApi.md#POST13iAction3iType1Post) | **Post** /POST/13?iAction&#x3D;3&amp;iType&#x3D;1 | Remove Real Server Monitor
[**POST13iAction4Post**](RealserverMonitorApi.md#POST13iAction4Post) | **Post** /POST/13?iAction&#x3D;4 | Remove Custom Monitor
[**POST13iAction5sendcsmPost**](RealserverMonitorApi.md#POST13iAction5sendcsmPost) | **Post** /POST/13?iAction&#x3D;5&amp;send&#x3D;csm | Upload Custom Monitor

# **GET13Get**
> ConfigMonitoring GET13Get(ctx, )
Real Server Monitor page data

This service is used to FETCH all data for Real Server Monitor page, it uses one API structures.    Config_Monitoring: Used to get all real server monitor details.    The nodes of JSON received, contains the given information-      MethodCombo: This node contains options of monitoring methods.     ConfigMonitoringGrid: This node contains monitoring grid data.     id: This node contains index of each real server monitor, that can be use to update or delete that monitor.     Name: This node contains name of real server monitor.     Description: This node contains description of real server monitor.     type: This node contains monitoring method used for this real server monitor.     url: This node contains page location url.     content: This node contains required content.     Username: This node contains username.     Password: This node contains password.     Threshold:      AppliedtoVS:AppliedtoVS.     dataset:dataset.     row:row. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConfigMonitoring**](Config_Monitoring.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET38Get**
> DeleteCustomMonitor GET38Get(ctx, )
Custom monitor combo

This service is used to FETCH all Custom Monitors to present them to user to delete from 'Delete Custom Monitor' dropdown, it uses one API structure-    DeleteCustomMonitor: Used to fetch custom monitors.    The nodes of JSON received, contains the given information-      delCSMCombo: This node contains options of custom monitors. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DeleteCustomMonitor**](DeleteCustomMonitor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST13iAction1iType1Post**
> ConfigMonitoring POST13iAction1iType1Post(ctx, )
Add Real Server Monitor

This service is used to add new real server monitor in real server monitor grid.    JSON that needs to be sent as request payload, Should contains given information-not required to send json. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConfigMonitoring**](Config_Monitoring.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST13iAction2iType1Post**
> ConfigMonitoring POST13iAction2iType1Post(ctx, body)
Update Real Server Monitor

This service is used to update details in real server monitor grid, you have to send id of realserver monitor with data, that needs to be updated, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-          Password : Password     Threshold :      content : Required content     description :      id : index(Received as id in GET/13 service) of real server monitor that needs to updated.     Name : Real server name.     Username : User Name.     type : id(Received as id in GET/13 service, node MethodCombo) of option selected from Monitoring method options.     Url :  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]UpdateRealRequestInner**](UpdateRealRequest_inner.md)| Object that needs to be sent to the server | 

### Return type

[**ConfigMonitoring**](Config_Monitoring.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST13iAction3iType1Post**
> ConfigMonitoring POST13iAction3iType1Post(ctx, body)
Remove Real Server Monitor

This service is used to remove row in real server monitor grid, you have to send id of real server monitor, that needs to be removed, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-            id: index(Received as id in GET/13 service) of real server monitor that needs to be deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveRealRequest**](RemoveRealRequest.md)| Object that needs to be sent to the server | 

### Return type

[**ConfigMonitoring**](Config_Monitoring.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST13iAction4Post**
> ConfigMonitoring POST13iAction4Post(ctx, body)
Remove Custom Monitor

This service is used to remove a custom monitor, you have to send id of custom monitor that needs to be removed, with POST request json.  JSON that needs to be sent as request payload, Should contains given information-      CustomMonitor: id(Received as id in GET/38 service) of custom monitor that needs to be deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveCustomMonitor**](RemoveCustomMonitor.md)| Object that needs to be sent to the server | 

### Return type

[**ConfigMonitoring**](Config_Monitoring.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST13iAction5sendcsmPost**
> ConfigMonitoring POST13iAction5sendcsmPost(ctx, )
Upload Custom Monitor

This service is used to upload custom monitor file, you have to send name in form field(UploadMonitorName) and file, that needs to be uploaded, with form submit.    We have to send given values to form submit .        UploadMonitorName: upload file name.     file: File needs to be Uploaded. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConfigMonitoring**](Config_Monitoring.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

