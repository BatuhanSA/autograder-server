FROM golang:latest

WORKDIR /autograder-server

COPY  . .

RUN ./scripts/build.sh 



