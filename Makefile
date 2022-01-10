#Challenge Makefile

start:
go run main.go

check:
go test ./src/database

#setup:
#if needed to setup the enviroment before starting it
