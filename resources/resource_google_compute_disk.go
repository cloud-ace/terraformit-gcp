package resources

import (
	"strings"
)

type google_compute_disk struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			CreationTimestamp           string            `json:"creation_timestamp"`
			Description                 string            `json:"description"`
			DiskEncryptionKey           []Params          `json:"disk_encryption_key"`
			ID                          string            `json:"id"`
			Image                       string            `json:"image"`
			LabelFingerprint            string            `json:"label_fingerprint"`
			Labels                      map[string]string `json:"labels"`
			LastAttachTimestamp         string            `json:"last_attach_timestamp"`
			LastDetachTimestamp         string            `json:"last_detach_timestamp"`
			Name                        string            `json:"name"`
			PhysicalBlockSizeBytes      string            `json:"physical_block_size_bytes"`
			Project                     string            `json:"project"`
			SelfLink                    string            `json:"self_link"`
			Size                        string            `json:"size"`
			Snapshot                    string            `json:"snapshot"`
			SourceImageEncryptionKey    []Params          `json:"source_image_encryption_key"`
			SourceSnapshotEncryptionKey []Params          `json:"source_snapshot_encryption_key"`
			SourceSnapshotID            string            `json:"source_snapshot_id"`
			Type                        string            `json:"type"`
			Users                       []string          `json:"users"`
			Zone                        string            `json:"zone"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}
type EncryptionKey struct {
	Number         string
	KmsKeySelfLink string
	RawKey         string
	Sha256         string
}

func (d *EncryptionKey) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "kms_key_self_link") {
		d.KmsKeySelfLink = value
	} else if strings.Contains(key, "raw_key") {
		d.RawKey = value
	} else if strings.Contains(key, "sha256") {
		d.Sha256 = value
	}
}

func (d *EncryptionKey) SetNumber(num string) {
	d.Number = num
}

func (d *EncryptionKey) GetNumber() string {
	return d.Number
}

func NewEncryptionKey(num string) Params {
	var s Params = new(EncryptionKey)
	s.SetNumber(num)
	return s
}

func Newgoogle_compute_disk(resourceData map[string]interface{}, resourceName string) *google_compute_disk {

	resource := new(google_compute_disk)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.Labels = CreateMap(attributes_map, "labels")
	resource.Primary.Attributes.Users = CreateStringSlice(attributes_map, attributesKey, "users")
	resource.Primary.Attributes.DiskEncryptionKey = CreateStructSlice(attributes_map, attributesKey, "disk_encryption_key", 1, NewEncryptionKey)
	resource.Primary.Attributes.SourceImageEncryptionKey = CreateStructSlice(attributes_map, attributesKey, "source_image_encryption_key", 1, NewEncryptionKey)
	resource.Primary.Attributes.SourceSnapshotEncryptionKey = CreateStructSlice(attributes_map, attributesKey, "source_snapshot_encryption_key", 1, NewEncryptionKey)
	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
