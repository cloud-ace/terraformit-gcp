package cloudasset

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	metadataFile      = "testdata/CloudAsset-metadata.golden"
	metadataAssetType = "cloudresourcemanager.googleapis.com/Project"
	metaByteSlice     [][]byte
	subcommands       = []string{"provider", "import", "resource"}
)

func TestReadMetaFile(t *testing.T) {
	t.Helper()

	meta_strings, err := ReadMetaFile(metadataFile)
	if err != nil {
		t.Fatalf("failed to call ReadMetaFile(): %v", err)
	}

	if meta_strings == nil {
		t.Fatalf("the result is not an expected value: nil")
	}

}

func TestMetaStringToStruct(t *testing.T) {
	t.Helper()

	meta_strings, err := ReadMetaFile(metadataFile)
	if err != nil {
		t.Fatalf("failed to call ReadMetaFile(): %v", err)
	}
	meta_structs, err := MetaStringToStruct(meta_strings)
	if err != nil {
		t.Fatalf("failed to call MetaStringToStruct(): %v", err)
	}

	if meta_structs == nil {
		t.Fatalf("the result is not an expected value: nil")
	} else if meta_structs[0].AssetType != metadataAssetType {
		t.Fatalf("the result is not an expected value: %v != %v", meta_structs[0].AssetType, metadataAssetType)
	}

}

func TestMetaByteToStruct(t *testing.T) {
	t.Helper()

	f, err := os.Open(metadataFile)
	if err != nil {
		t.Fatalf("failed to oepn File: %v", err)
	}
	testbuf, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("failed to read File: %v", err)
	}

	metaByteSlice := bytes.Split(testbuf, []byte("\n"))

	meta_structs, err := MetaByteToStruct(metaByteSlice)
	if err != nil {
		t.Fatalf("failed to call MetaByteToStruct(): %v", err)
	}

	if meta_structs == nil {
		t.Fatalf("the result is not an expected value: nil")
	} else if meta_structs[0].AssetType != metadataAssetType {
		t.Fatalf("the result is not an expected value: %v != %v", meta_structs[0].AssetType, metadataAssetType)
	}

}

func TestInsertIntoBuffer(t *testing.T) {
	t.Helper()
	meta_strings, err := ReadMetaFile(metadataFile)
	if err != nil {
		t.Fatalf("failed to call ReadMetaFile(): %v", err)
	}
	meta_structs, err := MetaStringToStruct(meta_strings)
	if err != nil {
		t.Fatalf("failed to call MetaStringToStruct(): %v", err)
	}

	var buf bytes.Buffer
	for _, subcommand := range subcommands {

		for _, v := range meta_structs {
			if err := InsertIntoBuffer(&v, &buf, ResouceNameMap, FuncMap, subcommand); err != nil {
				t.Fatalf("failed to call InsertIntoBuffer(): %v", err)
			}
		}
		if !strings.Contains(buf.String(), subcommand) {
			t.Fatalf("the result is not an expected value: %v doesnt contain  %v", buf.String(), "subcommand")
		}

	}

}
func TestCreateFileForImport(t *testing.T) {
	t.Helper()
	defer os.Remove(ResourceTfName)
	defer os.Remove(ImportShName)
	defer os.Remove(ProviderTfName)

	meta_strings, err := ReadMetaFile(metadataFile)
	if err != nil {
		t.Fatalf("failed to call ReadMetaFile(): %v", err)
	}
	meta_structs, err := MetaStringToStruct(meta_strings)
	if err != nil {
		t.Fatalf("failed to call MetaStringToStruct(): %v", err)
	}

	for _, v := range subcommands {
		if err := CreateFileForImport(meta_structs, v); err != nil {
			t.Fatalf("failed to call CreateFileForImport(): %v subcommand:%v", err, v)
		}
	}

}
