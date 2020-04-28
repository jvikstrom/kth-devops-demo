#!/bin/bash

## Change the environment to use minikube.
## The reason for this is to put the docker image in the local minikube registry.
eval $(minikube docker-env)

## Build server docker file.
docker build -t hello-server-image:latest -f ./docker/server.DOCKERFILE .
## Build client docker file.
docker build -t hello-client-image:latest -f ./docker/client.DOCKERFILE .
