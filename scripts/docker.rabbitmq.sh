#!/bin/bash

docker run -d -p 15691:15691 -p 15692:15692 -p 25672:25672 -p 4369:4369 -p 5671:5671 -p 5672:5672 \
      -v $PWD/docker/data/rabbitmq:/var/lib/rabbitm \
      -v $PWD/docker/config/influx:/etc/influxdb2 \
      -e RABBITMQ_DEFAULT_USER="omgind" \
      -e RABBITMQ_DEFAULT_PASS="123456" \
      --name rabbitmq
      rabbitmq:alpine
