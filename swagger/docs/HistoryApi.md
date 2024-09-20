# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET32Get**](HistoryApi.md#GET32Get) | **Get** /GET/32 | History page data
[**POST31iAction1iType2Post**](HistoryApi.md#POST31iAction1iType2Post) | **Post** /POST/31?iAction&#x3D;1&amp;iType&#x3D;2 | Update RRD status and time interval
[**POST31iAction1iType4Post**](HistoryApi.md#POST31iAction1iType4Post) | **Post** /POST/31?iAction&#x3D;1&amp;iType&#x3D;4 | Save backup
[**POST31iAction1iType5Post**](HistoryApi.md#POST31iAction1iType5Post) | **Post** /POST/31?iAction&#x3D;1&amp;iType&#x3D;5 | Delete RRD database
[**POST31iAction1iType6Post**](HistoryApi.md#POST31iAction1iType6Post) | **Post** /POST/31?iAction&#x3D;1&amp;iType&#x3D;6 | Restore RRD database
[**POST31iAction1iType7Post**](HistoryApi.md#POST31iAction1iType7Post) | **Post** /POST/31?iAction&#x3D;1&amp;iType&#x3D;7 | Save backup

# **GET32Get**
> HistoryData GET32Get(ctx, )
History page data

This service is used to FETCH History settings(RRD database enabled, backup and restore etc) data, it uses one API structure.    RRDtool: Used to get Backup, restore and other information of current RRD database.    The nodes of JSON received, contains the given information-      CRRD_Enabled: This node contains information that if RRD database enabled.     CRRD_Interval: This node contains interval for RRD.     CRRD_Last: This node contains last date of RRD.     CRRD_LastRed: This node contains last error date.     CRRD_LastMsg: This node contains last error messsage.     CRRD_BackupName: This node contains the backup options.     CRRD_DeleteName: This node contains the delete options.     CRRD_RestoreName: This node contains restore options.     CRRD_HasCurrent: This node contains information of if RRD has current database.     dataset:dataset. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HistoryData**](History_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST31iAction1iType2Post**
> HistoryData POST31iAction1iType2Post(ctx, body)
Update RRD status and time interval

This service is used to update RRD status and time interval, you have to send CRRD_Enabled=('ON' to enable otherwise 'OFF') and CRRD_Interval=(time in seconds 1 to 60) , with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CRRD_Enabled: set CRRD_Enabled='ON' to enable otherwise 'OFF'.     CRRD_Interval: CRRD Interval. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RrdUpdate**](RrdUpdate.md)| Object that needs to be sent to the server | 

### Return type

[**HistoryData**](History_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST31iAction1iType4Post**
> HistoryData POST31iAction1iType4Post(ctx, body)
Save backup

This service is used to save backup of RRD database, you have to send CRRD_BackupName=(select from available databases), with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CRRD_BackupName: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RrdSave**](RrdSave.md)| Object that needs to be sent to the server | 

### Return type

[**HistoryData**](History_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST31iAction1iType5Post**
> HistoryData POST31iAction1iType5Post(ctx, body)
Delete RRD database

This service is used to delete a RRD database, you have to send CRRD_DeleteName=(select from available databases), with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CRRD_DeleteName : Database name for delete. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RrdDelete**](RrdDelete.md)| Object that needs to be sent to the server | 

### Return type

[**HistoryData**](History_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST31iAction1iType6Post**
> HistoryData POST31iAction1iType6Post(ctx, body)
Restore RRD database

This service is used to restore a RRD database, you have to send CRRD_RestoreName=(select from available databases), with POST request json.    JSON that needs to be sent as request payload, Should contains given information-       CRRD_RestoreName : Restore database name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RrdRestore**](RrdRestore.md)| Object that needs to be sent to the server | 

### Return type

[**HistoryData**](History_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST31iAction1iType7Post**
> HistoryData POST31iAction1iType7Post(ctx, body)
Save backup

This service is used to save backup of RRD database, you have to send CRRD_BackupName=(select from available databases), with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CRRD_BackupName: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RrdSave**](RrdSave.md)| Object that needs to be sent to the server | 

### Return type

[**HistoryData**](History_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

