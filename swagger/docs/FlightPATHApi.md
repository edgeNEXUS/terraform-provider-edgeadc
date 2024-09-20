# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET11isPageLoadtrueGet**](FlightPATHApi.md#GET11isPageLoadtrueGet) | **Get** /GET/11?isPageLoad&#x3D;true | flightPATH page data
[**GET12Get**](FlightPATHApi.md#GET12Get) | **Get** /GET/12 | flightPATH combo options
[**POST11iAction10iType1FilterKeywordPost**](FlightPATHApi.md#POST11iAction10iType1FilterKeywordPost) | **Post** /POST/11?iAction&#x3D;10&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Add Action
[**POST11iAction11iType1FilterKeywordPost**](FlightPATHApi.md#POST11iAction11iType1FilterKeywordPost) | **Post** /POST/11?iAction&#x3D;11&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Update Action
[**POST11iAction12iType1FilterKeywordPost**](FlightPATHApi.md#POST11iAction12iType1FilterKeywordPost) | **Post** /POST/11?iAction&#x3D;12&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Remove action
[**POST11iAction1iType1Post**](FlightPATHApi.md#POST11iAction1iType1Post) | **Post** /POST/11?iAction&#x3D;1&amp;iType&#x3D;1 | add new flightPATH
[**POST11iAction2iType1FilterKeywordPost**](FlightPATHApi.md#POST11iAction2iType1FilterKeywordPost) | **Post** /POST/11?iAction&#x3D;2&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | update flightPATH
[**POST11iAction3iType1FilterKeywordPost**](FlightPATHApi.md#POST11iAction3iType1FilterKeywordPost) | **Post** /POST/11?iAction&#x3D;3&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Remove flightPATH
[**POST11iAction4iType1FilterKeywordPost**](FlightPATHApi.md#POST11iAction4iType1FilterKeywordPost) | **Post** /POST/11?iAction&#x3D;4&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Add new condition
[**POST11iAction5iType1FilterKeywordPost**](FlightPATHApi.md#POST11iAction5iType1FilterKeywordPost) | **Post** /POST/11?iAction&#x3D;5&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Update condition
[**POST11iAction6iType1FilterKeywordPost**](FlightPATHApi.md#POST11iAction6iType1FilterKeywordPost) | **Post** /POST/11?iAction&#x3D;6&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Remove condition
[**POST11iAction7iType1FilterKeywordPost**](FlightPATHApi.md#POST11iAction7iType1FilterKeywordPost) | **Post** /POST/11?iAction&#x3D;7&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | add evaluation
[**POST11iAction8iType1FilterKeywordPost**](FlightPATHApi.md#POST11iAction8iType1FilterKeywordPost) | **Post** /POST/11?iAction&#x3D;8&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Update Evaluation
[**POST11iAction9iType1FilterKeywordPost**](FlightPATHApi.md#POST11iAction9iType1FilterKeywordPost) | **Post** /POST/11?iAction&#x3D;9&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Remove evaluation

# **GET11isPageLoadtrueGet**
> FlightPathData GET11isPageLoadtrueGet(ctx, )
flightPATH page data

This service is used to FETCH all data for flightPATH page, it uses one API structures-    Config_FlightPath: Used to get all flightPATH and their Condition, evaluation and action settings details.      The nodes of JSON received, contains the given information        dataset: This node contains information of all flighPATHs.     fId: This node contains index of flightPATH row.     flightPathName: This node contains name of flightPATH.     flightPathDesc: This node contains description of flightPATH.     DiagnosticTracing: This node contains diagnostic tracing enabled or not.     flightPathInUse: This node contains all VS's IP address with which this flightPATH is in use.     conditions: This node contains information of all Conditions applied on a flightPATH.     cId: This node contains index of condition row.     condition: This node contains condition.     sense: This node contains sense.     check: This node contains check.     match: This node contains match.     condValue: This node contains Condition value.     values: This node contains information of all evaluation applied on a flightPATH.     vId: This node contains index of evaluation row.     variable: This node contains variable.     source: This node contains source.     detail: This node contains details.     valValue: This node contains evaluation value.     actions: This node contains information of all actions applied on a flightPATH.     aId: This node contains index of action row.     data: This node contains data.     action: This node contains action.     target: This node contains target.     conditionId:conditionId.     valuesId:valuesId.     actionId:actionId.     FilterKeyword:FilterKeyword. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FlightPathData**](flightPath_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET12Get**
> FlightPathCombo GET12Get(ctx, )
flightPATH combo options

This service is used to FETCH all data for Combo boxes of flightPATH page, we created a separate services to get combo box values because we need them just once when page loaded i. e. They are not changed on each update of flightPATH page, it uses one API structure     Config_FlightPath: Used to get all combo box(drop down) values used in flightPATH page-      The nodes of JSON received, contains the given information        SenseCombo: This node contains options of all senses.     CheckCombo: This node contains options of all Checks.     SourceCombo: This node contains options of all sources.     languageCombo: This node contains options of all languages.     conditionCombo: This node contains options of all conditions.     MatchCombo: This node contains options of all matches.     ConValueCombo: This node contains options of all condition value.     ActionCombo: This node contains options of all actions.     CountryCombo: This node contains options of all country.     AuthRuleCombo: This node contains options of authentication rules. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FlightPathCombo**](flightPath_Combo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction10iType1FilterKeywordPost**
> FlightPathRespData POST11iAction10iType1FilterKeywordPost(ctx, body)
Add Action

This service is used to add new action for a flightPATH in flightPATH action grid. you have to send id of flightPATH, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        fid: index(Received as fid in GET/11 service) of flightPATH for action needs to be added. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlightPathAddAct**](FlightPathAddAct.md)| Object that needs to be sent to the server | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction11iType1FilterKeywordPost**
> FlightPathRespData POST11iAction11iType1FilterKeywordPost(ctx, body)
Update Action

This service is used to update details in flightPATH action grid, you have to send id of flightPATH with id of row and data, that needs to be updated, with POST request json.     JSON that needs to be sent as request payload, Should contains given information-        fid: index(Received as fid in GET/11 service) of flightPATH for Action needs to be update.     aid: index(Received as aid in GET/11 service) of actions that needs to updated.     action: id(Received as id in GET/12 service, node ActionCombo) of option selected from Action options.     target: id(Received as id in GET/12 service, node ConValueCombo) of option selected from Target options.     data: Action data  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlightPathUpdateAct**](FlightPathUpdateAct.md)| Object that needs to be sent to the server | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction12iType1FilterKeywordPost**
> FlightPathRespData POST11iAction12iType1FilterKeywordPost(ctx, body)
Remove action

This service is used to remove row in flightPATH action grid, you have to send id of flightPATH with id of row, that needs to be removed, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        fid: index(Received as fid in GET/11 service) of flightPATH for action needs to be deleted.     aid: index(Received as aid in GET/11 service) of action that needs to be deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlightPathRemoveAct**](FlightPathRemoveAct.md)| Object that needs to be sent to the server | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction1iType1Post**
> FlightPathRespData POST11iAction1iType1Post(ctx, filterKeyword)
add new flightPATH

This service is used to add new flightPATH in flightPATH grid.   JSON that needs to be sent as request payload, Should contains given information- not required to send json. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filterKeyword** | **string**|  | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction2iType1FilterKeywordPost**
> FlightPathRespData POST11iAction2iType1FilterKeywordPost(ctx, body)
update flightPATH

This service is used to update details in flightPATH grid, you have to send id of row with data, that needs to be updated, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        fid : index(Received as fid in GET/11 service) of flightPATH that needs to be updated.     flightPathName : Flightpath Name.     flightPathDesc : Flightpath description. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlightPathUpdate**](FlightPathUpdate.md)| Object that needs to be sent to the server | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction3iType1FilterKeywordPost**
> FlightPathRespData POST11iAction3iType1FilterKeywordPost(ctx, body)
Remove flightPATH

This service is used to remove flightPATH rule from flightPATH grid, you have to send id of row, that needs to be removed, with POST request json.     JSON that needs to be sent as request payload, Should contains given information-        fid: index(Received as id in GET/11 service) of flightPATH that needs to be deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlightPathRemove**](FlightPathRemove.md)| Object that needs to be sent to the server | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction4iType1FilterKeywordPost**
> FlightPathRespData POST11iAction4iType1FilterKeywordPost(ctx, body)
Add new condition

This service is used to add new Condition for a flightPATH in flightPATH Condition grid. you have to send id of flightPATH, with POST request json.     JSON that needs to be sent as request payload, Should contains given information-.        fid: index(Received as fid in GET/11 service) of flightPATH for that condition needs to be added. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlightPathAddCon**](FlightPathAddCon.md)| Object that needs to be sent to the server | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction5iType1FilterKeywordPost**
> FlightPathRespData POST11iAction5iType1FilterKeywordPost(ctx, body)
Update condition

This service is used to update details in flightPATH condition grid, you have to send id of flightPATH with id of row and data, that needs to be updated, with POST request json.     JSON that needs to be sent as request payload, Should contains given information-        fid: index(Received as fid in GET/11 service) of flightPATH for that condition needs to be updated.     cid: index(Received as cid in GET/11 service) of conditions that needs to updated.     condition : id(Received as id in GET/12 service, node ActionCombo) of option selected from Condition options.     match: id(Received as id in GET/12 service, node ConValueCombo) of option selected from match options.     sense: id(Received as id in GET/12 service, node SenseCombo) of option selected from sense options.     check: id(Received as id in GET/12 service, node CheckCombo) of option selected from check options.     condValue:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlightPathUpdateCon**](FlightPathUpdateCon.md)| Object that needs to be sent to the server | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction6iType1FilterKeywordPost**
> FlightPathRespData POST11iAction6iType1FilterKeywordPost(ctx, body)
Remove condition

This service is used to remove row in flightPATH condition grid, you have to send id of flightPATH with id of row, that needs to be removed, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        fid: index(Received as fid in GET/11 service) of flightPATH for condition needs to be deleted.     cid: index(Received as cid in GET/11 service) of conditions that needs to delete. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlightPathRemoveCon**](FlightPathRemoveCon.md)| Object that needs to be sent to the server | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction7iType1FilterKeywordPost**
> FlightPathRespData POST11iAction7iType1FilterKeywordPost(ctx, body)
add evaluation

This service is used to add new evaluation for a flightPATH in flightPATH evaluation grid. you have to send id of flightPATH, with POST request json.     JSON that needs to be sent as request payload, Should contains given information-        fid: index(Received as fid in GET/11 service) of flightPATH for that evaluation needs to be added. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlightPathAddEva**](FlightPathAddEva.md)| Object that needs to be sent to the server | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction8iType1FilterKeywordPost**
> FlightPathRespData POST11iAction8iType1FilterKeywordPost(ctx, body)
Update Evaluation

This service is used to update details in flightPATH evaluation grid, you have to send id of flightPATH with id of row and data, that needs to be updated, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        fid: index(Received as fid in GET/11 service) of flightPATH for evaluation needs to be updated.     vid: index(Received as vid in GET/11 service) of values that needs to updated.     variable: variable     source: id(Received as id in GET/12 service, node SourceCombo) of option selected from Source options.     Detail: id(Received as id in GET/12 service, node ConValueCombo) of option selected from Detail options.     valValue: valValue 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlightPathUpdateEva**](FlightPathUpdateEva.md)| Object that needs to be sent to the server | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST11iAction9iType1FilterKeywordPost**
> FlightPathRespData POST11iAction9iType1FilterKeywordPost(ctx, body)
Remove evaluation

This service is used to remove row in flightPATH evaluation grid, you have to send id of flightPATH with id of row, that needs to be removed, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        fid: index(Received as fid in GET/11 service) of flightPATH for evaluation needs to be deleted.     cId: index(Received as vid in GET/11 service) of evaluation needs to be deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlightPathRemoveEva**](FlightPathRemoveEva.md)| Object that needs to be sent to the server | 

### Return type

[**FlightPathRespData**](flightPath_RespData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

