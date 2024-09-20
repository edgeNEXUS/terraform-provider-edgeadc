# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET26downloadconfpreDownloadCheckyesGet**](ConfigurationApi.md#GET26downloadconfpreDownloadCheckyesGet) | **Get** /GET/26?download&#x3D;conf&amp;preDownloadCheck&#x3D;yes | Download configuration
[**POST26iAction1iType1sendconfPost**](ConfigurationApi.md#POST26iAction1iType1sendconfPost) | **Post** /POST/26?iAction&#x3D;1&amp;iType&#x3D;1&amp;send&#x3D;conf | Upload configuration

# **GET26downloadconfpreDownloadCheckyesGet**
> GET26downloadconfpreDownloadCheckyesGet(ctx, )
Download configuration

This service is used to download Configuration file in text format, you have to send download=conf in query string. 

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

# **POST26iAction1iType1sendconfPost**
> ConfigurationResponse POST26iAction1iType1sendconfPost(ctx, )
Upload configuration

This service is used to upload Configuration, you have to send Configuration file and send=conf in query string, that needs to be uploaded, by using form submit. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConfigurationResponse**](Configuration_Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

