package resources

import (
	"strings"
)

type google_compute_firewall struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			Allow                 []Params `json:"allow"`
			CreationTimestamp     string   `json:"creation_timestamp"`
			Deny                  []Params `json:"deny"`
			Description           string   `json:"description"`
			DestinationRanges     []string `json:"destination_ranges"`
			Direction             string   `json:"direction"`
			Disabled              string   `json:"disabled"`
			ID                    string   `json:"id"`
			Name                  string   `json:"name"`
			Network               string   `json:"network"`
			Priority              string   `json:"priority"`
			Project               string   `json:"project"`
			SelfLink              string   `json:"self_link"`
			SourceRanges          []string `json:"source_ranges"`
			SourceServiceAccounts []string `json:"source_service_accounts"`
			SourceTags            []string `json:"source_tags"`
			TargetServiceAccounts []string `json:"target_service_accounts"`
			TargetTags            []string `json:"target_tags"`
			EnableLogging         string   `json:"enable_logging"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type AllowOrDeny struct {
	Number   string
	Protocol string
	Ports    []string
}

func (d *AllowOrDeny) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "protocol") {
		d.Protocol = value
	} else if strings.Contains(key, "ports") {
		d.Ports = append(d.Ports, value)
	}
}

func (d *AllowOrDeny) SetNumber(num string) {
	d.Number = num
}

func (d *AllowOrDeny) GetNumber() string {
	return d.Number
}

func NewAllowOrDeny(num string) Params {
	var s Params = new(AllowOrDeny)
	s.SetNumber(num)
	return s
}

func Newgoogle_compute_firewall(resourceData map[string]interface{}, resourceName string) *google_compute_firewall {

	resource := new(google_compute_firewall)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.DestinationRanges = CreateStringSlice(attributes_map, attributesKey, "destination_ranges")
	resource.Primary.Attributes.SourceRanges = CreateStringSlice(attributes_map, attributesKey, "source_ranges")
	resource.Primary.Attributes.SourceServiceAccounts = CreateStringSlice(attributes_map, attributesKey, "source_service_accounts")
	resource.Primary.Attributes.SourceTags = CreateStringSlice(attributes_map, attributesKey, "source_tags")
	resource.Primary.Attributes.TargetTags = CreateStringSlice(attributes_map, attributesKey, "target_tags")
	resource.Primary.Attributes.TargetServiceAccounts = CreateStringSlice(attributes_map, attributesKey, "target_service_accounts")
	resource.Primary.Attributes.Allow = CreateStructSlice(attributes_map, attributesKey, "allow", 1, NewAllowOrDeny)
	resource.Primary.Attributes.Deny = CreateStructSlice(attributes_map, attributesKey, "deny", 1, NewAllowOrDeny)

	//beta
	if resourceData["provider"] == "provider.google-beta" {
		resource.Primary.Attributes.EnableLogging = attributes_map["enable_logging"].(string)
	}
	return resource
}
