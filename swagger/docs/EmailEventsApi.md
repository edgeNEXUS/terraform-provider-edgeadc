# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Edgenexus/EdgeAPI/4.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GET17Get**](EmailEventsApi.md#GET17Get) | **Get** /GET/17 | Email events page data
[**GET17iAction9iType9CEETestPid38555CEETestEndtimeCEETestTempfilemailtestpi6QbdGet**](EmailEventsApi.md#GET17iAction9iType9CEETestPid38555CEETestEndtimeCEETestTempfilemailtestpi6QbdGet) | **Get** /GET/17?iAction&#x3D;9&amp;iType&#x3D;9&amp;CEE_TestPid&#x3D;38555&amp;CEE_TestEndtime&#x3D;&amp;CEE_TestTempfile&#x3D;mailtestpi6Qbd | Email events test
[**POST17iAction1iType1Post**](EmailEventsApi.md#POST17iAction1iType1Post) | **Post** /POST/17?iAction&#x3D;1&amp;iType&#x3D;1 | Update email event
[**POST17iAction1iType2Post**](EmailEventsApi.md#POST17iAction1iType2Post) | **Post** /POST/17?iAction&#x3D;1&amp;iType&#x3D;2 | Notifications details
[**POST17iAction1iType3Post**](EmailEventsApi.md#POST17iAction1iType3Post) | **Post** /POST/17?iAction&#x3D;1&amp;iType&#x3D;3 | Warning details
[**POST17iAction2iType1Post**](EmailEventsApi.md#POST17iAction2iType1Post) | **Post** /POST/17?iAction&#x3D;2&amp;iType&#x3D;1 | Test Break email event

# **GET17Get**
> []EmaileventsAll GET17Get(ctx, )
Email events page data

This service is used to FETCH E-MAIL Events data, it uses one API structure.    EmailEvents_All: Used to get Email events setings.      The nodes of JSON received, contains the given information        SecurityCombo: This node contains options of security such as SSL and TLS.     CEERemoteEmailAddress: This node contains Email address for Send E-mail.     CEELocalEmailAddress: This node contains Return E-Mail address.     CEEMailServer: This node contains Host Address.     CEEPort: This node contains Port.     CEETimeout: This node contains Timeout.     CEEAuthorisationRequired: This node contains '1', if authorization required.     CEESecurity: This node contains information of type of security such as SSL or TLS.     CEEMailServerLogon: This node contains Mail server account name     CEEEvIpServiceNoticeEnabled: This node have 1 if IP service notice enabled.      CEEEvIpServiceNoticeText: This node have ip service notice text.     CEEEvIpServiceAlertTextEnabled: This node have 1 if IP service alert text enabled.     CEEEvIpServiceAlertText: This node have ip service alert text.     CEEEvChannelNoticeEnabled: This node have 1 if VS notice enabled.      CEEEvChannelNoticeText: This node have VS notice text.     CEEEvChannelAlertEnabled: This node have 1 if VS alert text enabled.     CEEEvChannelAlertText: This node have VS alert text.     CEEEvContentServerNoticeEnabled: This node have 1 if real server notice enabled.      CEEEvContentServerNoticeText: This node have real server notice text.     CEEEvContentServerAlertEnabled: This node have 1 if real server alert text enabled.     CEEEvContentServerAlertText: This node have real server alert text.     CEEEvFlightPathEnabled: This node have 1 if flightpath notice enabled.      CEEEvFlightPathText: This node have flightpath notice text.     CEEGroupEnabled: This node have 1 if group notice enabled.      CEEGroupText: This node have group notice text.     CEEGroupTime: This node have group time.     CEEEvDiskSpaceEnabled: This node have 1 if disk space notice enabled.     CEEEvDiskSpaceText: This node have disk space notice text.     CEEEvLicenceRenevalEnabled: This node have 1 if Licence Reneval notice enabled.     CEEEvLicenceRenevalText: This node have Licence Reneval text.     CEEEvDiskSpacePercent: This node have disk space percentage.     CEEEvTestTempFile:     CEEEvTestProcessID:     CEEEvTestEndTime:     CEEEvTestStatus: 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EmaileventsAll**](Emailevents_All.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GET17iAction9iType9CEETestPid38555CEETestEndtimeCEETestTempfilemailtestpi6QbdGet**
> []EmaileventsAll GET17iAction9iType9CEETestPid38555CEETestEndtimeCEETestTempfilemailtestpi6QbdGet(ctx, )
Email events test

This service is used to FETCH E-MAIL Events data, it uses one API structure.    EmailEvents_All: Used to get Email events setings.      The nodes of JSON received, contains the given information        SecurityCombo: This node contains options of security such as SSL and TLS.     CEERemoteEmailAddress: This node contains Email address for Send E-mail.     CEELocalEmailAddress: This node contains Return E-Mail address.     CEEMailServer: This node contains Host Address.     CEEPort: This node contains Port.     CEETimeout: This node contains Timeout.     CEEAuthorisationRequired: This node contains '1', if authorization required.     CEESecurity: This node contains information of type of security such as SSL or TLS.     CEEMailServerLogon: This node contains Mail server account name     CEEEvIpServiceNoticeEnabled: This node have 1 if IP service notice enabled.      CEEEvIpServiceNoticeText: This node have ip service notice text.     CEEEvIpServiceAlertTextEnabled: This node have 1 if IP service alert text enabled.     CEEEvIpServiceAlertText: This node have ip service alert text.     CEEEvChannelNoticeEnabled: This node have 1 if VS notice enabled.      CEEEvChannelNoticeText: This node have VS notice text.     CEEEvChannelAlertEnabled: This node have 1 if VS alert text enabled.     CEEEvChannelAlertText: This node have VS alert text.     CEEEvContentServerNoticeEnabled: This node have 1 if real server notice enabled.      CEEEvContentServerNoticeText: This node have real server notice text.     CEEEvContentServerAlertEnabled: This node have 1 if real server alert text enabled.     CEEEvContentServerAlertText: This node have real server alert text.     CEEEvFlightPathEnabled: This node have 1 if flightpath notice enabled.      CEEEvFlightPathText: This node have flightpath notice text.     CEEGroupEnabled: This node have 1 if group notice enabled.      CEEGroupText: This node have group notice text.     CEEGroupTime: This node have group time.     CEEEvDiskSpaceEnabled: This node have 1 if disk space notice enabled.     CEEEvDiskSpaceText: This node have disk space notice text.     CEEEvLicenceRenevalEnabled: This node have 1 if Licence Reneval notice enabled.     CEEEvLicenceRenevalText: This node have Licence Reneval text.     CEEEvDiskSpacePercent: This node have disk space percentage.     CEEEvTestTempFile:     CEEEvTestProcessID:     CEEEvTestEndTime:     CEEEvTestStatus: 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EmaileventsAll**](Emailevents_All.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST17iAction1iType1Post**
> []EmaileventsAll POST17iAction1iType1Post(ctx, body)
Update email event

This service is used to update email events detaiils, you have to send email events details, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CEEAuthorisationRequired: set \"true\" if authorization required.     CEEEMailServerPassword: Mail server password.     CEELocalEmailAddress: E-Mail address.     CEEMailServer: Host Address.     CEEMailServerLogon: Mail server account name.     CEEPort: Port address.     CEERemoteEmailAddress: Email address for Send E-mail.     CEESecurity: security      CEETimeout: Timeout.     Test: set \"true\" if want to send test mail with update otherwise set \"false\" to only update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateEmailEvent**](UpdateEmailEvent.md)| Object that needs to be sent to the server | 

### Return type

[**[]EmaileventsAll**](Emailevents_All.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST17iAction1iType2Post**
> []EmaileventsAll POST17iAction1iType2Post(ctx, body)
Notifications details

This service is used to update notifications details, you have to send notifications details, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CEEEvChannelAlertEnabled: set \"true\" if VS alert text enabled.     CEEEvChannelAlertText: VS alert text.     CEEEvChannelNoticeEnabled: set \"true\" if channel notice enabled.      CEEEvChannelNoticeText: channel notice text.     CEEEvContentServerAlertEnabled: set \"true\" if real server alert text enabled.     CEEEvContentServerAlertText: real server alert text.     CEEEvContentServerNoticeEnabled:     CEEEvContentServerNoticeText:     CEEEvFlightPathEnabled: set \"true\" if flightpath notice enabled.      CEEEvFlightPathText: flightpath notice text.     CEEEvIpServiceAlertTextEnabled: set \"true\" if IP service alert text enabled.     CEEEvIpServiceAlertText: ip service alert text.     CEEEvIpServiceNoticeEnabled: set \"true\" if IP service notice enabled.      CEEEvIpServiceNoticeText: ip service notice text.     CEEGroupEnabled: set \"true\" if group notice enabled.      CEEGroupText: group notice text.     CEEGroupTime: group time. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationEmailEvent**](NotificationEmailEvent.md)| Object that needs to be sent to the server | 

### Return type

[**[]EmaileventsAll**](Emailevents_All.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST17iAction1iType3Post**
> []EmaileventsAll POST17iAction1iType3Post(ctx, body, iAction, iType)
Warning details

This service is used to update warnings details, you have to send warnings details, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CEEEvDiskSpaceEnabled: set \"true\" if disk space notice enabled.     CEEEvDiskSpaceText: disk space notice text.     CEEEvLicenceRenevalEnabled: set \"true\" if Licence Reneval notice enabled.     CEEEvLicenceRenevalText: Licence Reneval text.     CEEEvDiskSpacePercent: disk space percentage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WarningsEmailEvent**](WarningsEmailEvent.md)| Object that needs to be sent to the server | 
  **iAction** | **float64**| 1 | 
  **iType** | **float64**| 3 | 

### Return type

[**[]EmaileventsAll**](Emailevents_All.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POST17iAction2iType1Post**
> []EmaileventsAll POST17iAction2iType1Post(ctx, body)
Test Break email event

This service is used to update email events detaiils, you have to send email events details, with POST request json.    JSON that needs to be sent as request payload, Should contains given information-        CEEAuthorisationRequired: set \"true\" if authorization required.     CEEEMailServerPassword: Mail server password.     CEELocalEmailAddress: E-Mail address.     CEEMailServer: Host Address.     CEEMailServerLogon: Mail server account name.     CEEPort: Port address.     CEERemoteEmailAddress: Email address for Send E-mail.     CEESecurity: security      CEETimeout: Timeout.     Test: set \"true\" if want to send test mail with update otherwise set \"false\" to only update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateEmailEvent**](UpdateEmailEvent.md)| Object that needs to be sent to the server | 

### Return type

[**[]EmaileventsAll**](Emailevents_All.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

