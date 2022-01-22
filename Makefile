include .env

build:
	rm -rf bin
	go build -o bin/file-encryptor .

run-encrypt:
	go run ./ -path=test-data.txt -passphare=123 -encrypt

run-decrypt:
	go run ./ -path=test-data.txt -passphare=123 -decrypt