# Mailserver
a simple mailing service written in go.

## Setting up
creates mailserver directory and generates example `config.json`:

`$ git clone git@github.com:Parkietje/mailserver.git && cd mailserver && ./setup.sh`

## Configuration
in config.json fill in for your email provider:

- username: your accountname with the email service provider
- password: your password *
- server: [how to find your mail server](https://serversmtp.com/what-is-my-smtp/)
- port: [use port 587](https://www.mailgun.com/blog/which-smtp-port-understanding-ports-25-465-587/)

_(* if you are using google you might have to [generate 3rd party app password](https://support.google.com/accounts/answer/185833))_

## Run locally
after filling in `config.json` you can run the mailservice:

`$ ./run.sh`
