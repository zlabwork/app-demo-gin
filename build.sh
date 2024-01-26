#!/bin/bash
# @docs https://docs.docker.com/engine/reference/commandline/login/
# @docs https://docs.docker.com/engine/reference/commandline/image_push/

name="zlabwork/service"
tag=$(date "+%Y%m%d")-$(date "+%s")
echo "$name:$tag"

docker build -t "$name:$tag" -f Dockerfile .
docker push "$name:$tag"
