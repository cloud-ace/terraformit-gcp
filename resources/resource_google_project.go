package resources

type google_project struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			AppEngine         string            `json:"app_engine.#"`
			AutoCreateNetwork string            `json:"auto_create_network"`
			BillingAccount    string            `json:"billing_account"`
			FolderID          string            `json:"folder_id"`
			ID                string            `json:"id"`
			Labels            map[string]string `json:"labels.%"`
			Name              string            `json:"name"`
			Number            string            `json:"number"`
			OrgID             string            `json:"org_id"`
			ProjectID         string            `json:"project_id"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_project(resourceData map[string]interface{}, resourceName string) *google_project {

	resource := new(google_project)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	//attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))
	resource.Primary.Attributes.Labels = CreateMap(attributes_map, "labels")
	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
