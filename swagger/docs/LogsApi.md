# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET10Get**](LogsApi.md#GET10Get) | **Get** /GET/10 | Logs page data
[**GET10downloadsyslogiAction2iType1namew3c2014091907LogpreDownloadCheckyesGet**](LogsApi.md#GET10downloadsyslogiAction2iType1namew3c2014091907LogpreDownloadCheckyesGet) | **Get** /GET/10?download&#x3D;syslog&amp;iAction&#x3D;2&amp;iType&#x3D;1&amp;name&#x3D;w3c20140919-07.log&amp;preDownloadCheck&#x3D;yes | Download syslog text file
[**GET10downloadsyslogiAction2iType2namew3c2014091907LogpreDownloadCheckyesGet**](LogsApi.md#GET10downloadsyslogiAction2iType2namew3c2014091907LogpreDownloadCheckyesGet) | **Get** /GET/10?download&#x3D;syslog&amp;iAction&#x3D;2&amp;iType&#x3D;2&amp;name&#x3D;w3c20140919-07.log&amp;preDownloadCheck&#x3D;yes | Download w3c compressed file
[**GET10downloadw3ciAction1iType1namew3c2014091907LogpreDownloadCheckyesGet**](LogsApi.md#GET10downloadw3ciAction1iType1namew3c2014091907LogpreDownloadCheckyesGet) | **Get** /GET/10?download&#x3D;w3c&amp;iAction&#x3D;1&amp;iType&#x3D;1&amp;name&#x3D;w3c20140919-07.log&amp;preDownloadCheck&#x3D;yes | Download w3c text file
[**GET10downloadw3ciAction1iType2namew3c201409107LogpreDownloadCheckyesGet**](LogsApi.md#GET10downloadw3ciAction1iType2namew3c201409107LogpreDownloadCheckyesGet) | **Get** /GET/10?download&#x3D;w3c&amp;iAction&#x3D;1&amp;iType&#x3D;2&amp;name&#x3D;w3c2014091-07.log&amp;preDownloadCheck&#x3D;yes | Download w3c compressed file

# **GET10Get**
> MonitorLogging GET10Get(ctx, )
Logs page data

This service is used to FETCH all data for Logs page, it uses one API structure.    Monitor_Logging: Used to get list of Recent log files generated for both w3clogs and syslogs.    The nodes of JSON received, contains the given information      w3CLoggingList: This node contains list of all w3c logs available.     sysLoggingList: This node contains list of all syslogs available. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MonitorLogging**](Monitor_Logging.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET10downloadsyslogiAction2iType1namew3c2014091907LogpreDownloadCheckyesGet**
> MonitorLoggingDownload GET10downloadsyslogiAction2iType1namew3c2014091907LogpreDownloadCheckyesGet(ctx, )
Download syslog text file

This service is used to downlaod syslog file in text format(unzipped), you needs to send two query string params download=syslog and name=FILENAME. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MonitorLoggingDownload**](Monitor_Logging_Download.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET10downloadsyslogiAction2iType2namew3c2014091907LogpreDownloadCheckyesGet**
> MonitorLoggingDownload GET10downloadsyslogiAction2iType2namew3c2014091907LogpreDownloadCheckyesGet(ctx, )
Download w3c compressed file

This service is used to downlaod syslog file in Compressed format(GZIP), you needs to send two query string params download=w3syslog and name=FILENAME. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MonitorLoggingDownload**](Monitor_Logging_Download.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET10downloadw3ciAction1iType1namew3c2014091907LogpreDownloadCheckyesGet**
> MonitorLoggingDownload GET10downloadw3ciAction1iType1namew3c2014091907LogpreDownloadCheckyesGet(ctx, )
Download w3c text file

This service is used to downlaod w3c log file in text format(unzipped), you needs to send two query string params download=w3c and name=FILENAME. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MonitorLoggingDownload**](Monitor_Logging_Download.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET10downloadw3ciAction1iType2namew3c201409107LogpreDownloadCheckyesGet**
> MonitorLoggingDownload GET10downloadw3ciAction1iType2namew3c201409107LogpreDownloadCheckyesGet(ctx, )
Download w3c compressed file

This service is used to downlaod w3c log file in Compressed format(GZIP), you needs to send two query string params download=w3c and name=FILENAME. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MonitorLoggingDownload**](Monitor_Logging_Download.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

