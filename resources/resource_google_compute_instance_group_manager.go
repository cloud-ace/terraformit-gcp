package resources

type google_compute_instance_group_manager struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			BaseInstanceName    string   `json:"base_instance_name"`
			Description         string   `json:"description"`
			Fingerprint         string   `json:"fingerprint"`
			ID                  string   `json:"id"`
			InstanceGroup       string   `json:"instance_group"`
			InstanceTemplate    string   `json:"instance_template"`
			Name                string   `json:"name"`
			NamedPort           []Params `json:"named_port.#"`
			Project             string   `json:"project"`
			RollingUpdatePolicy string   `json:"rolling_update_policy.#"`
			SelfLink            string   `json:"self_link"`
			TargetPools         []string `json:"target_pools.#"`
			TargetSize          string   `json:"target_size"`
			UpdateStrategy      string   `json:"update_strategy"`
			Version             string   `json:"version.#"`
			WaitForInstances    string   `json:"wait_for_instances"`
			Zone                string   `json:"zone"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_compute_instance_group_manager(resourceData map[string]interface{}, resourceName string) *google_compute_instance_group_manager {

	resource := new(google_compute_instance_group_manager)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.TargetPools = CreateStringSlice(attributes_map, attributesKey, "target_pools")
	resource.Primary.Attributes.NamedPort = CreateStructSlice(attributes_map, attributesKey, "named_port", 1, NewNamedPort)

	//beta
	if resourceData["provider"] == "provider.google-beta" {
		//rolling update
		//version
	}
	return resource
}
