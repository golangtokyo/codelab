#!/bin/bash

declare -a codelabs;
codelabs=(
    "cloud-shell-go-setup" "16Xqaj-2RytzDzbIEgeOmYGQJiaPYlUlGhdZo6vaIej8" 
    "find-gophers"         "1fbpwn8iOLZAO5AYkcbdIDHAM6Dtn1lUwnsUSyNwScP4"
    "go-webassembly"       "12IXvrTgNIMUXotVcGkpqfu3WFhvtvfnzlvQfJIn0h6k"
)

for ((i = 0; i < ${#codelabs[@]}; i+=2)) {
    echo "generate ${codelabs[i]}..."
    echo claat export -prefix="/codelab/" ${codelabs[i+1]}
    claat export -prefix="/codelab/" ${codelabs[i+1]}
}
