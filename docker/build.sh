#!/bin/bash

gopath=$GOPATH
execBin=magicCenter
binpath=$gopath/bin/$execBin
imageID=""
imageName=muidea.ai/develop/$(echo $execBin | tr '[A-Z]' '[a-z]')
imageVersion=latest

function cleanUp()
{
    echo "cleanUp..."
    if [ -f log.txt ]; then
        rm -f log.txt
    fi

    if [ -f res.tar ]; then
        rm -f res.tar
    fi

    if [ -f $execBin ]; then
        rm -f $execBin
    fi
}

function prepareFile()
{
    echo "prepareFile..."
    cp $binpath ./
    if [ $? -ne 0 ]; then
        echo "prepareFile failed, copy magicCenter exception"
    fi

    cp -r ../static ./
    if [ $? -ne 0 ]; then
        echo "prepareFile failed, copy static exception"
    fi
    tar -caf res.tar static
    if [ $? -ne 0 ]; then
        echo "prepare file failed, compress failed exception."
        exit 1
    fi
    rm -rf static
}

function checkImage()
{
    echo "checkImage..."
    docker images | grep $1 | grep $2 > log.txt
    imageID=$(tail -1 log.txt|awk '{print $3}')
}

function buildImage()
{
    echo "buildImage..."
    docker build . > log.txt
    if [ $? -eq 0 ]; then
        echo "docker build success."
    else
        echo "docker build failed."
        exit 1
    fi

    imageID=$(tail -1 log.txt|awk '{print $3}')
}


function tagImage()
{
    echo "tag docker image..."
    docker tag $1 $2
    if [ $? -eq 0 ]; then
        echo "docker tag success."
    else
        echo "docker tag failed."
        exit 1
    fi
}

function rmiImage()
{
    echo "remove old docker image..."
    docker rmi $1:$2
    if [ $? -eq 0 ]; then
        echo "docker remove old image success."
    else
        echo "docker remove old image failed."
        exit 1
    fi
}

echo "build magicCenter docker image"

cleanUp

prepareFile

checkImage $imageName $imageVersion
if [ $imageID ]; then
    rmiImage $imageName $imageVersion
fi

buildImage

tagImage $imageID $imageName:$imageVersion

cleanUp