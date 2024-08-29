FROM golang:1.22.6-alpine3.20

RUN apk update
RUN apk add --no-cache bash git

WORKDIR /autograder-server

# TODO(Batu): Rework on the way that we do version.
COPY ./.git/HEAD  .git/
COPY ./.git/refs  .git/refs/ 
COPY ./.git/index .git/
COPY ./.git/objects .git/objects/
# TODO(Batu): Need to edit cmd/version/main.go so it looks for a VERSION.json that has info.
# RUN go run cmd/version/main.go --json > VERSION.json
# RUN rm -rf ./.git
# RUN apk del git 

COPY ./cmd/ cmd/
COPY ./internal/ internal/
COPY ./scripts/ scripts/
COPY ./testdata/ testdata/
COPY ./go.* .
COPY ./LICENSE .
COPY ./VERSION.txt .

COPY ./docker/entrypoint.sh .
RUN chmod 0775 entrypoint.sh

RUN ./scripts/build.sh 
RUN AUTOGRADER__DOCKER__DISABLE='true' ./scripts/run_tests.sh

RUN rm -rf ./bin

ENTRYPOINT ["./entrypoint.sh"]