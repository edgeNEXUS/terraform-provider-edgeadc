# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET18Get**](CacheApi.md#GET18Get) | **Get** /GET/18 | Cache page data with combo data
[**GET18isCacheRequesttruereqrefreshGet**](CacheApi.md#GET18isCacheRequesttruereqrefreshGet) | **Get** /GET/18?isCacheRequest&#x3D;true&amp;req&#x3D;refresh | Cache page data for Domain Combo
[**POST18iAction1iType1Post**](CacheApi.md#POST18iAction1iType1Post) | **Post** /POST/18?iAction&#x3D;1&amp;iType&#x3D;1 | Update global cache settings
[**POST18iAction2iType1Post**](CacheApi.md#POST18iAction2iType1Post) | **Post** /POST/18?iAction&#x3D;2&amp;iType&#x3D;1 | Add new host
[**POST18iAction2iType2Post**](CacheApi.md#POST18iAction2iType2Post) | **Post** /POST/18?iAction&#x3D;2&amp;iType&#x3D;2 | Remove host
[**POST18iAction2iType3Post**](CacheApi.md#POST18iAction2iType3Post) | **Post** /POST/18?iAction&#x3D;2&amp;iType&#x3D;3 | Update apply cache domain grid
[**POST18iAction2iType4Post**](CacheApi.md#POST18iAction2iType4Post) | **Post** /POST/18?iAction&#x3D;2&amp;iType&#x3D;4 | Add Domain(Other Domains Served)
[**POST18iAction2iType5Post**](CacheApi.md#POST18iAction2iType5Post) | **Post** /POST/18?iAction&#x3D;2&amp;iType&#x3D;5 | Remove Domain(Other Domains Served)
[**POST18iAction3iType1Post**](CacheApi.md#POST18iAction3iType1Post) | **Post** /POST/18?iAction&#x3D;3&amp;iType&#x3D;1 | Add new cache rule
[**POST18iAction3iType2Post**](CacheApi.md#POST18iAction3iType2Post) | **Post** /POST/18?iAction&#x3D;3&amp;iType&#x3D;2 | Remove cache rule
[**POST18iAction3iType3Post**](CacheApi.md#POST18iAction3iType3Post) | **Post** /POST/18?iAction&#x3D;3&amp;iType&#x3D;3 | Update cache rule

# **GET18Get**
> ConfigCache GET18Get(ctx, )
Cache page data with combo data

This service is used to FETCH all data for Cache page, it uses one API structures.    Config_Cache: Used to get all cache Rules, domains and other cache settings details.       The nodes of JSON received, contains the given information        CacheIncExcCombo: This node contains options of include and exclude.     CacheRuleHelpTypeCombo: This node contains options of various cache rule help to include or exclude.     CCHMaxCache: This node contains value of maximum cache size.     CCHMinCache: This node contains value of minimum cache size.     CCHHeuristicExpiry: This node contains value of heuristic expiry.     CCHCachableResponses: This node contains information of cacheable responses such as 302, 200 etc.     CCHTrimCache: This node contains trim cache information.     CCHCache304Stream: This node contains 304 stream information.     CacheOtherDomainCombo:This node contains options of various CacheOtherDomain.     CacheDomainGrid: This node contains data for cache domain grid.         name: This node contains information of name.     rule: This node contains information of rule.     CacheRulesGrid: This node contains data for cache rule grid.     id:This node contains id.     description: This node contains the value of description.     add:This node contains +.     cacheRuleText: This node contains the value of cacheRuleText.     CacheDomainCombo: This node contains data for cache other domain served. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConfigCache**](Config_Cache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET18isCacheRequesttruereqrefreshGet**
> ConfigCacheOtherDomainCombo GET18isCacheRequesttruereqrefreshGet(ctx, )
Cache page data for Domain Combo

This service is used to FETCH  data for Cache page Domain Combo, it uses one API structures.        CacheOtherDomainCombo:This node contains options of various CacheOtherDomain. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConfigCacheOtherDomainCombo**](Config_CacheOtherDomainCombo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST18iAction1iType1Post**
> ConfigCache POST18iAction1iType1Post(ctx, body)
Update global cache settings

 This service is used to update global cache setting.    JSON that needs to be sent as request payload, Should contains given information-        vButtonFlag: 'flag to update Global Cache Settings section(1 to check cache, 2 for clear cache and 3 for save global cache settings).'     CCHMaxCache: Maximum Cache Size.     CCHMinCache: Minimum Cache Size.     CCHHeuristicExpiry: Cache time.     CCHCachableResponses: Cacheable http response code.     CCHTrimCache: Cache checking time.     CCHCache304Stream: Cache fill counter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]UpdateCacheProp**](UpdateCache_prop.md)| Object that needs to be sent to the server | 

### Return type

[**ConfigCache**](Config_Cache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST18iAction2iType1Post**
> ConfigCache POST18iAction2iType1Post(ctx, body)
Add new host

 This service is used to add new host in Apply cache rule grid.    JSON that needs to be sent as request payload, Should contains given information- not required to send json. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]CacheAddRow**](CacheAdd_row.md)| Object that needs to be sent to the server | 

### Return type

[**ConfigCache**](Config_Cache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST18iAction2iType2Post**
> ConfigCache POST18iAction2iType2Post(ctx, body)
Remove host

 This service is used to remove a host from Apply cache domain grid, you have to send id of row, that needs to be removed, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        Id : index (Received as id in GET/18 service) of Cache Domain that needs to be deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveCache**](RemoveCache.md)| Object that needs to be sent to the server | 

### Return type

[**ConfigCache**](Config_Cache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST18iAction2iType3Post**
> ConfigCache POST18iAction2iType3Post(ctx, body)
Update apply cache domain grid

 This service is used to update details in Apply cache domain grid, you have to send id of row with data, that needs to be updated, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        id: index (Received as id in GET/18 service) of CacheDomainGrid that needs to updated.     name: name.     Rule: id(Received as id in GET/18 service, node CacheDomainCombo) of option selected from Caching Rulebase options. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]UpdateApplyCacheProp**](UpdateApplyCache_prop.md)| Object that needs to be sent to the server | 

### Return type

[**ConfigCache**](Config_Cache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST18iAction2iType4Post**
> ConfigCache POST18iAction2iType4Post(ctx, body)
Add Domain(Other Domains Served)

 This service is used Add Domain.    JSON that needs to be sent as request payload, Should contains given information-        name: Domain name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CacheAddDomain**](CacheAddDomain.md)| Object that needs to be sent to the server | 

### Return type

[**ConfigCache**](Config_Cache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST18iAction2iType5Post**
> ConfigCache POST18iAction2iType5Post(ctx, body)
Remove Domain(Other Domains Served)

 This service is used Remove Domain.    JSON that needs to be sent as request payload, Should contains given information-        name: Domain name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CacheAddDomain**](CacheAddDomain.md)| Object that needs to be sent to the server | 

### Return type

[**ConfigCache**](Config_Cache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST18iAction3iType1Post**
> ConfigCache POST18iAction3iType1Post(ctx, body)
Add new cache rule

 This service is used to add new rule in create cache rule grid.    JSON that needs to be sent as request payload, Should contains given information- not required to send json. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]CacheAddRow**](CacheAdd_row.md)| Object that needs to be sent to the server | 

### Return type

[**ConfigCache**](Config_Cache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST18iAction3iType2Post**
> ConfigCache POST18iAction3iType2Post(ctx, body)
Remove cache rule

 This service is used to remove rule from create cache rule grid, you have to send id of row, that needs to be removed, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        id: index(Received as id in GET/18 service) of Cache Rules that needs to be deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveCacheRule**](RemoveCacheRule.md)| Object that needs to be sent to the server | 

### Return type

[**ConfigCache**](Config_Cache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST18iAction3iType3Post**
> ConfigCache POST18iAction3iType3Post(ctx, body)
Update cache rule

 This service is used to update details in create cache rule grid, you have to send id of row with data, that needs to be updated, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        id: index(Received as id in GET/18 service) of CacheRulesGrid that needs to updated.     name: Rule name.     description: Rule description.     add:     cacheRuleText: Rule Condition. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]UpdateRuleCacheProp**](UpdateRuleCache_prop.md)| Object that needs to be sent to the service | 

### Return type

[**ConfigCache**](Config_Cache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

