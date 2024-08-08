#!bin/bash 

# Default port value
default_port=8080

# assigns default port number if a command line argument is not given 
port=${1:-$default_port}


mkdir -p /tmp/autograder-temp

docker build . -t autograder 

docker run -it --rm --name DOOD-parent -p $default_port:$default_port -v /var/run/docker.sock:/var/run/docker.sock -v /tmp/autograder-temp:/tmp/autograder-temp autograder bash