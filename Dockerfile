# docker build . -t test
# docker run --privileged --name dind --rm test    
# docker exec -it dind bash
FROM docker:27.0.3-dind-alpine3.20

RUN apk update && apk add --no-cache bash go openjdk8 g++ python3 py3-pip

WORKDIR /autograder-server

COPY cmd/ cmd/
COPY internal/ internal/
COPY scripts/ scripts/
COPY testdata/ testdata/
COPY go.* .
COPY LICENSE .
# COPY run_in_Docker .
COPY VERSION.txt .

RUN ./scripts/build.sh 

# # whats a better chmod code
# RUN chmod 777 run_in_Docker.sh

# # I tried CMD ENTRYPOINT worked better for me
# # I probaly missed something with CMD 
# ENTRYPOINT ["./run_in_Docker.sh"] 
