# ms-criptcoin-vote
<h1>Microserviço que expõe endpoints de votos para as criptomoedas conhecidas.</h1>



<h3>Para a solução do microserviço foram utlizados</h3>
* MongoDB para a base de dados das moedas e votos mongo-driver
* Kafka para visualização em tempo real dos votos nas moedas kafka-go
* JWT para segurança dos microserviços jwt-go
* logrus para logging
* Docker e docker-compose para conteinerização


para rodar a aplicação serão necessarios os passo:
1. docker-compose -f kafka-docker-compose.yml -up -d
2. docker-compose -f mongodb-docker-compose.yml -up -d
3. go build && ./ms-criptcoin-vote

ou dockerizar a aplicação
1. go build
2. docker build -t moisesba/ms-criptcoin-vote .
3. docker-compose -f ms-criptcoin-vote.yml -up -d


* somente para desenvolvimento
gerar serviços grpc
cd criptcoinvote && protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    criptcoinvote.proto && cd ..


cd criptcoinvote && protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    criptcoin.proto && cd ..

 