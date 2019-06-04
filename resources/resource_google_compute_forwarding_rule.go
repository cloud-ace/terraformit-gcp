package resources

type google_compute_forwarding_rule struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			AllPorts            string            `json:"all_ports"`
			BackendService      string            `json:"backend_service"`
			CreationTimestamp   string            `json:"creation_timestamp"`
			Description         string            `json:"description"`
			ID                  string            `json:"id"`
			IPAddress           string            `json:"ip_address"`
			IPProtocol          string            `json:"ip_protocol"`
			IPVersion           string            `json:"ip_version"`
			LoadBalancingScheme string            `json:"load_balancing_scheme"`
			Name                string            `json:"name"`
			Network             string            `json:"network"`
			NetworkTier         string            `json:"network_tier"`
			PortRange           string            `json:"port_range"`
			Ports               []string          `json:"ports"`
			Project             string            `json:"project"`
			Region              string            `json:"region"`
			SelfLink            string            `json:"self_link"`
			ServiceLabel        string            `json:"service_label"`
			ServiceName         string            `json:"service_name"`
			Subnetwork          string            `json:"subnetwork"`
			Target              string            `json:"target"`
			Labels              map[string]string `json:"labels"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_compute_forwarding_rule(resourceData map[string]interface{}, resourceName string) *google_compute_forwarding_rule {

	resource := new(google_compute_forwarding_rule)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.Ports = CreateStringSlice(attributes_map, attributesKey, "ports")
	resource.Primary.Attributes.Labels = CreateMap(attributes_map, "labels")

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
