#!/bin/bash

echo "Applying db migrations"
migrate -path db/migrations/dev -database "mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DB_NAME}" up

echo "Building app"
go build

echo "Running app"
./interest-points-api