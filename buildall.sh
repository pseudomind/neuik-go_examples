#!/usr/bin/env bash

exampleBin=''

if [ ! -d 'src/00-helloWorld' ]; then
	echo 'This script should only be run in the `neuik-go_examples` directory;'\
		'Aborting.'
	exit
fi

#------------------------------------------------------------------------------#
# Create a new bin directory (if there isn't one already)                      #
#------------------------------------------------------------------------------#
if [ ! -d 'bin' ]; then
	mkdir bin
fi
exampleBin="`pwd`/bin"

#------------------------------------------------------------------------------#
# Build and move over all of the individual example programs.                  #
#------------------------------------------------------------------------------#
cd src

ls -1 | while read example; do
	cd $example
	go build
	if [ -f "$example" ]; then
		mv $example $exampleBin
	fi
	cd ..
done

