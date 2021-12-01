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

 ------------------------------------------ com swagger e proxy reverso
 https://blog.csnet.me/blog/building-a-go-api-grpc-rest-and-openapi-swagger.1/

fonte: https://github.com/grpc-ecosystem/grpc-gateway


gRPC:
  go get google.golang.org/grpc

Go protoc plugin:
  go install github.com/golang/protobuf/protoc-gen-go@latest

grpc-gateway
  go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
  go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
  go install github.com/golang/protobuf/protoc-gen-go@latest  


export PATH=$PATH:~/go/bin/ && protoc -I. -I~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --grpc-gateway_out ./pb --swagger_out=:swagger   pb/device.proto



-------------------------------
https://blog.csnet.me/blog/building-a-go-api-grpc-rest-and-openapi-swagger.1/

fonte: https://github.com/grpc-ecosystem/grpc-gateway


gRPC:
  go get google.golang.org/grpc

Go protoc plugin:
  go install github.com/golang/protobuf/protoc-gen-go@latest

grpc-gateway
  go installgithub.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
  go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
  go install github.com/golang/protobuf/protoc-gen-go@latest  


export GOPATH=/home/moises/go
export PATH := $(GOPATH)/bin:$(PATH)
protoc -I. -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --grpc-gateway_out ./ --swagger_out=:swagger   criptcoinvote/*.proto
