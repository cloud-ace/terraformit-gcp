package resources

type google_compute_global_forwarding_rule struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			Description string            `json:"description"`
			ID          string            `json:"id"`
			IPAddress   string            `json:"ip_address"`
			IPProtocol  string            `json:"ip_protocol"`
			IPVersion   string            `json:"ip_version"`
			Labels      map[string]string `json:"labels"`
			Name        string            `json:"name"`
			PortRange   string            `json:"port_range"`
			Project     string            `json:"project"`
			SelfLink    string            `json:"self_link"`
			Target      string            `json:"target"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_compute_global_forwarding_rule(resourceData map[string]interface{}, resourceName string) *google_compute_global_forwarding_rule {

	resource := new(google_compute_global_forwarding_rule)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	//attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.Labels = CreateMap(attributes_map, "labels")

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
