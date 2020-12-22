#!/usr/bin/env bash

###########################
## Configuration
###########################

export BUILDERS=(
  "darwin/amd64"
  "linux/amd64"
  "windows/amd64"
)

###########################
## Setup
###########################

export TMP_DIRECTORY="$PWD"
cd "$(dirname "$0")/.." || exit 2

export __ROOT_DIRECTORY="$PWD"
export __COMMANDLINE_DIRECTORY="${__ROOT_DIRECTORY}/modules/cli"
export __RESULT_EXECUTION_DIRECTORY="${__ROOT_DIRECTORY}/resources/execution"
export __APPLICATION_NAME="money-pro-transformer"

###########################
## Builder
###########################

cd "${__COMMANDLINE_DIRECTORY}" || exit 2

for builder in "${BUILDERS[@]}"; do
  os="${builder%%/*}"
  arch="${builder##*/}"
  ext=""
  [[ $os == "windows" ]] && ext=".exe"

  export GOOS="$os"
  export GOARCH="$arch"

  result_directory="${__RESULT_EXECUTION_DIRECTORY}/${os}/${__APPLICATION_NAME}${ext}"

  if go build -o "$result_directory"; then
    echo "Built ${os} completed (${result_directory})"
  else
    echo "Built ${os} failure" >&2
  fi
done

cd "$TMP_DIRECTORY" || exit 2
unset TMP_DIRECTORY
