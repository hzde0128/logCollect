#!/bin/bash

start(){
  docker run --name zookeeper -d -p 2181:2181 \
    -e ALLOW_ANONYMOUS_LOGIN=yes zookeeper:3.4.14

  docker run --name kafka -d -p 9092:9092 -p 29092:29092 \
    -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181 \
    -e ALLOW_PLAINTEXT_LISTENER=yes \
    -e KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT \
    -e KAFKA_CFG_LISTENERS=PLAINTEXT://:9092,PLAINTEXT_HOST://:29092 \
    -e KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092,PLAINTEXT_HOST://localhost:29092 \
    bitnami/kafka:2.4.0

  docker run --name etcd -d -p 2379:2379 -e ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd:3.4.4
}

stop(){
  for container in zookeeper kafka etcd;do
    docker stop $container
    docker rm $container
  done  
}

case $1 in
start)
  start
  ;;
stop)
  stop
  ;;
restart)
  stop
  start
  ;;
esac

