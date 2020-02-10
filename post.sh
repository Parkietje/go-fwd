#!/bin/bash

#post an html form with example values
curl -X POST -k -H "Content-Type: application/x-www-form-urlencoded" -d "name=John Doe&email=example.org.com&message=Hey, this is a curl POST" http://localhost:8080