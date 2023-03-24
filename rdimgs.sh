#!/bin/bash

docker stop $(docker ps -a -q)

docker rm $(docker ps -a -q)

mkcert_image=$(docker images | grep mkcert | awk '{print $1}')

docker rmi $(docker images | grep -v $mkcert_image | awk '{print $3}')