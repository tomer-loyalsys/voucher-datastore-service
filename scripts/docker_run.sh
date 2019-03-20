#!/usr/bin/env bash

CD ../

version="$1"
name="voucher-datastore-service"

# delete build file
rm -f "$name"

# execute makefile to build and test the server
echo "****************************************************** makefile ******************************************************"
make -B version="$version" all
echo

# build the docker
echo "**************************************************** build docker ****************************************************"
docker build -t "$name" -f Dockerfile .
echo

# run docker
echo "***************************************************** docker run *****************************************************"
Docker run -it "$name" .
