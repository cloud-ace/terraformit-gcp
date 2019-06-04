package resources

import (
	"strings"
)

type google_storage_bucket struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			Cors          []Params          `json:"cors"`
			Encryption    []Params          `json:"encryption"`
			ForceDestroy  string            `json:"force_destroy"`
			ID            string            `json:"id"`
			Labels        map[string]string `json:"labels"`
			LifecycleRule []Params          `json:"lifecycle_rule"`
			Location      string            `json:"location"`
			Logging       []Params          `json:"logging"`
			Name          string            `json:"name"`
			Project       string            `json:"project"`
			RequesterPays string            `json:"requester_pays"`
			SelfLink      string            `json:"self_link"`
			StorageClass  string            `json:"storage_class"`
			URL           string            `json:"url"`
			Versioning    []Params          `json:"versioning"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type Cors struct {
	Number         string
	MaxAgeSeconds  string   `json:"max_age_seconds"`
	Method         []string `json:"method"`
	Origin         []string `json:"origin"`
	ResponseHeader []string `json:"response_header"`
}

func (d *Cors) SetNumber(num string) {
	d.Number = num
}
func (d *Cors) GetNumber() string {
	return d.Number
}
func NewCors(num string) Params {
	var s Params = new(Cors)
	s.SetNumber(num)
	return s
}
func (d *Cors) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "max_age_seconds") {
		d.MaxAgeSeconds = value
	} else if strings.Contains(key, "method") {
		d.Method = append(d.Method, value)
	} else if strings.Contains(key, "origin") {
		d.Origin = append(d.Origin, value)
	} else if strings.Contains(key, "response_header") {
		d.ResponseHeader = append(d.ResponseHeader, value)
	}
}

type LifecycleRule struct {
	Number    string
	Action    []Params `json:"lifecycle_rule.action"`
	Condition []Params `json:"lifecycle_rule.condition"`
}

func (d *LifecycleRule) SetNumber(num string) {
	d.Number = num
}
func (d *LifecycleRule) GetNumber() string {
	return d.Number
}
func NewLifecycleRule(num string) Params {
	var s Params = new(LifecycleRule)
	s.SetNumber(num)
	return s
}

func (d *LifecycleRule) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "action") {
		//lifecycle_rule.0.actionまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.Action = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewAction)
	} else if strings.Contains(key, "condition") {
		//lifecycle_rule.0.conditionまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.Condition = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewCondition)
	}
}

type Action struct {
	Number       string
	StorageClass string `json:"storage_class"`
	Type         string `json:"type"`
}

func (d *Action) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "storage_class") {
		d.StorageClass = value
	} else if strings.Contains(key, "type") {
		d.Type = value
	}
}

func (d *Action) SetNumber(num string) {
	d.Number = num
}

func (d *Action) GetNumber() string {
	return d.Number
}

func NewAction(num string) Params {
	var s Params = new(Action)
	s.SetNumber(num)
	return s
}

type Condition struct {
	Number              string
	Age                 string   `json:"age"`
	CreatedBefore       string   `json:"created_before"`
	IsLive              string   `json:"is_live"`
	MatchesStorageClass []string `json:"matches_storage_class"`
	NumNewerVersions    string   `json:"num_newer_versions"`
	WithState           string   `json:"with_state"`
}

func (d *Condition) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "age") && !strings.Contains(key, "storage") {
		d.Age = value
	} else if strings.Contains(key, "created_before") {
		d.CreatedBefore = value
	} else if strings.Contains(key, "is_live") {
		d.IsLive = value
	} else if strings.Contains(key, "matches_storage_class") {
		d.MatchesStorageClass = append(d.MatchesStorageClass, value)
	} else if strings.Contains(key, "num_newer_versions") {
		d.NumNewerVersions = value
	} else if strings.Contains(key, "with_state") {
		d.WithState = value
	}
}

func (d *Condition) SetNumber(num string) {
	d.Number = num
}

func (d *Condition) GetNumber() string {
	return d.Number
}

func NewCondition(num string) Params {
	var s Params = new(Condition)
	s.SetNumber(num)
	return s
}

type Encryption struct {
	Number            string
	DefaultKmsKeyName string `json:"default_kms_key_name"`
}

func (d *Encryption) SetNumber(num string) {
	d.Number = num
}
func (d *Encryption) GetNumber() string {
	return d.Number
}
func NewEncryption(num string) Params {
	var s Params = new(Encryption)
	s.SetNumber(num)
	return s
}

func (d *Encryption) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "default_kms_key_name") {
		d.DefaultKmsKeyName = value
	}
}

type Versioning struct {
	Number  string
	Enabled string `json:"enabled"`
}

func (d *Versioning) SetNumber(num string) {
	d.Number = num
}

func (d *Versioning) GetNumber() string {
	return d.Number
}

func NewVersioning(num string) Params {
	var s Params = new(Versioning)
	s.SetNumber(num)
	return s
}

func (d *Versioning) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "enabled") {
		d.Enabled = value
	}
}

type Logging struct {
	Number          string
	LogBucket       string `json:"log_bucket"`
	LogObjectPrefix string `json:"log_object_prefix"`
}

func (d *Logging) SetNumber(num string) {
	d.Number = num
}

func (d *Logging) GetNumber() string {
	return d.Number
}

func NewLogging(num string) Params {
	var s Params = new(Logging)
	s.SetNumber(num)
	return s
}

func (d *Logging) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "log_bucket") {
		d.LogBucket = value
	} else if strings.Contains(key, "log_object_prefix") {
		d.LogObjectPrefix = value
	}
}

func Newgoogle_storage_bucket(resourceData map[string]interface{}, resourceName string) *google_storage_bucket {

	resource := new(google_storage_bucket)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.Cors = CreateStructSlice(attributes_map, attributesKey, "cors", 1, NewCors)
	resource.Primary.Attributes.Encryption = CreateStructSlice(attributes_map, attributesKey, "encryption", 1, NewCors)
	resource.Primary.Attributes.Labels = CreateMap(attributes_map, "labels")
	resource.Primary.Attributes.LifecycleRule = CreateStructSlice(attributes_map, attributesKey, "lifecycle_rule", 1, NewLifecycleRule)
	resource.Primary.Attributes.Logging = CreateStructSlice(attributes_map, attributesKey, "logging", 1, NewLogging)
	resource.Primary.Attributes.Versioning = CreateStructSlice(attributes_map, attributesKey, "versioning", 1, NewVersioning)

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
