package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

var ()

func TestCreateFile(t *testing.T) {
	t.Helper()
	dir, _ := ioutil.TempDir("", "")
	defer os.RemoveAll(dir)
	value := "test"
	buf := bytes.NewBufferString(value)
	filename := dir + "/test"

	//Create test file written in "test"
	if err := CreateFile(filename, *buf); err != nil {
		t.Fatalf("failed to call CreateFile(): %v", err)
	}

	//Read test file
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("failed to oepn File: %v", err)
	}
	testbuf, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("failed to read File: %v", err)
	}
	result := string(testbuf)
	if result != value {
		t.Fatalf("the result is not an expected value: %v != %v", result, value)
	}

}
