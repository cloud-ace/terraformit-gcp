package resources

type google_compute_ssl_certificate struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			Certificate       string `json:"certificate"`
			CertificateID     string `json:"certificate_id"`
			CreationTimestamp string `json:"creation_timestamp"`
			Description       string `json:"description"`
			ID                string `json:"id"`
			Name              string `json:"name"`
			Project           string `json:"project"`
			SelfLink          string `json:"self_link"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_compute_ssl_certificate(resourceData map[string]interface{}, resourceName string) *google_compute_ssl_certificate {

	resource := new(google_compute_ssl_certificate)
	//attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	//attributesKey := CreateAttributesKey(attributes_map)
	//Struct
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
