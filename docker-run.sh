#!/bin/bash
#mkdir -p $PWD/logs
#run with volume attached
#docker run -d -p 8080:8080 -v $PWD/logs:/app/logs go-fwd

docker run -d -p 8080:8080 go-fwd:latest