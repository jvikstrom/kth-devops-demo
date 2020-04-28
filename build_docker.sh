#!/bin/bash

## Build server docker file.
docker build -t hello-server-image:latest -f ./docker/server.DOCKERFILE .
## Build client docker file.
docker build -t hello-client-image:latest -f ./docker/client.DOCKERFILE .
