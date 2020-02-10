#!/bin/bash

#get directory (works on UNIX)
cd $PWD

#build application executable
go build -o main .
echo "build succes!"

#run executable
echo "start execution"
./main