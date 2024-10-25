#!/bin/bash

goctl api plugin -plugin goctl-swagger="swagger -filename auth.json" -api auth_api.api -dir ./doc