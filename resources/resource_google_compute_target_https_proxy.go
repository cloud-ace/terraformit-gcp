package resources

type google_compute_target_https_proxy struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			CreationTimestamp string   `json:"creation_timestamp"`
			Description       string   `json:"description"`
			ID                string   `json:"id"`
			Name              string   `json:"name"`
			Project           string   `json:"project"`
			ProxyID           string   `json:"proxy_id"`
			QuicOverride      string   `json:"quic_override"`
			SelfLink          string   `json:"self_link"`
			SslCertificates   []string `json:"ssl_certificates"`
			SslPolicy         string   `json:"ssl_policy"`
			URLMap            string   `json:"url_map"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_compute_target_https_proxy(resourceData map[string]interface{}, resourceName string) *google_compute_target_https_proxy {

	resource := new(google_compute_target_https_proxy)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	//Struct
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.SslCertificates = CreateStringSlice(attributes_map, attributesKey, "ssl_certificates")

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
