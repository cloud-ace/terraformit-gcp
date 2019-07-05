# terraformit-gcp
terraformit-gcp is an open source command line tool for generating tf files and tfstate from existing GCP resources.  
Relieve the pain of coding tf of manually created GCP resources.

terraformit-gcp steps are as below.
1. terraformit-gcp get json data of existing GCP resources using [Cloud Asset API](https://cloud.google.com/resource-manager/docs/cloud-asset-inventory/reference/rest/) exportAssets method.

2. terraformit-gcp generates files for creating a tfstate(="terraform import") from the json data.

3. terraformit-gcp generates tf files from the tfstate. 

4. terraformit-gcp executes "terraform plan" command to check tf files are generated successfully.

## Version Table
terraformit-gcp does not support terraform 0.12.0 now.

| terraformit-gcp | go | terraform | google provider | google provider (beta) | 
|:-------:|:-------:|:-------:|:-------:|:-------:|
| v0.9.0 | v1.12 | v0.11.13 and v0.11.14 | v2.5.1 | v2.5.1 |

## To start using terraformit-gcp
Please follow these steps.

### Install commands
Install [terraform](https://learn.hashicorp.com/terraform/) or [tfenv](https://github.com/tfutils/tfenv)(Terraform version manager).  

Install [gcloud](https://cloud.google.com/sdk/gcloud/?hl=en) to create a credential.

Install [Go tools](https://golang.org/dl/) to use go command.

### Set gcloud authentication
Generate ~/.config/gcloud/application_default_credentials.json credential.  
Terraform command and google storage library use this credential.
```
gcloud init
```
or
```
gcloud auth login
```

### Install terraformit-gcp
Install terraformit-gcp.
git clone terraformit-gcp to your GOPATH.
```
export GO111MODULE=on
git clone https://github.com/cloud-ace/terraformit-gcp.git -b v0.9.1 ~/go/src/github.com/cloud-ace/terraformit-gcp
cd ~/go/src/github.com/cloud-ace/terraformit-gcp
go install
```

### Set path
Add GOPATH to PATH, if you need.  
(mac)
```
echo 'export GOPATH=$HOME/go' >> ~/.bash_profile
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bash_profile
source ~/.bash_profile
```

### Enable CloudAssetAPI
Enable [CloudAssetAPI](https://console.cloud.google.com/apis/library/cloudasset.googleapis.com).    

### Create bucket for storing CloudAssetAPI json data
Create bucket for storing CloudAssetAPI outputs.

### Generate and download credential for CloudAssetAPI
Genereate Oauth Client ID and download a credentials.  
Cloud Asset API only supports Oauth Client ID now.  
https://cloud.google.com/resource-manager/docs/cloud-asset-inventory/calling-api-with-local-machine-howto?hl=en#downloading_the_credential_file

### cd your pj directory
Change your terraform project directory.
```
cd "your terraform project directory"
```

### Create .terraformit-gcp.yaml in your project directory
Create .terraformit-gcp.yaml in your project directory and set your configuration.  
Please refer to the sample file(sample.terraformit-gcp.yaml) in this repository.
```
CloudAsset:
  # GCP project number
  project-number: "xxxxxxxx" 

  # bucket name. CloudAssetAPI MetadataFile is exported to this bucket.
  bucket: "xxxxxxxxx"

  # Oauth Client ID credential location
  credential: "/Users/xxxxx/Downloads/xxxxxx.json"

Terraform:
  # provider. "google" or "google-beta" should be set.
  provider: "google"

  # your workspace 
  workspace: "default"

  # buckend type "local" or "gcs" is supported now.
  # https://www.terraform.io/docs/backends/types/gcs.html
  backend-type: "local"

  # if you set "local" to backend-type, set "" to backend-location.
  # backend-location: ""
  # if you set "gcs" to backend-type ,set your bucket name to backend-location
  # backend-location: "bucketname"  
  backend-location: ""

  # Default Region
  gcp-provider-default-region: "asia-northeast1"

  # whether add Default resources("true") or remove("false").
  # set true or false. If you set "false", skip default resource.
  # Default service accounts are removed automatically because their name start with number("12233445@....") which cause an error. 
  resource-default-network: false
  resource-default-subnetwork: false
  resource-default-route: false
  resource-default-firewall: false
```

## terraformit-gcp Command
### terraformit-gcp plan
Following steps below are executed.  
1. create CloudAssetMetadata calling CloudAssetAPI
2. get CloudAssetMetadata from GCS
3. create ImportFiles
4. "terraform init"
5. "terraform workspace new"
6. "terraform import"(create tfstate)
7. create tffile
8. "terraform plan"

### terraformit-gcp create cloudasset
Following steps below are executed.  
1. create CloudAssetMetadata calling CloudAssetAPI

### terraformit-gcp create importfiles (-f ./xxx/xxxxx or gs://xxxxxx/xxxx)
Following steps below are executed.  
1. get CloudAssetMetadata from GCS or local(-f option)
2. create ImportFiles

### terraformit-gcp create tfstate
Following steps below are executed.  
1. "terraform init"
2. "terraform workspace new"
3. "terraform import" using importfiles

### terraformit-gcp create tffile (-f tfstatefile)
Following steps below are executed.  
1. create tffile

## Support Table
This command supports GCP resources which is supported by [Cloud Asset API](https://cloud.google.com/resource-manager/docs/cloud-asset-inventory/overview?hl=en).

✅:support  
🔧:will support
/ :no support

| CloudAssetAPI Name | CloudAssetAPI Support | terrafromResource name | terrafromResource Support | 
|:-------------:|:-------:|:-------------:|:--------:|
**Cloud Key Management Service**
| cloudkms.googleapis.com/KeyRing |  ✅ | google_kms_key_ring |  ✅ |
| cloudkms.googleapis.com/CryptoKey |  ✅ | google_kms_crypto_key |  ✅ |
| cloudkms.googleapis.com/CryptoKeyVersion | / | / | / |
**Resource Manager**
| cloudresourcemanager.googleapis.com/Organization | / | / | / |
| cloudresourcemanager.googleapis.com/Folder | / | google_folder | / |
| cloudresourcemanager.googleapis.com/Project |  ✅ | google_project |  ✅ |
**Compute Engine**
| compute.googleapis.com/Autoscaler |  ✅ | google_compute_autoscaler | ✅ |
| compute.googleapis.com/BackendBucket | ✅ | google_compute_backend_bucket | ✅ | 
| compute.googleapis.com/BackendService |  ✅ | google_compute_backend_service |  ✅ | 
| compute.googleapis.com/Disk |  ✅ | google_compute_disk |  ✅ | 
| compute.googleapis.com/Firewall |  ✅ | google_compute_firewall |  ✅ |
| compute.googleapis.com/ForwardingRule |  ✅(only support in default Region) | google_compute_forwarding_rule |  ✅ |
| compute.googleapis.com/GlobalForwardingRule |  ✅ | google_compute_global_forwarding_rule |  ✅ | 
| compute.googleapis.com/HealthCheck |  ✅ | google_compute_health_check |  ✅ | 
| compute.googleapis.com/HttpHealthCheck |  ✅ | google_compute_http_health_check |  ✅ | 
| compute.googleapis.com/HttpsHealthCheck | 🔧 | google_compute_https_health_check | 🔧 |
| compute.googleapis.com/Image |  ✅ | google_compute_image |  ✅ | 
| compute.googleapis.com/Instance |  ✅ | google_compute_instance |  ✅ | 
| compute.googleapis.com/InstanceGroup |  ✅ | google_compute_instance_group |  ✅ | 
| compute.googleapis.com/InstanceGroupManager |  ✅ | google_compute_instance_group_manager |  ✅ | 
| compute.googleapis.com/InstanceTemplate |  ✅ | google_compute_instance_template |  ✅ | 
| compute.googleapis.com/Network |  ✅ | google_compute_network |  ✅ | 
| compute.googleapis.com/Project | / | / | / | 
| compute.googleapis.com/RegionBackendService | 🔧 | google_compute_region_backend_service | 🔧 |
| compute.googleapis.com/Route |  ✅ | google_compute_route |  ✅ | 
| compute.googleapis.com/Router | 🔧 | google_compute_router | 🔧 |
| compute.googleapis.com/Snapshot |  ✅ | google_compute_snapshot |  ✅ |
| compute.googleapis.com/SslCertificate |  ✅ | google_compute_ssl_certificate(you need to set your private key manually) | ✅ |
| compute.googleapis.com/Subnetwork |  ✅ | google_compute_subnetwork |  ✅ | 
| compute.googleapis.com/TargetHttpProxy | ✅ | google_compute_target_http_proxy |  ✅ | 
| compute.googleapis.com/TargetHttpsProxy |  ✅ | google_compute_target_https_proxy |  ✅ |
| compute.googleapis.com/TargetInstance | / | / | / |
| compute.googleapis.com/TargetPool |  ✅(only support in default Region) | google_compute_target_pool |  ✅ |
| compute.googleapis.com/TargetTcpProxy | 🔧 | google_compute_target_tcp_proxy | 🔧 |
| compute.googleapis.com/TargetSslProxy | 🔧 | google_compute_target_ssl_proxy | 🔧 |
| compute.googleapis.com/TargetVpnGateway | 🔧 | google_compute_vpn_gateway | 🔧 |
| compute.googleapis.com/UrlMap |  ✅ | google_compute_url_map |  ✅ | 
| compute.googleapis.com/VpnTunnel | 🔧 | google_compute_vpn_tunnel | 🔧 |
**App Engine**
| appengine.googleapis.com/Application | 🔧 | google_app_engine_application(cannot delete app engine) | 🔧 |
| appengine.googleapis.com/Service | / | / | / |
| appengine.googleapis.com/Version | / | / | / |
**Google Kubernetes Engine** 
| container.googleapis.com/Cluster |  ✅ | google_container_cluster |  ✅ |
| container.googleapis.com/NodePool(beta) | 🔧 | google_container_node_pool | 🔧 |
**Cloud Billing**
| cloudbilling.googleapis.com/BillingAccount | / | / | / |
**Cloud Storage**
| storage.googleapis.com/Bucket |  ✅ | google_storage_bucket |  ✅ |
**Cloud DNS**
| dns.googleapis.com/ManagedZone |  ✅ | google_dns_managed_zone |  ✅ | 
| dns.googleapis.com/Policy |  ✅(only google-beta) | google_dns_policy |  ✅ |
**Cloud Spanner**
| spanner.googleapis.com/Instance | 🔧 | google_spanner_instance | 🔧 |
| spanner.googleapis.com/Database | 🔧 | google_spanner_database | 🔧 |
**BigQuery**
| bigquery.googleapis.com/Dataset| 🔧 | google_bigquery_dataset | 🔧 |
| bigquery.googleapis.com/Table | 🔧 | google_bigquery_table | 🔧 |
**Cloud Identity and Access Management** 
| iam.googleapis.com/Role | 🔧 | google_iam_member | 🔧 |
| iam.googleapis.com/ServiceAccount |  ✅ | google_service_account |  ✅ | 
**Cloud Pub/Sub**
| pubsub.googleapis.com/Topic |  ✅ | google_pubsub_subscription |  ✅ | 
| pubsub.googleapis.com/Subscription |  ✅ | google_pubsub_topic |  ✅ | 
**Cloud Dataproc**
| dataproc.googleapis.com/Cluster| 🔧 | google_dataproc_cluster | 🔧 |
| dataproc.googleapis.com/Job | 🔧 | google_dataproc_job | 🔧 |
**Cloud SQL** 
| sqladmin.googleapis.com/Instance |  ✅ | google_sql_database_instance |  ✅ |
**Cloud Bigtable**
| bigtableadmin.googleapis.com/Cluster| / | / | / |
| bigtableadmin.googleapis.com/Instance | 🔧 | google_bigtable_instance | 🔧 |
| bigtableadmin.googleapis.com/Table | 🔧 | google_bigtable_table | 🔧 |
**Google Kubernetes Engine**
| k8s.io/Node | / | / | / |
| k8s.io/Pod | / | / | / |
| k8s.io/Namespace | / | / | / |
| rbac.authorization.k8s.io/Role | / | / | / |
| rbac.authorization.k8s.io/RoleBinding | / | / | / |
| rbac.authorization.k8s.io/ClusterRole | / | / | / |
| rbac.authorization.k8s.io/RoleBinding | / | / | / |
