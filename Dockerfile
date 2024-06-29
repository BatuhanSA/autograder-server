# Started from Alpine
FROM alpine:latest

EXPOSE 8080

# Install necessary packages
RUN apk update && apk add --no-cache docker go

# install bash 
RUN apk add bash

WORKDIR /autograder-server

COPY . .

RUN ./scripts/build.sh 

# We need to change the permission
RUN chmod 777 run_in_Docker.sh

# I tried CMD ENTRYPOINT worked better for me
# I probaly missed something with CMD 
ENTRYPOINT ["./run_in_Docker.sh"] 