# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET47Get**](GlobalSettingsApi.md#GET47Get) | **Get** /GET/47 | Global Settings page data
[**POST47iAction1iType1Post**](GlobalSettingsApi.md#POST47iAction1iType1Post) | **Post** /POST/47?iAction&#x3D;1&amp;iType&#x3D;1 | Update SSL Cryptographic Library
[**POST47iAction1iType2Post**](GlobalSettingsApi.md#POST47iAction1iType2Post) | **Post** /POST/47?iAction&#x3D;1&amp;iType&#x3D;2 | Update Host Cache Timer
[**POST47iAction1iType3Post**](GlobalSettingsApi.md#POST47iAction1iType3Post) | **Post** /POST/47?iAction&#x3D;1&amp;iType&#x3D;3 | Update drain
[**POST47iAction1iType4Post**](GlobalSettingsApi.md#POST47iAction1iType4Post) | **Post** /POST/47?iAction&#x3D;1&amp;iType&#x3D;4 | Update authentication-setting

# **GET47Get**
> GlobSettingData GET47Get(ctx, )
Global Settings page data

This service is used to FETCH Global settings data like crypto library used, it uses one API structure.    AdvGlobalSetting: Used to get global settings.       The nodes of JSON received, contains the given information        GlobalSettingListString: This node contains options of cryptographic library can be used like SSL and YASSL.     CryptoLib: This node contains the value of current crypto lib used.     HostCacheTimer: This node contains the value of host cache timer(in seconds).     LdapTimeout: This node contains information of Authentication server time out.     DrainClearPersistence: This node contains the value of DrainClearPersistence. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GlobSettingData**](GlobSetting_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST47iAction1iType1Post**
> GlobSettingData POST47iAction1iType1Post(ctx, body)
Update SSL Cryptographic Library

This service is used to update cryptographic library used.    JSON that needs to be sent as request payload, Should contains given information-        CryptoLib: id(Received as id in GET/47 service, node GlobalSettingListString) of option selected from crypto lib options.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobUpdateCrypto**](GlobUpdateCrypto.md)| Object that needs to be sent to the server | 

### Return type

[**GlobSettingData**](GlobSetting_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST47iAction1iType2Post**
> GlobSettingData POST47iAction1iType2Post(ctx, body)
Update Host Cache Timer

This service is used to update Host Cache Timer.    JSON that needs to be sent as request payload, Should contains given information-        HostCacheTimer: This node contains the value of host cache timer(in seconds).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobUpdateHChacheTimer**](GlobUpdateHChacheTimer.md)| Object that needs to be sent to the server | 

### Return type

[**GlobSettingData**](GlobSetting_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST47iAction1iType3Post**
> GlobSettingData POST47iAction1iType3Post(ctx, body)
Update drain

This service is used to update drain.       JSON that needs to be sent as request paylad,  Should contains given information-        DrainClearPersistence: This node contains either 0 or 1.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobUpdateDrain**](GlobUpdateDrain.md)| Object that needs to be sent to the server | 

### Return type

[**GlobSettingData**](GlobSetting_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST47iAction1iType4Post**
> GlobSettingData POST47iAction1iType4Post(ctx, body)
Update authentication-setting

This service is used to update authentication setting.        JSON that needs to be sent as request payload,  Should contains given information--        LdapTimeout: This node contains the value of authentication server timeout.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobUpdateAuthenticationSetting**](GlobUpdateAuthenticationSetting.md)| Object that needs to be sent to the server | 

### Return type

[**GlobSettingData**](GlobSetting_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

