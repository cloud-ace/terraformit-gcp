package resources

type google_service_account struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			AccountID   string `json:"account_id"`
			DisplayName string `json:"display_name"`
			Email       string `json:"email"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			Project     string `json:"project"`
			UniqueID    string `json:"unique_id"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_service_account(resourceData map[string]interface{}, resourceName string) *google_service_account {

	resource := new(google_service_account)
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
