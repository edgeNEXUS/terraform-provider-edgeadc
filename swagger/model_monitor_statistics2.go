/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MonitorStatistics2 struct {
	// true
	Success string `json:"success,omitempty"`
	// images/light-hd_red.gif
	StatusImage string `json:"StatusImage,omitempty"`
	// get
	StatusText string `json:"StatusText,omitempty"`
	// images/jnpsStatusBg.gif
	GUISTATUS string `json:"GUISTATUS,omitempty"`
	// 13.0%
	CPU string `json:"CPU,omitempty"`
	// 46%
	Disk string `json:"Disk,omitempty"`
	// 86.8%( 1271.4MB of 1464.9MB)
	Memory string `json:"Memory,omitempty"`
	// 86.8%
	GraphMemory string `json:"GraphMemory,omitempty"`
	// 0
	FromCacheHits1 string `json:"FromCacheHits1,omitempty"`
	// -
	FromCacheHits2 string `json:"FromCacheHits2,omitempty"`
	// 0
	FromCacheBytes1 string `json:"FromCacheBytes1,omitempty"`
	// -
	FromCacheBytes2 string `json:"FromCacheBytes2,omitempty"`
	// 0
	FromServerHits1 string `json:"FromServerHits1,omitempty"`
	// -
	FromServerHits2 string `json:"FromServerHits2,omitempty"`
	// 0
	FromServerBytes1 string `json:"FromServerBytes1,omitempty"`
	// -
	FromServerBytes2 string `json:"FromServerBytes2,omitempty"`
	// 0 entries
	CacheContentsHits string `json:"CacheContentsHits,omitempty"`
	// 0
	CacheContentsBytes1 string `json:"CacheContentsBytes1,omitempty"`
	// 0.0%
	CacheContentsBytes2 string `json:"CacheContentsBytes2,omitempty"`
	// 11.2 MB
	RawInputBytes string `json:"RawInputBytes,omitempty"`
	// 16.6 MB
	RawOutputBytes string `json:"RawOutputBytes,omitempty"`
	// 6.75 kbps
	RawInputBytesSec string `json:"RawInputBytesSec,omitempty"`
	// 4.86 kbps
	RawOutputBytesSec string `json:"RawOutputBytesSec,omitempty"`
	// 0%
	ContentCompression string `json:"ContentCompression,omitempty"`
	// 0%
	UploadContentCompression string `json:"UploadContentCompression,omitempty"`
	// 0
	UploadContentUncompressed string `json:"UploadContentUncompressed,omitempty"`
	// 0
	UploadContentCompressed string `json:"UploadContentCompressed,omitempty"`
	// 0
	ContentUncompressed string `json:"ContentUncompressed,omitempty"`
	// 0
	ContentCompressed string `json:"ContentCompressed,omitempty"`
	// 0%
	OverallCompression string `json:"OverallCompression,omitempty"`
	// 0%
	UploadOverallCompression string `json:"UploadOverallCompression,omitempty"`
	// 1%
	InstantaneousPercent string `json:"InstantaneousPercent,omitempty"`
	// 0
	OverallUncompressed string `json:"OverallUncompressed,omitempty"`
	// 0
	UploadOverallUncompressed string `json:"UploadOverallUncompressed,omitempty"`
	// 0.00 Mbps (data)
	InstantaneousMbpsIn string `json:"InstantaneousMbpsIn,omitempty"`
	// 0
	OverallCompressed string `json:"OverallCompressed,omitempty"`
	// 0
	OverallCompressedCache string `json:"OverallCompressedCache,omitempty"`
	// 0.00 Mbps (data)
	InstantaneousMbpsOut string `json:"InstantaneousMbpsOut,omitempty"`
	// 0.00 Mbps (data)
	InstantaneousMbpsCache string `json:"InstantaneousMbpsCache,omitempty"`
	// 0.00 Mbps (data)
	InstantaneousMbpsTotal string `json:"InstantaneousMbpsTotal,omitempty"`
	// 0
	HitCount string `json:"HitCount,omitempty"`
	// 0
	InstantaneousHits string `json:"InstantaneousHits,omitempty"`
	// 0
	OverallConnections string `json:"OverallConnections,omitempty"`
	// 0
	InstantaneousConnections string `json:"InstantaneousConnections,omitempty"`
	// 0
	InstantaneousServerConnections string `json:"InstantaneousServerConnections,omitempty"`
	// 0
	MaxConnections string `json:"MaxConnections,omitempty"`
	// 0
	CurrentConnections string `json:"CurrentConnections,omitempty"`
	// 0
	TotalEntries string `json:"TotalEntries,omitempty"`
	// 0
	AllowedEntries string `json:"AllowedEntries,omitempty"`
	// 0
	NewEntries string `json:"NewEntries,omitempty"`
	// 0
	RevalidatedEntries string `json:"RevalidatedEntries,omitempty"`
	// 0
	ExpiredEntries string `json:"ExpiredEntries,omitempty"`
	// 0
	DeletedEntries string `json:"DeletedEntries,omitempty"`
	// ALB-X
	PageTitle string `json:"PageTitle,omitempty"`
	// 0
	LowConnectionsUsed string `json:"LowConnectionsUsed,omitempty"`
	// 0
	LowConnectionsFree string `json:"LowConnectionsFree,omitempty"`
	// 0
	MedConnectionsUsed string `json:"MedConnectionsUsed,omitempty"`
	// 0
	MedConnectionsFree string `json:"MedConnectionsFree,omitempty"`
	// 0
	HighConnectionsUsed string `json:"HighConnectionsUsed,omitempty"`
	// 0
	HighConnectionsFree string `json:"HighConnectionsFree,omitempty"`
	// 0
	LowMemoryUsed string `json:"LowMemoryUsed,omitempty"`
	// ALB-X
	LowMemoryFree string `json:"LowMemoryFree,omitempty"`
	// 0
	MedMemoryUsed string `json:"MedMemoryUsed,omitempty"`
	// 0
	MedMemoryFree string `json:"MedMemoryFree,omitempty"`
	// 0
	HighMemoryUsed string `json:"HighMemoryUsed,omitempty"`
	// 0
	HighMemoryFree string `json:"HighMemoryFree,omitempty"`
	// 0
	LowAverageBufFillSize string `json:"LowAverageBufFillSize,omitempty"`
	// 0
	MedAverageBufFillSize string `json:"MedAverageBufFillSize,omitempty"`
	// 0
	HighAverageBufFillSize string `json:"HighAverageBufFillSize,omitempty"`
	// 0
	PercentLowConnections string `json:"PercentLowConnections,omitempty"`
	// 0
	PercentMedConnections string `json:"PercentMedConnections,omitempty"`
	// 0
	PercentHighConnections string `json:"PercentHighConnections,omitempty"`
	// 0
	PercentLowMemory string `json:"PercentLowMemory,omitempty"`
	// 0
	PercentMedMemory string `json:"PercentMedMemory,omitempty"`
	// 0
	PercentHighMemory string `json:"PercentHighMemory,omitempty"`
}
