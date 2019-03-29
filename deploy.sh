#!/bin/bash

docker build -t go-deploy-poc .
readonly CONTAINER_NAME="gopoc"

if [[ $(docker ps | grep ${CONTAINER_NAME} | wc -l) -eq 1 ]]; 
then
    echo "A container is already running. It will be stopped, removed and rexecuted."
    docker stop ${CONTAINER_NAME}
    docker rm ${CONTAINER_NAME}
fi

docker run --name ${CONTAINER_NAME} -d -p 3030:3030 go-deploy-poc