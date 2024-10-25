#!/bin/bash
goctl api plugin -plugin goctl-swagger="swagger -filename auth.json -host 0.0.0.0:8888 -basepath / -schemes http,ws" -api auth_api.api -dir ./doc