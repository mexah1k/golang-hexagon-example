# Golang Hexagon Service Example + Kafka + MongoDB + Echo server
##Medium article: https://medium.com/@mike_polo/structuring-a-golang-project-hexagonal-architecture-43b4de480c14

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
