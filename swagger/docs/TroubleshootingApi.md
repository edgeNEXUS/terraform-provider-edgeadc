# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET21downloadsupportpid242604Get**](TroubleshootingApi.md#GET21downloadsupportpid242604Get) | **Get** /GET/21?download&#x3D;support&amp;pid&#x3D;242604 | Fire After Support File Request
[**GET23Get**](TroubleshootingApi.md#GET23Get) | **Get** /GET/23 | Capture details
[**GET23duration50packets999999adapteranydownloadcapturepreDownloadCheckyesGet**](TroubleshootingApi.md#GET23duration50packets999999adapteranydownloadcapturepreDownloadCheckyesGet) | **Get** /GET/23?duration&#x3D;50&amp;packets&#x3D;999999&amp;adapter&#x3D;any&amp;download&#x3D;capture&amp;preDownloadCheck&#x3D;yes | Download capture
[**GET41PingTempfilepingresultsxpc62LGet**](TroubleshootingApi.md#GET41PingTempfilepingresultsxpc62LGet) | **Get** /GET/41?PingTempfile&#x3D;pingresultsxpc62L | Ping Result
[**GET58Get**](TroubleshootingApi.md#GET58Get) | **Get** /GET/58 | Trace details
[**GET58downloadtraceGet**](TroubleshootingApi.md#GET58downloadtraceGet) | **Get** /GET/58?download&#x3D;trace | Trace download
[**POST23iAction4iType1Post**](TroubleshootingApi.md#POST23iAction4iType1Post) | **Post** /POST/23?iAction&#x3D;4&amp;iType&#x3D;1 | Start capture download process
[**POST24iAction5iType4Post**](TroubleshootingApi.md#POST24iAction5iType4Post) | **Post** /POST/24?iAction&#x3D;5&amp;iType&#x3D;4 | Ping
[**POST53iAction6iType1Post**](TroubleshootingApi.md#POST53iAction6iType1Post) | **Post** /POST/53?iAction&#x3D;6&amp;iType&#x3D;1 | Download support file
[**POST58iAction1iType1Post**](TroubleshootingApi.md#POST58iAction1iType1Post) | **Post** /POST/58?iAction&#x3D;1&amp;iType&#x3D;1 | Start/Stop Trace
[**POST58iAction1iType2Post**](TroubleshootingApi.md#POST58iAction1iType2Post) | **Post** /POST/58?iAction&#x3D;1&amp;iType&#x3D;2 | Clear Trace

# **GET21downloadsupportpid242604Get**
> ResultsAfterSupportFile GET21downloadsupportpid242604Get(ctx, )
Fire After Support File Request

This service is used to FETCH  Results after support file request.    The nodes of JSON received, contains the given information.        success:      DownloadStatus:  

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResultsAfterSupportFile**](resultsAfterSupport_File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET23Get**
> TroubleshootingData GET23Get(ctx, )
Capture details

This service is used to FETCH Capture details, it uses two API structure.    Services_Capture: Used to get Capture information like adapter used and number of packets etc.    Setup_Network: Used to get all available Adapters list string.      The nodes of JSON received, contains the given information        CaptureAdapterCombo: This node contains options of adapter available for capture.     CaptureAdapter: This node contains adapter selected by user for capture.     CapturePackets: This node contains number of packets to capture.     CaptureTime: This node contains capture time.     CaptureLeft: This node contains capture left to download.     SupportFileCombo: This node contains options of SupportFile.     CaptureAddress: This node contains CaptureAddress. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TroubleshootingData**](Troubleshooting_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET23duration50packets999999adapteranydownloadcapturepreDownloadCheckyesGet**
> GET23duration50packets999999adapteranydownloadcapturepreDownloadCheckyesGet(ctx, )
Download capture

This service is used to download capture file. you have to send Adapter,packets, duration and download=capture, in query string. This will work only if capture processing already started by above POST request.   Query string: for example duration=5&packets=99978&adapter=eth0&download=capture 

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

# **GET41PingTempfilepingresultsxpc62LGet**
> PingResults GET41PingTempfilepingresultsxpc62LGet(ctx, )
Ping Result

This service is used to FETCH Ping Result after ping request.    The nodes of JSON received, contains the given information.        success:      StatusImage:      StatusText:     PingResult:      PingTempfile:  

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PingResults**](ping_results.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET58Get**
> TroubleTraceDetail GET58Get(ctx, )
Trace details

This service is used to FETCH Trace details.    The nodes of JSON received, contains the given information.        AutoStopTime:      Nodes:      AutoStopRecords:     TracePath:      TraceConnections:      TraceCache:      TraceData:      flightpathCombo:     ServerMonitor:     MonitoringFault:     TraceStatus:     Purpose:     List: 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TroubleTraceDetail**](Trouble_TraceDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET58downloadtraceGet**
> GET58downloadtraceGet(ctx, )
Trace download

This service is used to download trace file. 

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

# **POST23iAction4iType1Post**
> []TroubleshootingData POST23iAction4iType1Post(ctx, body)
Start capture download process

This service is used to start capture download processing. you have to send Adapter, packets and duration, with post request JSON.    JSON that needs to be sent as request payload, Should contains given information-        CaptureHidden: capture flag.      Adapter : id(Received as id in GET/23 service, node CaptureAdapterCombo) of option selected from Adapter options.      Packets : number of packets needs to be captured.      Duration : Duration of capture.      CaptureAddress:CaptureAddress. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TroubleDownCapture**](TroubleDownCapture.md)| Object that needs to be sent to the server | 

### Return type

[**[]TroubleshootingData**](Troubleshooting_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST24iAction5iType4Post**
> TroublePingResp POST24iAction5iType4Post(ctx, body)
Ping

This service is used to ping any IP address. you have to send IP address, with post request JSON.    JSON that needs to be sent as request payload, Should contains given information-        ipAddress : IP address for Ping. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TroublePing**](TroublePing.md)| Object that needs to be sent to the server | 

### Return type

[**TroublePingResp**](Trouble_PingResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST53iAction6iType1Post**
> TroubleSupportResp POST53iAction6iType1Post(ctx, body)
Download support file

This service is used to download support file. you have to send days, with post request JSON.    JSON that needs to be sent as request payload, Should contains given information-        Days : days. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TroubleDownSupport**](TroubleDownSupport.md)| Object that needs to be sent to the server | 

### Return type

[**TroubleSupportResp**](Trouble_SupportResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST58iAction1iType1Post**
> TroubleTraceDetail POST58iAction1iType1Post(ctx, body)
Start/Stop Trace

This service is used to start/stop trace .    JSON that needs to be sent as request payload, Should contains given information-        AutoStopRecords:      AutoStopTime :      Nodes :      ServerMonitor :      TraceData:      Start:      TraceCache:      TraceConnections:      MonitoringFault:      Purpose :      TracePath: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TroubleTrace**](TroubleTrace.md)| Object that needs to be sent to the server | 

### Return type

[**TroubleTraceDetail**](Trouble_TraceDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST58iAction1iType2Post**
> []TroubleTraceDetail POST58iAction1iType2Post(ctx, body)
Clear Trace

This service is used to trace. you have to send Purpose with post request JSON.    JSON that needs to be sent as request payload, Should contains given information-        Purpose: Purpose text.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TroubleTraceClear**](TroubleTraceClear.md)| Object that needs to be sent to the server | 

### Return type

[**[]TroubleTraceDetail**](Trouble_TraceDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

