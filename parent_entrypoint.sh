#!/bin/bash 
# rightnow this is only for /bin 
# add for go run cmd/..../main.go
readonly ROOT_DIR="$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)"
readonly CMD_DIR="${ROOT_DIR}/cmd"

function main(){
    executable=$1 

    echo "Executable: $executable"
    if [ -d "/autograder-server/cmd/${executable}" ]; then
        echo $CMD_DIR
        # go run cmd/version/main.go
    else
        # make a better error message
        echo Invalid command
        exit 1
    fi
    shift 1
    go run ${CMD_DIR}/${executable}/main.go  "$@"
}

[[ "${BASH_SOURCE[0]}" == "${0}" ]] && main "$@"


