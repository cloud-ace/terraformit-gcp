package resources

import (
	"strings"
)

type google_dns_managed_zone struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			Description             string            `json:"description"`
			DNSName                 string            `json:"dns_name"`
			ForwardingConfig        []Params          `json:"forwarding_config"`
			ID                      string            `json:"id"`
			Labels                  map[string]string `json:"labels"`
			Name                    string            `json:"name"`
			NameServers             []string          `json:"name_servers"`
			PeeringConfig           []Params          `json:"peering_config"`
			PrivateVisibilityConfig []Params          `json:"private_visibility_config"`
			Project                 string            `json:"project"`
			Visibility              string            `json:"visibility"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type ForwardingConfig struct {
	Number            string
	TargetNameServers []Params `json:"target_name_servers"`
}

func (d *ForwardingConfig) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "target_name_servers") {
		Number := CreateNumberFromKey(key, 2)
		d.TargetNameServers = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewTargetNameServers)
	}
}

func (d *ForwardingConfig) SetNumber(num string) {
	d.Number = num
}

func (d *ForwardingConfig) GetNumber() string {
	return d.Number
}

func NewForwardingConfig(num string) Params {
	var s Params = new(ForwardingConfig)
	s.SetNumber(num)
	return s
}

type TargetNameServers struct {
	Number      string
	IPv4Address string `json:"ipv4_address"`
}

func (d *TargetNameServers) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "ipv4_address") {
		d.IPv4Address = value
	}
}

func (d *TargetNameServers) SetNumber(num string) {
	d.Number = num
}

func (d *TargetNameServers) GetNumber() string {
	return d.Number
}

func NewTargetNameServers(num string) Params {
	var s Params = new(TargetNameServers)
	s.SetNumber(num)
	return s
}

type PeeringConfig struct {
	Number        string
	TargetNetwork []Params `json:"target_network"`
}

func (d *PeeringConfig) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "target_network") {
		Number := CreateNumberFromKey(key, 2)
		d.TargetNetwork = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewTargetNetwork)
	}
}

func (d *PeeringConfig) SetNumber(num string) {
	d.Number = num
}

func (d *PeeringConfig) GetNumber() string {
	return d.Number
}

func NewPeeringConfig(num string) Params {
	var s Params = new(PeeringConfig)
	s.SetNumber(num)
	return s
}

type TargetNetwork struct {
	Number     string
	NetworkURL string `json:"network_url"`
}

func (d *TargetNetwork) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "network_url") {
		d.NetworkURL = value
	}
}

func (d *TargetNetwork) SetNumber(num string) {
	d.Number = num
}

func (d *TargetNetwork) GetNumber() string {
	return d.Number
}

func NewTargetNetwork(num string) Params {
	var s Params = new(TargetNetwork)
	s.SetNumber(num)
	return s
}

type PrivateVisibilityConfig struct {
	Number   string
	Networks []Params `json:"networks"`
}

func (d *PrivateVisibilityConfig) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "target_network") {
		Number := CreateNumberFromKey(key, 2)
		d.Networks = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewNetworks)
	}
}

func (d *PrivateVisibilityConfig) SetNumber(num string) {
	d.Number = num
}

func (d *PrivateVisibilityConfig) GetNumber() string {
	return d.Number
}

func NewPrivateVisibilityConfig(num string) Params {
	var s Params = new(PrivateVisibilityConfig)
	s.SetNumber(num)
	return s
}

type Networks struct {
	Number     string
	NetworkURL string `json:"network_url"`
}

func (d *Networks) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "network_url") {
		d.NetworkURL = value
	}
}

func (d *Networks) SetNumber(num string) {
	d.Number = num
}

func (d *Networks) GetNumber() string {
	return d.Number
}

func NewNetworks(num string) Params {
	var s Params = new(Networks)
	s.SetNumber(num)
	return s
}

func Newgoogle_dns_managed_zone(resourceData map[string]interface{}, resourceName string) *google_dns_managed_zone {

	resource := new(google_dns_managed_zone)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.NameServers = CreateStringSlice(attributes_map, attributesKey, "name_servers")
	resource.Primary.Attributes.PeeringConfig = CreateStructSlice(attributes_map, attributesKey, "peering_config", 1, NewPeeringConfig)
	resource.Primary.Attributes.ForwardingConfig = CreateStructSlice(attributes_map, attributesKey, "forwarding_config", 1, NewForwardingConfig)
	resource.Primary.Attributes.PrivateVisibilityConfig = CreateStructSlice(attributes_map, attributesKey, "private_visibility_config", 1, NewPrivateVisibilityConfig)

	return resource
}
