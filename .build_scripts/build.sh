#!/bin/sh

cd `dirname $0`
WORK_DIR=`pwd`

# build the app server
echo building the app server...
cd ${WORK_DIR}/../src/cmd/app_server
go build
echo the app server was built.

# build the auth server
echo building the auth server...
cd ${WORK_DIR}/../src/cmd/auth_server
go build
echo the auth server was built.

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
cp ./src/cmd/auth_server/auth_server ./dist/
echo collected.
