#!/bin/bash 

python3 -m venv env
source env/bin/activate
pip3 install autograder-py

# bash scripts/run_tests.sh > output_DooD.txt 2>&1

# go test ./internal/grader > output_DooD.txt 2>&1
