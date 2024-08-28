FROM golang:1.22.6-alpine3.20

RUN apk update
RUN apk add --no-cache bash git python3 

WORKDIR /autograder-server

RUN python3 -m venv env
RUN source env/bin/activate && pip3 install autograder-py

COPY ./.git/HEAD  .git/
COPY ./.git/refs  .git/refs/ 
COPY ./.git/index .git/
COPY ./.git/objects .git/objects/

COPY ./cmd/ cmd/
COPY ./internal/ internal/
COPY ./scripts/ scripts/
COPY ./testdata/ testdata/
COPY ./go.* .
COPY ./LICENSE .
COPY ./VERSION.txt .

COPY ./docker/entrypoint.sh .
RUN chmod 0775 entrypoint.sh

RUN ./scripts/build.sh && rm -rf ./bin
RUN source env/bin/activate && AUTOGRADER__DOCKER__DISABLE='true' ./scripts/run_tests.sh
RUN rm -rf ./env

ENTRYPOINT ["./entrypoint.sh"]