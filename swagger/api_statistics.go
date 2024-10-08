
/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type StatisticsApiService service
/*
StatisticsApiService Statistics page data
This service is used to FETCH all Statistics details, it uses one API structure.    Monitor_Statistics2: Used to get Statistics Information(cache, Compression, hit count memory used etc).    The nodes of JSON received, contains the given information-        GUISTATUS: This node contains information of GUI light.     CPU: This node contains CPU used.     Disk: This node contains DISK used.     Memory: This node contains MEMORY used.     GraphMemory: This node contains graph memory.     FromCacheHits1:      FromCacheHits2:      FromCacheBytes1:     FromCacheBytes2:      FromServerHits1:      FromServerHits2:     FromServerBytes1:      FromServerBytes2:     CacheContentsHits: This node contains cache content hits.     CacheContentsBytes: This node contains bytes of cached content.     RawInputBytes: This node contains raw input bytes.     RawOutputBytes: This node contains raw output bytes.     RawInputBytesSec: This node contains raw input bytes per second.     RawOutputBytesSec: This node contains raw output bytes per second.     ContentCompression: This node contains value of content compression.     UploadContentCompression: This node contains value of upload content compression.     UploadContentUncompressed:     ContentCompressed: This node contains value of content compressed.     ContentUncompressed     OverallCompression: This node contains value of overall compression.     UploadOverallCompression: This node contains value of overall uploaded compression.     InstantaneousPercent: This node contains instantaneous percent.     OverallUncompressed: This node contains value of overall uncompressed.     UploadOverallUncompressed: This node contains value of overall upload uncompressed.     InstantaneousMbpsIn: This node contains Instantaneous Mbps In.     OverallCompressed: This node contains value of overall compressed.     OverallCompressedCache: This node contains value of overall compressed cache.     InstantaneousMbpsOut: This node contains Instantaneous Mbps out.     InstantaneousMbpsCache: This node contains information of Instantaneous Mbps cache.     InstantaneousMbpsTotal: This node contains total of Instantaneous Mbps.     HitCount: This node contains value of Hit Count.     InstantaneousHits: This node contains Instantaneous Hits.     OverallConnections: This node contains Overall Connections.     InstantaneousConnections: This node contains information of Instantaneous Connections.     InstantaneousServerConnections: This node contains Instantaneous Server Connections.     MaxConnections: This node contains Maximum Connections.     CurrentConnections: This node contains information of Current Connections.     PageTitle     CacheContentsBytes1     CacheContentsBytes2     UploadContentCompressed     TotalEntries     AllowedEntries     NewEntries     RevalidatedEntries     ExpiredEntries     DeletedEntries     LowAverageBufFillSize : This node contains information of LowAverageBufFillSize.     MedAverageBufFillSize : This node contains information of MedAverageBufFillSize.     HighAverageBufFillSize : This node contains information of HighAverageBufFillSize.     PercentLowConnections : This node contains information of PercentLowConnections.     PercentMedConnections : This node contains information of PercentMedConnections.     PercentHighConnections : This node contains information of PercentHighConnections.     PercentLowMemory : This node contains information of PercentLowMemory.     PercentMedMemory : This node contains information of PercentMedMemory.     PercentHighMemory : This node contains information of PercentHighMemory.     LowMemoryFree : This node contains information of LowMemoryFree.     MedMemoryUsed : This node contains information of MedMemoryUsed.     MedMemoryFree : This node contains information of MedMemoryFree.     HighMemoryUsed : This node contains information of HighMemoryUsed.     HighMemoryFree : This node contains information of HighMemoryFree.     LowConnectionsUsed : This node contains information of LowConnectionsUsed.     LowConnectionsFree : This node contains information of LowConnectionsFree.     MedConnectionsFree : This node contains information of MedConnectionsFree.     HighConnectionsUsed : This node contains information of HighConnectionsUsed.     HighConnectionsFree : This node contains information of HighConnectionsFree.   
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return MonitorStatistics2
*/
func (a *StatisticsApiService) GET14reqrefreshGet(ctx context.Context) (MonitorStatistics2, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MonitorStatistics2
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/GET/14?req=refresh"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			
			localVarQueryParams.Add("GUID", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MonitorStatistics2
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
StatisticsApiService Reset
This service is used to Reset.    JSON that needs to be sent as request payload, Should contains given information-          reset:  
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body Object that needs to be sent to the server
@return MonitorStatistics2
*/
func (a *StatisticsApiService) POST14iAction1iType1Post(ctx context.Context, body StatisticsReset) (MonitorStatistics2, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MonitorStatistics2
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/POST/14?iAction=1&iType=1"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			
			localVarQueryParams.Add("GUID", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MonitorStatistics2
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
