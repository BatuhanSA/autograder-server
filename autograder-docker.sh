#!bin/bash 

readonly DEFAULT_PORT=8080
readonly DEFAULT_CMD=bash

# Usage message
usage() {
    echo "Usage: $0 [-p port_number] | [-x command] | [-h]"
    echo
    echo "Options:"
    echo "  -x            Specify the cmd that you want to run in the autograder-server (default is server)"
    echo "  -p PORT       Specify the port number (default is $DEFAULT_PORT)"
    echo "  -h            Show this help message"
    exit 1
}

function main() {

    local server_port=$DEFAULT_PORT
    local executable=$DEFAULT_CMD

    OPTSTRING=":p:x:h"

    while getopts ${OPTSTRING} opt; do
        case ${opt} in
            p)
            server_port=${OPTARG}
            ;;
            x)
            executable=${OPTARG}
            ;;
            h)
            usage
            ;;
            ?)
            usage
            exit 1
            ;;
        esac
    done

    # Exit immediately if a command exits with a non-zero status.
    set -e
    # If the script receives CTRL - c terminates
    trap exit SIGINT
    mkdir -p /tmp/autograder-temp

    docker build . -t autograder 
    docker run -it --rm --name DOOD-parent -p $server_port:$server_port -v /var/run/docker.sock:/var/run/docker.sock -v /tmp/autograder-temp:/tmp/autograder-temp autograder ${executable} -c web.port=${server_port}

}

[[ "${BASH_SOURCE[0]}" == "${0}" ]] && main "$@"
