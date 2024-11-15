#!/bin/bash
goctl api plugin -plugin goctl-swagger="swagger -filename user_api.json -host 0.0.0.0:8888 -basepath / -schemes http,ws" -api user_api.api -dir ./doc