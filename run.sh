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
    export $(cat $env_file | grep -v "#")

    # overwrite env
    if [ -z "$APP_ENV" ]; then
        echo "export .env file as default"
    elif [[ "$APP_ENV" != prod* ]]; then
        env_file=".env.$APP_ENV"
        if [[ -f $env_file ]]; then
          echo "export $env_file file"
          export $(cat $env_file | grep -v "#")
        fi
    fi
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
