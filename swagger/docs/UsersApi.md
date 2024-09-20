# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET33Get**](UsersApi.md#GET33Get) | **Get** /GET/33 | User page data
[**GET34Get**](UsersApi.md#GET34Get) | **Get** /GET/34 | User access logs
[**GET34downloadauditiAction1iType1preDownloadCheckyesGet**](UsersApi.md#GET34downloadauditiAction1iType1preDownloadCheckyesGet) | **Get** /GET/34?download&#x3D;audit&amp;iAction&#x3D;1&amp;iType&#x3D;1&amp;preDownloadCheck&#x3D;yes | Download access log in text
[**GET34downloadauditiAction1iType2preDownloadCheckyesGet**](UsersApi.md#GET34downloadauditiAction1iType2preDownloadCheckyesGet) | **Get** /GET/34?download&#x3D;audit&amp;iAction&#x3D;1&amp;iType&#x3D;2&amp;preDownloadCheck&#x3D;yes | Download access log in Compressed
[**POST33iAction1iType1Post**](UsersApi.md#POST33iAction1iType1Post) | **Post** /POST/33?iAction&#x3D;1&amp;iType&#x3D;1 | Add new user
[**POST33iAction1iType2Post**](UsersApi.md#POST33iAction1iType2Post) | **Post** /POST/33?iAction&#x3D;1&amp;iType&#x3D;2 | Update user
[**POST33iAction1iType3Post**](UsersApi.md#POST33iAction1iType3Post) | **Post** /POST/33?iAction&#x3D;1&amp;iType&#x3D;3 | Delete user

# **GET33Get**
> UsersData GET33Get(ctx, )
User page data

This service is used to FETCH USERs details, it uses one API structure.    Config_MultiUserAccess: Used to get all Users and their information(cluster user or local, access rights etc).    The nodes of JSON received, contains the given information-          Members: This node contains data of all users.       id: This node contains id of user.       OldPassword: This node contains oldpassword of users.       NewPassword: This node contains newpassword of users.       UserType: This node contains user type(such as cluster, local or clusterlocal).       UserName: This node contains username.       isAdmin: This node contains '1' if user have admin access rights.       isGUIW: This node contains '1' if user have GUI Write access rights.       isGUIR: This node contains '1' if user have GUI Read access rights.       isSSH: This node contains '1' if user have SSH access rights.       isAPI: This node contains '1' if user have API access rights.       isAddOn : This node contains '1' if user have AddOn access rights. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UsersData**](Users_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET34Get**
> UsersLogs GET34Get(ctx, )
User access logs

This service is used to FETCH User access log details, it uses one API structure.    Monitor_Logging2: Used to get last 50 access logs of current user.      The nodes of JSON received, contains the given information          Logs: This node contains data of last 50 access logs.     Row: This node contains one access log each. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UsersLogs**](Users_Logs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET34downloadauditiAction1iType1preDownloadCheckyesGet**
> GET34downloadauditiAction1iType1preDownloadCheckyesGet(ctx, )
Download access log in text

This service is used to download user access log file in text format, you have to send download=audit in query string. 

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

# **GET34downloadauditiAction1iType2preDownloadCheckyesGet**
> GET34downloadauditiAction1iType2preDownloadCheckyesGet(ctx, )
Download access log in Compressed

This service is used to download user access log file in compressed format(GZIP), you have to send download=audit in query string. 

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

# **POST33iAction1iType1Post**
> UsersData POST33iAction1iType1Post(ctx, body)
Add new user

This service is used to add new user, you have to send username, password and access rights needed for user, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        UserName : UserName.      NewPassword : New Password.      OldPassword : Old Password.      isAPI : set to '1' if user have API access rights.     isAddOn : set to '1' if user have AddOn access rights.      isAdmin : set to '1' if user have admin access rights.      isGUIR : set to '1' if user have GUI Read access rights.      isGUIW : set to '1' if user have GUI Write access rights.      isSSH : set to '1' if user have SSH access rights. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsersAdd**](UsersAdd.md)| Object that needs to be sent to the server | 

### Return type

[**UsersData**](Users_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST33iAction1iType2Post**
> UsersData POST33iAction1iType2Post(ctx, body)
Update user

This service is used to update user details.    JSON that needs to be sent as request payload, Should contains given information-        UserName : UserName.      NewPassword : New Password.      OldPassword : Old Password.      isAPI : set to '1' if user have API access rights.     isAddOn : set to '1' if user have AddOn access rights.      isAdmin : set to '1' if user have admin access rights.      isGUIR : set to '1' if user have GUI Read access rights.      isGUIW : set to '1' if user have GUI Write access rights.      isSSH : set to '1' if user have SSH access rights. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsersUpdate**](UsersUpdate.md)| Object that needs to be sent to the server | 

### Return type

[**UsersData**](Users_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST33iAction1iType3Post**
> UsersData POST33iAction1iType3Post(ctx, body)
Delete user

This service is used to delete a user, you have to send username, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        UserName : UserName 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsersDelete**](UsersDelete.md)| Object that needs to be sent to the server | 

### Return type

[**UsersData**](Users_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

