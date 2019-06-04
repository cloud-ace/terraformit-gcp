package resources

type google_compute_http_health_check struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			CheckIntervalSec   string `json:"check_interval_sec"`
			CreationTimestamp  string `json:"creation_timestamp"`
			Description        string `json:"description"`
			HealthyThreshold   string `json:"healthy_threshold"`
			Host               string `json:"host"`
			ID                 string `json:"id"`
			Name               string `json:"name"`
			Port               string `json:"port"`
			Project            string `json:"project"`
			RequestPath        string `json:"request_path"`
			SelfLink           string `json:"self_link"`
			TimeoutSec         string `json:"timeout_sec"`
			UnhealthyThreshold string `json:"unhealthy_threshold"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_compute_http_health_check(resourceData map[string]interface{}, resourceName string) *google_compute_http_health_check {

	resource := new(google_compute_http_health_check)
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
