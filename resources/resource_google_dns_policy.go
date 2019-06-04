package resources

import (
	"strings"
)

type google_dns_policy struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			AlternativeNameServerConfig []Params `json:"alternative_name_server_config"`
			Description                 string   `json:"description"`
			EnableInboundForwarding     string   `json:"enable_inbound_forwarding"`
			EnableLogging               string   `json:"enable_logging"`
			ID                          string   `json:"id"`
			Name                        string   `json:"name"`
			Networks                    []Params `json:"networks"`
			Project                     string   `json:"project"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type AlternativeNameServerConfig struct {
	Number            string
	TargetNameServers []string `json:"target_name_servers"`
}

func (d *AlternativeNameServerConfig) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "target_name_servers") {
		d.TargetNameServers = append(d.TargetNameServers, value)
	}
}

func (d *AlternativeNameServerConfig) SetNumber(num string) {
	d.Number = num
}

func (d *AlternativeNameServerConfig) GetNumber() string {
	return d.Number
}

func NewAlternativeNameServerConfig(num string) Params {
	var s Params = new(AlternativeNameServerConfig)
	s.SetNumber(num)
	return s
}

func Newgoogle_dns_policy(resourceData map[string]interface{}, resourceName string) *google_dns_policy {

	resource := new(google_dns_policy)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.AlternativeNameServerConfig = CreateStructSlice(attributes_map, attributesKey, "alternative_name_server_config", 1, NewAlternativeNameServerConfig)
	resource.Primary.Attributes.Networks = CreateStructSlice(attributes_map, attributesKey, "networks", 1, NewNetworks)

	return resource
}
