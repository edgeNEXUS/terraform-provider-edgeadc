# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET31isPageLoadtrueGet**](StatusApi.md#GET31isPageLoadtrueGet) | **Get** /GET/31?isPageLoad&#x3D;true | Status page data
[**GET36Get**](StatusApi.md#GET36Get) | **Get** /GET/36 | Status Layout data
[**POST36iAction1Post**](StatusApi.md#POST36iAction1Post) | **Post** /POST/36?iAction&#x3D;1 | Save status layout

# **GET31isPageLoadtrueGet**
> SetupIpServices GET31isPageLoadtrueGet(ctx, )
Status page data

This service is used to FETCH Status details for all Virtual Services (hit/second, request/second, cache % etc), it uses one API structure.    Setup_IPServices: Used to get Status of all Virtual Services and their Real Servers.    The nodes of JSON received, contains the given information-      dataset: This node contains collection of all IP VIP, VS and their real servers.     RowType: This node contains row information for coloring(such as total, subtotal etc).     interfaceImagePath: This node contains image path if it's a VIP.     interfaceStatusReason: This node contains status reason that shown on tooltip of image.     channelImagePath: This node contains image path if it's a VS.     channelStatusReason: This node contains status reason that shown on tooltip of image.     serviceName: This node contains name of VS.     VIPPort: This node contains value of PORT of VS.     hitsecnew: This node contains hits per second.     cacheperc: This node contains cache percentage.     compress: This node contains value off compression.     rsImagePath: This node contains image path of real server.     rsStatusReason: This node contains status reason that shown on tooltip of image..     realServer: This node contains ip address of real server.     realServerNotes: This node contains real server notes.     conn: This node contains value of connections.     data: This node contains value of data.     reqsec: This node contains value of request per second.     pool: This node contains value of pool size.     ipService:ipService.     sId:sId.     FilterKeyword: FilterKeyword. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SetupIpServices**](Setup_IPServices.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET36Get**
> SetupIpServicesLayout GET36Get(ctx, )
Status Layout data

This service is used to FETCH Status layout details for all columns.     The nodes of JSON received, contains the given information          layoutCheck: This node contains information if layout is selected by user or not(otherwise show default layout).      dataset: This node contains width information of all columns.      id: This node contains id that uniquely identify each column.      width: This node contains width of each column(if width is '0' then column is hidden). 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SetupIpServicesLayout**](Setup_IPServicesLayout.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST36iAction1Post**
> SetupIpServicesLayout POST36iAction1Post(ctx, body)
Save status layout

This service is used to update status layout details.    JSON that needs to be sent as request payload, Should contains given information-          id: id of to identify each column.       layoutcheck: if user selected to save layout information then '1' otherwise '0'.        width: width of each column, send '0' if needs to hide a column.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]StatusSaveRow**](StatusSave_row.md)| Object that needs to be sent to the server | 

### Return type

[**SetupIpServicesLayout**](Setup_IPServicesLayout.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

