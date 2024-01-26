#!/bin/bash
# @docs https://docs.docker.com/engine/reference/commandline/container_run/

docker run --name zlabwork-service \
    -v ./:/app \
    -v /var/shared:/var/shared \
    -p 8080:8080 \
    -d zlabwork/service:dev
