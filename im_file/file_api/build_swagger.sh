#!/bin/bash

goctl api plugin -plugin goctl-swagger="swagger -filename file_api.json" -api file_api.api -dir ./doc