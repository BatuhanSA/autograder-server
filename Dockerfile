FROM golang:1.22.3

WORKDIR /autograder-server

COPY  . .

RUN ./scripts/build.sh 

# whats a better chmod code
RUN chmod 777 run_in_Docker.sh

# I tried CMD ENTRYPOINT worked better for me
# I probaly missed something with CMD 
ENTRYPOINT ["./run_in_Docker.sh"] 
