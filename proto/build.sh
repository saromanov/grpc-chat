#!/usr/bin/env bash

set -e

echo "Generating protobuf for gRPC golang"
protoc -I./src --go_out=plugins=grpc:chat chat/chat.proto