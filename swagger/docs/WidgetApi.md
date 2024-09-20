# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET48EditWidgettrueWidgetNameIPServiceWidget1Get**](WidgetApi.md#GET48EditWidgettrueWidgetNameIPServiceWidget1Get) | **Get** /GET/48?EditWidget&#x3D;true&amp;WidgetName&#x3D;IPService-widget1 | Edit IPService Widget
[**GET51Get**](WidgetApi.md#GET51Get) | **Get** /GET/51 | Widget details
[**POST48iAction10iType1Post**](WidgetApi.md#POST48iAction10iType1Post) | **Post** /POST/48?iAction&#x3D;10&amp;iType&#x3D;1 | Remove Widget
[**POST48iAction1iType1Post**](WidgetApi.md#POST48iAction1iType1Post) | **Post** /POST/48?iAction&#x3D;1&amp;iType&#x3D;1 | Make Event Widget
[**POST48iAction2iType1Post**](WidgetApi.md#POST48iAction2iType1Post) | **Post** /POST/48?iAction&#x3D;2&amp;iType&#x3D;1 | Make IPService Widget
[**POST48iAction3iType1Post**](WidgetApi.md#POST48iAction3iType1Post) | **Post** /POST/48?iAction&#x3D;3&amp;iType&#x3D;1 | Make IPstatus Widget
[**POST48iAction4iType1Post**](WidgetApi.md#POST48iAction4iType1Post) | **Post** /POST/48?iAction&#x3D;4&amp;iType&#x3D;1 | Make System Widget
[**POST48iAction4iType2Post**](WidgetApi.md#POST48iAction4iType2Post) | **Post** /POST/48?iAction&#x3D;4&amp;iType&#x3D;2 | Make Interface Widget
[**POST48iAction5iType1Post**](WidgetApi.md#POST48iAction5iType1Post) | **Post** /POST/48?iAction&#x3D;5&amp;iType&#x3D;1 | Update Event Widget
[**POST48iAction6iType1Post**](WidgetApi.md#POST48iAction6iType1Post) | **Post** /POST/48?iAction&#x3D;6&amp;iType&#x3D;1 | Update IPService Widget
[**POST48iAction7iType1Post**](WidgetApi.md#POST48iAction7iType1Post) | **Post** /POST/48?iAction&#x3D;7&amp;iType&#x3D;1 | Update IPStatus Widget
[**POST48iAction8iType1Post**](WidgetApi.md#POST48iAction8iType1Post) | **Post** /POST/48?iAction&#x3D;8&amp;iType&#x3D;1 | Update System Widget
[**POST51iAction2iType1pageWidgetPost**](WidgetApi.md#POST51iAction2iType1pageWidgetPost) | **Post** /POST/51?iAction&#x3D;2&amp;iType&#x3D;1&amp;page&#x3D;Widget | Edit Event Widget
[**POST51iAction2iType3pageWidgetPost**](WidgetApi.md#POST51iAction2iType3pageWidgetPost) | **Post** /POST/51?iAction&#x3D;2&amp;iType&#x3D;3&amp;page&#x3D;Widget | Edit System Widget
[**POST51iAction2iType4pageWidgetPost**](WidgetApi.md#POST51iAction2iType4pageWidgetPost) | **Post** /POST/51?iAction&#x3D;2&amp;iType&#x3D;4&amp;page&#x3D;Widget | Edit IPStatus Widget

# **GET48EditWidgettrueWidgetNameIPServiceWidget1Get**
> EditIpServiceResp GET48EditWidgettrueWidgetNameIPServiceWidget1Get(ctx, )
Edit IPService Widget

This service is used for edit IPService widget. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EditIpServiceResp**](EditIPServiceResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET51Get**
> GetWidget GET51Get(ctx, )
Widget details

This service is used to FETCH widget details.    The nodes of JSON received, contains the given information      ConfiguredWidgetsComboStore: This node contains Configured Widgets options.     data: This node contains avaliable widgets.     PartID: This node contains part id of widget.      SectionName: This node contains section of widget.      Widget_Name: This node contains name of widget. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetWidget**](GetWidget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST48iAction10iType1Post**
> RemoveWidgetResp POST48iAction10iType1Post(ctx, body)
Remove Widget

This service is used for remove widget.    JSON that needs to be sent as request payload, Should contains given information-        Remove_Section_Name: section name of widget that needs to be removed.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveWidget**](RemoveWidget.md)| Object that needs to be sent to the server | 

### Return type

[**RemoveWidgetResp**](RemoveWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST48iAction1iType1Post**
> RemoveWidgetResp POST48iAction1iType1Post(ctx, body)
Make Event Widget

This service is used for make event widget.    JSON that needs to be sent as request payload, Should contains given information-        DashboardEvent_Filter: filter for widget.     DashboardEvent_Name: widget name.     Section_Name:widget section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventWidget**](EventWidget.md)| Object that needs to be sent to the server | 

### Return type

[**RemoveWidgetResp**](RemoveWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST48iAction2iType1Post**
> RemoveWidgetResp POST48iAction2iType1Post(ctx, body)
Make IPService Widget

This service is used for make ipservice widget.    JSON that needs to be sent as request payload, Should contains given information-        IPService_Columns:      IPService_Name:     IPService_Period:     IPService_Type:     IPService_VSRS:     Section_Name: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpServiceWidget**](IpServiceWidget.md)| Object that needs to be sent to the server | 

### Return type

[**RemoveWidgetResp**](RemoveWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST48iAction3iType1Post**
> RemoveWidgetResp POST48iAction3iType1Post(ctx, body)
Make IPstatus Widget

This service is used for make ipstatus widget.    JSON that needs to be sent as request payload, Should contains given information-        IPStatus_IP: filter IP.     IPStatus_Layout: widget layout.     IPStatus_Name:widget name.     Section_Name:widget section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpStatusWidget**](IpStatusWidget.md)| Object that needs to be sent to the server | 

### Return type

[**RemoveWidgetResp**](RemoveWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST48iAction4iType1Post**
> RemoveWidgetResp POST48iAction4iType1Post(ctx, body)
Make System Widget

This service is used for make system widget.    JSON that needs to be sent as request payload, Should contains given information-        Section_Name: widget section.     System_Name: widget name.     System_Type: system type. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SystemWidget**](SystemWidget.md)| Object that needs to be sent to the server | 

### Return type

[**RemoveWidgetResp**](RemoveWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST48iAction4iType2Post**
> RemoveWidgetResp POST48iAction4iType2Post(ctx, body)
Make Interface Widget

This service is used for make interface widget.    JSON that needs to be sent as request payload, Should contains given information-        Interface_Name: name of interface.      Section_Name: widget section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InterfaceWidget**](InterfaceWidget.md)| Object that needs to be sent to the server | 

### Return type

[**RemoveWidgetResp**](RemoveWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST48iAction5iType1Post**
> RemoveWidgetResp POST48iAction5iType1Post(ctx, body)
Update Event Widget

This service is used for update event widget after edit.    JSON that needs to be sent as request payload, Should contains given information-        DashboardEvent_Filter: filter keyword.      DashboardEvent_Name: name of widget.     Update_Section_Name: section name of widget.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateEventWidget**](UpdateEventWidget.md)| Object that needs to be sent to the server | 

### Return type

[**RemoveWidgetResp**](RemoveWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST48iAction6iType1Post**
> RemoveWidgetResp POST48iAction6iType1Post(ctx, body)
Update IPService Widget

This service is used for update IPService widget after edit.    JSON that needs to be sent as request payload, Should contains given information-        IPService_Columns:       IPService_Period: .      IPService_Type:      IPService_VSRS:      IPService_Name: name of widget.     Update_Section_Name: section name of widget.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateIpServiceWidget**](UpdateIpServiceWidget.md)| Object that needs to be sent to the server | 

### Return type

[**RemoveWidgetResp**](RemoveWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST48iAction7iType1Post**
> []RemoveWidgetResp POST48iAction7iType1Post(ctx, body)
Update IPStatus Widget

This service is used for update IPStatus widget after edit.    JSON that needs to be sent as request payload, Should contains given information-        IPStatus_IP: ip address.      IPStatus_Layout:      IPStatus_Name: name of widget.     Update_Section_Name: section name of widget.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateIpStatusWidget**](UpdateIpStatusWidget.md)| Object that needs to be sent to the server | 

### Return type

[**[]RemoveWidgetResp**](RemoveWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST48iAction8iType1Post**
> RemoveWidgetResp POST48iAction8iType1Post(ctx, body)
Update System Widget

This service is used for update system widget after edit.    JSON that needs to be sent as request payload, Should contains given information-        System_Name: name of widget.      System_Type:      Update_Section_Name: section name of widget.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSysWidget**](UpdateSysWidget.md)| Object that needs to be sent to the server | 

### Return type

[**RemoveWidgetResp**](RemoveWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction2iType1pageWidgetPost**
> EditEventResp POST51iAction2iType1pageWidgetPost(ctx, body)
Edit Event Widget

This service is used for edit event widget.    JSON that needs to be sent as request payload, Should contains given information-        Section_Name: section name of widget.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EditSysWidget**](EditSysWidget.md)| Object that needs to be sent to the server | 

### Return type

[**EditEventResp**](EditEventResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction2iType3pageWidgetPost**
> EditWidgetResp POST51iAction2iType3pageWidgetPost(ctx, body)
Edit System Widget

This service is used for edit system widget.    JSON that needs to be sent as request payload, Should contains given information-        Section_Name: section name of widget.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EditSysWidget**](EditSysWidget.md)| Object that needs to be sent to the server | 

### Return type

[**EditWidgetResp**](EditWidgetResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST51iAction2iType4pageWidgetPost**
> EditIpStatusResp POST51iAction2iType4pageWidgetPost(ctx, body)
Edit IPStatus Widget

This service is used for edit IPStatus widget.    JSON that needs to be sent as request payload, Should contains given information-        Section_Name: section name of widget.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EditSysWidget**](EditSysWidget.md)| Object that needs to be sent to the server | 

### Return type

[**EditIpStatusResp**](EditIPStatusResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

