package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cloud-ace/terraformit-gcp/cloudasset"
	"github.com/cloud-ace/terraformit-gcp/resources"
	"github.com/cloud-ace/terraformit-gcp/terraformUtil"
	"github.com/cloud-ace/terraformit-gcp/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Subcommands = []string{"cloudasset", "importfiles", "tfstate", "tffile"}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create files or tfstate",
	Long: `subcommand = [ cloudasset, importfiles, tfstate, tffile]
	terraformit-gcp create cloudasset
	terraformit-gcp create importfiles -f /xxx/xxx or gs://xxxx.xxxx/xxxxxx
	terraformit-gcp create tfstate
	terraformit-gcp create tffile`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			log.Printf("Please set what you create. Support: %v", Subcommands)
			os.Exit(1)
		}

		subcommand := args[0]

		switch subcommand {
		case "cloudasset":
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
			log.Println("################## create CloudAsset end ##################")

		case "importfiles":
			//Cloud Assetを取得
			log.Println("################## create importfiles start ##################")
			var metaStructs []cloudasset.Metadata
			if createOptions.file == "" {
				log.Println("you must set -f option. You can set localfile or gs://xxxx/xxxx")
				os.Exit(1)
			} else if strings.Contains(createOptions.file, "gs://") {
				//set credential
				if createOptions.secret != "" {
					os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", createOptions.secret)
				}

				//get metadata from gcs
				cloudAsset, err := cloudasset.ReadFileFromGCS(createOptions.file)
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

			} else {
				//get metadata from local file
				metaStrings, err := cloudasset.ReadMetaFile(createOptions.file)
				if err != nil {
					log.Printf("%+v\n", err)
					os.Exit(1)
				}
				metaStructs, err = cloudasset.MetaStringToStruct(metaStrings)
				if err != nil {
					log.Printf("%+v\n", err)
					os.Exit(1)
				}
			}

			//remove default resources
			metaStructs = cloudasset.RemoveDefaultResources(metaStructs, viper.GetStringMapString("Terraform")["resource-default-network"], viper.GetStringMapString("Terraform")["resource-default-subnetwork"], viper.GetStringMapString("Terraform")["resource-default-route"], viper.GetStringMapString("Terraform")["resource-default-firewall"])

			//create file
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
			if viper.GetStringMapString("Terraform")["workspace"] == "default" || viper.GetStringMapString("Terraform")["backend-type"] == "gcs" {
				if err := cloudasset.CreateFileForImport(metaStructs, "backend", viper.GetStringMapString("Terraform")["gcp-provider-default-region"], viper.GetStringMapString("Terraform")["backend-type"], viper.GetStringMapString("Terraform")["backend-location"], viper.GetStringMapString("Terraform")["provider"]); err != nil {
					log.Printf("%+v\n", err)
					os.Exit(1)
				}
			}

			log.Println("Importfiles created successfully")
			log.Println("################## create importfiles end ##################")

		case "tfstate":
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

		case "tffile":
			log.Println("################## create tffile start ##################")
			var tfstatePath string
			var err error
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

		default:
			log.Printf("wrong args: %v", subcommand)
			log.Printf("Please set what you create. Support: %v", Subcommands)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&createOptions.file, "file", "f", "", "CloudAssetAPI metadata file location or tfstate location")
	createCmd.Flags().StringVarP(&createOptions.secret, "secret", "s", "", "service account credential json file")
}
