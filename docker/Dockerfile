FROM golang:1.22.6-alpine3.20 AS build-stage

RUN apk update
RUN apk add --no-cache bash git

WORKDIR /autograder-server

COPY . ./full-copy

# Creat a full copy to a different directory 

COPY ./cmd/ cmd/
COPY ./internal/ internal/
COPY ./scripts/ scripts/
COPY ./testdata/ testdata/
COPY ./go.* ./
COPY ./LICENSE ./
COPY ./VERSION.txt ./

COPY ./docker/entrypoint.sh ./
RUN chmod 0775 entrypoint.sh

RUN cd full-copy && ./scripts/build.sh 
RUN cd full-copy && AUTOGRADER__DOCKER__DISABLE='true' ./scripts/run_tests.sh
RUN cd full-copy && go run ./cmd/version/main.go --out=VERSION.json && mv VERSION.json ..

RUN rm -rf ./full-copy
RUN apk del git

FROM scratch AS final-stage

ENV GOLANG_VERSION=1.22.6
ENV GOTOOLCHAIN=local
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
COPY --from=build-stage / /

WORKDIR /autograder-server

ENTRYPOINT ["./entrypoint.sh"]
