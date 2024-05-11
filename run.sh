#!/bin/bash

env=".env"
cfg="config/app.yaml"

if [[ ! -f "config" ]]; then
  mkdir -p "config"
fi

if [[ ! -f "storage/logs" ]]; then
  mkdir -p "storage/logs"
fi

if [[ ! -f "storage/data" ]]; then
  mkdir -p "storage/data"
fi

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
