# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET2Get**](DashboardApi.md#GET2Get) | **Get** /GET/2 | Show Event Widget details
[**GET48Get**](DashboardApi.md#GET48Get) | **Get** /GET/48 | Show Event Widget details
[**GET51isRefreshRequesttrueGet**](DashboardApi.md#GET51isRefreshRequesttrueGet) | **Get** /GET/51?isRefreshRequest&#x3D;true | Widget details
[**GET62Get**](DashboardApi.md#GET62Get) | **Get** /GET/62 | Show Event Widget details
[**POST51iAction1iType1Post**](DashboardApi.md#POST51iAction1iType1Post) | **Post** /POST/51?iAction&#x3D;1&amp;iType&#x3D;1 | Show Event Widget
[**POST51iAction1iType2Post**](DashboardApi.md#POST51iAction1iType2Post) | **Post** /POST/51?iAction&#x3D;1&amp;iType&#x3D;2 | Show IPService Widget
[**POST51iAction1iType3Post**](DashboardApi.md#POST51iAction1iType3Post) | **Post** /POST/51?iAction&#x3D;1&amp;iType&#x3D;3 | Show System Widget
[**POST51iAction1iType4Post**](DashboardApi.md#POST51iAction1iType4Post) | **Post** /POST/51?iAction&#x3D;1&amp;iType&#x3D;4 | Show IPStatus Widget
[**POST51iAction1iType5Post**](DashboardApi.md#POST51iAction1iType5Post) | **Post** /POST/51?iAction&#x3D;1&amp;iType&#x3D;5 | Remove Widget
[**POST51iAction1iType6Post**](DashboardApi.md#POST51iAction1iType6Post) | **Post** /POST/51?iAction&#x3D;1&amp;iType&#x3D;6 | Show Interface Widget
[**POST51iAction2iType1Post**](DashboardApi.md#POST51iAction2iType1Post) | **Post** /POST/51?iAction&#x3D;2&amp;iType&#x3D;1 | Default Event Widget
[**POST51iAction2iType2Post**](DashboardApi.md#POST51iAction2iType2Post) | **Post** /POST/51?iAction&#x3D;2&amp;iType&#x3D;2 | Default IPService Widget
[**POST51iAction2iType3Post**](DashboardApi.md#POST51iAction2iType3Post) | **Post** /POST/51?iAction&#x3D;2&amp;iType&#x3D;3 | Default System Widget
[**POST51iAction2iType4Post**](DashboardApi.md#POST51iAction2iType4Post) | **Post** /POST/51?iAction&#x3D;2&amp;iType&#x3D;4 | Default IPStatus Widget
[**POST51iAction3iType1Post**](DashboardApi.md#POST51iAction3iType1Post) | **Post** /POST/51?iAction&#x3D;3&amp;iType&#x3D;1 | Default Widget

# **GET2Get**
> DashEventWidget2 GET2Get(ctx, )
Show Event Widget details

This service is used for show event widget.      Headlines: This node contains Headlines.     CPU:CPU.     GraphMemory:GraphMemory.     PageTitle:PageTitle.     nature: nature.     colour:colour.     id:id.     message:message.      data:data. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DashEventWidget2**](DashEventWidget2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET48Get**
> EditIpServiceRespWidget1 GET48Get(ctx, )
Show Event Widget details

This service is used for show event widget.      period: period     systemData: This node contains systemData.     responseError: This node contains responseError.      timesServer: This node contains the value of timesServer.      axesCount: This node contains axesCount.     VSRS: VSRS.     RS: RS.     ipPort: ipPort.     responseWarning:responseWarning.     periodEnd: periodEnd.     type: reqsec.     periodEndNow: periodEndNow.     columns:columns.     timesUtc: timesUtc. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EditIpServiceRespWidget1**](EditIPServiceResp_widget1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET51isRefreshRequesttrueGet**
> GetDashboard GET51isRefreshRequesttrueGet(ctx, )
Widget details

This service is used to FETCH widget details.    The nodes of JSON received, contains the given information      ConfiguredWidgetsComboStore:     data: This node contains avaliable widgets.     PartID: This node contains partid of widget.      SectionName: This node contains section of widget.      Widget_Name: This node contains name of widget. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetDashboard**](GetDashboard.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET62Get**
> StatusDataRow GET62Get(ctx, )
Show Event Widget details

This service is used for show event widget.      interfaceImagePath: interfaceImagePath     data: This node contains avaliable widgets.     conn: This node contains conn of widget.      compress: This node contains the value of compress.      cacheperc: This node contains cacheperc.     rsStatusReason: rsStatusReason.     channelStatusReason: channelStatusReason.     pool: pool.     RowType:RowType.     VIPPort: VIPPort.     serviceName:serviceName.     realServerNotes:realServerNotes.     realServer: realServer.     rsImagePath:rsImagePath.     hitsecnew:hitsecnew.     channelImagePath: channelImagePath.     interfaceStatusReason: interfaceStatusReason.     reqsec: reqsec.     SeriesDataConn: SeriesDataConn.     SeriesDataDat:SeriesDataDat.     SeriesDataReq: SeriesDataReq. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StatusDataRow**](StatusData_row.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction1iType1Post**
> DashEventWidgetResp POST51iAction1iType1Post(ctx, body)
Show Event Widget

This service is used for show event widget.     JSON that needs to be sent as request payload, Should contains given information-        Section_Name: widget section name.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DashEventWidget**](DashEventWidget.md)| Object that needs to be sent to the server | 

### Return type

[**DashEventWidgetResp**](DashEventWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction1iType2Post**
> DashIpServiceWidgetResp POST51iAction1iType2Post(ctx, body)
Show IPService Widget

This service is used for show IPService widget.   JSON that needs to be sent as request payload, Should contains given information-        Section_Name: widget section name.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DashIpServiceWidget**](DashIpServiceWidget.md)| Object that needs to be sent to the server | 

### Return type

[**DashIpServiceWidgetResp**](DashIPServiceWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction1iType3Post**
> DashSystemWidgetResp POST51iAction1iType3Post(ctx, body)
Show System Widget

This service is used for show system widget.     JSON that needs to be sent as request payload, Should contains given information-        Section_Name: widget section name.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DashSystemWidget**](DashSystemWidget.md)| Object that needs to be sent to the server | 

### Return type

[**DashSystemWidgetResp**](DashSystemWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction1iType4Post**
> DashIpStatusWidgetResp POST51iAction1iType4Post(ctx, body)
Show IPStatus Widget

This service is used for show ipstatus widget.   JSON that needs to be sent as request payload, Should contains given information-        Section_Name: widget section name.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DashIpStatusWidget**](DashIpStatusWidget.md)| Object that needs to be sent to the server | 

### Return type

[**DashIpStatusWidgetResp**](DashIPStatusWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction1iType5Post**
> DashRemoveWidgetResp POST51iAction1iType5Post(ctx, body)
Remove Widget

This service is used for remove widget.    JSON that needs to be sent as request payload, Should contains given information-        Section_Name: widget section name.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DashInterfaceWidget**](DashInterfaceWidget.md)| Object that needs to be sent to the server | 

### Return type

[**DashRemoveWidgetResp**](DashRemoveWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction1iType6Post**
> DashInterfaceWidgetResp POST51iAction1iType6Post(ctx, body)
Show Interface Widget

This service is used for show interface widget.    JSON that needs to be sent as request payload, Should contains given information-        Section_Name: widget section name.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DashInterfaceWidget**](DashInterfaceWidget.md)| Object that needs to be sent to the server | 

### Return type

[**DashInterfaceWidgetResp**](DashInterfaceWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction2iType1Post**
> DefaultEventWidgetResp POST51iAction2iType1Post(ctx, body)
Default Event Widget

This service is used for show default Event widget.    JSON that needs to be sent as request payload, Should contains given information-        Section_Name: widget section name.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DefaultEventWidget**](DefaultEventWidget.md)| Object that needs to be sent to the server | 

### Return type

[**DefaultEventWidgetResp**](DefaultEventWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction2iType2Post**
> DefaultIpServiceWidgetResp POST51iAction2iType2Post(ctx, body)
Default IPService Widget

This service is used for show default IPService widget.    JSON that needs to be sent as request payload, Should contains given information-        Section_Name: widget section name.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DefaultIpServiceWidget**](DefaultIpServiceWidget.md)| Object that needs to be sent to the server | 

### Return type

[**DefaultIpServiceWidgetResp**](DefaultIPServiceWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction2iType3Post**
> DefaultSystemWidgetResp POST51iAction2iType3Post(ctx, body)
Default System Widget

This service is used for show default system widget.    JSON that needs to be sent as request payload, Should contains given information-        Section_Name: widget section name.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DefaultSystemWidget**](DefaultSystemWidget.md)| Object that needs to be sent to the server | 

### Return type

[**DefaultSystemWidgetResp**](DefaultSystemWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction2iType4Post**
> DefaultIpStatusWidgetResp POST51iAction2iType4Post(ctx, body)
Default IPStatus Widget

This service is used for show default IPStatus widget.    JSON that needs to be sent as request payload, Should contains given information-        Section_Name: widget section name.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DefaultIpStatusWidget**](DefaultIpStatusWidget.md)| Object that needs to be sent to the server | 

### Return type

[**DefaultIpStatusWidgetResp**](DefaultIPStatusWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction3iType1Post**
> DefaultWidgetResp POST51iAction3iType1Post(ctx, body)
Default Widget

This service is used for show default dashboard.    JSON that needs to be sent as request payload, Should contains given information-        Test:   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DefaultWidget**](DefaultWidget.md)| Object that needs to be sent to the server | 

### Return type

[**DefaultWidgetResp**](DefaultWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

