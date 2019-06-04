package resources

type google_compute_network struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			AutoCreateSubnetworks       string `json:"auto_create_subnetworks"`
			DeleteDefaultRoutesOnCreate string `json:"delete_default_routes_on_create"`
			Description                 string `json:"description"`
			GatewayIpv4                 string `json:"gateway_ipv4"`
			ID                          string `json:"id"`
			Ipv4Range                   string `json:"ipv4_range"`
			Name                        string `json:"name"`
			Project                     string `json:"project"`
			RoutingMode                 string `json:"routing_mode"`
			SelfLink                    string `json:"self_link"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_compute_network(resourceData map[string]interface{}, resourceName string) *google_compute_network {

	resource := new(google_compute_network)
	//attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	//attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
