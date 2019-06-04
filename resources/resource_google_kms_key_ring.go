package resources

type google_kms_key_ring struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			ID       string `json:"id"`
			Location string `json:"location"`
			Name     string `json:"name"`
			Project  string `json:"project"`
			SelfLink string `json:"self_link"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_kms_key_ring(resourceData map[string]interface{}, resourceName string) *google_kms_key_ring {

	resource := new(google_kms_key_ring)
	//attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	//attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform

	//beta
	return resource
}
