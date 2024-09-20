# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET25Get**](ProtocolApi.md#GET25Get) | **Get** /GET/25 | Protocol page data
[**GET44Get**](ProtocolApi.md#GET44Get) | **Get** /GET/44 | Server too busy file content
[**POST25iAction1iType1Post**](ProtocolApi.md#POST25iAction1iType1Post) | **Post** /POST/25?iAction&#x3D;1&amp;iType&#x3D;1 | Update &#x27;forward for&#x27;
[**POST25iAction2iType1Post**](ProtocolApi.md#POST25iAction2iType1Post) | **Post** /POST/25?iAction&#x3D;2&amp;iType&#x3D;1 | Update HTTP Compression
[**POST25iAction3iType1Post**](ProtocolApi.md#POST25iAction3iType1Post) | **Post** /POST/25?iAction&#x3D;3&amp;iType&#x3D;1 | Update current exclusion
[**POST25iAction4iType1sendprotocolPost**](ProtocolApi.md#POST25iAction4iType1sendprotocolPost) | **Post** /POST/25?iAction&#x3D;4&amp;iType&#x3D;1&amp;send&#x3D;protocol | Upload server too busy file
[**POST25iAction5iType1Post**](ProtocolApi.md#POST25iAction5iType1Post) | **Post** /POST/25?iAction&#x3D;5&amp;iType&#x3D;1 | Enable/Disable server too busy page
[**POST25iAction6iType1Post**](ProtocolApi.md#POST25iAction6iType1Post) | **Post** /POST/25?iAction&#x3D;6&amp;iType&#x3D;1 | Update Persistence Cookies

# **GET25Get**
> ProtocolData GET25Get(ctx, )
Protocol page data

This service is used to FETCH Protocol details, it uses four API structure.    Config_Performance: Used to get Performance settings details.    Config_ExclusionNew: Used to get current exclusion details.    Setup_XFF: Used to get Forwarding details.    Setup_ServerBusy: Used to get information if server too busy enabled.    The nodes of JSON received, contains the given information-          XFFCombo: This node contains options of X Forward For.     CompressAsYouGoCombo: This node contains options of compress.     SendServerTooBusyPage: This node contains '1' if server too busy page enabled.     XffOutputList: This node contains XFF output list.     XffHeader: This node contains XFF header.     CPF_INITIAL_PAGE_BUFFER_SIZE_K: This node contains initial page buffer size in KB.     CPF_MAX_CONTENT_LENGTH_K: This node contains maximum content length in KB.     CPF_INCREMENT_PAGE_BUFFER_SIZE_K: This node contains increment page buffer size in KB.     CPF_MinimumCompressionSize: This node contains minimum compression size.     CPF_DO_NOT_COMPRESS_CSS: This node contains '1' if do not compress css.     CPF_CompressionDisabled: This node contains '1' if compression disabled.     CPF_CompressAsYouGo: This node contains '1' if compress as you go enabled.     CurrentExclusions: This node contains current exclusions.     SameSite: This node contains SameSite.     Secure: This node contains either 0 or 1.     HttpOnly: This node contains either 0 or 1.     SameSiteListString: This node contains list of SameSiteListString. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProtocolData**](Protocol_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET44Get**
> ProtocolServer GET44Get(ctx, )
Server too busy file content

This service is used to FETCH Server too busy file content, it uses one API structure .    Setup_ServerBusy: Used to get server too busy file content.      The nodes of JSON received, contains the given information        File: This node contains all content of server too busy file. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProtocolServer**](Protocol_Server.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST25iAction1iType1Post**
> ProtocolData POST25iAction1iType1Post(ctx, body)
Update 'forward for'

This service is used to update 'Forward for' settings, you have to send data, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        ForwardedForOutput : id(Received as id in GET/25 service, node XFFCombo) of option selected from Forworded-for Output options.      ForwardedForHeader : Forworded-for Header. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProtocolUpdate**](ProtocolUpdate.md)| Object that needs to be sent to the server | 

### Return type

[**ProtocolData**](Protocol_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST25iAction2iType1Post**
> ProtocolData POST25iAction2iType1Post(ctx, body)
Update HTTP Compression

This service is used to update HTTP Compression settings, you have to send data, with POST request json.  JSON that needs to be sent as request payload, Should contains given information-      InitialThreadMemory : Initial Thread Memory.      MaximumThreadMemory : Maxmimum Thread Memory.      IncrementMemory : Increment Memory.      MinimumCompressionSize : Minimum compression size.      SafeMode : set 'true' if safe mode enabled.      DisableCompression : set 'true' if disabled compression.      CompressAsYouGo : id(Received as id in GET/25 service, node CompressAsYouGoCombo) of option selected from compress as you Go options. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProtocolUpdateHttp**](ProtocolUpdateHttp.md)| Object that needs to be sent to the server | 

### Return type

[**ProtocolData**](Protocol_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST25iAction3iType1Post**
> ProtocolData POST25iAction3iType1Post(ctx, body)
Update current exclusion

This service is used to update Current exclusion settings, you have to send CurrentExclusions, with POST request json.  JSON that needs to be sent as request payload, Should contains given information-      CurrentExclusions: current exclusions for example *.gif. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProtocolUpdateConclusion**](ProtocolUpdateConclusion.md)| Object that needs to be sent to the server | 

### Return type

[**ProtocolData**](Protocol_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST25iAction4iType1sendprotocolPost**
> ProtocolData POST25iAction4iType1sendprotocolPost(ctx, file)
Upload server too busy file

This service is used to upload Server too busy file, you have to send file and send=protocol in query string, that needs to be uploaded, by using form submit.    We have to send given values to form submit.        file: Upload html file. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 

### Return type

[**ProtocolData**](Protocol_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST25iAction5iType1Post**
> ProtocolData POST25iAction5iType1Post(ctx, body)
Enable/Disable server too busy page

This service is used to enable/disable 'server too busy page', you have to send data, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        HiddenSendServerTooBusyPage : set 'true' to enable server too busy page. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProtocolUpdateEneble**](ProtocolUpdateEneble.md)| Object that needs to be sent to the server | 

### Return type

[**ProtocolData**](Protocol_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST25iAction6iType1Post**
> ProtocolData POST25iAction6iType1Post(ctx, body)
Update Persistence Cookies

This service is used to update Persistence Cookies, you have to send Persistence Cookies, with POST request json.  JSON that needs to be sent as request payload, Should contain given information-      SameSite: This node contains SameSite.     Secure: This node contains either 0 or 1.     HttpOnly: This node contains either 0 or 1. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProtocolPersistenceCookies**](ProtocolPersistenceCookies.md)| Object that needs to be sent to the server | 

### Return type

[**ProtocolData**](Protocol_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

