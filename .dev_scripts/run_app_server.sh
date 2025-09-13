#!/bin/sh

# .lite-sns-docker/docker-compose.dev.yml を使って起動した環境で
# アプリケーションサーバーを起動するスクリプト。
# (DB初期化も行う)

cd `dirname $0`
WORK_DIR=`pwd`
cd ${WORK_DIR}/../src/cmd/db_initializer
go run main.go
cd ${WORK_DIR}/../src/cmd/app_server
go run main.go
