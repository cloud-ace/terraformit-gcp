package resources

type google_compute_target_pool struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			BackupPool      string   `json:"backup_pool"`
			Description     string   `json:"description"`
			FailoverRatio   string   `json:"failover_ratio"`
			HealthChecks    []string `json:"health_checks"`
			ID              string   `json:"id"`
			Instances       []string `json:"instances"`
			Name            string   `json:"name"`
			Project         string   `json:"project"`
			Region          string   `json:"region"`
			SelfLink        string   `json:"self_link"`
			SessionAffinity string   `json:"session_affinity"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_compute_target_pool(resourceData map[string]interface{}, resourceName string) *google_compute_target_pool {

	resource := new(google_compute_target_pool)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))
	resource.Primary.Attributes.HealthChecks = CreateStringSlice(attributes_map, attributesKey, "health_checks")
	resource.Primary.Attributes.Instances = CreateStringSlice(attributes_map, attributesKey, "instances")
	//attributes need to transform

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
