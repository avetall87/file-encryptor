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

	fmt.Println("Run tests ...")
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

func TestRewriteFile_wrongFilename(t *testing.T) {
	err := rewriteFileData("non-existent-file", []byte(REWRITE_FILE_DATA_ETALON))
	if err == nil {
		t.Error("Rewrite wrong file - ", err)
	}
}

func TestValidateParameterData(t *testing.T) {
	err_encrypt := validateParameterData("some_file_name", "some_pass", true, false)
	if err_encrypt != nil {
		t.Error("Validate parameter failed for encrypt!")
	}

	err_decrypt := validateParameterData("some_file_name", "some_pass", false, true)
	if err_decrypt != nil {
		t.Error("Validate parameter failed for decrypt!")
	}
}

func TestValidateParameterData_empty_path(t *testing.T) {
	err := validateParameterData("", "some_pass", true, false)
	if err == nil {
		t.Error("Validate parameter failed for empty_path!")
	}

	if err.Error() != "ошибка:\n- не указан путь к файлу, параметер - path\n" {
		t.Error("Validate parameter failed for empty_path - wrong error text!")
	}
}

func TestValidateParameterData_empty_passpaher(t *testing.T) {
	err := validateParameterData("some_file_name", "", true, false)
	if err == nil {
		t.Error("Validate parameter failed for empty_passpaher!")
	}

	if err.Error() != "ошибка:\n- не указан пароль, параметер - passphare\n" {
		t.Error("Validate parameter failed for empty_passpaher - wrong error text!")
	}
}

func TestValidateParameterData_encrypt_and_decrypt_together_true(t *testing.T) {
	err := validateParameterData("some_file_name", "some_pass", true, true)
	if err == nil {
		t.Error("Validate parameter failed for encrypt_and_decrypt_together_true!")
	}

	if err.Error() != "ошибка:\n- одновременно указаны флаги encrypt и decrytp (можно указать только один флаг)\n" {
		t.Error("Validate parameter failed for encrypt_and_decrypt_together_true - wrong error text!")
	}
}

func TestValidateParameterData_encrypt_and_decrypt_together_false(t *testing.T) {
	err := validateParameterData("some_file_name", "some_pass", false, false)
	if err == nil {
		t.Error("Validate parameter failed for encrypt_and_decrypt_together_false!")
	}

	if err.Error() != "ошибка:\n- не указан ни один из флагов encrypt или decrytp (можно указать только один флаг)\n" {
		t.Error("Validate parameter failed for encrypt_and_decrypt_together_false - wrong error text!")
	}
}
