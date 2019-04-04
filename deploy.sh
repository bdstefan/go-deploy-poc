#!/bin/bash

readonly CONTAINER_NAME="gopoc"
readonly EXPOSE_PORT=3030

if [[ $(docker ps | grep ${CONTAINER_NAME} | wc -l) -eq 1 ]]; 
then
    echo "A ${CONTAINER_NAME} container already exists. It will be stopped if is running, will be removed and rexecuted."
    docker-compose down
fi

docker-compose up -d --build

echo "The app is available at http://localhost:${EXPOSE_PORT}/liveness"