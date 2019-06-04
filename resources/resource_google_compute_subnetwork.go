package resources

import (
	"strings"
)

type google_compute_subnetwork struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			CreationTimestamp     string   `json:"creation_timestamp"`
			Description           string   `json:"description"`
			EnableFlowLogs        string   `json:"enable_flow_logs"`
			Fingerprint           string   `json:"fingerprint"`
			GatewayAddress        string   `json:"gateway_address"`
			ID                    string   `json:"id"`
			IPCidrRange           string   `json:"ip_cidr_range"`
			Name                  string   `json:"name"`
			Network               string   `json:"network"`
			PrivateIPGoogleAccess string   `json:"private_ip_google_access"`
			Project               string   `json:"project"`
			Region                string   `json:"region"`
			SecondaryIPRange      []Params `json:"secondary_ip_range"`
			SelfLink              string   `json:"self_link"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type SecondaryIPRange struct {
	Number        string
	Ip_cidr_range string
	Range_name    string
}

func (d *SecondaryIPRange) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "ip_cidr_range") {
		d.Ip_cidr_range = value
	} else if strings.Contains(key, "range_name") {
		d.Range_name = value
	}
}

func (d *SecondaryIPRange) SetNumber(num string) {
	d.Number = num
}

func (d *SecondaryIPRange) GetNumber() string {
	return d.Number
}

func NewSecondaryIPRange(num string) Params {
	var s Params = new(SecondaryIPRange)
	s.SetNumber(num)
	return s
}

func Newgoogle_compute_subnetwork(resourceData map[string]interface{}, resourceName string) *google_compute_subnetwork {

	resource := new(google_compute_subnetwork)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.SecondaryIPRange = CreateStructSlice(attributes_map, attributesKey, "secondary_ip_range", 1, NewSecondaryIPRange)
	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
