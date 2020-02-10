#!/bin/bash

#get directory (works on UNIX)
cd $PWD

#generate an example config if no config exists
[ -f config.json ] || printf '{\n\t"authentication": {\n\t\t"username" : "example@email.com",\n\t\t"password" : "yourPassword",\n\t\t"server" : "smtp.email.com",\n\t\t"port" : "587"\n\t}\n}' > config.json