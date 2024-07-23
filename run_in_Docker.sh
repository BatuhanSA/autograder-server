#!/bin/bash 
executable=$1 
cd bin/
shift 1
./$executable "$@"
