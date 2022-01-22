package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

const (
	NOT_ENCRYPTED_DATA_ETALON = "some data\n123\n456"
	REWRITE_FILE_DATA_ETALON  = "456"
	REWRITE_FILE_NAME         = "test-data/rewrite-file"
)

func init() {
	fmt.Println("Prepare testing data ...")

	// int file data for ewrite-file
	ioutil.WriteFile(REWRITE_FILE_NAME, []byte("123"), fs.ModeAppend)

}

func TestReadFileData(t *testing.T) {
	data, err := readFileData("test-data/not-encrypted")
	if err != nil {
		t.Error("File not readed - ", err)
		return
	}

	if data != NOT_ENCRYPTED_DATA_ETALON {
		t.Error("Readed data no equal etalon data - ", err)
	}
}

func TestReadFileData_wrongFilename(t *testing.T) {
	_, err := readFileData("test-data/non-existent-file")

	if err == nil {
		t.Error("Read wrong file - ", err)
	}
}

func TestRewriteFile(t *testing.T) {
	err := rewriteFileData(REWRITE_FILE_NAME, []byte(REWRITE_FILE_DATA_ETALON))
	if err != nil {
		t.Error("File data not rewrited - ", err)
	}
}
