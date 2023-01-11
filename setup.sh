#!/bin/bash

for i in {1..25}
do
    if [ ! -d $i ] 
    then
        # Create exercise dir
        mkdir $i
        touch $i/README.md

        # Create empty input files
        touch $i/input.txt
        touch $i/input_final.txt

        # Set up go module
        cat << EOF > $i/go.mod
module github.com/noahssarcastic/advent2022/$i

go 1.19
EOF
        cat << EOF > $i/main.go
package main

func main() {

}
EOF
        go work use ./$i
    fi
done
