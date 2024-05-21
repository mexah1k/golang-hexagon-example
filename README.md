# Golang Hexagon Service Example + Kafka + MongoDB + Echo server

## Introduction
How to run the service:

1. start mongoDB, kafka and zookeeper
```bash
docker-compose -f docker-compose.infra.yaml up
```

2. start the service
```bash
make start
```

## Description

URL Analysis Service that calculate the number of meta tags in a given URL.
