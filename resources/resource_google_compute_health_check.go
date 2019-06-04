package resources

import "strings"

type google_compute_health_check struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			CheckIntervalSec   string   `json:"check_interval_sec"`
			CreationTimestamp  string   `json:"creation_timestamp"`
			Description        string   `json:"description"`
			HealthyThreshold   string   `json:"healthy_threshold"`
			ID                 string   `json:"id"`
			Name               string   `json:"name"`
			Project            string   `json:"project"`
			SelfLink           string   `json:"self_link"`
			HTTPSHealthCheck   []Params `json:"https_health_check.port"`
			HTTPHealthCheck    []Params `json:"http_health_check.port"`
			SslHealthCheck     []Params `json:"ssl_health_check.port"`
			TCPHealthCheck     []Params `json:"tcp_health_check.port"`
			TimeoutSec         string   `json:"timeout_sec"`
			Type               string   `json:"type"`
			UnhealthyThreshold string   `json:"unhealthy_threshold"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type HealthCheck struct {
	Number      string
	Port        string `json:"port"`
	ProxyHeader string `json:"proxy_header"`
	RequestPath string `json:"request_path"`
	Response    string `json:"response"`
	Host        string `json:"host"`
	Request     string `json:"request_path"`
}

func (d *HealthCheck) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "port") {
		d.Port = value
	} else if strings.Contains(key, "proxy_header") {
		d.ProxyHeader = value
	} else if strings.Contains(key, "request_path") {
		d.RequestPath = value
	} else if strings.Contains(key, "response") {
		d.Response = value
	} else if strings.Contains(key, "host") {
		d.Host = value
		//request_pathが先に評価される
	} else if strings.Contains(key, "request") {
		d.Request = value
	}
}

func (d *HealthCheck) SetNumber(num string) {
	d.Number = num
}

func (d *HealthCheck) GetNumber() string {
	return d.Number
}

func NewHealthCheck(num string) Params {
	var s Params = new(HealthCheck)
	s.SetNumber(num)
	return s
}

func Newgoogle_compute_health_check(resourceData map[string]interface{}, resourceName string) *google_compute_health_check {

	resource := new(google_compute_health_check)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.HTTPSHealthCheck = CreateStructSlice(attributes_map, attributesKey, "https_health_check", 1, NewHealthCheck)
	resource.Primary.Attributes.HTTPHealthCheck = CreateStructSlice(attributes_map, attributesKey, "http_health_check", 1, NewHealthCheck)
	resource.Primary.Attributes.SslHealthCheck = CreateStructSlice(attributes_map, attributesKey, "ssl_health_check", 1, NewHealthCheck)
	resource.Primary.Attributes.TCPHealthCheck = CreateStructSlice(attributes_map, attributesKey, "tcp_health_check", 1, NewHealthCheck)
	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
