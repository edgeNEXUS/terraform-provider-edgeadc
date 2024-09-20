# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET35Get**](LogoutApi.md#GET35Get) | **Get** /GET/35 | Logout user

# **GET35Get**
> JetLogoutResp GET35Get(ctx, )
Logout user

This service is used to logout a user. GUID is sent with each request so when ALB receive this service, it logout the user associated with it's GUID. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JetLogoutResp**](jet_LogoutResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

