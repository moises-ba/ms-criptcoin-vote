package main

import (
	"moises-ba/ms-criptcoin-vote/config"
	pb "moises-ba/ms-criptcoin-vote/criptcoinvote"
	"moises-ba/ms-criptcoin-vote/log"
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
	mongoVoterDB := mongoClient.Database(utils.GetEnv(config.MONGO_QRCODE_BD, "criptcoinVotesDB"))

	//repositories
	voterRepository := repository.NewVoterMongoRepository(mongoVoterDB)

	//services
	voterService := service.NewService(voterRepository)

	//controllers
	voterController := controller.NewVoteController(voterService)

	//iniciando o servicos
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Logger().Fatalf("Falha ao escutar porta: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterCriptCoinVoterServer(s, voterController)

	log.Logger().Printf("Servidor escutando em %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Logger().Fatalf("Falha ao iniciar o servidor: %v", err)
	}

}
