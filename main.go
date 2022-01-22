package main

import (
	"file-encryptor/handler"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
)

const appVersion string = "File-hasher;1.0;AES256\n"

func readFileData(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении файла - %v", fileName)
	}
	return string(bytes), nil
}

func rewriteFileData(fileName string, data []byte) error {
	err := ioutil.WriteFile(fileName, data, fs.ModeAppend)
	if err != nil {
		return fmt.Errorf("ошибка записи закодированных данных в файл - %v", err)
	}

	return nil
}

func validateParameterData(path string, passphare string, encrypt bool, decrytp bool) error {
	var text string = ""

	if path == "" {
		text += "- не указан путь к файлу, параметер - path\n"
	}

	if passphare == "" {
		text += "- не указан пароль, параметер - passphare\n"
	}

	if encrypt && decrytp {
		text += "- одновременно указаны флаги encrypt и decrytp (можно указать только один флаг)\n"
	} else if !encrypt && !decrytp {
		text += "- не указан ни один из флагов encrypt или decrytp (можно указать только один флаг)\n"
	}

	if text != "" {
		return fmt.Errorf("ошибка:\n%v", text)
	} else {
		return nil
	}
}

func encryptAndSave(passphare string, str string, path string) {
	aesData, err := handler.EncryptAES([]byte(handler.GetMD5Hash(passphare)), []byte(str))
	if err != nil {
		log.Fatal("ошибка: не получилось зашифровать данные из файла - ", err)
	}
	rewriteFileData(path, []byte(appVersion+aesData))
}

func decryptAndSave(passphare string, str string, path string) {
	i := strings.Index(str, "\n")
	processedData := str[i:]

	data, err := handler.DecryptAES([]byte(handler.GetMD5Hash(passphare)), processedData)

	if err != nil {
		log.Fatal("ошибка: не получилось расшифровать данные из файла - ", err)
	}
	rewriteFileData(path, []byte(data))
}

func main() {

	path := flag.String("path", "", "путь к файлу")
	passphare := flag.String("passphare", "", "пароль")
	encrypt := flag.Bool("encrypt", false, "зашифровать файл (признак)")
	decrypt := flag.Bool("decrypt", false, "расшифровать файл (признак)")

	flag.Parse()

	// Валидация параметров командной строки
	err := validateParameterData(*path, *passphare, *encrypt, *decrypt)
	if err != nil {
		log.Fatal(err)
	}

	// читаем данные файла в переменную
	fData, err := readFileData(*path)
	if err != nil {
		log.Fatalf("ошибка: не удалось прочитать данные из файла: %v", err)
	}

	if *encrypt {
		encryptAndSave(*passphare, fData, *path)
	}

	if *decrypt {
		decryptAndSave(*passphare, fData, *path)
	}
}
