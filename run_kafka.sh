#!/bin/bash
docker network create app-tier --driver bridge

docker run -d --restart=always --name zookeeper-server --network app-tier -e ALLOW_ANONYMOUS_LOGIN=yes bitnami/zookeeper:latest


docker run -d --restart=always  --name kafka-server  --network app-tier  -p 9092:9092 -e ALLOW_PLAINTEXT_LISTENER=yes  -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper-server:2181 -e KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://192.168.3.242:9092  bitnami/kafka:latest


docker run -d --restart=always  --name kafka-map --network app-tier  -p 9001:8080 -v /opt/kafka-map/data:/usr/local/kafka-map/data  -e DEFAULT_USERNAME=admin  -e DEFAULT_PASSWORD=admin --restart always dushixiang/kafka-map:latest
