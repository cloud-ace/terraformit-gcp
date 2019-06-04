package resources

import (
	"strings"
)

type google_container_cluster struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			AdditionalZones                []string          `json:"additional_zones"`
			AddonsConfig                   []Params          `json:"addons_config"`
			ClusterAutoscaling             []Params          `json:"cluster_autoscaling"`
			ClusterIpv4Cidr                string            `json:"cluster_ipv4_cidr"`
			Description                    string            `json:"description"`
			EnableKubernetesAlpha          string            `json:"enable_kubernetes_alpha"`
			EnableLegacyAbac               string            `json:"enable_legacy_abac"`
			Endpoint                       string            `json:"endpoint"`
			ID                             string            `json:"id"`
			InitialNodeCount               string            `json:"initial_node_count"`
			InstanceGroupUrls              []string          `json:"instance_group_urls"`
			IPAllocationPolicy             []Params          `json:"ip_allocation_policy"`
			Location                       string            `json:"location"`
			LoggingService                 string            `json:"logging_service"`
			MaintenancePolicy              []Params          `json:"maintenance_policy"`
			MasterAuth                     []Params          `json:"master_auth"`
			MasterAuthorizedNetworksConfig []Params          `json:"master_authorized_networks_config"`
			MasterVersion                  string            `json:"master_version"`
			MonitoringService              string            `json:"monitoring_service"`
			Name                           string            `json:"name"`
			Network                        string            `json:"network"`
			NetworkPolicy                  []Params          `json:"network_policy"`
			NodeConfig                     []Params          `json:"node_config"`
			NodeLocations                  []string          `json:"node_locations"`
			NodePool                       []Params          `json:"node_pool"`
			NodeVersion                    string            `json:"node_version"`
			PrivateClusterConfig           []Params          `json:"private_cluster_config"`
			Project                        string            `json:"project"`
			ResourceLabels                 map[string]string `json:"resource_labels"`
			Subnetwork                     string            `json:"subnetwork"`
			Zone                           string            `json:"zone"`
			Region                         string            `json:"region"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type AddonsConfig struct {
	Number                   string
	HorizontalPodAutoscaling []Params `json:"horizontal_pod_autoscaling"`
	HTTPLoadBalancing        []Params `json:"http_load_balancing"`
	KubernetesDashboard      []Params `json:"kubernetes_dashboard"`
	NetworkPolicyConfig      []Params `json:"network_policy_config"`
}

func (d *AddonsConfig) SetNumber(num string) {
	d.Number = num
}
func (d *AddonsConfig) GetNumber() string {
	return d.Number
}
func NewAddonsConfig(num string) Params {
	var s Params = new(AddonsConfig)
	s.SetNumber(num)
	return s
}

func (d *AddonsConfig) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "horizontal_pod_autoscaling") {
		//lifecycle_rule.0.actionまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.HorizontalPodAutoscaling = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewHorizontalPodAutoscaling)
	} else if strings.Contains(key, "http_load_balancing") {
		//lifecycle_rule.0.conditionまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.HTTPLoadBalancing = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewHTTPLoadBalancing)
	} else if strings.Contains(key, "kubernetes_dashboard") {
		//lifecycle_rule.0.conditionまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.KubernetesDashboard = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewKubernetesDashboard)
	} else if strings.Contains(key, "network_policy_config") {
		//lifecycle_rule.0.conditionまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.NetworkPolicyConfig = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewNetworkPolicyConfig)
	}
}

type HorizontalPodAutoscaling struct {
	Number   string
	Disabled string `json:"disabled"`
}

func (d *HorizontalPodAutoscaling) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "disabled") {
		d.Disabled = value
	}
}

func (d *HorizontalPodAutoscaling) SetNumber(num string) {
	d.Number = num
}

func (d *HorizontalPodAutoscaling) GetNumber() string {
	return d.Number
}

func NewHorizontalPodAutoscaling(num string) Params {
	var s Params = new(HorizontalPodAutoscaling)
	s.SetNumber(num)
	return s
}

type HTTPLoadBalancing struct {
	Number   string
	Disabled string `json:"disabled"`
}

func (d *HTTPLoadBalancing) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "disabled") {
		d.Disabled = value
	}
}

func (d *HTTPLoadBalancing) SetNumber(num string) {
	d.Number = num
}

func (d *HTTPLoadBalancing) GetNumber() string {
	return d.Number
}

func NewHTTPLoadBalancing(num string) Params {
	var s Params = new(HTTPLoadBalancing)
	s.SetNumber(num)
	return s
}

type KubernetesDashboard struct {
	Number   string
	Disabled string `json:"disabled"`
}

func (d *KubernetesDashboard) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "disabled") {
		d.Disabled = value
	}
}

func (d *KubernetesDashboard) SetNumber(num string) {
	d.Number = num
}

func (d *KubernetesDashboard) GetNumber() string {
	return d.Number
}

func NewKubernetesDashboard(num string) Params {
	var s Params = new(KubernetesDashboard)
	s.SetNumber(num)
	return s
}

type NetworkPolicyConfig struct {
	Number   string
	Disabled string `json:"disabled"`
}

func (d *NetworkPolicyConfig) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "disabled") {
		d.Disabled = value
	}
}

func (d *NetworkPolicyConfig) SetNumber(num string) {
	d.Number = num
}

func (d *NetworkPolicyConfig) GetNumber() string {
	return d.Number
}

func NewNetworkPolicyConfig(num string) Params {
	var s Params = new(NetworkPolicyConfig)
	s.SetNumber(num)
	return s
}

type ClusterAutoscaling struct {
	Number         string
	Enabled        string   `json:"enabled"`
	ResourceLimits []Params `json:"resource_limits"`
}

func (d *ClusterAutoscaling) SetNumber(num string) {
	d.Number = num
}
func (d *ClusterAutoscaling) GetNumber() string {
	return d.Number
}
func NewClusterAutoscaling(num string) Params {
	var s Params = new(ClusterAutoscaling)
	s.SetNumber(num)
	return s
}

func (d *ClusterAutoscaling) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "enabled") {
		d.Enabled = value
	} else if strings.Contains(key, "resource_limits") {
		//lifecycle_rule.0.conditionまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.ResourceLimits = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewResourceLimits)
	}
}

type ResourceLimits struct {
	Number       string
	ResourceType string `json:"resource_type"`
	Minimum      string `json:"minimum"`
	Maximum      string `json:"maximum"`
}

func (d *ResourceLimits) SetNumber(num string) {
	d.Number = num
}
func (d *ResourceLimits) GetNumber() string {
	return d.Number
}
func NewResourceLimits(num string) Params {
	var s Params = new(ResourceLimits)
	s.SetNumber(num)
	return s
}

func (d *ResourceLimits) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "resource_type") {
		d.ResourceType = value
	} else if strings.Contains(key, "minimum") {
		d.Minimum = value
	} else if strings.Contains(key, "maximum") {
		d.Maximum = value
	}
}

type IPAllocationPolicy struct {
	Number                     string
	ClusterIpv4CidrBlock       string `json:"cluster_ipv4_cidr_block"`
	ClusterSecondaryRangeName  string `json:"cluster_secondary_range_name"`
	CreateSubnetwork           string `json:"create_subnetwork"`
	NodeIpv4CidrBlock          string `json:"node_ipv4_cidr_block"`
	ServicesIpv4CidrBlock      string `json:"services_ipv4_cidr_block"`
	ServicesSecondaryRangeName string `json:"services_secondary_range_name"`
	SubnetworkName             string `json:"subnetwork_name"`
	UseIPAliases               string `json:"use_ip_aliases"`
}

func (d *IPAllocationPolicy) SetNumber(num string) {
	d.Number = num
}
func (d *IPAllocationPolicy) GetNumber() string {
	return d.Number
}
func NewIPAllocationPolicy(num string) Params {
	var s Params = new(IPAllocationPolicy)
	s.SetNumber(num)
	return s
}

func (d *IPAllocationPolicy) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "cluster_ipv4_cidr_block") {
		d.ClusterIpv4CidrBlock = value
	} else if strings.Contains(key, "cluster_secondary_range_name") {
		d.ClusterSecondaryRangeName = value
	} else if strings.Contains(key, "create_subnetwork") {
		d.CreateSubnetwork = value
	} else if strings.Contains(key, "node_ipv4_cidr_block") {
		d.NodeIpv4CidrBlock = value
	} else if strings.Contains(key, "services_ipv4_cidr_block") {
		d.ServicesIpv4CidrBlock = value
	} else if strings.Contains(key, "services_secondary_range_name") {
		d.ServicesSecondaryRangeName = value
	} else if strings.Contains(key, "subnetwork_name") {
		d.SubnetworkName = value
	} else if strings.Contains(key, "use_ip_aliases") {
		d.UseIPAliases = value
	}
}

type MaintenancePolicy struct {
	Number                 string
	DailyMaintenanceWindow []Params `json:"daily_maintenance_window"`
}

func (d *MaintenancePolicy) SetNumber(num string) {
	d.Number = num
}
func (d *MaintenancePolicy) GetNumber() string {
	return d.Number
}
func NewMaintenancePolicy(num string) Params {
	var s Params = new(MaintenancePolicy)
	s.SetNumber(num)
	return s
}

func (d *MaintenancePolicy) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "daily_maintenance_window") {
		//lifecycle_rule.0.conditionまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.DailyMaintenanceWindow = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewDailyMaintenanceWindow)
	}
}

type DailyMaintenanceWindow struct {
	Number    string
	Duration  string `json:"duration"`
	StartTime string `json:"start_time"`
}

func (d *DailyMaintenanceWindow) SetNumber(num string) {
	d.Number = num
}
func (d *DailyMaintenanceWindow) GetNumber() string {
	return d.Number
}
func NewDailyMaintenanceWindow(num string) Params {
	var s Params = new(DailyMaintenanceWindow)
	s.SetNumber(num)
	return s
}

func (d *DailyMaintenanceWindow) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "duration") {
		d.Duration = value
	} else if strings.Contains(key, "start_time") {
		d.StartTime = value
	}
}

type MasterAuth struct {
	Number                  string
	ClientCertificate       string   `json:"client_certificate"`
	ClientCertificateConfig []Params `json:"client_certificate_config"`
	ClientKey               string   `json:"client_key"`
	ClusterCaCertificate    string   `json:"cluster_ca_certificate"`
	Password                string   `json:"password"`
	Username                string   `json:"username"`
}

func (d *MasterAuth) SetNumber(num string) {
	d.Number = num
}
func (d *MasterAuth) GetNumber() string {
	return d.Number
}
func NewMasterAuth(num string) Params {
	var s Params = new(MasterAuth)
	s.SetNumber(num)
	return s
}

func (d *MasterAuth) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "client_certificate") {
		d.ClientCertificate = value
	} else if strings.Contains(key, "client_certificate_config") {
		//TODO fill
		//d.ClientCertificateConfig = value
	} else if strings.Contains(key, "client_key") {
		d.ClientKey = value
	} else if strings.Contains(key, "cluster_ca_certificate") {
		d.ClusterCaCertificate = value
	} else if strings.Contains(key, "password") {
		d.Password = value
	} else if strings.Contains(key, "username") {
		d.Username = value
	}
}

type ClientCertificateConfig struct {
	//TODO fill
	Number string
}

type MasterAuthorizedNetworksConfig struct {
	Number     string
	CidrBlocks []Params `json:"cidr_blocks"`
}

func (d *MasterAuthorizedNetworksConfig) SetNumber(num string) {
	d.Number = num
}
func (d *MasterAuthorizedNetworksConfig) GetNumber() string {
	return d.Number
}
func NewMasterAuthorizedNetworksConfig(num string) Params {
	var s Params = new(MasterAuthorizedNetworksConfig)
	s.SetNumber(num)
	return s
}

func (d *MasterAuthorizedNetworksConfig) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "cidr_blocks") {
		Number := CreateNumberFromKey(key, 2)
		d.CidrBlocks = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewCidrBlocks)
	}
}

type CidrBlocks struct {
	Number      string
	CidrBlock   string `json:"cidr_block"`
	DisplayName string `json:"display_name"`
}

func (d *CidrBlocks) SetNumber(num string) {
	d.Number = num
}
func (d *CidrBlocks) GetNumber() string {
	return d.Number
}
func NewCidrBlocks(num string) Params {
	var s Params = new(CidrBlocks)
	s.SetNumber(num)
	return s
}

func (d *CidrBlocks) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "cidr_block") {
		d.CidrBlock = value
	} else if strings.Contains(key, "display_name") {
		d.DisplayName = value
	}
}

type NetworkPolicy struct {
	Number   string
	Enabled  string `json:"enabled"`
	Provider string `json:"provider"`
}

func (d *NetworkPolicy) SetNumber(num string) {
	d.Number = num
}
func (d *NetworkPolicy) GetNumber() string {
	return d.Number
}
func NewNetworkPolicy(num string) Params {
	var s Params = new(NetworkPolicy)
	s.SetNumber(num)
	return s
}

func (d *NetworkPolicy) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "enabled") {
		d.Enabled = value
	} else if strings.Contains(key, "provider") {
		d.Provider = value
	}
}

type NodeConfig struct {
	Number                 string
	DiskSizeGb             string            `json:"disk_size_gb"`
	DiskType               string            `json:"disk_type"`
	GuestAccelerator       []Params          `json:"guest_accelerator"`
	ImageType              string            `json:"image_type"`
	Labels                 map[string]string `json:"labels"`
	MachineType            string            `json:"machine_type"`
	Metadata               map[string]string `json:"metadata"`
	OauthScopes            []string          `json:"oauth_scopes"`
	Preemptible            string            `json:"preemptible"`
	ServiceAccount         string            `json:"service_account"`
	Tags                   []string          `json:"tags"`
	Taint                  []Params          `json:"taint"`
	WorkloadMetadataConfig string            `json:"workload_metadata_config"`
}

func (d *NodeConfig) SetNumber(num string) {
	d.Number = num
}
func (d *NodeConfig) GetNumber() string {
	return d.Number
}
func NewNodeConfig(num string) Params {
	var s Params = new(NodeConfig)
	s.SetNumber(num)
	return s
}

func (d *NodeConfig) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "disk_size_gb") {
		d.DiskSizeGb = value
	} else if strings.Contains(key, "disk_type") {
		d.DiskType = value
	} else if strings.Contains(key, "image_type") {
		d.ImageType = value
	} else if strings.Contains(key, "guest_accelerator") {
		//lifecycle_rule.0.conditionまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.GuestAccelerator = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewGuestAccelerator)
	} else if strings.Contains(key, "labels") {
		Number := CreateNumberFromKey(key, 2)
		d.Labels = CreateMap(attributes_map, Number)
	} else if strings.Contains(key, "machine_type") {
		d.MachineType = value
	} else if strings.Contains(key, "metadata") {
		Number := CreateNumberFromKey(key, 2)
		d.Metadata = CreateMap(attributes_map, Number)
	} else if strings.Contains(key, "oauth_scopes") {
		d.OauthScopes = append(d.OauthScopes, value)
	} else if strings.Contains(key, "preemptible") {
		d.Preemptible = value
	} else if strings.Contains(key, "service_account") {
		d.ServiceAccount = value
	} else if strings.Contains(key, "tags") {
		d.Tags = append(d.Tags, value)
	} else if strings.Contains(key, "taint") {
		//lifecycle_rule.0.conditionまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.Taint = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewTaint)
	} else if strings.Contains(key, "workload_metadata_config") {
		d.WorkloadMetadataConfig = value
	}
}

type Taint struct {
	Number string
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

func (d *Taint) SetNumber(num string) {
	d.Number = num
}
func (d *Taint) GetNumber() string {
	return d.Number
}
func NewTaint(num string) Params {
	var s Params = new(Taint)
	s.SetNumber(num)
	return s
}

func (d *Taint) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "key") {
		d.Key = value
	} else if strings.Contains(key, "value") {
		d.Value = value
	} else if strings.Contains(key, "effect") {
		d.Effect = value
	}
}

type NodePool struct {
	Number            string
	Autoscaling       []Params `json:"autoscaling"`
	InitialNodeCount  string   `json:"initial_node_count"`
	InstanceGroupUrls []string `json:"instance_group_urls"`
	Management        []Params `json:"management"`
	MaxPodsPerNode    string   `json:"max_pods_per_node"`
	Name              string   `json:"name"`
	NamePrefix        string   `json:"name_prefix"`
	NodeConfig        []Params `json:"node_config"`
	NodeCount         string   `json:"node_count"`
	Version           string   `json:"version"`
}

func (d *NodePool) SetNumber(num string) {
	d.Number = num
}
func (d *NodePool) GetNumber() string {
	return d.Number
}
func NewNodePool(num string) Params {
	var s Params = new(NodePool)
	s.SetNumber(num)
	return s
}

func (d *NodePool) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "autoscaling") {
		Number := CreateNumberFromKey(key, 2)
		d.Autoscaling = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewAutoscaling)
	} else if strings.Contains(key, "initial_node_count") {
		d.InitialNodeCount = value
	} else if strings.Contains(key, "instance_group_urls") {
		d.InstanceGroupUrls = append(d.InstanceGroupUrls, value)
	} else if strings.Contains(key, "management") {
		Number := CreateNumberFromKey(key, 2)
		d.Management = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewManagement)
	} else if strings.Contains(key, "max_pods_per_node") {
		d.MaxPodsPerNode = value
	} else if strings.Contains(key, "name") {
		d.Name = value
	} else if strings.Contains(key, "name_prefix") {
		d.NamePrefix = value
	} else if strings.Contains(key, "node_config") {
		Number := CreateNumberFromKey(key, 2)
		d.NodeConfig = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewNodeConfig)
	} else if strings.Contains(key, "node_count") {
		d.NodeCount = value
	} else if strings.Contains(key, "version") {
		d.Version = value
	}
}

type Autoscaling struct {
	Number       string
	MaxNodeCount string `json:"max_node_count"`
	MinNodeCount string `json:"min_node_count"`
}

func (d *Autoscaling) SetNumber(num string) {
	d.Number = num
}
func (d *Autoscaling) GetNumber() string {
	return d.Number
}
func NewAutoscaling(num string) Params {
	var s Params = new(Autoscaling)
	s.SetNumber(num)
	return s
}

func (d *Autoscaling) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "max_node_count") {
		d.MaxNodeCount = value
	} else if strings.Contains(key, "min_node_count") {
		d.MinNodeCount = value
	}
}

type Management struct {
	Number      string
	AutoRepair  string `json:"auto_repair"`
	AutoUpgrade string `json:"auto_upgrade"`
}

func (d *Management) SetNumber(num string) {
	d.Number = num
}
func (d *Management) GetNumber() string {
	return d.Number
}
func NewManagement(num string) Params {
	var s Params = new(Management)
	s.SetNumber(num)
	return s
}

func (d *Management) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "auto_repair") {
		d.AutoRepair = value
	} else if strings.Contains(key, "auto_upgrade") {
		d.AutoUpgrade = value
	}
}

type PrivateClusterConfig struct {
	Number                string
	EnablePrivateEndpoint string `json:"enable_private_endpoint"`
	EnablePrivateNodes    string `json:"enable_private_nodes"`
	MasterIpv4CidrBlock   string `json:"master_ipv4_cidr_block"`
	PrivateEndpoint       string `json:"private_endpoint"`
	PublicEndpoint        string `json:"public_endpoint"`
}

func (d *PrivateClusterConfig) SetNumber(num string) {
	d.Number = num
}
func (d *PrivateClusterConfig) GetNumber() string {
	return d.Number
}
func NewPrivateClusterConfig(num string) Params {
	var s Params = new(PrivateClusterConfig)
	s.SetNumber(num)
	return s
}

func (d *PrivateClusterConfig) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "enable_private_endpoint") {
		d.EnablePrivateEndpoint = value
	} else if strings.Contains(key, "enable_private_nodes") {
		d.EnablePrivateNodes = value
	} else if strings.Contains(key, "master_ipv4_cidr_block") {
		d.MasterIpv4CidrBlock = value
	} else if strings.Contains(key, "private_endpoint") {
		d.PrivateEndpoint = value
	} else if strings.Contains(key, "public_endpoint") {
		d.PublicEndpoint = value
	}
}

func Newgoogle_container_cluster(resourceData map[string]interface{}, resourceName string) *google_container_cluster {

	resource := new(google_container_cluster)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.AdditionalZones = CreateStringSlice(attributes_map, attributesKey, "additional_zones")
	resource.Primary.Attributes.InstanceGroupUrls = CreateStringSlice(attributes_map, attributesKey, "instance_group_urls")
	resource.Primary.Attributes.NodeLocations = CreateStringSlice(attributes_map, attributesKey, "node_locations")
	resource.Primary.Attributes.ResourceLabels = CreateMap(attributes_map, "resource_labels")

	resource.Primary.Attributes.AddonsConfig = CreateStructSlice(attributes_map, attributesKey, "addons_config", 1, NewAddonsConfig)
	resource.Primary.Attributes.ClusterAutoscaling = CreateStructSlice(attributes_map, attributesKey, "cluster_autoscaling", 1, NewClusterAutoscaling)
	resource.Primary.Attributes.IPAllocationPolicy = CreateStructSlice(attributes_map, attributesKey, "ip_allocation_policy", 1, NewIPAllocationPolicy)
	resource.Primary.Attributes.MaintenancePolicy = CreateStructSlice(attributes_map, attributesKey, "maintenance_policy", 1, NewMaintenancePolicy)
	resource.Primary.Attributes.MasterAuth = CreateStructSlice(attributes_map, attributesKey, "master_auth", 1, NewMasterAuth)
	resource.Primary.Attributes.MasterAuthorizedNetworksConfig = CreateStructSlice(attributes_map, attributesKey, "master_authorized_networks_config", 1, NewMasterAuthorizedNetworksConfig)
	resource.Primary.Attributes.NetworkPolicy = CreateStructSlice(attributes_map, attributesKey, "network_policy", 1, NewNetworkPolicy)
	resource.Primary.Attributes.NodeConfig = CreateStructSlice(attributes_map, attributesKey, "node_config", 1, NewNodeConfig)
	resource.Primary.Attributes.NodePool = CreateStructSlice(attributes_map, attributesKey, "node_pool", 1, NewNodePool)
	resource.Primary.Attributes.PrivateClusterConfig = CreateStructSlice(attributes_map, attributesKey, "private_cluster_config", 1, NewPrivateClusterConfig)
	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
