package resources

type google_compute_snapshot struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			CreationTimestamp     string            `json:"creation_timestamp"`
			Description           string            `json:"description"`
			DiskSizeGb            string            `json:"disk_size_gb"`
			ID                    string            `json:"id"`
			LabelFingerprint      string            `json:"label_fingerprint"`
			Labels                map[string]string `json:"labels"`
			Name                  string            `json:"name"`
			Project               string            `json:"project"`
			SelfLink              string            `json:"self_link"`
			SnapshotEncryptionKey []Params          `json:"snapshot_encryption_key"`
			SnapshotID            string            `json:"snapshot_id"`
			SourceDisk            string            `json:"source_disk"`
			SourceDiskLink        string            `json:"source_disk_link"`
			StorageBytes          string            `json:"storage_bytes"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type SnapshotEncryptionKey struct {
	Number string
	RawKey string
}

func (s *SnapshotEncryptionKey) SetRawKey(RawKey string) {
	s.RawKey = RawKey
}
func NewSnapshotEncryptionKey(num string) *SnapshotEncryptionKey {
	s := new(SnapshotEncryptionKey)
	s.Number = num
	return s
}

func Newgoogle_compute_snapshot(resourceData map[string]interface{}, resourceName string) *google_compute_snapshot {

	resource := new(google_compute_snapshot)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.Labels = CreateMap(attributes_map, "labels")
	resource.Primary.Attributes.SnapshotEncryptionKey = CreateStructSlice(attributes_map, attributesKey, "snapshot_encryption_key", 1, NewEncryptionKey)

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
