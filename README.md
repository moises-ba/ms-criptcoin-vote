# ms-criptcoin-vote
Microserviço que expõe endpoints de votos para as criptomoedas conhecidas.


gerar serviços grpc
cd criptcoinvote && protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    criptcoinvote.proto && cd ..


cd criptcoinvote && protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    criptcoin.proto && cd ..


containers docker---
kafka

https://hub.docker.com/r/bitnami/kafka/
$ curl -sSL https://raw.githubusercontent.com/bitnami/bitnami-docker-kafka/master/docker-compose.yml > docker-compose.yml
$ docker-compose up -d

mongo.yml
version: '2'

services:

  mongo:
    image: mongo
    restart: always
    ports:
      - 27017:27017
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: example


-----
