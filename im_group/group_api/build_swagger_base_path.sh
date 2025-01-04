#!/bin/bash
goctl api plugin -plugin goctl-swagger="swagger -filename group_api.json -host 0.0.0.0:8888 -basepath / -schemes http,ws" -api group_api.api -dir ./doc