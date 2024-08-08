#!bin/bash 

mkdir -p /tmp/autograder-temp

docker build . -t autograder

docker run -it --rm --name DOOD-parent -v /var/run/docker.sock:/var/run/docker.sock -v /tmp/autograder-temp:/tmp/autograder-temp autograder bash 
