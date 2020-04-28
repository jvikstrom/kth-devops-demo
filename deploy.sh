#!/bin/bash
### Note this only works for the initial deployment.
### If any changes are made to the images the deployments must either be updated
### to use the new image or deleted and redeployed. Else, the image won't change.

./build_docker.sh
## Deploy the Kubernetes configs.
cd k8s
./deploy.sh
cd ..

