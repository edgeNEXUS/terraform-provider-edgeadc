# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET40Get**](AuthenticationApi.md#GET40Get) | **Get** /GET/40 | Authentication page data
[**POST38iAction1iType1Post**](AuthenticationApi.md#POST38iAction1iType1Post) | **Post** /POST/38?iAction&#x3D;1&amp;iType&#x3D;1 | Add Authentication Server
[**POST38iAction1iType2Post**](AuthenticationApi.md#POST38iAction1iType2Post) | **Post** /POST/38?iAction&#x3D;1&amp;iType&#x3D;2 | Update Authentication Server
[**POST38iAction1iType3Post**](AuthenticationApi.md#POST38iAction1iType3Post) | **Post** /POST/38?iAction&#x3D;1&amp;iType&#x3D;3 | Remove Authentication Server
[**POST38iAction2iType1Post**](AuthenticationApi.md#POST38iAction2iType1Post) | **Post** /POST/38?iAction&#x3D;2&amp;iType&#x3D;1 | Add Authentication Rule
[**POST38iAction2iType2Post**](AuthenticationApi.md#POST38iAction2iType2Post) | **Post** /POST/38?iAction&#x3D;2&amp;iType&#x3D;2 | Update Authentication Rule
[**POST38iAction2iType3Post**](AuthenticationApi.md#POST38iAction2iType3Post) | **Post** /POST/38?iAction&#x3D;2&amp;iType&#x3D;3 | Remove Authentication Rule
[**POST38iAction3iType1sendformPost**](AuthenticationApi.md#POST38iAction3iType1sendformPost) | **Post** /POST/38?iAction&#x3D;3&amp;iType&#x3D;1&amp;send&#x3D;form | Upload Authentication Form
[**POST38iAction4iType1Post**](AuthenticationApi.md#POST38iAction4iType1Post) | **Post** /POST/38?iAction&#x3D;4&amp;iType&#x3D;1 | Remove Form

# **GET40Get**
> AuthenticationData GET40Get(ctx, )
Authentication page data

This service is used to FETCH all data for Authentication page(grid data and combo box options), it uses three API structures    Authentication_Server: Used to get rows details for Authentication server grid.    Authentication_Rule: Used to get rows details for Authentication rule grid.    Authentication_Form: Used to get all avalable authentication forms.      The nodes of JSON received, contains the given information        AuthServersCombo: This node contains options of various authentication server configured.     AuthMethodCombo: This node contains options of various authentication methods such as LDAP and RADIUS.     ClientAuthMethodsCombo: This node contains options of various client authentication methods such as Forms, Basic etc.     ServerAuthMethodsCombo: This node contains options of various server authentication methods such as NTLM and BASIC.     AuthFormsCombo: This node contains options of various forms uploaded.     LoginFormatListString: This node contains options of various login formats such as username only, username and domain etc.     PreAuthServerGrid: This node contains authentication server grid data.     id: This node contains index of each server which needed to update or delete that server.     Name: This node contains name of server.     Description: This node contains description of server.     AuthMethod: This node contains authentication method used for this server.     Domain: This node contains domain.     Address: This node contains ip address.     Port: This node contains port.     SearchBase: This node contains search base.     LoginFormat: This node contains login format.     Passphrase: This node contains authentication server password.     DeadTime: This node contains authentication server dead time.     Searchcondition: This node contains authentication searchcondition .     PreAuthRuleGrid: This node contains authentication rule grid data.     id: This node contains index of each rule which needed to update or delete that rule.     Name: This node contains name of rule.     Description: This node contains description of rule.     RootDomain:      AuthServer: This node contains authentication server used for this rule.     ClientAuthMethod: This node contains client authentication method.     ServerAuthMethod: This node contains server authentication method.     AuthForm: This node contains authentication form used for this rule.     Message: This node contains message.     dataset:dataset.     row:row.     Timeout: This node contains timeout. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthenticationData**](Authentication_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST38iAction1iType1Post**
> AuthenticationData POST38iAction1iType1Post(ctx, body)
Add Authentication Server

This service is used to add new Authentication server in Authentication server grid. you have to send data of new Row, that needs to be added, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-           Address: IP address.       AuthMethod: Value(Received as Value in GET/40 service, node AuthMethodCombo) of option selected from Authentication method options.       DeadTime:        Description: Authentication server Description.       Domain: Domain Name.       LoginFormat: id(Received as id in GET/40 service, node LoginFormatListString) of option selected from Login Format options.       Name: Username.       Passphrase: Password.       Port: Port number.       SearchBase:        Searchcondition: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddAuthentication**](AddAuthentication.md)| Object that needs to be sent to the service | 

### Return type

[**AuthenticationData**](Authentication_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST38iAction1iType2Post**
> AuthenticationData POST38iAction1iType2Post(ctx, body)
Update Authentication Server

This service is used to update details in Authentication server grid, you have to send id of Authentication server with data, that needs to be updated, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-            Address: IP address.         AuthMethod: Value(Received as Value in GET/40 service, node AuthMethodCombo) of option selected from Authentication method options.         DeadTime:          Description: Authentication server Description.         Domain: Domain Name.         LoginFormat: id(Received as id in GET/40 service, node LoginFormatListString) of option selected from Login Format options.         Name: Username.         Passphrase: Password.         Port: Port number.         SearchBase:          Searchcondition:         id: index(Received as id in GET/40 service) of PreAuthServerGrid that needs to updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateAuthentication**](UpdateAuthentication.md)| Object that needs to be sent to the service | 

### Return type

[**AuthenticationData**](Authentication_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST38iAction1iType3Post**
> AuthenticationData POST38iAction1iType3Post(ctx, body)
Remove Authentication Server

This service is used to remove row in Authentication server grid, you have to send id of Authentication server, that needs to be removed, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        Id: index(Received as id in GET/40 service) of Authentication Server that needs to be deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveAuthentication**](RemoveAuthentication.md)| Object that needs to be sent to the service | 

### Return type

[**AuthenticationData**](Authentication_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST38iAction2iType1Post**
> AuthenticationData POST38iAction2iType1Post(ctx, body)
Add Authentication Rule

This service is used to add new Authentication rule in Authentication rule grid. you have to send data of new Row, that needs to be added, with POST request json.     JSON that needs to be sent as request payload, Should contains given information-          AuthForm: id(Received as id in GET/40 service, node AuthFormsCombo) of option selected from Form options.      AuthServer: id(Received as id in GET/40 service, node AuthServersCombo) of option selected from Authentication server options.      ClientAuthMethod: id(Received as id in GET/40 service, node ClientAuthMethodsCombo) of option selected from client Authentication options.      Description:       Message: Message.      Name: Authentication server name      RootDomain:      ServerAuthMethod: id(Received as id in GET/40 service, node ServerAuthMethodsCombo) of option selected from server Authentication options.      Timeout: Authentication server name 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RuleAuthentication**](RuleAuthentication.md)| Object that needs to be sent to the service | 

### Return type

[**AuthenticationData**](Authentication_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST38iAction2iType2Post**
> AuthenticationData POST38iAction2iType2Post(ctx, body)
Update Authentication Rule

This service is used to Update Authentication Rule in Authentication rule grid. you have to send data of new Row, that needs to be added, with POST request json.     JSON that needs to be sent as request payload, Should contains given information-          AuthForm: id(Received as id in GET/40 service, node AuthFormsCombo) of option selected from Form options.      AuthServer: id(Received as id in GET/40 service, node AuthServersCombo) of option selected from Authentication server options.      ClientAuthMethod: id(Received as id in GET/40 service, node ClientAuthMethodsCombo) of option selected from client Authentication options.      Description:       Message: Message.      Name: Authentication server name      RootDomain:      ServerAuthMethod: id(Received as id in GET/40 service, node ServerAuthMethodsCombo) of option selected from server Authentication options.      Timeout: Authentication server name      id: index(Received as id in GET/40 service) of PreAuthRuleGrid that needs to updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RuleUpdateAuthentication**](RuleUpdateAuthentication.md)| Object that needs to be sent to the service | 

### Return type

[**AuthenticationData**](Authentication_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST38iAction2iType3Post**
> AuthenticationData POST38iAction2iType3Post(ctx, body)
Remove Authentication Rule

This service is used to remove row in Authentication rule grid, you have to send id of Authentication rule, that needs to be removed, with POST request json.     JSON that needs to be sent as request payload, Should contains given information-          id: index(Received as id in GET/40 service) of Authentication rule that needs to be deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RuleDeleteAuthentication**](RuleDeleteAuthentication.md)| Object that needs to be sent to the service | 

### Return type

[**AuthenticationData**](Authentication_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST38iAction3iType1sendformPost**
> AuthenticationData POST38iAction3iType1sendformPost(ctx, )
Upload Authentication Form

This service is used to upload Authentication form file, you have to send name in form field(UploadFormName) and file with send=form in query string, that needs to be uploaded, with form submit.     We have to send given values to form submit          UploadFormName: Upload Form Name.      file: File needs to be Uploaded. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthenticationData**](Authentication_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST38iAction4iType1Post**
> AuthenticationData POST38iAction4iType1Post(ctx, body)
Remove Form

This service is used to Remove Form. you have to send id, that needs to be remove, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-           id: id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveAuthenticationForm**](RemoveAuthenticationForm.md)| Object that needs to be sent to the service | 

### Return type

[**AuthenticationData**](Authentication_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

