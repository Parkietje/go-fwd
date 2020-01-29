#!/bin/bash

#get directory (works on UNIX)
cd $PWD

#generate an example config if no config exists
[ -f test.json ] || printf '{\n\t"authorization": {\n\t\t"username" : "example@email.com",\n\t\t"password" : "yourPassword",\n\t\t"server" : "smtp.email.com",\n\t\t"port" : "587"\n\t}\n}' > test.json