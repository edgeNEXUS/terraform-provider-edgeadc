# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET24pagePowerGet**](PowerApi.md#GET24pagePowerGet) | **Get** /GET/24?page&#x3D;Power | Software version
[**GET28Get**](PowerApi.md#GET28Get) | **Get** /GET/28 | Software version
[**POST24iAction5iType1Post**](PowerApi.md#POST24iAction5iType1Post) | **Post** /POST/24?iAction&#x3D;5&amp;iType&#x3D;1 | Restart ALB
[**POST24iAction5iType2Post**](PowerApi.md#POST24iAction5iType2Post) | **Post** /POST/24?iAction&#x3D;5&amp;iType&#x3D;2 | Reboot ALB
[**POST24iAction5iType3Post**](PowerApi.md#POST24iAction5iType3Post) | **Post** /POST/24?iAction&#x3D;5&amp;iType&#x3D;3 | Power OFF ALB

# **GET24pagePowerGet**
> PowerVersion GET24pagePowerGet(ctx, )
Software version

This service is used to FETCH software version details, it uses one API structure.    Do_Restart: Used to get software version details. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PowerVersion**](Power_Version.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET28Get**
> PowerCheckStatus GET28Get(ctx, )
Software version

This service is used to FETCH software version details, it uses one API structure. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PowerCheckStatus**](Power_CheckStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST24iAction5iType1Post**
> []PowerRestart POST24iAction5iType1Post(ctx, )
Restart ALB

This service is used to quickly stop and start essential jetNEXUS ALB services.    JSON that needs to be sent as request payload, Should contains given information- not required to send json.        Warning - This will cause a brief break in current connections. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PowerRestart**](Power_Restart.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST24iAction5iType2Post**
> PowerReboot POST24iAction5iType2Post(ctx, )
Reboot ALB

This service is used to re-initialise all jetNEXUS ALB services.    JSON that needs to be sent as request payload, Should contains given information- not required to send json.        Warning - This will suspend your Connections and Services for about 2 minutes. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PowerReboot**](Power_Reboot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST24iAction5iType3Post**
> PowerPowerOff POST24iAction5iType3Post(ctx, )
Power OFF ALB

This service is used to completely halt jetNEXUS ALB.    JSON that needs to be sent as request payload, Should contains given information- not required to send json.        Warning - This will suspend your web site and require a hardware power on. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PowerPowerOff**](Power_PowerOff.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

