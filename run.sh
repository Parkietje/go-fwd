#!/bin/bash

#get directory (works on UNIX)
cd $PWD/cmd

#build application executable
go build -o $PWD/main
echo "build succes!"

#run executable
echo "start execution"
./main