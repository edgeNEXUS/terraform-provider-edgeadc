# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POST32pageloginPost**](LoginApi.md#POST32pageloginPost) | **Post** /POST/32?page&#x3D;login | Login request

# **POST32pageloginPost**
> JetLoginResp POST32pageloginPost(ctx, body)
Login request

This service is used to login user to ALB that allow him to use other services.    JSON that needs to be sent as request payload, Should contains given information-    username:\"password\" for example admin: \"jetnexus\"    The nodes of JSON received, contains the given information-        LoginStatus: This node contains information of user login status.     UserName: This node contains username.     GUID: This node contains a unique id(GUID) that needs to be setted in browser cookie(as GUID=) and should be sent with each web service request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JetLogin**](JetLogin.md)| Object that needs to be sent to the server | 

### Return type

[**JetLoginResp**](jet_LoginResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

