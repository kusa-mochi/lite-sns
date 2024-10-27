#!/bin/sh

cd `dirname $0`
WORK_DIR=`pwd`
DIST_DIR=${WORK_DIR}/../dist
cd $DIST_DIR

# kill the runnning auth server
pkill ./auth_server

# set permissions
chmod +x ./auth_server

# run the auth server
./auth_server &

# # run the web server for clients
