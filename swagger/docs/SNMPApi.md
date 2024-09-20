# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET22Get**](SNMPApi.md#GET22Get) | **Get** /GET/22 | SNMP page data
[**POST22iAction3iType1Post**](SNMPApi.md#POST22iAction3iType1Post) | **Post** /POST/22?iAction&#x3D;3&amp;iType&#x3D;1 | Update SNMP details

# **GET22Get**
> SnmpData GET22Get(ctx, )
SNMP page data

This service is used to FETCH SNMP details, it uses one API structure.    Setup_SNMP: Used to get SNMPV1V2Checked, SNMPV3Checked and Community string.    The nodes of JSON received, contains the given information          SNMPV1V2Checked: This node contains '1' if SNMP V1/2c is enabled.       SNMPCommunityString: This node contains SNMP Community String.       SNMPV3Checked: This node contains '1' if SNMP V3 is enabled. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SnmpData**](SNMP_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST22iAction3iType1Post**
> SnmpData POST22iAction3iType1Post(ctx, body)
Update SNMP details

This service is used to update SNMP details, you have to send SNMP data, with POST request json.    JSON that needs to be sent as request payload, Should contains given information        SNMPV1V2Checked: set 'true' if SNMP V1/2c is enabled.      SNMPCommunityString : SNMP Community String.      SNMPV3Checked : set 'true' if SNMP V3 is enabled.      oldPassPhrase : new password      newPassPhrase : Old password      confirmNewPassPhrase : confirm password 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SnmpUpdate**](SnmpUpdate.md)| Object that needs to be sent to the server | 

### Return type

[**SnmpData**](SNMP_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

