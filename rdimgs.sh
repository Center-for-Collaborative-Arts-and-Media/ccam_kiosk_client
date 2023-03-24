#!/bin/bash

docker compose down

mkcert_image=$(docker images | grep mkcert | awk '{print $1}')

docker rmi $(docker images | grep -v $mkcert_image | awk '{print $3}')