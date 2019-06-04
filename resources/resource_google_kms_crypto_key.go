package resources

import "strings"

type google_kms_crypto_key struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			ID              string   `json:"id"`
			KeyRing         string   `json:"key_ring"`
			Name            string   `json:"name"`
			RotationPeriod  string   `json:"rotation_period"`
			SelfLink        string   `json:"self_link"`
			VersionTemplate []Params `json:"version_template"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type VersionTemplate struct {
	Number          string
	Algorithm       string `json:"algorithm"`
	ProtectionLevel string `json:"protection_level"`
}

func (d *VersionTemplate) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "algorithm") {
		d.Algorithm = value
	} else if strings.Contains(key, "protection_level") {
		d.ProtectionLevel = value
	}
}

func (d *VersionTemplate) SetNumber(num string) {
	d.Number = num
}

func (d *VersionTemplate) GetNumber() string {
	return d.Number
}

func NewVersionTemplate(num string) Params {
	var s Params = new(VersionTemplate)
	s.SetNumber(num)
	return s
}

func Newgoogle_kms_crypto_key(resourceData map[string]interface{}, resourceName string) *google_kms_crypto_key {

	resource := new(google_kms_crypto_key)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.VersionTemplate = CreateStructSlice(attributes_map, attributesKey, "version_template", 1, NewVersionTemplate)

	return resource
}
