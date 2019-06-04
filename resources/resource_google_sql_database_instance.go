package resources

import (
	"strings"
)

type google_sql_database_instance struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			ConnectionName             string   `json:"connection_name"`
			DatabaseVersion            string   `json:"database_version"`
			FirstIPAddress             string   `json:"first_ip_address"`
			ID                         string   `json:"id"`
			IPAddress                  []Params `json:"ip_address"`
			MasterInstanceName         string   `json:"master_instance_name"`
			Name                       string   `json:"name"`
			PrivateIPAddress           string   `json:"private_ip_address"`
			Project                    string   `json:"project"`
			PublicIPAddress            string   `json:"public_ip_address"`
			Region                     string   `json:"region"`
			ReplicaConfiguration       []Params `json:"replica_configuration"`
			SelfLink                   string   `json:"self_link"`
			ServerCaCert               []Params `json:"server_ca_cert"`
			ServiceAccountEmailAddress string   `json:"service_account_email_address"`
			Settings                   []Params `json:"settings"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}
type IPAddress struct {
	Number       string
	IPAddress    string `json:"ip_address"`
	TimeToRetire string `json:"time_to_retire"`
	Type         string `json:"type"`
}

func (d *IPAddress) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "type") {
		d.Type = value
	} else if strings.Contains(key, "time_to_retire") {
		d.TimeToRetire = value
	} else if strings.Contains(key, "ip_address") {
		d.IPAddress = value
	}
}

func (d *IPAddress) SetNumber(num string) {
	d.Number = num
}

func (d *IPAddress) GetNumber() string {
	return d.Number
}

func NewIPAddress(num string) Params {
	var s Params = new(IPAddress)
	s.SetNumber(num)
	return s
}

type ReplicaConfiguration struct {
	Number                  string
	CaCertificate           string `json:"ca_certificate"`
	ClientCertificate       string `json:"client_certificate"`
	ClientKey               string `json:"client_key"`
	ConnectRetryInterval    string `json:"connect_retry_interval"`
	DumpFilePath            string `json:"dump_file_path"`
	FailoverTarget          string `json:"failover_target"`
	MasterHeartbeatPeriod   string `json:"master_heartbeat_period"`
	Password                string `json:"password"`
	SslCipher               string `json:"ssl_cipher"`
	Username                string `json:"username"`
	VerifyServerCertificate string `json:"verify_server_certificate"`
}

func (d *ReplicaConfiguration) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "ca_certificate") {
		d.CaCertificate = value
	} else if strings.Contains(key, "client_certificate") {
		d.ClientCertificate = value
	} else if strings.Contains(key, "client_key") {
		d.ClientKey = value
	} else if strings.Contains(key, "connect_retry_interval") {
		d.ConnectRetryInterval = value
	} else if strings.Contains(key, "dump_file_path") {
		d.DumpFilePath = value
	} else if strings.Contains(key, "failover_target") {
		d.FailoverTarget = value
	} else if strings.Contains(key, "master_heartbeat_period") {
		d.MasterHeartbeatPeriod = value
	} else if strings.Contains(key, "password") {
		d.Password = value
	} else if strings.Contains(key, "ssl_cipher") {
		d.SslCipher = value
	} else if strings.Contains(key, "username") {
		d.Username = value
	} else if strings.Contains(key, "verify_server_certificate") {
		d.VerifyServerCertificate = value
	}
}

func (d *ReplicaConfiguration) SetNumber(num string) {
	d.Number = num
}

func (d *ReplicaConfiguration) GetNumber() string {
	return d.Number
}

func NewReplicaConfiguration(num string) Params {
	var s Params = new(ReplicaConfiguration)
	s.SetNumber(num)
	return s
}

type ServerCaCert struct {
	Number          string
	Cert            string `json:"cert"`
	CommonName      string `json:"common_name"`
	CreateTime      string `json:"create_time"`
	ExpirationTime  string `json:"expiration_time"`
	Sha1Fingerprint string `json:"sha1_fingerprint"`
}

func (d *ServerCaCert) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "common_name") {
		d.CommonName = value
	} else if strings.Contains(key, "create_time") {
		d.CreateTime = value
	} else if strings.Contains(key, "expiration_time") {
		d.ExpirationTime = value
	} else if strings.Contains(key, "sha1_fingerprint") {
		d.Sha1Fingerprint = value
	} else if strings.Contains(key, "cert") {
		d.Cert = value
	}
}

func (d *ServerCaCert) SetNumber(num string) {
	d.Number = num
}

func (d *ServerCaCert) GetNumber() string {
	return d.Number
}

func NewServerCaCert(num string) Params {
	var s Params = new(ServerCaCert)
	s.SetNumber(num)
	return s
}

type Settings struct {
	Number                    string
	ActivationPolicy          string            `json:"activation_policy"`
	AuthorizedGaeApplications []string          `json:"authorized_gae_applications"`
	AvailabilityType          string            `json:"availability_type"`
	BackupConfiguration       []Params          `json:"backup_configuration"`
	CrashSafeReplication      string            `json:"crash_safe_replication"`
	DatabaseFlags             []Params          `json:"database_flags"`
	DiskAutoresize            string            `json:"disk_autoresize"`
	DiskSize                  string            `json:"disk_size"`
	DiskType                  string            `json:"disk_type"`
	IPConfiguration           []Params          `json:"ip_configuration"`
	LocationPreference        []Params          `json:"location_preference"`
	MaintenanceWindow         []Params          `json:"maintenance_window"`
	PricingPlan               string            `json:"pricing_plan"`
	ReplicationType           string            `json:"replication_type"`
	Tier                      string            `json:"tier"`
	UserLabels                map[string]string `json:"user_labels"`
	Version                   string            `json:"version"`
}

func (d *Settings) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "activation_policy") {
		d.ActivationPolicy = value
	} else if strings.Contains(key, "availability_type") {
		d.AvailabilityType = value
	} else if strings.Contains(key, "crash_safe_replication") {
		d.CrashSafeReplication = value
	} else if strings.Contains(key, "disk_autoresize") {
		d.DiskAutoresize = value
	} else if strings.Contains(key, "disk_size") {
		d.DiskSize = value
	} else if strings.Contains(key, "disk_type") {
		d.DiskType = value
	} else if strings.Contains(key, "pricing_plan") {
		d.PricingPlan = value
	} else if strings.Contains(key, "replication_type") {
		d.ReplicationType = value
	} else if strings.Contains(key, "tier") {
		d.Tier = value
	} else if strings.Contains(key, "version") {
		d.Version = value
	} else if strings.Contains(key, "authorized_gae_applications") {
		d.AuthorizedGaeApplications = append(d.AuthorizedGaeApplications, value)
	} else if strings.Contains(key, "backup_configuration") {
		Number := CreateNumberFromKey(key, 2)
		d.BackupConfiguration = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewBackupConfiguration)
	} else if strings.Contains(key, "database_flags") {
		Number := CreateNumberFromKey(key, 2)
		d.DatabaseFlags = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewDatabaseFlags)
	} else if strings.Contains(key, "ip_configuration") {
		Number := CreateNumberFromKey(key, 2)
		d.IPConfiguration = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewIPConfiguration)
	} else if strings.Contains(key, "location_preference") {
		Number := CreateNumberFromKey(key, 2)
		d.LocationPreference = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewLocationPreference)
	} else if strings.Contains(key, "maintenance_window") {
		Number := CreateNumberFromKey(key, 2)
		d.MaintenanceWindow = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewMaintenanceWindow)
	} else if strings.Contains(key, "user_labels") {
		Number := CreateNumberFromKey(key, 2)
		d.UserLabels = CreateMap(attributes_map, Number)
	}

}

func (d *Settings) SetNumber(num string) {
	d.Number = num
}

func (d *Settings) GetNumber() string {
	return d.Number
}

func NewSettings(num string) Params {
	var s Params = new(Settings)
	s.SetNumber(num)
	return s
}

type BackupConfiguration struct {
	Number           string
	BinaryLogEnabled string `json:"binary_log_enabled"`
	Enabled          string `json:"enabled"`
	StartTime        string `json:"start_time"`
}

func (d *BackupConfiguration) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "binary_log_enabled") {
		d.BinaryLogEnabled = value
	} else if strings.Contains(key, "enabled") {
		d.Enabled = value
	} else if strings.Contains(key, "start_time") {
		d.StartTime = value
	}
}

func (d *BackupConfiguration) SetNumber(num string) {
	d.Number = num
}

func (d *BackupConfiguration) GetNumber() string {
	return d.Number
}

func NewBackupConfiguration(num string) Params {
	var s Params = new(BackupConfiguration)
	s.SetNumber(num)
	return s
}

type DatabaseFlags struct {
	Number string
	Name   string `json:"name"`
	Value  string `json:"value"`
}

func (d *DatabaseFlags) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "name") {
		d.Name = value
	} else if strings.Contains(key, "value") {
		d.Value = value
	}
}

func (d *DatabaseFlags) SetNumber(num string) {
	d.Number = num
}

func (d *DatabaseFlags) GetNumber() string {
	return d.Number
}

func NewDatabaseFlags(num string) Params {
	var s Params = new(DatabaseFlags)
	s.SetNumber(num)
	return s
}

type IPConfiguration struct {
	Number             string
	AuthorizedNetworks []Params `json:"authorized_networks"`
	Ipv4Enabled        string   `json:"ipv4_enabled"`
	PrivateNetwork     string   `json:"private_network"`
	RequireSsl         string   `json:"require_ssl"`
}

func (d *IPConfiguration) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "authorized_networks") {
		Number := CreateNumberFromKey(key, 4)
		d.AuthorizedNetworks = CreateStructSlice(attributes_map, attributesKey, Number, 5, NewAuthorizedNetworks)
	} else if strings.Contains(key, "ipv4_enabled") {
		d.Ipv4Enabled = value
	} else if strings.Contains(key, "private_network") {
		d.PrivateNetwork = value
	} else if strings.Contains(key, "require_ssl") {
		d.RequireSsl = value
	}
}

func (d *IPConfiguration) SetNumber(num string) {
	d.Number = num
}

func (d *IPConfiguration) GetNumber() string {
	return d.Number
}

func NewIPConfiguration(num string) Params {
	var s Params = new(IPConfiguration)
	s.SetNumber(num)
	return s
}

type AuthorizedNetworks struct {
	Number         string
	ExpirationTime string `json:"expiration_time"`
	Name           string `json:"name"`
	Value          string `json:"value"`
}

func (d *AuthorizedNetworks) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "name") {
		d.Name = value
	} else if strings.Contains(key, "value") {
		d.Value = value
	} else if strings.Contains(key, "expiration_time") {
		d.ExpirationTime = value
	}
}

func (d *AuthorizedNetworks) SetNumber(num string) {
	d.Number = num
}

func (d *AuthorizedNetworks) GetNumber() string {
	return d.Number
}

func NewAuthorizedNetworks(num string) Params {
	var s Params = new(AuthorizedNetworks)
	s.SetNumber(num)
	return s
}

type LocationPreference struct {
	Number               string
	FollowGaeApplication string `json:"follow_gae_application"`
	Zone                 string `json:"zone"`
}

func (d *LocationPreference) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "follow_gae_application") {
		d.FollowGaeApplication = value
	} else if strings.Contains(key, "zone") {
		d.Zone = value
	}
}

func (d *LocationPreference) SetNumber(num string) {
	d.Number = num
}

func (d *LocationPreference) GetNumber() string {
	return d.Number
}

func NewLocationPreference(num string) Params {
	var s Params = new(LocationPreference)
	s.SetNumber(num)
	return s
}

type MaintenanceWindow struct {
	Number      string
	Day         string `json:"day"`
	Hour        string `json:"hour"`
	UpdateTrack string `json:"update_track"`
}

func (d *MaintenanceWindow) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "day") {
		d.Day = value
	} else if strings.Contains(key, "hour") {
		d.Hour = value
	} else if strings.Contains(key, "update_track") {
		d.UpdateTrack = value
	}
}

func (d *MaintenanceWindow) SetNumber(num string) {
	d.Number = num
}

func (d *MaintenanceWindow) GetNumber() string {
	return d.Number
}

func NewMaintenanceWindow(num string) Params {
	var s Params = new(MaintenanceWindow)
	s.SetNumber(num)
	return s
}

func Newgoogle_sql_database_instance(resourceData map[string]interface{}, resourceName string) *google_sql_database_instance {

	resource := new(google_sql_database_instance)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.IPAddress = CreateStructSlice(attributes_map, attributesKey, "ip_address", 1, NewIPAddress)
	resource.Primary.Attributes.ReplicaConfiguration = CreateStructSlice(attributes_map, attributesKey, "replica_configuration", 1, NewReplicaConfiguration)
	resource.Primary.Attributes.ServerCaCert = CreateStructSlice(attributes_map, attributesKey, "server_ca_cert", 1, NewServerCaCert)
	resource.Primary.Attributes.Settings = CreateStructSlice(attributes_map, attributesKey, "settings", 1, NewSettings)

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
