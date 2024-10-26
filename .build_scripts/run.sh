#!/bin/sh

cd `dirname $0`
WORK_DIR=`pwd`
DIST_DIR=${WORK_DIR}/../dist
cd $DIST_DIR

# set permissions
chmod +x ./app_server
chmod +x ./auth_server

# run the app server
./app_server &

# run the auth server
./auth_server &

# run the web server for clients
