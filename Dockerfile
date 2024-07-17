FROM docker:dind

RUN apk update && apk add --no-cache go bash openjdk8 g++ python3 py3-pip


WORKDIR /autograder-server

COPY  . .

RUN ./scripts/build.sh 

