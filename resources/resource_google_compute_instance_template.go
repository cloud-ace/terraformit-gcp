package resources

import (
	"strings"
)

type google_compute_instance_template struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			CanIPForward        string            `json:"can_ip_forward"`
			Description         string            `json:"description"`
			Disk                []Params          `json:"disk.#"`
			ID                  string            `json:"id"`
			InstanceDescription string            `json:"instance_description"`
			Labels              map[string]string `json:"labels"`
			MachineType         string            `json:"machine_type"`
			Metadata            map[string]string `json:"metadata"`
			MetadataFingerprint string            `json:"metadata_fingerprint"`
			MinCPUPlatform      string            `json:"min_cpu_platform"`
			Name                string            `json:"name"`
			NetworkInterface    []Params          `json:"network_interface"`
			Project             string            `json:"project"`
			Scheduling          []Params          `json:"scheduling"`
			SelfLink            string            `json:"self_link"`
			ServiceAccount      []Params          `json:"service_account"`
			Tags                []string          `json:"tags"`
			TagsFingerprint     string            `json:"tags_fingerprint"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type NetworkInterface struct {
	//paramsの2重構造は対応できないから、配列で持って要素で関連づけ
	Number                            string
	AccessConfigAssignedNatIP         []string `json:"assigned_nat_ip"`
	AccessConfigNatIP                 []string `json:"nat_ip"`
	AccessConfigNetworkTier           []string `json:"network_tier"`
	Address                           string   `json:"address"`
	AliasIPRangeIPCidrRange           []string `json:"alias_ip_range"`
	AliasIPRangeIPSubnetworkRangeName []string `json:"alias_ip_range.#"`
	Network                           string   `json:"network"`
	NetworkIP                         string   `json:"network_ip"`
	Subnetwork                        string   `json:"subnetwork"`
	SubnetworkProject                 string   `json:"subnetwork_project"`
}

func (d *NetworkInterface) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "access_config") && strings.Contains(key, "assigned_nat_ip") {
		d.AccessConfigAssignedNatIP = append(d.AccessConfigAssignedNatIP, value)
	} else if strings.Contains(key, "access_config") && strings.Contains(key, "nat_ip") && !strings.Contains(key, "assigned_nat_ip") {
		d.AccessConfigNatIP = append(d.AccessConfigNatIP, value)
	} else if strings.Contains(key, "access_config") && strings.Contains(key, "network_tier") {
		d.AccessConfigNetworkTier = append(d.AccessConfigNetworkTier, value)
	} else if strings.Contains(key, "address") {
		d.Address = value
	} else if strings.Contains(key, "alias_ip_range") && strings.Contains(key, "ip_cidr_range") {
		d.AliasIPRangeIPCidrRange = append(d.AliasIPRangeIPCidrRange, value)
	} else if strings.Contains(key, "alias_ip_range") && !strings.Contains(key, "subnetwork_range_name") {
		d.AliasIPRangeIPSubnetworkRangeName = append(d.AliasIPRangeIPSubnetworkRangeName, value)
	} else if strings.Contains(key, "network_ip") {
		d.NetworkIP = value
	} else if strings.Contains(key, "subnetwork") && !strings.Contains(key, "subnetwork_project") {
		d.Subnetwork = value
	} else if strings.Contains(key, "subnetwork_project") {
		d.SubnetworkProject = value
	} else if strings.Contains(key, "network") && !strings.Contains(key, "subnetwork") && !strings.Contains(key, "network_tier") && !strings.Contains(key, "network_ip") {
		d.Network = value
	}
}

func (d *NetworkInterface) SetNumber(num string) {
	d.Number = num
}

func (d *NetworkInterface) GetNumber() string {
	return d.Number
}

func NewNetworkInterface(num string) Params {
	var s Params = new(NetworkInterface)
	s.SetNumber(num)
	return s
}

type NetworkInterfaceAliasIPRange struct {
	Number string
}

type Disk struct {
	Number                string
	DiskAutoDelete        string           `json:"disk.auto_delete"`
	DiskBoot              string           `json:"disk.boot"`
	DiskDeviceName        string           `json:"disk.device_name"`
	DiskDiskEncryptionKey []*EncryptionKey `json:"disk.disk_encryption_key.#"`
	DiskDiskName          string           `json:"disk.disk_name"`
	DiskDiskSizeGb        string           `json:"disk.disk_size_gb"`
	DiskDiskType          string           `json:"disk.disk_type"`
	DiskInterface         string           `json:"disk.interface"`
	DiskMode              string           `json:"disk.mode"`
	DiskSource            string           `json:"disk.source"`
	DiskSourceImage       string           `json:"disk.source_image"`
	DiskType              string           `json:"disk.type"`
}

func (d *Disk) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "auto_delete") {
		d.DiskAutoDelete = value
	} else if strings.Contains(key, "boot") {
		d.DiskBoot = value
	} else if strings.Contains(key, "device_name") {
		d.DiskDeviceName = value
	} else if strings.Contains(key, "disk_encryption_key") && !strings.Contains(key, "#") {
		d.DiskDiskEncryptionKey = append(d.DiskDiskEncryptionKey, NewDiskEncryptionKeyforDisk(value))
	} else if strings.Contains(key, "disk_name") {
		d.DiskDiskName = value
	} else if strings.Contains(key, "disk_size_gb") {
		d.DiskDiskSizeGb = value
	} else if strings.Contains(key, "disk_type") {
		d.DiskDiskType = value
	} else if strings.Contains(key, "interface") {
		d.DiskInterface = value
	} else if strings.Contains(key, "mode") {
		d.DiskMode = value
	} else if strings.Contains(key, "source_image") {
		d.DiskSourceImage = value
	} else if strings.Contains(key, "source") {
		d.DiskSource = value
	} else if strings.Contains(key, "type") {
		d.DiskType = value
	}
}

func (d *Disk) SetNumber(num string) {
	d.Number = num
}

func (d *Disk) GetNumber() string {
	return d.Number
}

func NewDisk(num string) Params {
	var s Params = new(Disk)
	s.SetNumber(num)
	return s
}

//bimyou
func NewDiskEncryptionKeyforDisk(v string) *EncryptionKey {
	s := new(EncryptionKey)
	s.KmsKeySelfLink = v
	return s
}

type Scheduling struct {
	Number            string
	AutomaticRestart  string `json:"automatic_restart"`
	OnHostMaintenance string `json:"on_host_maintenance"`
	Preemptible       string `json:"preemptible"`
}

func (d *Scheduling) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "automatic_restart") {
		d.AutomaticRestart = value
	} else if strings.Contains(key, "on_host_maintenance") {
		d.OnHostMaintenance = value
	} else if strings.Contains(key, "preemptible") {
		d.Preemptible = value
	}
}

func (d *Scheduling) SetNumber(num string) {
	d.Number = num
}

func (d *Scheduling) GetNumber() string {
	return d.Number
}

func NewScheduling(num string) Params {
	var s Params = new(Scheduling)
	s.SetNumber(num)
	return s
}

type ServiceAccount struct {
	Number string
	Email  string   `json:"service_account.email"`
	Scopes []string `json:"service_account.scopes"`
}

func (d *ServiceAccount) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "email") {
		d.Email = value
	} else if strings.Contains(key, "scopes") {
		d.Scopes = append(d.Scopes, value)
	}
}

func (d *ServiceAccount) SetNumber(num string) {
	d.Number = num
}

func (d *ServiceAccount) GetNumber() string {
	return d.Number
}

func NewServiceAccount(num string) Params {
	var s Params = new(ServiceAccount)
	s.SetNumber(num)
	return s
}

func Newgoogle_compute_instance_template(resourceData map[string]interface{}, resourceName string) *google_compute_instance_template {

	resource := new(google_compute_instance_template)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.Labels = CreateMap(attributes_map, "labels")
	resource.Primary.Attributes.Metadata = CreateMap(attributes_map, "metadata.")
	resource.Primary.Attributes.NetworkInterface = CreateStructSlice(attributes_map, attributesKey, "network_interface", 1, NewNetworkInterface)
	resource.Primary.Attributes.Scheduling = CreateStructSlice(attributes_map, attributesKey, "scheduling", 1, NewScheduling)
	resource.Primary.Attributes.ServiceAccount = CreateStructSlice(attributes_map, attributesKey, "service_account", 1, NewServiceAccount)
	resource.Primary.Attributes.Tags = CreateStringSlice(attributes_map, attributesKey, "tags")
	resource.Primary.Attributes.Disk = CreateStructSlice(attributes_map, attributesKey, "disk", 1, NewDisk)

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
