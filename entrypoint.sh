#!/bin/bash

set -e

echo "=============================================================================================="
echo "To check any other containers log,"
echo "run \"docker compose up --build --force-recreate --always-recreate-deps --renew-anon-volumes\""
echo "=============================================================================================="

readonly MAX_TRIES=120

attempt_num=0
until curl -I -s http://mysql-router-http:8080 -k >/dev/null; do
  echo "Waiting for mysql router(http) ($attempt_num/$MAX_TRIES)"
  sleep $((attempt_num++))
  if ((attempt_num == MAX_TRIES)); then
    exit 1
  fi
done

attempt_num=0
until curl -I -s https://mysql-router-https:8443 -k >/dev/null; do
  echo "Waiting for mysql router(https) ($attempt_num/$MAX_TRIES)"
  sleep $((attempt_num++))
  if ((attempt_num == MAX_TRIES)); then
    exit 1
  fi
done

echo "mysql router is ready."

#
# start go test
#

cd /go/src/mysqlrouter-go
go test -v .

