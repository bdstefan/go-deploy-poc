#!/bin/bash

docker build -t go-deploy-poc .
readonly CONTAINER_NAME="gopoc"
readonly EXPOSE_PORT=3030

if [[ $(docker ps -a | grep ${CONTAINER_NAME} | wc -l) -eq 1 ]]; 
then
    echo "A ${CONTAINER_NAME} container already exists. It will be stopped if is running, will be removed and rexecuted."
    docker stop ${CONTAINER_NAME} > /dev/null
    docker rm ${CONTAINER_NAME} > /dev/null
fi

docker run --name ${CONTAINER_NAME} -d -p ${EXPOSE_PORT}:3030 go-deploy-poc > /dev/null

echo "The app is available at http://localhost:${EXPOSE_PORT}/liveness"