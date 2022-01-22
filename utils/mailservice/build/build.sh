#!/bin/bash

gittoken=$GITTOKEN

echo "$gittoken"

docker build \
    -t mailservice \
    --build-arg GITTOKEN="$gittoken" \
    -f ./Dockerfile ..