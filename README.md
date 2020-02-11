# go-fwd
a simple mail forwarding service written in go.

## 1) Setting up
First, clone repo and generate example `config.json`:

`$ git clone git@github.com:Parkietje/go-fwd.git && cd go-fwd && ./setup.sh`


## 2) Configuration
Secondly, fill in `config.json`:

- username: your accountname with the email service provider
- password: your password (*)
- server: [how to find your provider's mail server](https://serversmtp.com/what-is-my-smtp/)
- port: [use port 587](https://www.mailgun.com/blog/which-smtp-port-understanding-ports-25-465-587/)

_(* if you are using google you might have to [generate 3rd party app password](https://support.google.com/accounts/answer/185833))_

## 3) Build and test
After filling in `config.json` you can test go-fwd by running it on localhost:

`$ ./run.sh`

In another terminal, make a POST request to the server:

`$ ./post.sh`

## 4) Dockerize
Build a docker image which runs the go-fwd executable:

`$ ./docker-build.sh`

Run the container locally:

`$ ./docker-run.sh`


