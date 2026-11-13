#!/bin/sh
set -eu
name=helpdesk-lab1-check
docker build -t helpdesk-go:lab-1 .
docker run --rm -d --name "$name" -p 127.0.0.1:18081:8080 helpdesk-go:lab-1
trap 'docker stop "$name" >/dev/null' EXIT
i=0
until curl -fsS http://127.0.0.1:18081/healthz >/dev/null; do
  i=$((i+1))
  test "$i" -lt 30
  sleep 1
done
BASE_URL=http://127.0.0.1:18081 sh tests/smoke.sh
