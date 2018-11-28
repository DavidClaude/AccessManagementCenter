#!/bin/bash
rm -rf userHttpEntrance
export BASEDIR=`pwd`
export GOPATH=$BASEDIR/../../
#export GOROOT=/home/cancui2/workspace/goroot/go1.9.2/go
export GOROOT=/usr/local/go
chmod -R +x $GOROOT
export PATH=$GOROOT/bin:$PATH
go version
go build  -o userHttpEntrance
echo 'finished'
