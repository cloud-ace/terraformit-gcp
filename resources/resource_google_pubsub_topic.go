package resources

type google_pubsub_topic struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			ID      string            `json:"id"`
			Labels  map[string]string `json:"labels"`
			Name    string            `json:"name"`
			Project string            `json:"project"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_pubsub_topic(resourceData map[string]interface{}, resourceName string) *google_pubsub_topic {

	resource := new(google_pubsub_topic)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	//attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.Labels = CreateMap(attributes_map, "labels")

	return resource
}
