package terraformUtil

import (
	"bytes"
	"os/exec"
	"path"
	"strings"

	"github.com/cloud-ace/terraformit-gcp/utils"
	"github.com/spf13/viper"
	"golang.org/x/xerrors"
)

func TerraformCmd(subcommand ...string) error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("terraform", subcommand...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	//subcommandごとにエラー処理
	if err != nil {
		return xerrors.Errorf("Error: %v \ncommand:terraform %v \n%s", err, subcommand, stderr.String())
	}
	return nil
}

func ExecuteImportSh() error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("sh", path.Base(utils.ImportShName))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		return xerrors.Errorf("Import Error: %v\n%s", err, stderr.String())
	}
	if strings.Contains(stdout.String(), "Error") || strings.Contains(stderr.String(), "Error") {
		return xerrors.Errorf("Error! %v\n%s", err, stderr.String())
	}
	return nil
}

func GetTfstatePath() (string, error) {
	workspace := viper.GetStringMapString("Terraform")["workspace"]
	if workspace == "" {
		return "", xerrors.New("Error: workspace is not set")
	}
	switch workspace {
	case "default":
		return "./" + utils.TfstateName, nil
	default:
		return "./" + "terraform.tfstate.d/" + workspace + "/" + utils.TfstateName, nil
	}
}
