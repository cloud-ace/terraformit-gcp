package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cloud-ace/terraformit-gcp/cloudasset"
	"github.com/cloud-ace/terraformit-gcp/resources"
	"github.com/cloud-ace/terraformit-gcp/terraformUtil"
	"github.com/cloud-ace/terraformit-gcp/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "execute all steps",
	Long: `Following steps below are executed.
	create CloudAssetMetadata calling CloudAssetAPI
	get CloudAssetMetadata from GCS
	create ImportFiles
	"terraform init"
	"terraform workspace new"
	"terraform import"(create tfstate)
	create tffile
	"terraform plan"`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 0 {
			log.Printf("args are not required")
			os.Exit(1)
		}

		//Cloud Assetを作成
		log.Println("################## create CloudAsset start ##################")
		objectName, err := cloudasset.CreateCloudAsset(viper.GetStringMapString("CloudAsset")["credential"], viper.GetStringMapString("CloudAsset")["bucket"], viper.GetStringMapString("CloudAsset")["project-number"], utils.ObjectNamePrefix)
		if err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
		fmt.Println("CloudAsset created successfully.")
		fmt.Println("CloudAsset object:", objectName)
		URI := "gs://" + viper.GetStringMapString("CloudAsset")["bucket"] + "/" + objectName
		fmt.Println("CloudAsset URI:", URI)
		time.Sleep(8 * time.Second)
		log.Println("################## create CloudAsset end ##################")

		//Cloud Assetを取得
		log.Println("################## create importfiles start ##################")
		var metaStructs []cloudasset.Metadata
		//リモートファイルから取得
		cloudAsset, err := cloudasset.ReadFileFromGCS(URI)
		if err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
		sliceCloudAsset := bytes.Split(cloudAsset, []byte("\n"))
		//delete last slice because no data is included
		sliceCloudAsset = sliceCloudAsset[:len(sliceCloudAsset)-1]
		metaStructs, err = cloudasset.MetaByteToStruct(sliceCloudAsset)
		if err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}

		//remove default resources
		metaStructs = cloudasset.RemoveDefaultResources(metaStructs, viper.GetStringMapString("Terraform")["resource-default-network"], viper.GetStringMapString("Terraform")["resource-default-subnetwork"], viper.GetStringMapString("Terraform")["resource-default-route"], viper.GetStringMapString("Terraform")["resource-default-firewall"])

		// terraform init/importのためのファイルを作成
		if err := cloudasset.CreateFileForImport(metaStructs, "provider", viper.GetStringMapString("Terraform")["gcp-provider-default-region"], viper.GetStringMapString("Terraform")["backend-type"], viper.GetStringMapString("Terraform")["backend-location"], viper.GetStringMapString("Terraform")["provider"]); err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
		if err := cloudasset.CreateFileForImport(metaStructs, "import", viper.GetStringMapString("Terraform")["gcp-provider-default-region"], viper.GetStringMapString("Terraform")["backend-type"], viper.GetStringMapString("Terraform")["backend-location"], viper.GetStringMapString("Terraform")["provider"]); err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
		if err := cloudasset.CreateFileForImport(metaStructs, "resource", viper.GetStringMapString("Terraform")["gcp-provider-default-region"], viper.GetStringMapString("Terraform")["backend-type"], viper.GetStringMapString("Terraform")["backend-location"], viper.GetStringMapString("Terraform")["provider"]); err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
		if viper.GetStringMapString("Terraform")["backend-type"] == "gcs" {
			if err := cloudasset.CreateFileForImport(metaStructs, "backend", viper.GetStringMapString("Terraform")["gcp-provider-default-region"], viper.GetStringMapString("Terraform")["backend-type"], viper.GetStringMapString("Terraform")["backend-location"], viper.GetStringMapString("Terraform")["provider"]); err != nil {
				log.Printf("%+v\n", err)
				os.Exit(1)
			}
		}
		log.Println("Importfiles created successfully")
		log.Println("################## create importfiles end ##################")

		// //terraform init
		log.Println("################## create tfstate start ##################")
		if err := terraformUtil.TerraformCmd("init"); err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
		//workspace作成
		workspace := viper.GetStringMapString("Terraform")["workspace"]
		if workspace != "default" {
			if err := terraformUtil.TerraformCmd("workspace", "new", workspace); err != nil {
				log.Printf("%+v\n", err)
				os.Exit(1)
			}
			if err := terraformUtil.TerraformCmd("workspace", "select", workspace); err != nil {
				log.Printf("%+v\n", err)
				os.Exit(1)
			}
		}
		//terraform.sh -- terraform import
		if err := terraformUtil.ExecuteImportSh(); err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
		log.Println("tfstate created successfully")
		log.Println("################## create tfstate end ##################")

		log.Println("################## create tffile start ##################")
		var tfstatePath string
		var tfstateBytes []byte
		if createOptions.file != "" {
			tfstatePath = createOptions.file
			log.Printf("get tfstate  %v\n", tfstatePath)
			tfstateBytes, err = resources.ReadTfstateFile(tfstatePath)
			if err != nil {
				log.Printf("%+v\n", err)
				os.Exit(1)
			}
		} else {
			if viper.GetStringMapString("Terraform")["backend-type"] == "gcs" {
				url := "gs://" + viper.GetStringMapString("Terraform")["backend-location"] + "/" + viper.GetStringMapString("Terraform")["workspace"] + ".tfstate"
				log.Printf("get tfstate from %v\n", url)
				tfstateBytes, err = cloudasset.ReadFileFromGCS(url)
				if err != nil {
					log.Printf("%+v\n", err)
					os.Exit(1)
				}
			} else if viper.GetStringMapString("Terraform")["backend-type"] == "local" {

				tfstatePath, err = terraformUtil.GetTfstatePath(viper.GetStringMapString("Terraform")["workspace"])
				if err != nil {
					log.Printf("%+v\n", err)
				}
				log.Printf("get tfstate  %v\n", tfstatePath)

				tfstateBytes, err = resources.ReadTfstateFile(tfstatePath)
				if err != nil {
					log.Printf("%+v\n", err)
					os.Exit(1)
				}
			} else {
				log.Printf("wrong backend-type %v\n", viper.GetStringMapString("Terraform")["backend-type"])
				os.Exit(1)
			}

		}

		tfstateStruct, err := resources.TfstateByteToStruct(tfstateBytes)
		if err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
		if err := resources.CreateTf(tfstateStruct); err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
		//terraform fmt
		if err := terraformUtil.TerraformCmd("fmt", utils.TfName); err != nil {
			log.Printf("%+v\n", err)
		}
		log.Println("tf file created successfully")
		log.Println("################## create tfstate end ##################")

		//terraform plan
		log.Println("################## terraform plan start ##################")
		if err := terraformUtil.TerraformCmd("plan"); err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
		log.Println("plan sccess")
		log.Println("################## terraform plan end ##################")
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
