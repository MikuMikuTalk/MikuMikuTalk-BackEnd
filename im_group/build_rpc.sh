#!/bin/bash
goctl rpc protoc ./group_rpc.proto --go_out=./group_rpc/types --go-grpc_out=./group_rpc/types --zrpc_out=./group_rpc/