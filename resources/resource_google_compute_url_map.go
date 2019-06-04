package resources

import (
	"strings"
)

type google_compute_url_map struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			CreationTimestamp string   `json:"creation_timestamp"`
			DefaultService    string   `json:"default_service"`
			Description       string   `json:"description"`
			Fingerprint       string   `json:"fingerprint"`
			HostRule          []Params `json:"host_rule"`
			ID                string   `json:"id"`
			MapID             string   `json:"map_id"`
			Name              string   `json:"name"`
			PathMatcher       []Params `json:"path_matcher"`
			Project           string   `json:"project"`
			SelfLink          string   `json:"self_link"`
			Test              []Params `json:"test"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type PathMatcher struct {
	Number         string
	DefaultService string `json:"default_service"`
	Description    string `json:"description"`
	Name           string `json:"name"`
	PathRule       []Params
}

func (d *PathMatcher) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "default_service") {
		d.DefaultService = value
	} else if strings.Contains(key, "path_rule") && strings.Contains(key, "service") {
		//path_matcher.#.path_ruleまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.PathRule = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewPathRule)

	} else if strings.Contains(key, "description") {
		d.Description = value
	} else if strings.Contains(key, "name") {
		d.Name = value
	}
}

func (d *PathMatcher) SetNumber(num string) {
	d.Number = num
}

func (d *PathMatcher) GetNumber() string {
	return d.Number
}

func NewPathMatcher(num string) Params {
	var s Params = new(PathMatcher)
	s.SetNumber(num)
	return s
}

type PathRule struct {
	Number  string
	Paths   []string
	Service string `json:".path_rule.service"`
}

func (d *PathRule) SetNumber(num string) {
	d.Number = num
}

func (d *PathRule) GetNumber() string {
	return d.Number
}

func NewPathRule(num string) Params {
	var s Params = new(PathRule)
	s.SetNumber(num)
	return s
}

func (d *PathRule) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "service") {
		d.Service = value
	} else if strings.Contains(key, "paths") {
		d.Paths = append(d.Paths, value)
	}
}

type HostRule struct {
	Number      string
	Description string   `json:"description"`
	Hosts       []string `json:"hosts"`
	PathMatcher string   `json:"path_matcher"`
}

func (d *HostRule) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "description") {
		d.Description = value
	} else if strings.Contains(key, "hosts") {
		d.Hosts = append(d.Hosts, value)
	} else if strings.Contains(key, "path_matcher") {
		d.PathMatcher = value
	}
}

func (d *HostRule) SetNumber(num string) {
	d.Number = num
}

func (d *HostRule) GetNumber() string {
	return d.Number
}

func NewHostRule(num string) Params {
	var s Params = new(HostRule)
	s.SetNumber(num)
	return s
}

type Test struct {
	Number      string
	Description string `json:"description"`
	Host        string `json:"host"`
	Path        string `json:"path"`
	Service     string `json:"service"`
}

func (d *Test) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "description") {
		d.Description = value
	} else if strings.Contains(key, "host") {
		d.Host = value
	} else if strings.Contains(key, "path") {
		d.Path = value
	} else if strings.Contains(key, "service") {
		d.Service = value
	}
}

func (d *Test) SetNumber(num string) {
	d.Number = num
}

func (d *Test) GetNumber() string {
	return d.Number
}

func NewTest(num string) Params {
	var s Params = new(HostRule)
	s.SetNumber(num)
	return s
}

func Newgoogle_compute_url_map(resourceData map[string]interface{}, resourceName string) *google_compute_url_map {

	resource := new(google_compute_url_map)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.HostRule = CreateStructSlice(attributes_map, attributesKey, "host_rule", 1, NewHostRule)
	resource.Primary.Attributes.PathMatcher = CreateStructSlice(attributes_map, attributesKey, "path_matcher", 1, NewPathMatcher)

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
