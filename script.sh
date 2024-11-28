#!/bin/bash

docker image build -f Dockerfile -t aawimage .

docker container run -p 8080:8080 --detach --name aawcontainer aawimage
