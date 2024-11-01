#!/bin/bash

env_file=".env"
yaml_file="config/app.yaml"

# build directory
function check_dir() {
    if [[ ! -f "config" ]]; then
      mkdir -p "config"
    fi

    if [[ ! -f "storage/logs" ]]; then
      mkdir -p "storage/logs"
    fi

    if [[ ! -f "storage/data" ]]; then
      mkdir -p "storage/data"
    fi
}

# check environment
function export_env() {
    if [[ ! -f $env_file ]]; then
      echo "$env_file is not exits"
      exit
    fi

    if [[ ! -f $yaml_file ]]; then
      echo "$yaml_file is not exits"
      exit
    fi

    # export env
    echo "export $env_file"
    export $(cat $env_file | grep -v "#")

    # overwrite env
    env_local=".env.local"
    if [[ -f $env_local ]]; then
      echo "export $env_local"
      export $(cat $env_local | grep -v "#")
    fi
}

# TODO:: .env.prod .env.test .env.dev
function export_env_files() {
    for file in .env.*; do
      if [[ -f "$file" ]]; then
        echo "Reading file: $file"
      fi
    done
}

function run_app() {
    go run main.go
}

function main() {
    check_dir
    export_env
    run_app
}

main
