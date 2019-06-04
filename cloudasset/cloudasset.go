package cloudasset

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"

	"cloud.google.com/go/storage"
	"github.com/cloud-ace/terraformit-gcp/utils"
	"github.com/google/oauth2l/sgauth"
	"github.com/google/oauth2l/util"
	"golang.org/x/xerrors"
)

//ReadFileFromGCS gets contents of any file in GCS using gsutil cat.
// func ReadFileFromGCS(gsUrl string) ([]byte, error) {
// 	var stdout bytes.Buffer
// 	var stderr bytes.Buffer

// 	cmd := exec.Command("gsutil", "cat", gsUrl)
// 	cmd.Stdout = &stdout
// 	cmd.Stderr = &stderr
// 	err := cmd.Run()
// 	if err != nil {
// 		return nil, xerrors.Errorf("Error: %v \n%s", err, stderr.String())
// 	}
// 	return []byte(strings.TrimRight(stdout.String(), "\n")), nil
// }
func ReadFileFromGCS(URI string) ([]byte, error) {

	//get bucket and objcetPath from URI
	slice := strings.Split(URI, "/")
	bucket := slice[2]
	path := slice[3]
	fmt.Println(bucket, path)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	r, err := client.Bucket(bucket).Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

//Oauth2l
var (
	scopePrefix = "https://www.googleapis.com/auth/"
	cmds        = []string{"fetch", "header", "info", "test"}
)

//Oauth2l
func readJSON(file string) (string, error) {
	if file != "" {
		secretBytes, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		return string(secretBytes), nil
	}
	return "", nil
}

//Oauth2l
func defaultAuthorizeFlowHandler(authorizeUrl string) (string, error) {
	// Print the url on console, let user authorize and paste the token back.
	fmt.Printf("Go to the following link in your browser:\n\n   %s\n\n", authorizeUrl)
	fmt.Println("Enter verification code: ")
	var code string
	fmt.Scanln(&code)
	return code, nil
}

//Oauth2l
func parseScopes(scopes []string) string {
	for i := 0; i < len(scopes); i++ {
		if !strings.Contains(scopes[i], "//") {
			scopes[i] = scopePrefix + scopes[i]
		}
	}
	return strings.Join(scopes, " ")
}

//Oauth2l
func captureStdout(f func(*sgauth.Settings, string), settings *sgauth.Settings, format string) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	f(settings, format)

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

//CreateCloudAsset create CloudAssetMetadata
func CreateCloudAsset(credential, bucket, projectId, cloudAssetPrefix string) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	//OauthToken取得
	jsonFile := credential
	scope := []string{"cloud-platform"}
	format := "json"
	json, err := readJSON(jsonFile)
	if err != nil {
		return "", xerrors.Errorf("Error: %v \n", err)
	}
	settings := &sgauth.Settings{
		CredentialsJSON:  json,
		Scope:            parseScopes(scope),
		OAuthFlowHandler: defaultAuthorizeFlowHandler,
		State:            "state",
	}
	util.Header(settings, format)
	authorization := captureStdout(util.Header, settings, format)

	//create CloudAssetAPI curl command
	//time
	t := time.Now()
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	date := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + "-" + strconv.Itoa(hour) + "-" + strconv.Itoa(minute) + "-" + strconv.Itoa(second)
	authorization = strings.TrimRight(authorization, "\n")
	contentType := "Content-Type: application/json"
	outputURL := "gs://" + bucket + "/" + cloudAssetPrefix + "-" + date
	data := `{"contentType":"RESOURCE","outputConfig":{"gcsDestination":{"uri":"` + outputURL + `"}}}`
	apiURL := "https://cloudasset.googleapis.com/v1/projects/" + projectId + ":exportAssets"
	cmd := exec.Command("curl", "--verbose", "-H", authorization, "-H", contentType, "-d", data, apiURL)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return "", xerrors.Errorf("Error: %v \n%s", err, stderr.String())
	}
	if strings.Contains(stdout.String(), "error") {
		return "", xerrors.Errorf("Error: Failed to create CloudAssetAPI \n%s", stdout.String())
	}

	return cloudAssetPrefix + "-" + date, nil

}

//Metadata is CloudAssetAPIMetadata Struct
type Metadata struct {
	Name      string `json:"name"`
	AssetType string `json:"asset_type"`
	Resource  struct {
		Version              string `json:"version"`
		DiscoveryDocumentURI string `json:"discovery_document_uri"`
		DiscoveryName        string `json:"discovery_name"`
		Parent               string `json:"parent"`
		Data                 struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			ProjectId string `json:"projectid"`
			Region    string `json:"region"`
			Zone      string `json:"zone"`
			Email     string `json:"email"`
			Location  string `json:"location"`
		} `json:"data"`
	} `json:"resource"`
	ResourceMetadata `json:"resource_metadata"`
	DefaultRegion    string `json:"default_region"`
	BackendType      string `json:"backend_type "`
	BackendLocation  string `json:"backend_location"`
	Provider         string `json:"provider"`
}

//ResourceMetadata includes Resource Type
type ResourceMetadata struct {
	Name         string
	ResourceType string
}

//ResouceNameMap CloudAssetAPI AssetType -- Terraform Resource
var ResouceNameMap = map[string]ResourceMetadata{
	//Cloud Key Management Service
	"cloudkms.googleapis.com/KeyRing":   {"google_kms_key_ring", "project"},
	"cloudkms.googleapis.com/CryptoKey": {"google_kms_crypto_key", "project"},
	//Resource Manager
	"cloudresourcemanager.googleapis.com/Project": {"google_project", "project"},
	//Compute Engine
	"compute.googleapis.com/Autoscaler":     {"google_compute_autoscaler", "project/zone"},
	"compute.googleapis.com/BackendService": {"google_compute_backend_service", "project"},
	"compute.googleapis.com/BackendBucket":  {"google_compute_backend_bucket", "project"},
	"compute.googleapis.com/Disk":           {"google_compute_disk", "project/zone"},
	"compute.googleapis.com/Firewall":       {"google_compute_firewall", "project"},
	//Forwarding Rule support only defaultRegion probably because of TerraformBug
	"compute.googleapis.com/ForwardingRule":       {"google_compute_forwarding_rule", "project"},
	"compute.googleapis.com/GlobalForwardingRule": {"google_compute_global_forwarding_rule", "project"},
	"compute.googleapis.com/HealthCheck":          {"google_compute_health_check", "project"},
	"compute.googleapis.com/HttpHealthCheck":      {"google_compute_http_health_check", "project"},
	"compute.googleapis.com/Image":                {"google_compute_image", "project"},
	"compute.googleapis.com/Instance":             {"google_compute_instance", "project/zone"},
	"compute.googleapis.com/InstanceGroup":        {"google_compute_instance_group", "zone"},
	"compute.googleapis.com/InstanceGroupManager": {"google_compute_instance_group_manager", "project/zone"},
	"compute.googleapis.com/InstanceTemplate":     {"google_compute_instance_template", "project"},
	"compute.googleapis.com/Network":              {"google_compute_network", "project"},
	"compute.googleapis.com/Route":                {"google_compute_route", "project"},
	"compute.googleapis.com/Snapshot":             {"google_compute_snapshot", "project"},
	"compute.googleapis.com/SslCertificate":       {"google_compute_ssl_certificate", "project"},
	"compute.googleapis.com/Subnetwork":           {"google_compute_subnetwork", "region"},
	"compute.googleapis.com/TargetHttpProxy":      {"google_compute_target_http_proxy", "project"},
	"compute.googleapis.com/TargetHttpsProxy":     {"google_compute_target_https_proxy", "project"},
	//TargetPool support only defaultRegion probably because of TerraformBug
	"compute.googleapis.com/TargetPool": {"google_compute_target_pool", "project"},
	"compute.googleapis.com/UrlMap":     {"google_compute_url_map", "project"},
	//Google Kubernetes Engine
	"container.googleapis.com/Cluster": {"google_container_cluster", "location"},
	//Cloud Storage
	"storage.googleapis.com/Bucket": {"google_storage_bucket", "project"},
	//Cloud DNS
	"dns.googleapis.com/ManagedZone": {"google_dns_managed_zone", "project"},
	//only beta support
	"dns.googleapis.com/Policy": {"google_dns_policy", "project"},
	//Cloud Identity and Access Management
	"iam.googleapis.com/ServiceAccount": {"google_service_account", "project"},
	//Cloud Pub/Sub
	"pubsub.googleapis.com/Topic":        {"google_pubsub_topic", "project"},
	"pubsub.googleapis.com/Subscription": {"google_pubsub_subscription", "project"},
	//Cloud SQL
	"sqladmin.googleapis.com/Instance": {"google_sql_database_instance", "project"},
}

//ReadMetaFile reads CloudAssetAPI file and return Metadata[]string
func ReadMetaFile(sourcefile string) ([]string, error) {
	f, err := os.Open(sourcefile)
	if err != nil {
		return []string{}, xerrors.Errorf("Error: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var meta_strings []string
	for scanner.Scan() {
		meta_strings = append(meta_strings, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return []string{}, xerrors.Errorf("Error: %w", err)
	}

	return meta_strings, nil
}

//MetaStringToStruct convert String to Struct
func MetaStringToStruct(ss []string) ([]Metadata, error) {
	var metas []Metadata

	for _, v := range ss {
		bytes := []byte(v)
		var meta Metadata
		err := json.Unmarshal(bytes, &meta)
		if err != nil {
			return []Metadata{}, xerrors.Errorf("Error: %w", err)
		}
		metas = append(metas, meta)
	}

	return metas, nil
}

//MetaStringToStruct convert String to Struct
func MetaByteToStruct(bb [][]byte) ([]Metadata, error) {
	var metas []Metadata

	for _, v := range bb {
		bytes := []byte(v)
		var meta Metadata
		err := json.Unmarshal(bytes, &meta)
		if err != nil {
			return []Metadata{}, xerrors.Errorf("MetaByteToStruct Error: %w", err)
		}
		metas = append(metas, meta)
	}

	return metas, nil
}

//InsertIntoBuffer meta_struct -> buffer
func InsertIntoBuffer(meta_struct *Metadata, buf *bytes.Buffer, resourceNameMap map[string]ResourceMetadata, funcMap template.FuncMap, subCommand string) error {

	//skip resource managerのjsonはskipする
	var templateName string

	switch subCommand {
	case "resource":
		templateName = utils.GetGoPath() + "/src/" + utils.RepositoryDirectory + "/cloudasset/templates/resource.tf.tmpl"
	case "import":
		templateName = utils.GetGoPath() + "/src/" + utils.RepositoryDirectory + "/cloudasset/templates/import.sh.tmpl"
	case "provider":
		templateName = utils.GetGoPath() + "/src/" + utils.RepositoryDirectory + "/cloudasset/templates/provider.tf.tmpl"
	case "backend":
		templateName = utils.GetGoPath() + "/src/" + utils.RepositoryDirectory + "/cloudasset/templates/backend.tf.tmpl"
	}

	meta_struct.ResourceMetadata = resourceNameMap[meta_struct.AssetType]
	// skip resource which is not supported.
	if meta_struct.ResourceMetadata.Name == "" {
		if subCommand == "resource" {
			fmt.Println("Skip not supported resource:", meta_struct.AssetType)
			return nil
		} else if subCommand == "import" {
			return nil
		}
	}

	name := path.Base(templateName)
	t, err := template.New(name).Funcs(funcMap).ParseFiles(templateName)
	if err != nil {
		return xerrors.Errorf("Error ParseFiles: %w", err)
	}

	if err := t.Execute(buf, meta_struct); err != nil {
		return xerrors.Errorf("Error execution: %w", err)
	}

	return nil
}

func TrimMetaStruct(meta_structs []Metadata) []Metadata {

	var noDuplicateMetaStructs []Metadata

	for _, v := range meta_structs {

		if v.AssetType == "compute.googleapis.com/InstanceTemplate" && !strings.Contains(v.Name, "global") {
			//nothing else
		} else {
			noDuplicateMetaStructs = append(noDuplicateMetaStructs, v)
		}

	}
	return noDuplicateMetaStructs
}

//RemoveDefaultResources default resources.
func RemoveDefaultResources(meta_structs []Metadata, defaultNetwork, defaultSubnetwork, defaultRoute, defaultFirewall string) []Metadata {
	var removeDefaultStructs []Metadata

	for _, v := range meta_structs {

		//skip default service account.
		if v.AssetType == "iam.googleapis.com/ServiceAccount" {
			if IsServiceAccountDefault(v.Resource.Data.Email) {
				fmt.Println("Skip default service account:", v.Name)
			} else {
				removeDefaultStructs = append(removeDefaultStructs, v)
			}
		} else if defaultNetwork == "false" && IsDefaultNetwork(v) {
			fmt.Println("Skip default resource:", v.Name)
		} else if defaultSubnetwork == "false" && IsDefaultSubnetwork(v) {
			fmt.Println("Skip default resource:", v.Name)
		} else if defaultRoute == "false" && IsDefaultRoute(v) {
			fmt.Println("Skip default resource:", v.Name)
		} else if defaultFirewall == "false" && IsDefaultFirewall(v) {
			fmt.Println("Skip default resource:", v.Name)
		} else if strings.Contains(v.Name, "aef-default") {
			fmt.Println("Skip app-engine default resource:", v.Name)
		} else {
			removeDefaultStructs = append(removeDefaultStructs, v)
		}

	}
	return removeDefaultStructs
}

func IsDefaultNetwork(meta_struct Metadata) bool {
	if meta_struct.Resource.Data.Name == "default" {
		return true
	}
	return false
}
func IsDefaultSubnetwork(meta_struct Metadata) bool {
	if meta_struct.Resource.Data.Name == "default" {
		return true
	}
	return false
}
func IsDefaultRoute(meta_struct Metadata) bool {
	if strings.Contains(meta_struct.Resource.Data.Name, "default-route-") {
		return true
	}
	return false
}

func IsDefaultFirewall(meta_struct Metadata) bool {
	if strings.Contains(meta_struct.Resource.Data.Name, "default-allow-") || strings.Contains(meta_struct.Resource.Data.Name, "default-deny-") {
		return true
	}
	return false
}

//CreateFileForImport creates files for terraform import command
func CreateFileForImport(meta_structs []Metadata, subCommand, defaultRegion, backendType, backendLocation, provider string) error {

	//重複を削除
	noDuplicateMetaStructs := TrimMetaStruct(meta_structs)

	var buf bytes.Buffer
	if subCommand == "provider" {
		//get project name from first resource from resource except for cloudresourcemanager.googleapis.com
		for _, v := range meta_structs {
			if strings.Contains(v.AssetType, "cloudresourcemanager.googleapis.com") {
			} else {
				v.DefaultRegion = defaultRegion
				noDuplicateMetaStructs = []Metadata{v}
				break
			}
		}

	} else if subCommand == "backend" {
		//structsとしてbucket名を渡す
		noDuplicateMetaStructs = []Metadata{Metadata{BackendType: backendType, BackendLocation: backendLocation}}
	}

	for _, v := range noDuplicateMetaStructs {
		v.Provider = provider
		if err := InsertIntoBuffer(&v, &buf, ResouceNameMap, utils.FuncMap, subCommand); err != nil {
			return xerrors.Errorf("Error: %w", err)
		}
	}

	switch subCommand {
	case "resource":
		if err := utils.CreateFile(utils.ResourceTfName, buf); err != nil {
			return xerrors.Errorf("Error: %w", err)
		}
	case "import":
		if err := utils.CreateFile(utils.ImportShName, buf); err != nil {
			return xerrors.Errorf("Error: %w", err)
		}
	case "provider":
		if err := utils.CreateFile(utils.ProviderTfName, buf); err != nil {
			return xerrors.Errorf("Error: %w", err)
		}
	case "backend":
		if err := utils.CreateFile(utils.BackendTfName, buf); err != nil {
			return xerrors.Errorf("Error: %w", err)
		}
	}
	return nil
}

func IsServiceAccountDefault(email string) bool {
	if _, err := strconv.Atoi(email[0:1]); err == nil {
		return true
	}
	return false
}
