#!/bin/bash 
# rightnow this is only for /bin 
# add for go run cmd/..../main.go
executable=$1 
cd bin/
shift 1
./$executable "$@"
