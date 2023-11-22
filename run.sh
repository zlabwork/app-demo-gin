#!/bin/bash

env=".env"
cfg="config/app.yaml"

if [[ ! -f $env ]]; then
  echo "$env is not exits"
  exit
fi

if [[ ! -f $cfg ]]; then
  echo "$cfg is not exits"
  exit
fi

export $(cat $env | grep -v "#")

go run main.go
