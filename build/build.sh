#!/usr/bin/env bash

version=v0.7.0
lua_version=v0.3.0
docker_tag=0.7.0
echo $version
echo docker_tag
# Delete every Docker containers
# Must be run first because images are attached to containers
docker rm -f $(docker ps -a -q)

# Delete every Docker image
docker rmi -f $(docker images -q)

#versionTag=$(git describe --tags `git rev-list --tags --max-count=1`)
#dockerVersion=$(echo $versionTag | cut -c2-)

# HEAD build: commit=$(git rev-parse --short HEAD 2>/dev/null)
docker build --rm --build-arg "commit=tags/${version}" --build-arg "lua_commit=tags/${lua_version}" --build-arg "version=${version}(${lua_version})" -t "skillzint/nakama:${docker_tag}" .

image_id=$(docker images --filter=reference="skillzint/nakama:${docker_tag}" --format "{{.ID}}")
echo image_id

docker tag "${image_id}" "skillzint/nakama:${docker_tag}"
#docker tag "${image_id}" "skillzint/nakama:latest"
docker push "skillzint/nakama:${docker_tag}"
#docker push "skillzint/nakama:latest"
