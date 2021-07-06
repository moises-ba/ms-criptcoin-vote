package main

import (
	"moises-ba/ms-criptcoin-vote/config"
	pb "moises-ba/ms-criptcoin-vote/criptcoinvote"
	"moises-ba/ms-criptcoin-vote/log"
	"moises-ba/ms-criptcoin-vote/messaging"
	"moises-ba/ms-criptcoin-vote/repository"
	"moises-ba/ms-criptcoin-vote/server/controller"
	"moises-ba/ms-criptcoin-vote/service"
	"moises-ba/ms-criptcoin-vote/utils"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":9090"
)

func main() {

	//inicializando a conexão com o mongo
	mongoClient, funcDisconnect, err := repository.ConnectMongo()
	if err != nil {
		log.Logger().Fatal("Falha ao conectar no mongo.", err)
	}
	defer funcDisconnect()

	//databases
	mongoCriptcoinDB := mongoClient.Database(utils.GetEnv(config.MONGO_QRCODE_BD, "criptcoinDB"))

	//repositories
	voterRepository := repository.NewVoterMongoRepository(mongoCriptcoinDB)
	criptCoinRepository := repository.NewCriptCoinMongoRepository(mongoCriptcoinDB)

	//services
	voterService := service.NewService(voterRepository, messaging.NewKafkaProducer())
	criptCoinService := service.NewCriptCoinService(criptCoinRepository)

	//controllers
	voterController := controller.NewVoteController(voterService)
	criptCoinController := controller.NewCriptCoinController(criptCoinService)

	//iniciando o servicos
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Logger().Fatalf("Falha ao escutar porta: %v", err)
	}
	s := grpc.NewServer()

	//registrando controllers grpc
	pb.RegisterCriptCoinVoterApiServer(s, voterController)
	pb.RegisterCriptCoinApiServer(s, criptCoinController)

	log.Logger().Printf("Servidor escutando em %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Logger().Fatalf("Falha ao iniciar o servidor: %v", err)
	}

}
