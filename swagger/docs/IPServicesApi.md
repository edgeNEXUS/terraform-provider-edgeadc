# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET29Get**](IPServicesApi.md#GET29Get) | **Get** /GET/29 | IP Services page combo options
[**GET9isPageLoadtrueGet**](IPServicesApi.md#GET9isPageLoadtrueGet) | **Get** /GET/9?isPageLoad&#x3D;true | IP Services page data
[**POST9iAction2iType1FilterKeywordPost**](IPServicesApi.md#POST9iAction2iType1FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;2&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Update Virtual Service
[**POST9iAction2iType2FilterKeywordPost**](IPServicesApi.md#POST9iAction2iType2FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;2&amp;iType&#x3D;2&amp;FilterKeyword&#x3D; | Update Real Server
[**POST9iAction2iType3FilterKeywordPost**](IPServicesApi.md#POST9iAction2iType3FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;2&amp;iType&#x3D;3&amp;FilterKeyword&#x3D; | Update Basic Tab Information
[**POST9iAction2iType4FilterKeywordPost**](IPServicesApi.md#POST9iAction2iType4FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;2&amp;iType&#x3D;4&amp;FilterKeyword&#x3D; | Update Advance Tab Information
[**POST9iAction3iType1FilterKeywordPost**](IPServicesApi.md#POST9iAction3iType1FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;3&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Add/Copy Virtual Service
[**POST9iAction3iType2FilterKeywordPost**](IPServicesApi.md#POST9iAction3iType2FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;3&amp;iType&#x3D;2&amp;FilterKeyword&#x3D; | Add Port
[**POST9iAction3iType3FilterKeywordPost**](IPServicesApi.md#POST9iAction3iType3FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;3&amp;iType&#x3D;3&amp;FilterKeyword&#x3D; | Add/Copy Real Server
[**POST9iAction3iType4FilterKeywordPost**](IPServicesApi.md#POST9iAction3iType4FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;3&amp;iType&#x3D;4&amp;FilterKeyword&#x3D; | Remove Virtual Service
[**POST9iAction3iType5FilterKeywordPost**](IPServicesApi.md#POST9iAction3iType5FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;3&amp;iType&#x3D;5&amp;FilterKeyword&#x3D; | Remove Real Server
[**POST9iAction4iType1FilterKeywordPost**](IPServicesApi.md#POST9iAction4iType1FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;4&amp;iType&#x3D;1&amp;FilterKeyword&#x3D; | Move flightPATH Available to Applied
[**POST9iAction4iType2FilterKeywordPost**](IPServicesApi.md#POST9iAction4iType2FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;4&amp;iType&#x3D;2&amp;FilterKeyword&#x3D; | Move flightPATH Applied to Available
[**POST9iAction4iType3FilterKeywordPost**](IPServicesApi.md#POST9iAction4iType3FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;4&amp;iType&#x3D;3&amp;FilterKeyword&#x3D; | Move flightPATH UP in Applied
[**POST9iAction4iType4FilterKeywordPost**](IPServicesApi.md#POST9iAction4iType4FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;4&amp;iType&#x3D;4&amp;FilterKeyword&#x3D; | Move flightPATH DOWN in Applied
[**POST9iAction4iType5FilterKeywordPost**](IPServicesApi.md#POST9iAction4iType5FilterKeywordPost) | **Post** /POST/9?iAction&#x3D;4&amp;iType&#x3D;5&amp;FilterKeyword&#x3D; | DragDrop flightPATH in Applied grid

# **GET29Get**
> IpServicescombo GET29Get(ctx, )
IP Services page combo options

This service is used to FETCH all data for Combo boxes of IP services page, we created a separate services to get combo box values because we need them just once when page loaded i. e. They are not changed on each update of IP services page, it uses one API structure    Setup_IPServices: Used to get all combo box(drop down) values used in IP services page.      JSON returned, Contains given information-        ServiceTypeCombo: This node contains options of all available service types.      SSLCertCombo: This node contains options of all available SSL Certificates that can be applied on VS.      CacheRuleCombo: This node contains options of all cache types.      MonitorPolicyCombo: This node contains options of all Monitoring policy available, you can create or upload custom monitor in Real server monitor page.     SSLClientCertCombo: This node contains options of all available SSL Certificates that can be applied on a Real server.      LBPolicyCombo: This node contains options of all Load Balancing policy      LBPolicyFTPCombo: This node contains options of all Load Balancing policy when service type selected is FTP.      ConnectivityCombo: This node contains options of connectivity that can be selected for a VS.      CSActivityCombo: This node contains options of all activities that can be selected for a real server.      AccelerateCombo: This node contains all acceleration options.      FlightPathSelectionList: This node contains all selected flightpaths.      SSLCipherListString: This node contains options of all Cipher.     SNICertificateListString: This node contains options of all SNI Certificate.     SSLRenegotiationString: This node contains options of all SSL Renegotiation.     SecurityLogString: This node contains options of all Security Log.     isNoVIP: 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IpServicescombo**](IPServicescombo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET9isPageLoadtrueGet**
> IpServices GET9isPageLoadtrueGet(ctx, )
IP Services page data

This service is used to FETCH all data for IP services page, it uses two API structures    Setup_IPServices: Used to get all virtual services and thier real server's inforamtion.    Setup_Cluster: Used to get information, if ALB is clustered with other ALB.    JSON returned, Contains given information-         ipService: This node contains information of a Virtual Services, their real server and other settings.        sId: This node contains information of index of each VS.        InterfaceID: This node contains index of each interface(VIP) which needed to update or delete that VIP.        ChannelID: This node contains index of each channel(VS) which needed to update or delete that VS.        ipAddr: This node contains IP Address.        subnetMask: This node contains subnet mask.        serviceName: This node contains service name.        localPortEnabledChecked:  This node contains true if VS is enabled.         interfaceImagePath:This node contains image path for VIP to show status lights.        interfaceStatusReason: This node contains status reason to show as tooltip of status image for VIP.        channelImagePath: This node contains image path for VS to show status lights.        channelStatusReason: This node contains status reason to show as tooltip of status image for VS.        primaryChecked: This node contains \"active/passive\" if ALB is clustered otherwise true if check for manual mode and standalone for standalone mode in clutering.        serviceType: This node contains service type.        maxConn: This node contains maximum connection allowed.        contentServerGroupName: This node contains real server group name.        enableConnectionPool: This node contains true if connection pool is enabled.        connectionPoolSize: This node contains connection pool size.        serverMonitoring: This node contains monitoring policy.        loadBalancingPolicy: This node contains load balancing policy.        connectivity: This node contains connectivity used.        acceleration: This node contains accelaration.        sslCertificate: This node contains server SSL certificate.        sslClientCertificate: This node contains client ssl certificate.        cachingRule: This node contains caching strategy.        monitoringInterval: This node contains monitoring interval.        monitoringTimeout: This node contains monitoring timeout.        ConnectionTimeout: This node contains connection timeout.        InCount: This node contains In Count.        offlinonfailure: This node contains information of Switch to offline on failure.       OutCount: This node contains Out Count.        sslRenegotiation: This node contains on/off value.        SecurityLog: This node contains on/off value.        CipherName: This node contains Cipher Name.        flightPath: This node contains list of all flightpath.        fId: This node contains information of index of each flightpath.        flightPathSelected: If this node contains 0 then it's not applied to this VS, otherwise it shows priority of applied flightpath on that VS.        Name: This node contains name of flightpath.        contentServer: This node contains list of all real servers.        cId: This node contains information of index of each real server.        imagePath: This node contains image path for real server to show status lights.        statusReason: This node contains status reason to show as tooltip of status image for RS.       CSActivity:         CSIPAddr: This node contains IP Address of real server.        CSPort: This node contains port of real server.        CSNotes: This node contains notes of real server.        WeightFactor: This node contains weight factor of real server.       WeightFactorCalculated: This node contains calculated weight factor of real server.       ServerKey: This node contains config ID.       FilterKeyword: FilterKeyword.       InterfaceKey:InterfaceKey.       ChannelKey:ChannelKey.       MaskState:MaskState.       SNIDefaultCertificateName:SNIDefaultCertificateName.       CServerId:CServerId. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IpServices**](IPServices.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction2iType1FilterKeywordPost**
> IpServicesPost POST9iAction2iType1FilterKeywordPost(ctx, body)
Update Virtual Service

This service is used to UPDATE Virtual Service, In case of update after COPY request send \"CopyVIP\":\"1\" (It'll copy all real server and settings of a VS to a new VS), with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CopyVIP: set to 1 if it's update request after copy VIP, otherwise 0.      editedInterface: index(Received as InterfaceID in GET/9 service) of interface that needs to updated.      editedChannel: index(Received as ChannelID in GET/9 service) of channel that needs to updated.      IpAddr: IP address of VS.      localPortEnabledChecked: if enabled checked for VS then send 'true' otherwise 'false'.      primaryChecked: send 'true' to set primary, if manual mode in clustering.      interfaceImagePath: interface image path.      port: PORT of VS.      serviceName: Service name.      serviceType: id(Received as id in GET/29 service, node ServiceTypeCombo) of option selected from service type options.      subnetMask: Subnet mask. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CopyIp**](CopyIp.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction2iType2FilterKeywordPost**
> IpServicesPost POST9iAction2iType2FilterKeywordPost(ctx, body)
Update Real Server

This service is used to update Real Server, you have to send Interface id and channel id and server id with data, that needs to be updated, with POST request json.     JSON that needs to be sent as request payload, Should contains given information         editedInterface: index(Received as InterfaceID in GET/9 service) of interface that's real server needs to be updated.      editedChannel: index(Received as ChannelID in GET/9 service) of channel that's real server needs to be updated.      cId: index(Received as cId in GET/9 service) of server that needs to be updated.      CSActivity: Activity of real server(online or offline etc)      CSIPAddr: IP address of RS.      WeightFactor: Weight Factor of RS.      imagePath: imagePath of RS.      statusReason: statusReason of RS.      contentServerGroupName: group name of RS.      CSPort: PORT of RS.      CSNotes: Server notes.      ServerId: ServerId. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateServer**](UpdateServer.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction2iType3FilterKeywordPost**
> IpServicesPost POST9iAction2iType3FilterKeywordPost(ctx, body)
Update Basic Tab Information

This service is used to update Basic Tab Information, you have to send Interface id and channel id with data, that needs to be updated, with POST request json.    JSON that needs to be sent as request payload, Should contains given information        editedInterface: index(Received as InterfaceID in GET/9 service) of interface that's Basic information needs to be updated      editedChannel: index(Received as ChannelID in GET/9 service) of channel that's Basic information needs to be updated.      serverMonitoring: id(Received as id in GET/29 service, node MonitorPolicyCombo) of option selected from Server monitoring options.      acceleration: id(Received as id in GET/29 service, node AccelerateCombo) of option selected from accelaration options.      loadBalancingPolicy: id(Received as id in GET/29 service, node LBPolicyCombo) of option selected from Load balancing options.      sslCertificate: id(Received as id in GET/29 service, node SSLCertCombo) of option selected from VS certificate options.      sslClientCertificate: id(Received as id in GET/29 service, node SSLClientCertCombo) of option selected from RS certificate options.      cachingRule: id(Received as id in GET/29 service, node CacheRuleCombo) of option selected from Caching options.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBasicTab**](UpdateBasicTab.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction2iType4FilterKeywordPost**
> IpServicesPost POST9iAction2iType4FilterKeywordPost(ctx, body)
Update Advance Tab Information

This service is used to update Advance Tab Information, you have to send Interface id and channel id with data, that needs to be updated, with POST request json.    JSON that needs to be sent as request payload, Should contains given information        editedInterface: index(Received as InterfaceID in GET/9 service) of interface that's advance information needs to be updated.      editedChannel: index(Received as ChannelID in GET/9 service) of channel that's advance information needs to be updated.      connectivity: id(Received as id in GET/29 service, node ConnectivityCombo) of option selected from Connectivity options.      SNIDefaultCertificateName: send 'true' if enabling connection pool.      maxConn: Maximum connection.      sslResumption:      ConnectionTimeout:      monitoringInterval:      monitoringTimeout:      InCount:      OutCount:      CipherName: id(Received as id in GET/29 service, node SSLCipherListString) of option selected from Cipher options.      securityLog:id(Received as id in GET/29 service, node SecurityLogString).      sslRenegotiation:     offlinonfailure: This node contains information of Switch to offline on failure. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateAdvanceTab**](UpdateAdvanceTab.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction3iType1FilterKeywordPost**
> IpServicesPost POST9iAction3iType1FilterKeywordPost(ctx, body)
Add/Copy Virtual Service

This service is used to ADD/Copy Virtual Service. JSON that needs to be sent as request payload, Should contains given information-        CopyVIP: set to 1 if it's update request after copy VIP, otherwise 0.      editedInterface: index(Received as InterfaceID in GET/9 service) of interface that needs to add/copy.      editedChannel: index(Received as ChannelID in GET/9 service) of channel that needs to copy.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CopyIpAddress**](CopyIpAddress.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction3iType2FilterKeywordPost**
> IpServicesPost POST9iAction3iType2FilterKeywordPost(ctx, body)
Add Port

This service is used to Add PORT, you have to send Interface id and channel id , to which real server needs to be added, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        editedInterface: index(Received as InterfaceID in GET/9 service) of interface to which real server needs to be added.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddServer**](AddServer.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction3iType3FilterKeywordPost**
> IpServicesPost POST9iAction3iType3FilterKeywordPost(ctx, body)
Add/Copy Real Server

This service is used to Add/Copy Real Server to a VS, you have to send Interface id and channel id , to which real server needs to be added, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        editedInterface: index(Received as InterfaceID in GET/9 service) of interface to which real server needs to be added.      editedChannel: index(Received as ChannelID in GET/9 service) of channel to which real server needs to be added. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddServer**](AddServer.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction3iType4FilterKeywordPost**
> IpServicesPost POST9iAction3iType4FilterKeywordPost(ctx, body, filterKeyword)
Remove Virtual Service

This service is used to REMOVE Virtual Service, you have to send Interface id and channel id , that needs to be removed, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        editedInterface: index(Received as InterfaceID in GET/9 service) of interface that needs to removed.     editedChannel: index(Received as ChannelID in GET/9 service) of channel that needs to removed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveIp**](RemoveIp.md)| Json data that needs to be sent to the server. | 
  **filterKeyword** | **string**|  | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction3iType5FilterKeywordPost**
> IpServicesPost POST9iAction3iType5FilterKeywordPost(ctx, body)
Remove Real Server

This service is used to remove Real Server, you have to send Interface id and channel id and server id with data, that needs to be remove, with POST request json.    JSON that needs to be sent as request payload, Should contains given information        editedInterface: index(Received as InterfaceID in GET/9 service) of interface that's real server needs to be removed.      editedChannel: index(Received as ChannelID in GET/9 service) of channel that's real server needs to be removed.      cId: index(Received as cId in GET/9 service) of server that needs to be removed.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveServer**](RemoveServer.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction4iType1FilterKeywordPost**
> IpServicesPost POST9iAction4iType1FilterKeywordPost(ctx, body)
Move flightPATH Available to Applied

This service is used to move flightPATH Available to Applied in IP services page, you have to send Interface id and channel id with data, that needs to be updated, with POST request json. JSON that needs to be sent as request payload, Should contains given information       editedInterface: index(Received as InterfaceID in GET/9 service) of interface that's flightPATH information needs to be updated.      editedChannel: index(Received as ChannelID in GET/9 service) of channel that's flightPATH information needs to be updated.      flightPathName: name of flightpath.      flightPathDropId:      flightPathDropName:      position:flightpath rule drop position.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoveFlightToSelected**](MoveFlightToSelected.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction4iType2FilterKeywordPost**
> IpServicesPost POST9iAction4iType2FilterKeywordPost(ctx, body)
Move flightPATH Applied to Available

This service is used to move flightPATH Applied to Available in IP services page, you have to send Interface id and channel id with data, that needs to be updated, with POST request json.     JSON that needs to be sent as request payload, Should contains given information         editedInterface: index(Received as InterfaceID in GET/9 service) of interface that's flightPATH information needs to be updated.      editedChannel: index(Received as ChannelID in GET/9 service) of channel that's flightPATH information needs to be updated.      flightPathName: name of flightpath.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoveFlightToAvailable**](MoveFlightToAvailable.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction4iType3FilterKeywordPost**
> IpServicesPost POST9iAction4iType3FilterKeywordPost(ctx, body)
Move flightPATH UP in Applied

This service is used to move flightPATH Up (change Applied grid Data index) in flightPATH Tab of IP services page, you have to send Interface id and channel id with data, that needs to be updated, with POST request json.    JSON that needs to be sent as request payload, Should contains given information        editedInterface: index(Received as InterfaceID in GET/9 service) of interface that's flightPATH information needs to be updated.      editedChannel: index(Received as ChannelID in GET/9 service) of channel that's flightPATH information needs to be updated.      flightPathName: name of flightpath.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoveFlightUp**](MoveFlightUp.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction4iType4FilterKeywordPost**
> IpServicesPost POST9iAction4iType4FilterKeywordPost(ctx, body)
Move flightPATH DOWN in Applied

This service is used to move flightPATH down (change Applied grid Data index) in flightPATH Tab of IP services page, you have to send Interface id and channel id with data, that needs to be updated, with POST request json.    JSON that needs to be sent as request payload, Should contains given information        editedInterface: index(Received as InterfaceID in GET/9 service) of interface that's flightPATH information needs to be updated.      editedChannel: index(Received as ChannelID in GET/9 service) of channel that's flightPATH information needs to be updated.      flightPathName: name of flightpath.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoveFlightDown**](MoveFlightDown.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST9iAction4iType5FilterKeywordPost**
> IpServicesPost POST9iAction4iType5FilterKeywordPost(ctx, body)
DragDrop flightPATH in Applied grid

This service is used to DragDrop flightPATH in Applied grid , you have to send Interface id and channel id with data, that needs to be updated, with POST request json. JSON that needs to be sent as request payload, Should contains given information       editedInterface: index(Received as InterfaceID in GET/9 service) of interface that's flightPATH information needs to be updated.      editedChannel: index(Received as ChannelID in GET/9 service) of channel that's flightPATH information needs to be updated.      flightPathDragName:       flightPathDropId:      flightPathDropName:      flightPathDragId:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DragDropFlight**](DragDropFlight.md)| Json data that needs to be sent to the server. | 

### Return type

[**IpServicesPost**](IPServicesPost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

