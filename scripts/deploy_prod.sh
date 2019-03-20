#!/usr/bin/env bash

CD ../

version="$1"
name="voucher-datastore-service"
dockerPath="gcr.io/loyalsys-main/"$name":"$version""

echo "$dockerPath"

if [[ -z "$version" ]]
then
    echo "first argument should be version" && exit 500
fi

run_make() {
    # delete build file
    rm -f "$name"

    make -B version="$version" all

    [[ $? != 0 ]] && \
        echo "run makefile failed" && exit 100
    echo "run makefile success"
}

docker_build() {
    docker build -t "$dockerPath" -f Dockerfile .
    [[ $? != 0 ]] && \
        echo "Docker image build failed !" && exit 101
    echo "Docker image build success"
}

push_to_google() {
     gcloud docker -- push "$dockerPath"
     [[ $? != 0 ]] && \
         echo "push to google failed !" && exit 103
     echo "push to google success"
}

# execute makefile to build and test the server
echo "****************************************************** makefile ******************************************************"
run_make

# build the docker
echo "**************************************************** build docker ****************************************************"
docker_build

# push to google cloud
echo "***************************************************** push to gcs ****************************************************"
push_to_google

echo Done!