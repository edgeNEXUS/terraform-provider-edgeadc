# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET30Get**](ClusteringApi.md#GET30Get) | **Get** /GET/30 | Cluster details
[**POST30iAction1iType1Post**](ClusteringApi.md#POST30iAction1iType1Post) | **Post** /POST/30?iAction&#x3D;1&amp;iType&#x3D;1 | Update cluster failover time
[**POST30iAction1iType2Post**](ClusteringApi.md#POST30iAction1iType2Post) | **Post** /POST/30?iAction&#x3D;1&amp;iType&#x3D;2 | Update cluster state
[**POST30iAction2iType1Post**](ClusteringApi.md#POST30iAction2iType1Post) | **Post** /POST/30?iAction&#x3D;2&amp;iType&#x3D;1 | Move up ALB in cluster grid
[**POST30iAction2iType2Post**](ClusteringApi.md#POST30iAction2iType2Post) | **Post** /POST/30?iAction&#x3D;2&amp;iType&#x3D;2 | Move down ALB in cluster grid
[**POST30iAction2iType3Post**](ClusteringApi.md#POST30iAction2iType3Post) | **Post** /POST/30?iAction&#x3D;2&amp;iType&#x3D;3 | Move ALB from cluster to unclaimed
[**POST30iAction2iType4Post**](ClusteringApi.md#POST30iAction2iType4Post) | **Post** /POST/30?iAction&#x3D;2&amp;iType&#x3D;4 | Move ALB from unclaimed to cluster

# **GET30Get**
> SetupCluster GET30Get(ctx, )
Cluster details

This service is used to FETCH cluster details , it uses one API structure.   Setup_Cluster: Used to get cluster details like failover time, cluster members and unclustered servers ready to cluster(Unpartnered).       The nodes of JSON received, contains the given information            clusterState: This node contains the state of cluster such as if it's clustered, manual or standby.      Failover: This node contains value of failover time.      Members: This node contains information of all clustered members.      UnPartenered: This node contains information of ALB that are ready(available) to be clustered.      Clustermember: This node contains cluster member.      ChannelStatusReason: This node contains channel Status Reason.      Status:This node contains the value of status.      Id: This node contains the value of id.      Unclaimeddevices: This node contains the value of unclaimed devices.      dataset:dataset. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SetupCluster**](Setup_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST30iAction1iType1Post**
> SetupCluster POST30iAction1iType1Post(ctx, body)
Update cluster failover time

This service is used to update cluster Failover time(wait time), you have to send Cluster FailOvertimer=(time in miliseconds), with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        FailOvertimer: It's the amount of time an ALB should wait before takeover the cluster if primary ALB failed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FailoverCluster**](FailoverCluster.md)| Object that needs to be sent to the server | 

### Return type

[**SetupCluster**](Setup_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST30iAction1iType2Post**
> SetupCluster POST30iAction1iType2Post(ctx, body)
Update cluster state

This service is used to update cluster state(Cluster, Manual or Standalone), you have to send ClusterState=(1,2 or 3), with POST request json.   1. Cluster   2. Manual   3. Stand-alone           ClusterState: radio box value. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCluster**](UpdateCluster.md)| Object that needs to be sent to the server | 

### Return type

[**SetupCluster**](Setup_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST30iAction2iType1Post**
> SetupCluster POST30iAction2iType1Post(ctx, body)
Move up ALB in cluster grid

This service is used to move up a server in cluster grid, you have to send ip address of server as ipAddr=IP address of server, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        ipAddr: IP address 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoveupCluster**](MoveupCluster.md)| Object that needs to be sent to the server | 

### Return type

[**SetupCluster**](Setup_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST30iAction2iType2Post**
> SetupCluster POST30iAction2iType2Post(ctx, body)
Move down ALB in cluster grid

This service is used to move down a server in cluster grid, you have to send ip address of server as ipAddr=IP address of server, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        ipAddr: IP address 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoveupCluster**](MoveupCluster.md)| Object that needs to be sent to the server | 

### Return type

[**SetupCluster**](Setup_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST30iAction2iType3Post**
> SetupCluster POST30iAction2iType3Post(ctx, body)
Move ALB from cluster to unclaimed

This service is used to move a server from cluster to unclaimed, you have to send ip address of server as ipAddr=IP address of server, with POST request json .    JSON that needs to be sent as request payload, Should contains given information-      ipAddr: IP address 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoveupCluster**](MoveupCluster.md)| Object that needs to be sent to the server | 

### Return type

[**SetupCluster**](Setup_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST30iAction2iType4Post**
> SetupCluster POST30iAction2iType4Post(ctx, body)
Move ALB from unclaimed to cluster

This service is used to move a server from unclaimed to cluster, you  have to send ip address of server as ipAddr=IP address of server, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        ipAddr: IP address 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoveupCluster**](MoveupCluster.md)| Object that needs to be sent to the server | 

### Return type

[**SetupCluster**](Setup_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

