# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET1Get**](NetworkApi.md#GET1Get) | **Get** /GET/1 | Appliance data
[**GET3Get**](NetworkApi.md#GET3Get) | **Get** /GET/3 | Network page combo
[**GET5Get**](NetworkApi.md#GET5Get) | **Get** /GET/5 | Hardware data
[**GET6Get**](NetworkApi.md#GET6Get) | **Get** /GET/6 | Network page data
[**POST1iAction2iType1Post**](NetworkApi.md#POST1iAction2iType1Post) | **Post** /POST/1?iAction&#x3D;2&amp;iType&#x3D;1 | Remove adapter
[**POST1iAction3iType1Post**](NetworkApi.md#POST1iAction3iType1Post) | **Post** /POST/1?iAction&#x3D;3&amp;iType&#x3D;1 | Update adapter on specific conditions
[**POST1iAction4iType1Post**](NetworkApi.md#POST1iAction4iType1Post) | **Post** /POST/1?iAction&#x3D;4&amp;iType&#x3D;1 | Add or update adapter
[**POST5iAction1iType1Post**](NetworkApi.md#POST5iAction1iType1Post) | **Post** /POST/5?iAction&#x3D;1&amp;iType&#x3D;1 | Update interface grid
[**POST5iAction2iType1Post**](NetworkApi.md#POST5iAction2iType1Post) | **Post** /POST/5?iAction&#x3D;2&amp;iType&#x3D;1 | Update bonding grid
[**POST5iAction3iType1Post**](NetworkApi.md#POST5iAction3iType1Post) | **Post** /POST/5?iAction&#x3D;3&amp;iType&#x3D;1 | Remove hardware bonding
[**POST5iAction4iType1Post**](NetworkApi.md#POST5iAction4iType1Post) | **Post** /POST/5?iAction&#x3D;4&amp;iType&#x3D;1 | Add bonding
[**POST6iAction1iType1Post**](NetworkApi.md#POST6iAction1iType1Post) | **Post** /POST/6?iAction&#x3D;1&amp;iType&#x3D;1 | Update static route
[**POST6iAction2iType1Post**](NetworkApi.md#POST6iAction2iType1Post) | **Post** /POST/6?iAction&#x3D;2&amp;iType&#x3D;1 | Delete static route
[**POST6iAction3iType1Post**](NetworkApi.md#POST6iAction3iType1Post) | **Post** /POST/6?iAction&#x3D;3&amp;iType&#x3D;1 | Update default gateway, ALB name and DNS
[**POST6iAction4iType1Post**](NetworkApi.md#POST6iAction4iType1Post) | **Post** /POST/6?iAction&#x3D;4&amp;iType&#x3D;1 | Advanced Network
[**POST6iAction5iType1Post**](NetworkApi.md#POST6iAction5iType1Post) | **Post** /POST/6?iAction&#x3D;5&amp;iType&#x3D;1 | Update SNAT
[**POST6iAction6iType1Post**](NetworkApi.md#POST6iAction6iType1Post) | **Post** /POST/6?iAction&#x3D;6&amp;iType&#x3D;1 | Delete snat
[**POST6iAction7iType1Post**](NetworkApi.md#POST6iAction7iType1Post) | **Post** /POST/6?iAction&#x3D;7&amp;iType&#x3D;1 | Add SNAT

# **GET1Get**
> ApplianceData GET1Get(ctx, )
Appliance data

This service is used to FETCH adapter details grid rows details, it uses two API structure.    Setup_ALB_Addresses_New: Used to get adapter details grid rows details.    Setup_FailOver: Used to get Failover enabled details.    The nodes of JSON received, contains the given information          AdapterListString: This node contains options of adapters.       FailoverEnabled: This node contains \"ON\" if failover enabled.       FailoverTimer: This node contains value of failover timer.       ApplianceGrid: This node contains data of appliance grid.       id: This node contains unique id for each appliance grid row(index), that used for update and delete of row.       ethtype: This node contains eth name.       address: This node contains address of appliance.       mask: This node contains subnet mask.       gateway: This node contains gateway.       desc: This node contains description of appliance.       RestEnabled: This node contains '1' if REST services is enabled on this appliance.       webconsoleChecked: This node contains '1' if appliance address is current ALB address.       ethVlan:       RpFilter: 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApplianceData**](Appliance_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET3Get**
> NetworkCombo GET3Get(ctx, )
Network page combo

This service is used to FETCH hardware page combo box options list, it uses three API structure.    Config_Ethernet: Used to get mode(Full duplex, half duplex etc), bond and speed list strings.    Config_Bonding: Used to get ethernet bonding list string.    Setup_ALB_Addresses_New: Used to get VLAN available or not.      The nodes of JSON received, contains the given information        EthDuplexListString: This node contains options of ethernet duplex.     HideVlan: This node contains 'true' if hide VLAN enabled.     EthSpeedListString: This node contains options of ethernet speed.     BondModeListString: This node contains options of bond modes.     NoBonding: This node contains 'true' if 'no' bonding.     EthBondListString: This node contains options of bonds available. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NetworkCombo**](Network_Combo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET5Get**
> HardwareData GET5Get(ctx, )
Hardware data

This service is used to FETCH hardware and bond grid rows details, it uses two API structure.    Config_Etherne: Used to get hardware interfaces grid rows details.    Config_Bonding: Used to get bonding grid rows details.      The nodes of JSON received, contains the given information        HardwareGrid: This node contains data for hardware grid.     id: This node contains unique id for each hardware row(index), that used for update and delete of row.     duplex: This node contains value of duplex.     speed: This node contains value of speed.     ethName: This node contains ethernet name.     ethBond: This node contains ethernet bond.     BondGrid: This node contains data of bond grid.     ethStatus:     id: This node contains unique id for each bond row(index), that used for update and delete of row.     bondName: This node contains bond name.     bondMode: This node contains bond mode.     AdapterListString: This node contains options of adapters. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HardwareData**](Hardware_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET6Get**
> NetworkData GET6Get(ctx, )
Network page data

This service is used to FETCH Network details, it uses four API structure.    Setup_Network: Used to get Default gateway and static route details.    Setup_Nagle: Used to get Server nagle and client nagle details.    Setup_ServerRef: Used to get ALB name.    Setup_DNS_Server: Used to get DNS server details.    The nodes of JSON received, contains the given information-      AdapterOptionsListString: This node contains adapters options.     DefaultGateway: This node contains default gateway.     DefaultGatewayIPv6:      ServerNagle: This node contains Server Nagle.     ClientNagle: This node contains Client Nagle.     ServerRefField: This node contains ALB name.     DNSServerField: This node contains DNS server '1'.     DNSServerField2: This node contains DNS server 2.     DefaultStatusImage: This node contains status image of gateway.     NetworkStaticRoutesGrid: This node contains Network Static Routes Grid data.     CurrentTable: This node contains Current Table data.     SNATGrid: Used to get SNAT Grid details.     Id: Used to get Id details.     Interface: This node contains the value of Interface.     Sourceip: This node contains the value of Sourceip.     Destinationip: This node contains the value od Destinationip.     protocol: Used to get Protocol details.     SourcePort: Used to get Source Port details.     Destinationport: Used to get Destination port details.     Snattoip: Used to get Snat to ip details.     Snattoport: This node contains the value of Snat to port .     Notes: Used to get Nodes details.     Error: This node contains the value of Error.     ProtocolCombo: This node contains options for SNAT protocol. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NetworkData**](Network_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST1iAction2iType1Post**
> ApplianceData POST1iAction2iType1Post(ctx, body)
Remove adapter

This service is used to remove adapter grid details, you have to send adapter grid row id, with POST request json.    JSON that needs to be sent as request payload, Should contains given information- not required to send json.        id: index(Received as id in GET/1 service) of Appliance that needs to delete 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveAdapter**](RemoveAdapter.md)| Object that needs to be sent to the server | 

### Return type

[**ApplianceData**](Appliance_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST1iAction3iType1Post**
> ApplianceData POST1iAction3iType1Post(ctx, body)
Update adapter on specific conditions

This service is used to update row in adapter grid details, you have to send adapter row data with row id, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-     ethName:      address:   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SpecificAdapterUpdate**](SpecificAdapterUpdate.md)| Object that needs to be sent to the server | 

### Return type

[**ApplianceData**](Appliance_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST1iAction4iType1Post**
> ApplianceData POST1iAction4iType1Post(ctx, body)
Add or update adapter

This service is used to update or add new row in adapter grid details, you have to send adapter row data with row id, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        RestEnabled:      address: Appliance IP address.      desc: Appliance IP address.      gateway: Appliance gateway.      ethVlan: Appliance VLAN address.      ethtype: id(Received as id in GET/1 service, node AdapterListString) of option selected from ethernet options.      id: index(Received as id in GET/1 service) of Appliance that needs to updated(if add request then send 'maxid+1').      mask: Appliance subnet mask.      webconsoleChecked: set \"true\" if you want to change current ALB address to IP address of this Appliance.     RpFilter:   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]AddAdapterOpt**](Add_adapter_opt.md)| Object that needs to be sent to the server | 

### Return type

[**ApplianceData**](Appliance_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST5iAction1iType1Post**
> HardwareData POST5iAction1iType1Post(ctx, body)
Update interface grid

This service is used to update hardware interfaces details, you have to send hardware row data with row id, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        id: index(Received as id in GET/5 service) of HardwareGrid that needs to updated.      Duplex: id(Received as id in GET/3 service, node EthDuplexListString) of option selected from Duplex options.      Speed: id(Received as id in GET/3 service, node EthSpeedListString) of option selected from Speed options.      EthName: ethernet name.      ethBond: id(Received as id in GET/3 service, node EthBondListString) of option selected from Bonding options. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateInterface**](UpdateInterface.md)| Object that needs to be sent to the server | 

### Return type

[**HardwareData**](Hardware_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST5iAction2iType1Post**
> HardwareData POST5iAction2iType1Post(ctx, body)
Update bonding grid

This service is used to update hardware bonding details, you have to send hardware bonding row data(bondMode) with row id, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        bondMode: id(Received as id in GET/3 service, node BondModeListString) of option selected from Bond Mode options.      id: index(Received as id in GET/5 service) of Bond that needs to updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBonding**](UpdateBonding.md)| Object that needs to be sent to the server | 

### Return type

[**HardwareData**](Hardware_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST5iAction3iType1Post**
> HardwareData POST5iAction3iType1Post(ctx, body)
Remove hardware bonding

This service is used to remove hardware bonding details, you have to send hardware bonding grid row id, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        id: index(Received as id in GET/5 service) of Bond that needs to delete. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveBonding**](RemoveBonding.md)| Object that needs to be sent to the server | 

### Return type

[**HardwareData**](Hardware_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST5iAction4iType1Post**
> HardwareData POST5iAction4iType1Post(ctx, )
Add bonding

This service is used to add new row to hardware bonding grid.    JSON that needs to be sent as request payload, Should contains given information- not required to send json. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HardwareData**](Hardware_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST6iAction1iType1Post**
> NetworkData POST6iAction1iType1Post(ctx, body)
Update static route

This service is used to update network static route details, you have to send static route details data with row id, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        destination: destination IP address.     gateway: route gatway.     mask: static route mask.     adapter: id(Received as id in GET/16 service, node AdapterOptionsListString) of option selected from Adapter options.     id: index(Received as id in GET/6 service) of Static Routes that needs to delete.     isValidationRequired:     active: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateStatic**](UpdateStatic.md)| Object that needs to be sent to the server | 

### Return type

[**NetworkData**](Network_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST6iAction2iType1Post**
> NetworkData POST6iAction2iType1Post(ctx, body)
Delete static route

This service is used to delete network static route details, you have to send static route details row id that needs to be deleted, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-      id: index(Received as id in GET/6 service) of Static Routes that needs to delete. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteStatic**](DeleteStatic.md)| Object that needs to be sent to the server | 

### Return type

[**NetworkData**](Network_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST6iAction3iType1Post**
> NetworkData POST6iAction3iType1Post(ctx, body)
Update default gateway, ALB name and DNS

This service is used to update default gateway, ALB name and DNS server address details, you have to send data, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        DNSServerField: DSN server '1' address.     DNSServerField2: DSN server 2 address.     DefaultGateway: Defaults gateway.     ServerRefField: ALB Name.     DefaultGatewayIPv6:     isValidationRequired: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateDefault**](UpdateDefault.md)| Object that needs to be sent to the server | 

### Return type

[**NetworkData**](Network_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST6iAction4iType1Post**
> NetworkData POST6iAction4iType1Post(ctx, body)
Advanced Network

This service is used to update Advanced Network setting, you have to send POST request json.    JSON that needs to be sent as request payload, Should contains given information-        ClientNagle: 'true/false' value.      ServerNagle: 'true/false' value. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdvanceNetwork**](AdvanceNetwork.md)| Object that needs to be sent to the server | 

### Return type

[**NetworkData**](Network_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST6iAction5iType1Post**
> NetworkData POST6iAction5iType1Post(ctx, body)
Update SNAT

This service is used to update network SNAT details, you have to send SNAT details, with POST request json.    JSON that needs to be sent as request payload, Should contain given information-        destinationport: destination port.     interface: interface.     id: This node contains row id .     notes: snat notes.     protocol: id of option selected from protocol options.     snattoip: snattoip.     snattoport: snattoport.     sourceip: sourceip.     sourceport: sourceport.     destinationip: destinationip 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSnat**](UpdateSnat.md)| Object that needs to be sent to the server | 

### Return type

[**NetworkData**](Network_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST6iAction6iType1Post**
> NetworkData POST6iAction6iType1Post(ctx, body)
Delete snat

This service is used to delete network snat details, you have to send snat details row id that needs to be deleted, with POST request json.    JSON that needs to be sent as request payload, Should contain given information-        id: index(Received as id in api/network-basic service, node SNATGid) of Network SNAT that needs to delete. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteSnat**](DeleteSnat.md)| Object that needs to be sent to the server | 

### Return type

[**NetworkData**](Network_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST6iAction7iType1Post**
> NetworkData POST6iAction7iType1Post(ctx, body)
Add SNAT

This service is used to add network SNAT details, you have to send SNAT details, with POST request json.    JSON that needs to be sent as request payload, Should contain given information-        destinationport: destination port.     interface: interface.     notes: snat notes.     protocol: id of option selected from protocol options.     snattoip: snattoip.     snattoport: snattoport.     sourceip: sourceip.     sourceport: sourceport.     destinationip: destinationip 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddSnat**](AddSnat.md)| Object that needs to be sent to the server | 

### Return type

[**NetworkData**](Network_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

