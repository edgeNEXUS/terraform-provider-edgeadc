# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET54Get**](AddOnsApi.md#GET54Get) | **Get** /GET/54 | Container details
[**GET54Post**](AddOnsApi.md#GET54Post) | **Post** /GET/54 | Play
[**GET54downloaddockerConfigID5Get**](AddOnsApi.md#GET54downloaddockerConfigID5Get) | **Get** /GET/54?download&#x3D;docker&amp;ConfigID&#x3D;5 | Export
[**POST26iAction12iType1Post**](AddOnsApi.md#POST26iAction12iType1Post) | **Post** /POST/26?iAction&#x3D;12&amp;iType&#x3D;1 | Update container
[**POST54iAction4iType1Post**](AddOnsApi.md#POST54iAction4iType1Post) | **Post** /POST/54?iAction&#x3D;4&amp;iType&#x3D;1 | Public Ip
[**POST54iAction7iType1Post**](AddOnsApi.md#POST54iAction7iType1Post) | **Post** /POST/54?iAction&#x3D;7&amp;iType&#x3D;1 | Pause
[**POST54iAction7iType2Post**](AddOnsApi.md#POST54iAction7iType2Post) | **Post** /POST/54?iAction&#x3D;7&amp;iType&#x3D;2 | Unpause
[**POST54iAction8iType2Post**](AddOnsApi.md#POST54iAction8iType2Post) | **Post** /POST/54?iAction&#x3D;8&amp;iType&#x3D;2 | Stop
[**POST54iAction8iType3Post**](AddOnsApi.md#POST54iAction8iType3Post) | **Post** /POST/54?iAction&#x3D;8&amp;iType&#x3D;3 | Remove add-on
[**POST54iAction9iType1ConfigID5senddockerPost**](AddOnsApi.md#POST54iAction9iType1ConfigID5senddockerPost) | **Post** /POST/54?iAction&#x3D;9&amp;iType&#x3D;1&amp;ConfigID&#x3D;5&amp;send&#x3D;docker | Import

# **GET54Get**
> AddOnsData GET54Get(ctx, )
Container details

This service is used to FETCH all container details.    The nodes of JSON received, contains the given information-      AddOns: This node contains containers details.     id: This node contains container is.     AppName: This node contains app name.     ParentImage:      DockerImage:      TimeStarted:     TimeStopped:     InternalIP:      AddOnID:      AddOnGUI:      ExternalIP:      ExternalPort:      ExternalAdapter:      sIcon:      isPaused:      AddonImport:      dataset: dataset.     row:row .     isRestore: This node contains either true or false. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AddOnsData**](AddOns_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET54Post**
> AddOnsData GET54Post(ctx, body)
Play

This service is used to play the container.    JSON that needs to be sent as request payload, Should contains given information-        AddOnName:      ConfigID:     ExternalIPAddress:     ExternalPort:     AddOnID:     Repository: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddOnsPlay**](AddOnsPlay.md)| Json data that needs to be sent to the server. | 

### Return type

[**AddOnsData**](AddOns_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET54downloaddockerConfigID5Get**
> GET54downloaddockerConfigID5Get(ctx, )
Export

This service is used to export the file. 

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

# **POST26iAction12iType1Post**
> AddOnsUpdateData POST26iAction12iType1Post(ctx, body)
Update container

This service is used to update the container.    JSON that needs to be sent as request payload, Should contains given information-        Repository:      AddOnName:     ExternalIPAddress:     ExternalPort:     ConfigID: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddOnsUpdatePostData**](AddOnsUpdatePostData.md)| Json object that needs to be added to the store | 

### Return type

[**AddOnsUpdateData**](AddOns_UpdateData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST54iAction4iType1Post**
> AddOnsPublicIpResponse POST54iAction4iType1Post(ctx, body)
Public Ip

This service is used to get Public Ip of the container.    JSON that needs to be sent as request payload, Should contains given information-        AddOnGUI:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddOnsPublicIp**](AddOnsPublicIp.md)| Json data that needs to be sent to the server. | 

### Return type

[**AddOnsPublicIpResponse**](AddOns_PublicIP_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST54iAction7iType1Post**
> AddOnsData POST54iAction7iType1Post(ctx, body)
Pause

This service is used to pause the container.    JSON that needs to be sent as request payload, Should contains given information-        AddOnID:      ConfigID: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddOnsPause**](AddOnsPause.md)| Json data that needs to be sent to the server. | 

### Return type

[**AddOnsData**](AddOns_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST54iAction7iType2Post**
> AddOnsData POST54iAction7iType2Post(ctx, body)
Unpause

This service is used to unpause the container.    JSON that needs to be sent as request payload, Should contains given information-        AddOnName:      ConfigID:     ExternalIPAddress:     ExternalPort:     AddOnID:     Repository: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddOnsPlay**](AddOnsPlay.md)| Json data that needs to be sent to the server. | 

### Return type

[**AddOnsData**](AddOns_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST54iAction8iType2Post**
> AddOnsData POST54iAction8iType2Post(ctx, body)
Stop

This service is used to stop the container.    JSON that needs to be sent as request payload, Should contains given information-        AddOnID:      ConfigID:     Repository: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddOnsStop**](AddOnsStop.md)| Json data that needs to be sent to the server. | 

### Return type

[**AddOnsData**](AddOns_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST54iAction8iType3Post**
> AddOnsData POST54iAction8iType3Post(ctx, body)
Remove add-on

This service is used to remove the container.    JSON that needs to be sent as request payload, Should contains given information-        ConfigID:      AddOnID: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddOnsRemove**](AddOnsRemove.md)| Json object that needs to be added to the store | 

### Return type

[**AddOnsData**](AddOns_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST54iAction9iType1ConfigID5senddockerPost**
> AddOnsData POST54iAction9iType1ConfigID5senddockerPost(ctx, )
Import

This service is used to import the file using form submit. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AddOnsData**](AddOns_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

