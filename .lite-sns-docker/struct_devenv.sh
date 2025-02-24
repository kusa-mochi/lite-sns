#!/bin/sh

cd /project
go mod tidy

cd /project/src/cmd/db_initializer
go run main.go
