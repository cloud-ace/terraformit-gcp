package resources

import "strings"

type google_compute_instance_group struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			Description string   `json:"description"`
			ID          string   `json:"id"`
			Instances   []string `json:"instances."`
			Name        string   `json:"name"`
			NamedPort   []Params `json:"named_port.#"`
			Network     string   `json:"network"`
			Project     string   `json:"project"`
			SelfLink    string   `json:"self_link"`
			Size        string   `json:"size"`
			Zone        string   `json:"zone"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type NamedPort struct {
	Number string
	Name   string
	Port   string
}

func (d *NamedPort) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, ".name") {
		d.Name = value
	} else if strings.Contains(key, ".port") {
		d.Port = value
	}
}

func (d *NamedPort) SetNumber(num string) {
	d.Number = num
}

func (d *NamedPort) GetNumber() string {
	return d.Number
}

func NewNamedPort(num string) Params {
	var s Params = new(NamedPort)
	s.SetNumber(num)
	return s
}

func Newgoogle_compute_instance_group(resourceData map[string]interface{}, resourceName string) *google_compute_instance_group {

	resource := new(google_compute_instance_group)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.NamedPort = CreateStructSlice(attributes_map, attributesKey, "named_port", 1, NewNamedPort)

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
