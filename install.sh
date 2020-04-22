#!/bin/sh
set -ex

cd go
export GOPATH=$PWD

# export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk1.8.0_144.jdk/Contents/Home/

#编译
go install  main

echo OK