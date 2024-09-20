# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET16Get**](LoggingApi.md#GET16Get) | **Get** /GET/16 | Logging page data
[**POST16iAction1iType1Post**](LoggingApi.md#POST16iAction1iType1Post) | **Post** /POST/16?iAction&#x3D;1&amp;iType&#x3D;1 | Update logging levels
[**POST16iAction2iType1Post**](LoggingApi.md#POST16iAction2iType1Post) | **Post** /POST/16?iAction&#x3D;2&amp;iType&#x3D;1 | Update remote log storage
[**POST16iAction3iType1Post**](LoggingApi.md#POST16iAction3iType1Post) | **Post** /POST/16?iAction&#x3D;3&amp;iType&#x3D;1 | Update remote syslog server
[**POST16iAction4iType1Post**](LoggingApi.md#POST16iAction4iType1Post) | **Post** /POST/16?iAction&#x3D;4&amp;iType&#x3D;1 | Clear log
[**POST16iAction5iType1Post**](LoggingApi.md#POST16iAction5iType1Post) | **Post** /POST/16?iAction&#x3D;5&amp;iType&#x3D;1 | Update syslog

# **GET16Get**
> LoggingData GET16Get(ctx, )
Logging page data

This service is used to FETCH Logging details, it uses four API structure.    Config_Logging: Used to get Logging level details.    Config_StatsInHeader: Used to get list of JETNEXUS HEADERS.    Config_RemoteLogStorage: Used to get Remote log storage details.    Config_RemoteSysLogServer: Used to get Remote Syslog Server details      The nodes of JSON received, contains the given information          W3CCombo: This node contains options of w3c logging.     W3CFCombo: This node contains options of jetnexus w3c logging.     RemoteLogStorage: This node have 1 if remote log storage enabled.     RemoteLogIP: This node contains information, to whom this licence is issued.     RemoteLogShare: This node contains information of contact person.     RemoteLogDir: This node contains date when licence was issued.     RemoteLogUser: This node contains ALB name.     W3cLoggingModeList: This node contains information, to whom this licence is issued.     LogClearListString: This node contains information of contact person.     SecurityLogCombo:     SecurityLog:     XffLoggingList:     RsysLog1IP:     RsysLog2IP:     RsysLog1PROTOCOL:     RsysLog2PROTOCOL:     RsysServerPort1:     RsysServerPort2:     RsysServer1enabled:     RsysServer2enabled:     RsysProtocolCombo:     AdvancedW3C: This node contains either 0 or 1.     LogMessageLevel: Id of option selected from message level.     LogMessageLevelCombo: 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LoggingData**](Logging_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST16iAction1iType1Post**
> LoggingData POST16iAction1iType1Post(ctx, body)
Update logging levels

This service is used to update Logging level details, you have to send logging level data, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        W3cLoggingModeList: id(Received as id in GET/16 service, node W3CCombo) of option selected from W3C Logging options.     XffLoggingList: id(Received as id in GET/16 service, node W3CFCombo) of option selected from jetnexus w3c Logging options.     securityLog: on/off value.     AdvancedW3C:. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateLogging**](UpdateLogging.md)| Object that needs to be sent to the server | 

### Return type

[**LoggingData**](Logging_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST16iAction2iType1Post**
> LoggingData POST16iAction2iType1Post(ctx, body)
Update remote log storage

This service is used to update remote log storage details, you have to send remote log settings, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        RemoteLogStorage: 'true/false'.     RemoteLogIP: IP address.     RemoteLogShare: share name.      RemoteLogDir: Remote Log Directory.     RemoteLogUser: username.     RemoteLogPassword: Password. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateRemotelog**](UpdateRemotelog.md)| Object that needs to be sent to the server | 

### Return type

[**LoggingData**](Logging_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST16iAction3iType1Post**
> LoggingData POST16iAction3iType1Post(ctx, body)
Update remote syslog server

This service is used to update remote syslog server details, you have to send remote syslog server settings(IP,PORT etc), with POST request json.    JSON that needs to be sent as request payload, Should contains given information-           RsysLog1IP: Syslog server '1' IP address.     RsysLog1PROTOCOL: id(Received as id in GET/16 service, node W3RsysProtocolCombo) of option selected from Protocol options.     RsysLog2IP: Syslog server '2' IP address.     RsysLog2PROTOCOL: id(Received as id in GET/16 service, node W3RsysProtocolCombo) of option selected from Protocol options.     RsysServer1enabled: 'true/false'.     RsysServer2enabled: 'true/false'.     RsysServerPort1: Syslog server '1' Port.     RsysServerPort2: Syslog server '2' Port. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateRemotesyslog**](UpdateRemotesyslog.md)| Object that needs to be sent to the server | 

### Return type

[**LoggingData**](Logging_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST16iAction4iType1Post**
> LoggingData POST16iAction4iType1Post(ctx, body)
Clear log

This service is used to Clear log, you have to send log files name.    JSON that needs to be sent as request payload, Should contains given information-        LogClear: id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClearLog**](ClearLog.md)| Object that needs to be sent to the server | 

### Return type

[**LoggingData**](Logging_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST16iAction5iType1Post**
> LoggingData POST16iAction5iType1Post(ctx, body)
Update syslog

This service is used to update syslog.    ON that needs to be sent as request paylad,  Should contains given information-        LogMessageLevel: This node contains id of LogMessageLevel.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LoggingDataSyslog**](LoggingDataSyslog.md)| Object that needs to be sent to the server | 

### Return type

[**LoggingData**](Logging_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

