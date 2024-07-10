FROM docker:dind

RUN apk update && apk add --no-cache go bash


WORKDIR /autograder-server

COPY  . .

RUN ./scripts/build.sh 

