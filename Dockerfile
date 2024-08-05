# docker build . -t testing
# docker run -it --name DooD-Testing --rm -v /var/run/docker.sock:/var/run/docker.sock testing

FROM golang:1.23rc1-alpine3.20

RUN apk update && apk add --no-cache bash go docker-cli openjdk8 g++ python3-dev py3-pip git libffi-dev


WORKDIR /autograder-server

COPY setup.sh .
RUN bash setup.sh

COPY cmd/ cmd/
COPY internal/ internal/
COPY scripts/ scripts/
COPY testdata/ testdata/
COPY go.* .
COPY LICENSE .
# Added for tracking changes that I do ONLY FOR DEV/TESTING
COPY .git/ .git/
COPY .gitignore .
# This is going to be ultimalty the file that we run in docker
# COPY run_in_Docker .
COPY VERSION.txt .
# COPY test.sh test.sh

#RUN ls -la

# RUN ./scripts/build.sh 

# # whats a better chmod code
# RUN chmod 777 run_in_Docker.sh

# # I tried CMD ENTRYPOINT worked better for me
# # I probaly missed something with CMD 
# ENTRYPOINT ["./run_in_Docker.sh"] 
