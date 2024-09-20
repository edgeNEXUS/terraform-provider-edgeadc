# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET21Get**](DateTimeApi.md#GET21Get) | **Get** /GET/21 | Date and time page data
[**POST21iAction1iType1Post**](DateTimeApi.md#POST21iAction1iType1Post) | **Post** /POST/21?iAction&#x3D;1&amp;iType&#x3D;1 | Update current date time
[**POST21iAction1iType2Post**](DateTimeApi.md#POST21iAction1iType2Post) | **Post** /POST/21?iAction&#x3D;1&amp;iType&#x3D;2 | Update server time details

# **GET21Get**
> SetupSetTime GET21Get(ctx, )
Date and time page data

This service is used to FETCH date and time data, it uses three API structure.    Setup_SetTime: Used to get current time.    Setup_TimeServer: Used to get URL of time server to synchronize.    Setup_TimeChecks: Used to get other Synchrinization settings(updated at, period and timezone).    The nodes of JSON received, contains the given information-           NTPTimeZoneCombo:       NTPTypeCombo: This node contains options of NTP type such as SNTP, TCP, UDP.       NTPType: This node contains the id of item that user selected from NTP Type combo.       DateTime: This node contains current date and time.       NtpTimeServer: This node contains NTP time server for time synchronization.       NtpTimeUpdate: This node contains NTP time update at.       NtpPeriodUpdate: This node contains NTP time update period.       NtpTimeZone:       Enabled: This node contains NTP 'true/false' value 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SetupSetTime**](Setup_SetTime.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST21iAction1iType1Post**
> SetupSetTime POST21iAction1iType1Post(ctx, body)
Update current date time

This service is used to update current date and time, you have to send date time as TimeOfDay=date(DD/MM/YYYY) H|M|S, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        NtpTimeZone: id(Received as id in GET/21 service, node NTPTimeZoneCombo) of option selected from Time Zone options.     TimeOfDay: current System Time.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SetupUpdateTime**](SetupUpdateTime.md)| Object that needs to be sent to the server | 

### Return type

[**SetupSetTime**](Setup_SetTime.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST21iAction1iType2Post**
> SetupSetTime POST21iAction1iType2Post(ctx, body)
Update server time details

This service is used to update server time detaiils for date and time sync, you have to send server time details, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        NTPType: id(Received as id in GET/21 service, node NTPTypeCombo) of option selected from NTP Type options.     Timeserver: Timeserver url.     Updated: Update time.      Updateperiod: Update Period.     Enabled: 0 or 1. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SetupUpdateServer**](SetupUpdateServer.md)| Object that needs to be sent to the server | 

### Return type

[**SetupSetTime**](Setup_SetTime.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

