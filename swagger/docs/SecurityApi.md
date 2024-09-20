# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**15iAction1iType2Post**](SecurityApi.md#15iAction1iType2Post) | **Post** /15?iAction&#x3D;1&amp;iType&#x3D;2 | Enable SSH
[**GET15Get**](SecurityApi.md#GET15Get) | **Get** /GET/15 | Securty page data
[**POST15iAction1iType4Post**](SecurityApi.md#POST15iAction1iType4Post) | **Post** /POST/15?iAction&#x3D;1&amp;iType&#x3D;4 | Enable and update REST Services
[**POST15iAction1iType5Post**](SecurityApi.md#POST15iAction1iType5Post) | **Post** /POST/15?iAction&#x3D;1&amp;iType&#x3D;5 | Update GUI port and ssl cert

# **15iAction1iType2Post**
> SecurityData 15iAction1iType2Post(ctx, body)
Enable SSH

This service is used to Enable, Disable SSH, SSH status(\"true to \"Enabled) information.    JSON that needs to be sent as request payload, Should contains given information-        SSH :set \"true\" to enable ssh. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecurityEnableSsh**](SecurityEnableSsh.md)| Object that needs to be sent to the server | 

### Return type

[**SecurityData**](Security_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET15Get**
> SecurityData GET15Get(ctx, )
Securty page data

This service is used to FETCH security information(GUI PORT and SSL certificate used, REST Service Information) and ssl cert combobox data, it uses four API structure.    Setup_ALB_Addresses_New: Used to get IP address of all available adapters to those we can bind REST services.    Config_Security_SSH: Used to get SSH enable information.    Config_Security_CP_GUI3: Used to get the current PORT and SSL certificate used for GUI webserver and list of all available certificates that can be applied.    Config_Security_CP_GUI3_REST: Used to get IP address and ssl certificate used to which REST services is binded.    The nodes of JSON received, contains the given information-          EnableSSH: This node contains '1' if ssh is enabled.       SCP_CertificateGUI3: This node contains ssl certificate used for GUI.       SCP_PortGUI3: This node contains port of GUI.       SCP_SecurePortGUI3: This node contains port of GUI if it's using ssl(https).       SCP_CertificateRest: This node contains ssl certificate used for REST.       SCP_PortRest: This node contains port of REST.       SCP_IPRest: This node contains ip address of REST.       SCP_EnabledRest: This node contains '1' if REST enabled.       RestIPList: This node contains comma separated list of IP address of adapters that have REST enabled.       SCP_CertificateListStringGUI3: This node contains options of SSL certificates available. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SecurityData**](Security_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST15iAction1iType4Post**
> SecurityData POST15iAction1iType4Post(ctx, body)
Enable and update REST Services

This service is used to Enable, Disable or Update details of REST Services(binded to other PORT or IP), you have to send IP address, PORT, SSL Certificate and REST status(Enabled or Disabled) information, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        SCP_RESTCertificate : id(Received as id in GET/15 service, node SCP_CertificateListStringGUI3) of option selected from SSL Certificate options.      SCP_RESTEnabled : set \"yes\" to enable REST.      SCP_RESTIP : Rest API IP address.      SCP_RESTPort : REST Port Address. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecurityEnableRest**](SecurityEnableRest.md)| Object that needs to be sent to the server | 

### Return type

[**SecurityData**](Security_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST15iAction1iType5Post**
> []SecurityData POST15iAction1iType5Post(ctx, body)
Update GUI port and ssl cert

This service is used to update GUI PORT and SSL Certificate, you have to send PORT and SSL Certificate information with change accepted confimation, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        SCP_CertificateGUI3: ssl certificate used for GUI.      SCP_ChangeAccepted: set \"1\" to confirm changes.      SCP_PortGUI3: port of GUI.      SCP_SecurePortGUI3: port of GUI if it's using ssl(https).          NOTE: While testing GUI by switching between http and https, then take care about protocol changes in URL as well, for example if you have ALB running on https and you changed to http, then to use REST Services all URL should have protocol http(not https as default given in example URLs). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecurityUpdateGui**](SecurityUpdateGui.md)| Object that needs to be sent to the server | 

### Return type

[**[]SecurityData**](Security_Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

