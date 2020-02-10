#!/bin/bash

#single stage
#docker build -f Dockerfile -t go-fwd:latest .

#multistage
docker build -t go-fwd-multi:latest -f Dockerfile.multistage .