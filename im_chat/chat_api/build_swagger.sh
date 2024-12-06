#!/bin/bash

goctl api plugin -plugin goctl-swagger="swagger -filename chat_api.json" -api chat_api.api -dir ./doc