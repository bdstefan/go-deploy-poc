#!/bin/bash

docker build -t go-deploy-poc .
readonly CONTAINER_NAME="gopoc"
readonly EXPOSE_PORT=3030

if [[ $(docker ps -a | grep ${CONTAINER_NAME} | wc -l) -eq 1 ]]; 
then
    echo "A container is already running. It will be stopped, removed and rexecuted."
    docker stop ${CONTAINER_NAME}
    docker rm ${CONTAINER_NAME}
fi

docker run --name ${CONTAINER_NAME} -d -p ${EXPOSE_PORT}:3030 go-deploy-poc

echo "The app is available at http://localhost:${EXPOSE_PORT}/10#change-up-limit-here"