@echo off
cd ./src
go build -o ./bin/compact.exe
cd ./bin
.\compact.exe ./test.hello