#!/usr/bin/env sh

protoc --go_out=plugins=grpc:. $1