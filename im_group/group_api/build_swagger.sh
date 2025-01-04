#!/bin/bash

goctl api plugin -plugin goctl-swagger="swagger -filename group_api.json" -api group_api.api -dir ./doc