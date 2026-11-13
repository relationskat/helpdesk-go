#!/bin/sh
set -eu
base=${BASE_URL:-http://127.0.0.1:8080}
curl -fsS "$base/healthz" | tr -d '\n' | grep -q '"status":"ok"'
created=$(curl -fsS -X POST "$base/tickets" -H 'Content-Type: application/json' -d '{"title":"Lab 1 smoke check"}')
printf '%s\n' "$created" | grep -q '"status":"open"'
id=$(printf '%s' "$created" | sed -n 's/.*"id":\([0-9][0-9]*\).*/\1/p')
test -n "$id"
curl -fsS "$base/tickets" | grep -q "\"id\":$id,\"title\":\"Lab 1 smoke check\""
echo 'PASS: health, create ticket, find created ticket'
