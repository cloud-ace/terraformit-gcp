package resources

import (
	"strings"
)

type google_compute_instance struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			AttachedDisk                    []Params          `json:"attached_disk.#"`
			BootDisk                        string            `json:"boot_disk.#"`
			BootDiskAutoDelete              string            `json:"boot_disk.auto_delete"`
			BootDiskDeviceName              string            `json:"boot_disk.device_name"`
			BootDiskDiskEncryptionKeyRaw    string            `json:"boot_disk.disk_encryption_key_raw"`
			BootDiskDiskEncryptionKeySha256 string            `json:"boot_disk.disk_encryption_key_sha256"`
			BootDiskInitializeParamsImage   string            `json:"boot_disk.initialize_params.1.image"`
			BootDiskInitializeParamsSize    string            `json:"boot_disk.initialize_params.1.size"`
			BootDiskInitializeParamsType    string            `json:"boot_disk.initialize_params.1.type"`
			BootDiskSource                  string            `json:"boot_disk.source"`
			CanIPForward                    string            `json:"can_ip_forward"`
			CPUPlatform                     string            `json:"cpu_platform"`
			DeletionProtection              string            `json:"deletion_protection"`
			GuestAccelerator                []Params          `json:"guest_accelerator.#"`
			Hostname                        string            `json:"hostname"`
			ID                              string            `json:"id"`
			InstanceID                      string            `json:"instance_id"`
			LabelFingerprint                string            `json:"label_fingerprint"`
			Labels                          map[string]string `json:"labels"`
			MachineType                     string            `json:"machine_type"`
			Metadata                        map[string]string `json:"metadata"`
			MetadataFingerprint             string            `json:"metadata_fingerprint"`
			MetadataStartupScript           string            `json:"metadata_startup_script"`
			MinCPUPlatform                  string            `json:"min_cpu_platform"`
			Name                            string            `json:"name"`
			NetworkInterface                []Params          `json:"network_interface"`
			Project                         string            `json:"project"`
			Scheduling                      []Params          `json:"scheduling"`
			ScratchDisk                     []Params          `json:"scratch_disk"`
			SelfLink                        string            `json:"self_link"`
			ServiceAccount                  []Params          `json:"service_account"`
			Tags                            []string          `json:"tags"`
			TagsFingerprint                 string            `json:"tags_fingerprint"`
			Zone                            string            `json:"zone"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type AttachedDisk struct {
	Number              string
	DeviceName          string `json:"device_name"`
	EncryptionKeyRaw    string `json:"disk_encryption_key_raw"`
	EncryptionKeySha256 string `json:"disk_encryption_key_sha256"`
	Mode                string `json:"mode"`
	Source              string `json:"source"`
}

func (d *AttachedDisk) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "device_name") {
		d.DeviceName = value
	} else if strings.Contains(key, "disk_encryption_key_raw") {
		d.EncryptionKeyRaw = value
	} else if strings.Contains(key, "disk_encryption_key_sha256") {
		d.EncryptionKeySha256 = value
	} else if strings.Contains(key, "mode") {
		d.Mode = value
	} else if strings.Contains(key, "source") {
		d.Source = value
	}
}

func (d *AttachedDisk) SetNumber(num string) {
	d.Number = num
}

func (d *AttachedDisk) GetNumber() string {
	return d.Number
}

func NewAttachedDisk(num string) Params {
	var s Params = new(AttachedDisk)
	s.SetNumber(num)
	return s
}

type GuestAccelerator struct {
	Number string
	Type   string `json:"type"`
	Count  string `json:"count"`
}

func (d *GuestAccelerator) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "type") {
		d.Type = value
	} else if strings.Contains(key, "count") {
		d.Count = value
	}
}

func (d *GuestAccelerator) SetNumber(num string) {
	d.Number = num
}

func (d *GuestAccelerator) GetNumber() string {
	return d.Number
}

func NewGuestAccelerator(num string) Params {
	var s Params = new(GuestAccelerator)
	s.SetNumber(num)
	return s
}

type ScratchDisk struct {
	Number    string
	Interface string `json:"interface"`
}

func (d *ScratchDisk) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "interface") {
		d.Interface = value
	}
}

func (d *ScratchDisk) SetNumber(num string) {
	d.Number = num
}

func (d *ScratchDisk) GetNumber() string {
	return d.Number
}

func NewScratchDisk(num string) Params {
	var s Params = new(ScratchDisk)
	s.SetNumber(num)
	return s
}

func Newgoogle_compute_instance(resourceData map[string]interface{}, resourceName string) *google_compute_instance {

	resource := new(google_compute_instance)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.AttachedDisk = CreateStructSlice(attributes_map, attributesKey, "attached_disk", 1, NewAttachedDisk)
	//BootDiskは１つという前提
	resource.Primary.Attributes.BootDiskAutoDelete = attributes_map["boot_disk.0.auto_delete"].(string)
	resource.Primary.Attributes.BootDiskDeviceName = attributes_map["boot_disk.0.device_name"].(string)
	resource.Primary.Attributes.BootDiskDiskEncryptionKeyRaw = attributes_map["boot_disk.0.disk_encryption_key_raw"].(string)
	resource.Primary.Attributes.BootDiskDiskEncryptionKeySha256 = attributes_map["boot_disk.0.disk_encryption_key_sha256"].(string)
	resource.Primary.Attributes.BootDiskInitializeParamsImage = attributes_map["boot_disk.0.initialize_params.0.image"].(string)
	resource.Primary.Attributes.BootDiskInitializeParamsSize = attributes_map["boot_disk.0.initialize_params.0.size"].(string)
	resource.Primary.Attributes.BootDiskInitializeParamsType = attributes_map["boot_disk.0.initialize_params.0.type"].(string)
	resource.Primary.Attributes.BootDiskSource = attributes_map["boot_disk.0.source"].(string)
	resource.Primary.Attributes.GuestAccelerator = CreateStructSlice(attributes_map, attributesKey, "guest_accelerator", 1, NewGuestAccelerator)
	resource.Primary.Attributes.Labels = CreateMap(attributes_map, "labels")
	resource.Primary.Attributes.Metadata = CreateMap(attributes_map, "metadata.")
	resource.Primary.Attributes.NetworkInterface = CreateStructSlice(attributes_map, attributesKey, "network_interface", 1, NewNetworkInterface)
	resource.Primary.Attributes.Scheduling = CreateStructSlice(attributes_map, attributesKey, "scheduling", 1, NewScheduling)
	resource.Primary.Attributes.ScratchDisk = CreateStructSlice(attributes_map, attributesKey, "scratch_disk", 1, NewScratchDisk)
	resource.Primary.Attributes.ServiceAccount = CreateStructSlice(attributes_map, attributesKey, "service_account", 1, NewServiceAccount)
	resource.Primary.Attributes.Tags = CreateStringSlice(attributes_map, attributesKey, "tags")

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
