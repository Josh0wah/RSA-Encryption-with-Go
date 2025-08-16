**Usage**
This program utilizes CLI arguments, so

To encrypt a text file run:
go run main.go -e message-to-encrypt.txt
Backup the keys in keys.txt so they aren't lost the next time you encrypt a file

To decrypt a text file run:
go run main.go -d message-to-encrypt.txt d n
