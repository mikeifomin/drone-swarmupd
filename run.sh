#!/bin/bash

export PLUGIN_TOKEN=tst
export PLUGIN_URL=http://localhost:3097
export PLUGIN_SERVICE_NAME=test
export PLUGIN_NEW_TAG=dev

go run main.go
