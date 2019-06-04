package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"golang.org/x/xerrors"
)

//File name
const (
	RepositoryDirectory = "github.com/cloud-ace/terraformit-gcp"
	ProviderTfName      = "terraformit-gcp-provider.tf"
	ResourceTfName      = "terraformit-gcp-resource.tf"
	BackendTfName       = "terraformit-gcp-backend.tf"
	TfName              = "terraformit-gcp-resource.tf"
	ImportShName        = "terraformit-gcp-import.sh"
	TfstateName         = "terraform.tfstate"
	ObjectNamePrefix    = "CloudAssetMetadata"
)

func CreateFile(fileName string, buf bytes.Buffer) error {
	file, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return xerrors.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	file.Write(buf.Bytes())
	if err != nil {
		return xerrors.Errorf("Error: %w", err)
	}
	fmt.Println(fileName, " created.")
	return nil
}

//templateに渡すfunction
var FuncMap = template.FuncMap{
	"plus": func(a int) int { return a + 1 },
	"createRegion": func(s string) string {

		slice := strings.Split(s, "/")
		return slice[8]
	},
	"createZone": func(s string) string {

		slice := strings.Split(s, "/")
		return slice[8]
	},
	"createProject": func(s string) string {

		slice := strings.Split(s, "/")
		return slice[4]
	},
	"convertSlashAndDotToDash": func(s string) string {
		tmp := strings.Replace(s, "/", "-", -1)
		return strings.Replace(tmp, ".", "-", -1)
	},
	"convertsingleTodouble": func(s string) string {
		return strings.Replace(s, "`", "\"", -1)
	},
	"escapeDoubleQuote": func(s string) string {
		return strings.Replace(s, "\"", "\\\"", -1)
	},
	"removeAtmark": func(s string) string {
		return strings.Replace(s, "@", "-", -1)
		// ss := strings.Split(s, "@")
		// return ss[0]
	},
	"getLastWord": func(s string) string {

		ss := strings.Split(s, "/")
		return ss[len(ss)-1]
	},
}

func GetGoPath() string {
	GoPath := os.Getenv("GOPATH")
	if GoPath == "" {
		GoPath = os.Getenv("HOME") + "/go"
	}
	return GoPath
}
