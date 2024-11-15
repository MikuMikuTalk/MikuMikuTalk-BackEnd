#!/bin/bash

goctl api plugin -plugin goctl-swagger="swagger -filename user_api.json" -api user_api.api -dir ./doc