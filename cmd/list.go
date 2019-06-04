package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cloud-ace/terraformit-gcp/cloudasset"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `list`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			log.Printf("no arg")
			os.Exit(1)
		}

		subcommand := args[0]

		switch subcommand {
		case "cloudasset":
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
			supporttrimDuplicate := map[string]bool{}
			noSupporttrimDuplicate := map[string]bool{}
			for _, v := range metaStructs {
				v.ResourceMetadata = cloudasset.ResouceNameMap[v.AssetType]
				if v.ResourceMetadata.Name == "" {
					if noSupporttrimDuplicate[v.AssetType] == false {
						noSupporttrimDuplicate[v.AssetType] = true
					}
				} else {
					if supporttrimDuplicate[v.AssetType] == false {
						supporttrimDuplicate[v.AssetType] = true
					}
				}
			}
			fmt.Println("Support:")
			for k, _ := range supporttrimDuplicate {
				fmt.Println(k)
			}
			fmt.Println("\nNoSupport:")
			for k, _ := range noSupporttrimDuplicate {
				fmt.Println(k)
			}

		default:
			log.Printf("wrong subcommand")
			os.Exit(1)
		}
	},
}

func init() {

	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&createOptions.file, "file", "f", "", "CloudAssetAPI metadata file")
}
