#!/bin/bash

# Tests that all source files run in go

DIRS=`ls -d */`
TOTAL=`find | grep .go | wc | awk ' { print $1 } '`
COUNT=0
IMG_URL="https:\/\/img.shields.io\/badge\/Go%20Examples%20Working-"
COLOR="red"


git config --global user.email "data@altamontco.com"
git config --global user.name "Automatic Linting"

for dir in $DIRS; do
    echo "TOTAL FOUND: $TOTAL"
    for fname in $dir/*.go; do
        echo "Starting $fname..."
        go run $fname 
        if [ $? -gt 0 ]; then
            echo "FAIL - $fname"
            MULTIPLY=$((COUNT * 100))
            PERCENTAGE=$((MULTIPLY/TOTAL))

            echo "Completed Percentage: $PERCENTAGE"

            if [ $PERCENTAGE -gt 50 ]; then
                echo "More than 50% completed - orange"
                $COLOR="orange"
            fi

            if [ $PERCENTAGE -gt 75 ]; then
                echo "More than 75% completed - yellow"
                $COLOR="yellow"
            fi
            # If we've failed, we need to update README properly

            sed "s/^\!\[Go\].*$/![Go]($IMG_URL$COUNT%2F$TOTAL-$COLOR)/g" README.md > temp.md

            mv temp.md README.md

            git add .
            git commit -m "docs: automatic update to badge"
            git push

            exit 0;
        fi
        COUNT=`expr $COUNT + 1`
    done
done

# If we're here, we need to update README properly

sed "s/^\!\[Go\].*$/![Go]($IMG_URL$COUNT%2F$COUNT-green)/g" README.md > temp.md

mv temp.md README.md


git add .
git commit -m "docs: automatic update to badge"
git push



