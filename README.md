# file-encryptor
CLI утилита для шифрования файлов на GO
----
В папке example находятся зашифрованные и расшифрованные файлы:
- encrypted-data.txt - файл с зашифрованными данными
- decrypted-data.txt - файл с расшифрованными данными

Примеры запуска:
- ./file-encryptor -path=test-data.txt -passphare=123 -encrypt - данная команда зашифрует указанный файл  с паролем 123
- ./file-encryptor -path=test-data.txt -passphare=123 -decrypt - данная команда расшифрует указанный файл с паролем 123, пароль должен быть тот который указывался при шифровании файла