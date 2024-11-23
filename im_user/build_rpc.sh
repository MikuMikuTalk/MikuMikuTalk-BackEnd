#!/bin/bash
goctl rpc protoc ./user_rpc/user_rpc.proto --go_out=./user_rpc/types --go-grpc_out=./user_rpc/types --zrpc_out=./user_rpc/