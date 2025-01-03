#!/bin/bash
goctl api plugin -plugin goctl-swagger="swagger -filename file_api.json -host 0.0.0.0:8888 -basepath / -schemes http,ws" -api ../file_api.api -dir ./doc