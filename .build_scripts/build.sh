#!/bin/sh

cd `dirname $0`
WORK_DIR=`pwd`

# download dependencies for the servers
echo start go mod tidy...
cd ${WORK_DIR}/..
go mod tidy
echo end go mod tidy.

# build the app server
echo building the app server...
cd ${WORK_DIR}/../src/cmd/app_server
go build
echo the app server was built.

# build the client
echo building the client...
echo the client was build.

# remove the dist dir
echo removing the dist dir...
cd ${WORK_DIR}/..
rm -rf ./dist/
echo the dist dir was removed.

# collect the programs
echo collecting the programs...
cd ${WORK_DIR}/..
mkdir -p dist
cp ./src/cmd/app_server/app_server ./dist/
cp ./src/cmd/auth_server/auth_server ./src/cmd/server_common/server_config.json ./dist/
echo collected.
