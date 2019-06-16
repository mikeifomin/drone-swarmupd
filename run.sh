#!/bin/bash

export DRONE_TOKEN=tst
export DRONE_URL=http://localhost:3097
export DRONE_SERVICE_NAME=test
export DRONE_NEW_TAG=dev

go run main.go
