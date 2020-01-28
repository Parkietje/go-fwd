#!/bin/bash

#works on UNIX
cd $PWD/cmd
go build -o $PWD/cmd
echo "build succes!"
echo "start execution"
./cmd
