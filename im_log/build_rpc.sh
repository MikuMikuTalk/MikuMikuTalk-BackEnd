#!/bin/bash
goctl rpc protoc ./chat_rpc/chat_rpc.proto --go_out=./chat_rpc/types --go-grpc_out=./chat_rpc/types --zrpc_out=./chat_rpc/