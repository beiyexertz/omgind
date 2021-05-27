#!/usr/bin/env bash

docker run -d -p 8086:8086 \
      -v $PWD/docker/data/influxdb:/var/lib/influxdb2 \
      -v $PWD/docker/config/influx:/etc/influxdb2 \
      -e DOCKER_INFLUXDB_INIT_MODE=setup \
      -e DOCKER_INFLUXDB_INIT_USERNAME="omgind" \
      -e DOCKER_INFLUXDB_INIT_PASSWORD="123456" \
      -e DOCKER_INFLUXDB_INIT_ORG="omgind" \
      -e DOCKER_INFLUXDB_INIT_BUCKET="omgind" \
      -e DOCKER_INFLUXDB_INIT_RETENTION=1w \
      -e DOCKER_INFLUXDB_INIT_ADMIN_TOKEN="123456" \
      --name influxdb \
      influxdb:2.0-alpine
