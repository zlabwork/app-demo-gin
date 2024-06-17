#!/bin/bash

env_file=".env"
cfg_file="config/app.yaml"

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

    if [[ ! -f $cfg_file ]]; then
      echo "$cfg_file is not exits"
      exit
    fi

    # export env
    export $(cat $env_file | grep -v "#")

    # overwrite env
    if [ -z "$APP_ENV" ]; then
        echo "use .env file as default"
    elif [[ "$APP_ENV" != prod* ]]; then
        env_file=".env.$APP_ENV"
        if [[ ! -f $env_file ]]; then
          echo "$env_file is not exits"
          exit
        fi
        echo "use $env_file file"
        export $(cat $env_file | grep -v "#")
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
