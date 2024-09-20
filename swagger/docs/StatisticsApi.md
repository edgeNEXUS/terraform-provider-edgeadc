# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET14reqrefreshGet**](StatisticsApi.md#GET14reqrefreshGet) | **Get** /GET/14?req&#x3D;refresh | Statistics page data
[**POST14iAction1iType1Post**](StatisticsApi.md#POST14iAction1iType1Post) | **Post** /POST/14?iAction&#x3D;1&amp;iType&#x3D;1 | Reset

# **GET14reqrefreshGet**
> MonitorStatistics2 GET14reqrefreshGet(ctx, )
Statistics page data

This service is used to FETCH all Statistics details, it uses one API structure.    Monitor_Statistics2: Used to get Statistics Information(cache, Compression, hit count memory used etc).    The nodes of JSON received, contains the given information-        GUISTATUS: This node contains information of GUI light.     CPU: This node contains CPU used.     Disk: This node contains DISK used.     Memory: This node contains MEMORY used.     GraphMemory: This node contains graph memory.     FromCacheHits1:      FromCacheHits2:      FromCacheBytes1:     FromCacheBytes2:      FromServerHits1:      FromServerHits2:     FromServerBytes1:      FromServerBytes2:     CacheContentsHits: This node contains cache content hits.     CacheContentsBytes: This node contains bytes of cached content.     RawInputBytes: This node contains raw input bytes.     RawOutputBytes: This node contains raw output bytes.     RawInputBytesSec: This node contains raw input bytes per second.     RawOutputBytesSec: This node contains raw output bytes per second.     ContentCompression: This node contains value of content compression.     UploadContentCompression: This node contains value of upload content compression.     UploadContentUncompressed:     ContentCompressed: This node contains value of content compressed.     ContentUncompressed     OverallCompression: This node contains value of overall compression.     UploadOverallCompression: This node contains value of overall uploaded compression.     InstantaneousPercent: This node contains instantaneous percent.     OverallUncompressed: This node contains value of overall uncompressed.     UploadOverallUncompressed: This node contains value of overall upload uncompressed.     InstantaneousMbpsIn: This node contains Instantaneous Mbps In.     OverallCompressed: This node contains value of overall compressed.     OverallCompressedCache: This node contains value of overall compressed cache.     InstantaneousMbpsOut: This node contains Instantaneous Mbps out.     InstantaneousMbpsCache: This node contains information of Instantaneous Mbps cache.     InstantaneousMbpsTotal: This node contains total of Instantaneous Mbps.     HitCount: This node contains value of Hit Count.     InstantaneousHits: This node contains Instantaneous Hits.     OverallConnections: This node contains Overall Connections.     InstantaneousConnections: This node contains information of Instantaneous Connections.     InstantaneousServerConnections: This node contains Instantaneous Server Connections.     MaxConnections: This node contains Maximum Connections.     CurrentConnections: This node contains information of Current Connections.     PageTitle     CacheContentsBytes1     CacheContentsBytes2     UploadContentCompressed     TotalEntries     AllowedEntries     NewEntries     RevalidatedEntries     ExpiredEntries     DeletedEntries     LowAverageBufFillSize : This node contains information of LowAverageBufFillSize.     MedAverageBufFillSize : This node contains information of MedAverageBufFillSize.     HighAverageBufFillSize : This node contains information of HighAverageBufFillSize.     PercentLowConnections : This node contains information of PercentLowConnections.     PercentMedConnections : This node contains information of PercentMedConnections.     PercentHighConnections : This node contains information of PercentHighConnections.     PercentLowMemory : This node contains information of PercentLowMemory.     PercentMedMemory : This node contains information of PercentMedMemory.     PercentHighMemory : This node contains information of PercentHighMemory.     LowMemoryFree : This node contains information of LowMemoryFree.     MedMemoryUsed : This node contains information of MedMemoryUsed.     MedMemoryFree : This node contains information of MedMemoryFree.     HighMemoryUsed : This node contains information of HighMemoryUsed.     HighMemoryFree : This node contains information of HighMemoryFree.     LowConnectionsUsed : This node contains information of LowConnectionsUsed.     LowConnectionsFree : This node contains information of LowConnectionsFree.     MedConnectionsFree : This node contains information of MedConnectionsFree.     HighConnectionsUsed : This node contains information of HighConnectionsUsed.     HighConnectionsFree : This node contains information of HighConnectionsFree.   

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MonitorStatistics2**](Monitor_Statistics2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST14iAction1iType1Post**
> MonitorStatistics2 POST14iAction1iType1Post(ctx, body)
Reset

This service is used to Reset.    JSON that needs to be sent as request payload, Should contains given information-          reset:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StatisticsReset**](StatisticsReset.md)| Object that needs to be sent to the server | 

### Return type

[**MonitorStatistics2**](Monitor_Statistics2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

