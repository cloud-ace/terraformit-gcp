package resources

type google_compute_image struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			ArchiveSizeBytes  string            `json:"archive_size_bytes"`
			CreationTimestamp string            `json:"creation_timestamp"`
			Description       string            `json:"description"`
			DiskSizeGb        string            `json:"disk_size_gb"`
			Family            string            `json:"family"`
			ID                string            `json:"id"`
			LabelFingerprint  string            `json:"label_fingerprint"`
			Labels            map[string]string `json:"labels"`
			Licenses          []string          `json:"licenses"`
			Name              string            `json:"name"`
			Project           string            `json:"project"`
			SelfLink          string            `json:"self_link"`
			SourceDisk        string            `json:"source_disk"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

func Newgoogle_compute_image(resourceData map[string]interface{}, resourceName string) *google_compute_image {

	resource := new(google_compute_image)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.Labels = CreateMap(attributes_map, "labels")
	resource.Primary.Attributes.Licenses = CreateStringSlice(attributes_map, attributesKey, "licenses")
	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
