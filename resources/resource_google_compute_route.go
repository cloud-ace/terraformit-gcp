package resources

type google_compute_route struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			Description      string   `json:"description"`
			DestRange        string   `json:"dest_range"`
			ID               string   `json:"id"`
			Name             string   `json:"name"`
			Network          string   `json:"network"`
			NextHopGateway   string   `json:"next_hop_gateway"`
			NextHopInstance  string   `json:"next_hop_instance"`
			NextHopIP        string   `json:"next_hop_ip"`
			NextHopNetwork   string   `json:"next_hop_network"`
			NextHopVpnTunnel string   `json:"next_hop_vpn_tunnel"`
			Priority         string   `json:"priority"`
			Project          string   `json:"project"`
			SelfLink         string   `json:"self_link"`
			Tags             []string `json:"tags"`
			EnableLogging    string   `json:"enable_logging"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_compute_route(resourceData map[string]interface{}, resourceName string) *google_compute_route {

	resource := new(google_compute_route)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.Tags = CreateStringSlice(attributes_map, attributesKey, "tags")
	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
