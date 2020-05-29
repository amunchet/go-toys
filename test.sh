#!/bin/sh

# Tests that all source files run in go

DIRS=`ls -d */`


for dir in $DIRS; do
    for fname in $dir/*.go; do
        echo "Starting $fname..."
        go run $fname 
        if [ $? -gt 0 ]; then
            echo "FAIL - $fname"
            exit 1;
        fi
    done
done

