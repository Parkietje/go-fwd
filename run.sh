#!/bin/bash

#works on UNIX
cd $PWD
[ -f config.json ] || printf '{\n"authorization": {\n\t"username" : "example@email.com",\n\t"password" : "yourPassword",\n\t"server" : "smtp-relay.gmail.com:587"\n\t}\n}' > test.json

cd cmd
go build -o $PWD/cmd
echo "build succes!"
echo "start execution"
./cmd