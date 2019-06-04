package resources

import (
	"strings"
)

type google_compute_backend_service struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			AffinityCookieTTLSec         string   `json:"affinity_cookie_ttl_sec"`
			Backend                      []Params `json:"backend"`
			CdnPolicy                    []Params `json:"cdn_policy"`
			ConnectionDrainingTimeoutSec string   `json:"connection_draining_timeout_sec"`
			CreationTimestamp            string   `json:"creation_timestamp"`
			Description                  string   `json:"description"`
			EnableCdn                    string   `json:"enable_cdn"`
			Fingerprint                  string   `json:"fingerprint"`
			HealthChecks                 []string `json:"health_checks"`
			Iap                          []Params `json:"iap"`
			ID                           string   `json:"id"`
			LoadBalancingScheme          string   `json:"load_balancing_scheme"`
			Name                         string   `json:"name"`
			PortName                     string   `json:"port_name"`
			Project                      string   `json:"project"`
			Protocol                     string   `json:"protocol"`
			SecurityPolicy               string   `json:"security_policy"`
			SelfLink                     string   `json:"self_link"`
			SessionAffinity              string   `json:"session_affinity"`
			TimeoutSec                   string   `json:"timeout_sec"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type Backend struct {
	Number                    string
	BalancingMode             string `json:"balancing_mode"`
	CapacityScaler            string `json:"capacity_scaler"`
	Description               string `json:"description"`
	Group                     string `json:"group"`
	MaxConnections            string `json:"max_connections"`
	MaxConnectionsPerInstance string `json:"max_connections_per_instance"`
	MaxRate                   string `json:"max_rate"`
	MaxRatePerInstance        string `json:"max_rate_per_instance"`
	MaxUtilization            string `json:"max_utilization"`
}

func (d *Backend) SetNumber(num string) {
	d.Number = num
}
func (d *Backend) GetNumber() string {
	return d.Number
}
func NewBackend(num string) Params {
	var s Params = new(Backend)
	s.SetNumber(num)
	return s
}
func (d *Backend) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "balancing_mode") {
		d.BalancingMode = value
	} else if strings.Contains(key, "capacity_scaler") {
		d.CapacityScaler = value
	} else if strings.Contains(key, "description") {
		d.Description = value
	} else if strings.Contains(key, "group") {
		d.Group = value
	} else if strings.Contains(key, "max_connections_per_instance") {
		d.MaxConnectionsPerInstance = value
	} else if strings.Contains(key, "max_connection") {
		d.MaxConnections = value
	} else if strings.Contains(key, "max_rate_per_instance") {
		d.MaxRatePerInstance = value
	} else if strings.Contains(key, "max_rate") {
		d.MaxRate = value
	} else if strings.Contains(key, "max_utilization") {
		d.MaxUtilization = value
	}
}

type CdnPolicy struct {
	Number                  string
	CacheKeyPolicy          []Params `json:"cache_key_policy"`
	SignedUrlCacheMaxAgeSec string   `json:"signed_url_cache_max_age_sec"`
}

func (d *CdnPolicy) SetNumber(num string) {
	d.Number = num
}
func (d *CdnPolicy) GetNumber() string {
	return d.Number
}
func NewCdnPolicy(num string) Params {
	var s Params = new(CdnPolicy)
	s.SetNumber(num)
	return s
}
func (d *CdnPolicy) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "cache_key_policy") {
		Number := CreateNumberFromKey(key, 2)
		d.CacheKeyPolicy = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewCacheKeyPolicy)
	} else if strings.Contains(key, "signed_url_cache_max_age_sec") {
		d.SignedUrlCacheMaxAgeSec = value
	}
}

type CacheKeyPolicy struct {
	Number               string
	IncludHost           string `json:"include_host"`
	IncludProtocol       string `json:"include_protocol"`
	IncludQueryString    string `json:"include_query_string"`
	QueryStringBlackList string `json:"query_string_blacklist"`
	QueryStringWhiteList string `json:"query_string_whitelist"`
}

func (d *CacheKeyPolicy) SetNumber(num string) {
	d.Number = num
}
func (d *CacheKeyPolicy) GetNumber() string {
	return d.Number
}
func NewCacheKeyPolicy(num string) Params {
	var s Params = new(CacheKeyPolicy)
	s.SetNumber(num)
	return s
}
func (d *CacheKeyPolicy) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "include_host") {
		d.IncludHost = value
	} else if strings.Contains(key, "include_protocol") {
		d.IncludProtocol = value
	} else if strings.Contains(key, "include_query_string") {
		d.IncludQueryString = value
	} else if strings.Contains(key, "query_string_blacklist") {
		d.QueryStringBlackList = value
	} else if strings.Contains(key, "query_string_whitelist") {
		d.QueryStringWhiteList = value
	}
}

type Iap struct {
	Number                   string
	Oauth2ClientId           string `json:"oauth2_client_id"`
	Oauth2ClientSecret       string `json:"oauth2_client_secret"`
	Oauth2ClientSecretSha256 string `json:"oauth2_client_secret_sha256"`
}

func (d *Iap) SetNumber(num string) {
	d.Number = num
}
func (d *Iap) GetNumber() string {
	return d.Number
}
func NewIap(num string) Params {
	var s Params = new(Iap)
	s.SetNumber(num)
	return s
}
func (d *Iap) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "oauth2_client_id") {
		d.Oauth2ClientId = value
	} else if strings.Contains(key, "oauth2_client_secret_sha256") {
		d.Oauth2ClientSecretSha256 = value
	} else if strings.Contains(key, "oauth2_client_secret") {
		d.Oauth2ClientSecret = value
	}
}

func Newgoogle_compute_backend_service(resourceData map[string]interface{}, resourceName string) *google_compute_backend_service {

	resource := new(google_compute_backend_service)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.HealthChecks = CreateStringSlice(attributes_map, attributesKey, "health_checks")
	resource.Primary.Attributes.Backend = CreateStructSlice(attributes_map, attributesKey, "backend", 1, NewBackend)
	resource.Primary.Attributes.CdnPolicy = CreateStructSlice(attributes_map, attributesKey, "cdn_policy", 1, NewCdnPolicy)
	resource.Primary.Attributes.Iap = CreateStructSlice(attributes_map, attributesKey, "iap", 1, NewIap)

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
