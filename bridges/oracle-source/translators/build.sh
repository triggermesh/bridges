#!/bin/bash

# Perform the docker build/push of the translator

REPO=cab105
IMAGE_SLACK=${REPO}/oracledemotrans-slack
IMAGE_ZENDESK=${REPO}/oracledemotrans-zendesk

docker build --build-arg TARGET=slack -t $IMAGE_SLACK:latest .
docker build --build-arg TARGET=zendesk -t $IMAGE_ZENDESK:latest .

docker push $IMAGE_SLACK:latest
docker push $IMAGE_ZENDESK:latest
