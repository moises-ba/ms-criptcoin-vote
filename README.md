# ms-criptcoin-vote
Microserviço que expõe endpoints de votos para as criptomoedas conhecidas.


gerar serviços grpc

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    criptcoinvote.proto
