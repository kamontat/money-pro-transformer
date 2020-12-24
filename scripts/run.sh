#!/usr/bin/env bash

###########################
## Setup
###########################

export TMP_DIRECTORY="$PWD"
cd "$(dirname "$0")/.." || exit 2

export __ROOT_DIRECTORY="$PWD"
export __COMMANDLINE_DIRECTORY="${__ROOT_DIRECTORY}/modules/cli"
export __APPLICATION_NAME="commandline"

# Build application
cd "${__COMMANDLINE_DIRECTORY}" || exit 2
go build || exit 4

# Run built application
cd "${__ROOT_DIRECTORY}" || exit 2
"${__COMMANDLINE_DIRECTORY}/${__APPLICATION_NAME}" -rootDir="${__ROOT_DIRECTORY}" "$@" || exit $?

###########################
## Cleanup
###########################

rm "${__COMMANDLINE_DIRECTORY}/${__APPLICATION_NAME}"

cd "$TMP_DIRECTORY" || exit 3
unset TMP_DIRECTORY
