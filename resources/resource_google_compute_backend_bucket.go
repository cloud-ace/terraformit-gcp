package resources

type google_compute_backend_bucket struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			BucketName        string   `json:"bucket_name"`
			CdnPolicy         []Params `json:"cdn_policy"`
			CreationTimestamp string   `json:"creation_timestamp"`
			Description       string   `json:"description"`
			EnableCdn         string   `json:"enable_cdn"`
			ID                string   `json:"id"`
			Name              string   `json:"name"`
			Project           string   `json:"project"`
			SelfLink          string   `json:"self_link"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_compute_backend_bucket(resourceData map[string]interface{}, resourceName string) *google_compute_backend_bucket {

	resource := new(google_compute_backend_bucket)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.CdnPolicy = CreateStructSlice(attributes_map, attributesKey, "cdn_policy", 1, NewCdnPolicy)

	return resource
}
