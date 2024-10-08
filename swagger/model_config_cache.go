/*
 * Edgenexus REST API
 *
 * EdgeADC REST Web Services User Guide
 *
 * API version: 4.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ConfigCache struct {
	// true
	Success string `json:"success,omitempty"`
	// images/light-hd_red.gif
	StatusImage string `json:"StatusImage,omitempty"`
	// get
	StatusText string `json:"StatusText,omitempty"`
	CacheIncExcCombo *ConfigCacheCacheIncExcCombo `json:"CacheIncExcCombo,omitempty"`
	CacheRuleHelpTypeCombo *ConfigCacheCacheRuleHelpTypeCombo `json:"CacheRuleHelpTypeCombo,omitempty"`
	// 50
	CCHMaxCache string `json:"CCHMaxCache,omitempty"`
	// 30
	CCHMinCache string `json:"CCHMinCache,omitempty"`
	// 11/00:00
	CCHHeuristicExpiry string `json:"CCHHeuristicExpiry,omitempty"`
	// 200 203 301 304 410
	CCHCachableResponses string `json:"CCHCachableResponses,omitempty"`
	// 6/01:00
	CCHTrimCache string `json:"CCHTrimCache,omitempty"`
	// 25
	CCHCache304Stream string `json:"CCHCache304Stream,omitempty"`
	CacheOtherDomainCombo *ConfigCacheCacheOtherDomainCombo `json:"CacheOtherDomainCombo,omitempty"`
	CacheDomainGrid *ConfigCacheCacheDomainGrid `json:"CacheDomainGrid,omitempty"`
	CacheRulesGrid *ConfigCacheCacheRulesGrid `json:"CacheRulesGrid,omitempty"`
	CacheDomainCombo *ConfigCacheCacheDomainCombo `json:"CacheDomainCombo,omitempty"`
}
