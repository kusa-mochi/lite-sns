#!/bin/sh

# .lite-sns-docker/docker-compose.dev.yml を使って起動した環境で
# lite-sns クライアントを起動するスクリプト。

cd `dirname $0`
WORK_DIR=`pwd`
cd ${WORK_DIR}/../src/cmd/client/lite-sns-client
npm run dev
