# Run docker run -ti --rm -v /var/run/docker.sock:/var/run/docker.sock docker /bin/ash
# Start from Alpine Linux
FROM alpine:latest

EXPOSE 80
# Install necessary packages
RUN apk update && \
    apk add --no-cache \
        docker \
        go

# Define environment variables for Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

# Optionally, set up your Go workspace directory
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Display Go version and check Docker installation
RUN go version && \
    docker --version

RUN apk add bash

WORKDIR /autograder-server

COPY . .

RUN ./scripts/build.sh 

RUN chmod 777 run_in_Docker.sh

# I tried CMD ENTRYPOINT worked better for me
# I probaly missed something with CMD 
ENTRYPOINT ["./run_in_Docker.sh"] 