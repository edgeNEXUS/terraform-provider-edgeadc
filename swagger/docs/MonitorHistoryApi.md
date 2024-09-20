# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET37Get**](MonitorHistoryApi.md#GET37Get) | **Get** /GET/37 | History page data
[**POST37iAction1iType2Post**](MonitorHistoryApi.md#POST37iAction1iType2Post) | **Post** /POST/37?iAction&#x3D;1&amp;iType&#x3D;2 | Update user preferences of History

# **GET37Get**
> MonitorHistory GET37Get(ctx, )
History page data

This service is used to FETCH all data for history page, it uses one API structure .    Monitor_History:Used to get History data as per user preferences(i.e. which database and which metrices are selected and how much timespan of data).    The nodes of JSON received, contains the given information        responseError: This node contains error message from RRD.     responseWarning: This node contains warning message.     type: This node contains te type of database selected(such as System, Virtual Service, Real Server).     periodEndNow: This node contains the starting point of time from where we needs to show data.     periodEnd: This node contains the starting point of time from where we needs to show data.     period: This node contains the time scale for which we needs to show data(last minute or hour or day or week etc).     columns: This node contains the metrices selected for which we needs to show data.     axesCount: This node contains number of metrices selected.     timesUtc: This node contains UTC time of each database entry.     timesServer:This node contains UTC time of server .     systemData: This node contains data for each metrices if system is selected as database.     name:      dataset: This node contains collection of all VS and RS available to select and show from RRD database.     VSRS: This node contains collection of all VS and RS.     ipPort: This node contains ip address and port of each VS.     RS: This node contains ip address and port of each RS.     MaxValue: This node contains maximum matrix data value. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MonitorHistory**](Monitor_History.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST37iAction1iType2Post**
> MonitorHistory POST37iAction1iType2Post(ctx, body)
Update user preferences of History

This service is used to update user preferences for history page, you have to send data as database(System, VS and RS) and metrices(CPU, Bytes in etc) selected and timespan(to show data of last minute or hour or day etc), that needs to be updated, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        RequestColumn: Name of metrices for which you needs to show graph data.      RequestVsrs: id(Received as id in GET/19 service, node VSRS) of option selected from VS/RS options.     Period: option selected from Period options such as minute, hour, day, week etc.     PeriodEndNow: Number of periods in minutes or days(for ex- data for last 20 minutes)     type: Type of database selected(from System, VS and RS).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateMonitorHistory**](UpdateMonitorHistory.md)| Object that needs to be sent to the server | 

### Return type

[**MonitorHistory**](Monitor_History.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

