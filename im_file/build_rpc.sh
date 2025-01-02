#!/bin/bash
goctl rpc protoc ./file_rpc.proto --go_out=./file_rpc/types --go-grpc_out=./file_rpc/types --zrpc_out=./file_rpc/