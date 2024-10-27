#!/bin/sh

cd `dirname $0`
WORK_DIR=`pwd`
DIST_DIR=${WORK_DIR}/../dist
cd $DIST_DIR

# kill the runnning app server
pkill ./app_server

# set permissions
chmod +x ./app_server

# run the app server
./app_server &

# # run the web server for clients
