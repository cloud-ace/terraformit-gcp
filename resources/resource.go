package resources

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/cloud-ace/terraformit-gcp/utils"
	"golang.org/x/xerrors"
)

//Module Struct
type Module struct {
	Path    []string `json:"path"`
	Outputs struct {
	} `json:"outputs"`
	Resources interface{} `json:"resources"`
	DependsOn []string    `json:"depends_on"`
}

//Tfstate Struct
type Tfstate struct {
	Version          int      `json:"version"`
	TerraformVersion string   `json:"terraform_version"`
	Serial           int      `json:"serial"`
	Lineage          string   `json:"lineage"`
	Modules          []Module `json:"modules"`
}

//Params interface for Parameters
type Params interface {
	SetParameter(string, string, map[string]interface{}, []string)
	SetNumber(string)
	GetNumber() string
}

//CreateAttributesKey creates attributeskey_slice from attributes_map
func CreateAttributesKey(attributes_map map[string]interface{}) []string {
	var attributesKey []string
	for k, _ := range attributes_map {
		attributesKey = append(attributesKey, k)
	}
	sort.Strings(attributesKey)
	return attributesKey
}

func CreateNumberFromKey(key string, index int) string {
	slice := strings.Split(key, ".")
	index = index + 1
	return strings.Join(slice[:index], ".")
}

//CreateStructSlice creates Parameter struct slice ex) Allow
func CreateStructSlice(attributes_map map[string]interface{}, attributesKey []string, name string, slicePoint int, fn func(string) Params) []Params {
	var resource_slice []Params
	var resource_numbers []string
	for k, v := range attributesKey {

		//初めの文字が与えられた物と一緒だったら
		keySlice := strings.Split(attributesKey[k], ".")
		nameSlice := strings.Split(name, ".")

		if keySlice[0] == nameSlice[0] && strings.Contains(attributesKey[k], name+".") && !strings.Contains(attributesKey[k], name+".#") {
			Number := CreateNumberFromKey(attributesKey[k], slicePoint)
			if strings.Contains(attributesKey[k], Number) {
				//同じナンバーがあったらappendしない
				if !arrayContains(resource_numbers, Number) {
					resource_numbers = append(resource_numbers, Number)
					resource_slice = append(resource_slice, fn(Number))
				}
			}
			for _, vv := range resource_slice {
				number := vv.GetNumber()
				if number == Number && !strings.Contains(attributesKey[k], "#") {
					vv.SetParameter(attributesKey[k], attributes_map[v].(string), attributes_map, attributesKey)
				}
			}
		}
	}
	return resource_slice
}

func arrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

//InterfaceToStrings converts Interface{} to []string
func InterfaceToStrings(i interface{}) []string {
	length := len(i.([]interface{}))
	slice := make([]string, length, length)
	for i, v := range i.([]interface{}) {
		slice[i] = v.(string)
	}
	return slice
}

//CreateMap creates Parameter map ex)Label,Metadata
func CreateMap(attributes_map map[string]interface{}, name string) map[string]string {
	tmp_map := map[string]string{}
	for k, v := range attributes_map {

		//初めの文字が与えられた物と一緒だったら
		keySlice := strings.Split(k, ".")
		nameSlice := strings.Split(name, ".")

		if keySlice[0] == nameSlice[0] && strings.Contains(k, name) {
			if strings.Contains(k, "%") {
				//nothing
			} else {
				key := strings.Split(k, ".")
				tmp_map[key[len(key)-1]] = v.(string)
			}
		}
	}
	return tmp_map
}

//CreateStringSlice creates slice ex)Users,Ports,SourceRanges
func CreateStringSlice(attributes_map map[string]interface{}, attributesKey []string, name string) []string {
	var slice []string
	for k, v := range attributesKey {
		if strings.Contains(attributesKey[k], name+".") && !strings.Contains(attributesKey[k], "#") {
			slice = append(slice, attributes_map[v].(string))
		}
	}
	return slice
}

//InsertTfstateIntoBuffer templateのデータをバッファに書き込み
func InsertTfstateIntoBuffer(s string, resourceData map[string]interface{}, buf *bytes.Buffer, funcMap template.FuncMap) error {

	//create resource instance
	var resource interface{}

	//get only resource id(ex. google_compute_project)
	slice := strings.Split(s, ".")
	resourceId := slice[0]

	//get resource id(ex. instance-1)
	resourceName := slice[1]

	//create Template Object
	templateName := utils.GetGoPath() + "/src/" + utils.RepositoryDirectory + "/resources/templates/" + resourceId + ".tf.tmpl"
	name := path.Base(templateName)
	t, err := template.New(name).Funcs(funcMap).ParseFiles(templateName)
	if err != nil {
		return xerrors.Errorf("Error: %w", err)
	}

	switch resourceId {
	//Cloud Key Management Service
	case "google_kms_key_ring":
		resource = Newgoogle_kms_key_ring(resourceData, resourceName)
	case "google_kms_crypto_key":
		resource = Newgoogle_kms_crypto_key(resourceData, resourceName)
	//resource manager
	case "google_project":
		resource = Newgoogle_project(resourceData, resourceName)
	//compute
	case "google_compute_autoscaler":
		resource = Newgoogle_compute_autoscaler(resourceData, resourceName)
	case "google_compute_backend_service":
		resource = Newgoogle_compute_backend_service(resourceData, resourceName)
	case "google_compute_backend_bucket":
		resource = Newgoogle_compute_backend_bucket(resourceData, resourceName)
	case "google_compute_disk":
		resource = Newgoogle_compute_disk(resourceData, resourceName)
	case "google_compute_firewall":
		resource = Newgoogle_compute_firewall(resourceData, resourceName)
	case "google_compute_forwarding_rule":
		resource = Newgoogle_compute_forwarding_rule(resourceData, resourceName)
	case "google_compute_global_forwarding_rule":
		resource = Newgoogle_compute_global_forwarding_rule(resourceData, resourceName)
	case "google_compute_health_check":
		resource = Newgoogle_compute_health_check(resourceData, resourceName)
	case "google_compute_http_health_check":
		resource = Newgoogle_compute_http_health_check(resourceData, resourceName)
	case "google_compute_image":
		resource = Newgoogle_compute_image(resourceData, resourceName)
	case "google_compute_instance":
		resource = Newgoogle_compute_instance(resourceData, resourceName)
	case "google_compute_instance_group":
		resource = Newgoogle_compute_instance_group(resourceData, resourceName)
	case "google_compute_instance_group_manager":
		resource = Newgoogle_compute_instance_group_manager(resourceData, resourceName)
	case "google_compute_instance_template":
		resource = Newgoogle_compute_instance_template(resourceData, resourceName)
	case "google_compute_network":
		resource = Newgoogle_compute_network(resourceData, resourceName)
	case "google_compute_route":
		resource = Newgoogle_compute_route(resourceData, resourceName)
	case "google_compute_snapshot":
		resource = Newgoogle_compute_snapshot(resourceData, resourceName)
	case "google_compute_ssl_certificate":
		resource = Newgoogle_compute_ssl_certificate(resourceData, resourceName)
	case "google_compute_subnetwork":
		resource = Newgoogle_compute_subnetwork(resourceData, resourceName)
	case "google_compute_target_http_proxy":
		resource = Newgoogle_compute_target_http_proxy(resourceData, resourceName)
	case "google_compute_target_https_proxy":
		resource = Newgoogle_compute_target_https_proxy(resourceData, resourceName)
	case "google_compute_target_pool":
		resource = Newgoogle_compute_target_pool(resourceData, resourceName)
	case "google_compute_url_map":
		resource = Newgoogle_compute_url_map(resourceData, resourceName)
	//Google Kubernetes Engine
	case "google_container_cluster":
		resource = Newgoogle_container_cluster(resourceData, resourceName)
	//Cloud Storage
	case "google_storage_bucket":
		resource = Newgoogle_storage_bucket(resourceData, resourceName)
	//Cloud DNS
	case "google_dns_managed_zone":
		resource = Newgoogle_dns_managed_zone(resourceData, resourceName)
	case "google_dns_policy":
		resource = Newgoogle_dns_policy(resourceData, resourceName)
	//Cloud Identity and Access Management
	case "google_service_account":
		resource = Newgoogle_service_account(resourceData, resourceName)
	//Cloud Pub/Sub
	case "google_pubsub_subscription":
		resource = Newgoogle_pubsub_subscription(resourceData, resourceName)
	case "google_pubsub_topic":
		resource = Newgoogle_pubsub_topic(resourceData, resourceName)
	//Cloud SQL
	case "google_sql_database_instance":
		resource = Newgoogle_sql_database_instance(resourceData, resourceName)
	default:
		return xerrors.Errorf("Error: %v is not supported", resourceId)
	}

	if err := t.Execute(buf, resource); err != nil {
		return xerrors.Errorf("Error: %w", err)
	}
	return nil
}

//ReadTfstateFile reads tfstatefile and return []byte
func ReadTfstateFile(sourcefile string) ([]byte, error) {

	f, err := os.Open(sourcefile)
	if err != nil {
		return []byte{}, xerrors.Errorf("Error: %w", err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return []byte{}, xerrors.Errorf("Error: %w", err)
	}
	return bs, nil
}

//TfstateByteToStruct converts byte to struct
func TfstateByteToStruct(bytes []byte) (Tfstate, error) {
	var ts Tfstate
	err := json.Unmarshal(bytes, &ts)
	if err != nil {
		return Tfstate{}, xerrors.Errorf("Error: %w", err)
	}
	return ts, nil
}

//GetResourceName gets all resourceNames from Tfstate
func GetResourceName(tfstate_struct Tfstate) []string {
	Resources := tfstate_struct.Modules[0].Resources
	var resourceNames []string
	for k, _ := range Resources.(map[string]interface{}) {
		resourceNames = append(resourceNames, k)
	}
	return resourceNames
}

//CreateTf creates tf file
func CreateTf(tfstate_struct Tfstate) error {

	resourceNames := GetResourceName(tfstate_struct)
	sort.Strings(resourceNames)
	var buf bytes.Buffer

	//insert resourceData into buffer
	for _, v := range resourceNames {
		//interface->map
		resourceData, _ := tfstate_struct.Modules[0].Resources.(map[string]interface{})[v].(map[string]interface{})
		if err := InsertTfstateIntoBuffer(v, resourceData, &buf, utils.FuncMap); err != nil {
			return xerrors.Errorf("Error: %w", err)
		}
	}
	utils.CreateFile(utils.TfName, buf)

	return nil
}

//MapToStruct map -> struct
func MapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
