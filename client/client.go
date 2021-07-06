package main

import (
	"context"
	pb "moises-ba/ms-criptcoin-vote/criptcoinvote"
	"moises-ba/ms-criptcoin-vote/log"

	"google.golang.org/grpc"
)

func inserirMoeda(conn *grpc.ClientConn, id, name, description string) error {

	clientCiptcoin := pb.NewCriptCoinApiClient(conn)
	clientCiptcoinIn := pb.CriptCoin{Id: id, Name: name, Description: description}
	insertReply, err := clientCiptcoin.Insert(context.Background(), &clientCiptcoinIn)
	if err != nil {
		log.Logger().Fatalf("Falha ao inserir %v", err)
	}

	log.Logger().Info("Registro inserido :" + insertReply.Message)
	return nil
}

func updateMoeda(conn *grpc.ClientConn, id, name, description string) error {
	clientCiptcoin := pb.NewCriptCoinApiClient(conn)
	clientCiptcoinIn := pb.CriptCoin{Id: id, Name: name, Description: description}
	insertReply, err := clientCiptcoin.Update(context.Background(), &clientCiptcoinIn)
	if err != nil {
		log.Logger().Fatalf("Falha ao Alterar %v", err)
	}

	log.Logger().Info("Registro Alterado :" + insertReply.Message)
	return nil
}

func deleteMoeda(conn *grpc.ClientConn, id, name, description string) error {
	clientCiptcoin := pb.NewCriptCoinApiClient(conn)
	clientCiptcoinIn := pb.CriptCoin{Id: id, Name: name, Description: description}
	insertReply, err := clientCiptcoin.Delete(context.Background(), &clientCiptcoinIn)
	if err != nil {
		log.Logger().Fatalf("Falha ao Alterar %v", err)
	}

	log.Logger().Info("Registro Alterado :" + insertReply.Message)
	return nil
}

func main() {

	// dail server
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Logger().Fatalf("Impossivel conectar no servidor %v", err)
	}

	//inserirMoeda(conn, "eos", "Eos", "Moeda eos")

	//updateMoeda(conn, "btc", "Bitcoin", "Moeda bitcoin Alterado")

	deleteMoeda(conn, "btc", "Bitcoin", "Moeda bitcoin Alterado")

	/*
		// create stream
		client := pb.NewVoteStreamClient(conn)
		in := &pb.Request{Id: 1}
		stream, err := client.FetchResponse(context.Background(), in)
		if err != nil {
			log.Fatalf("openn stream error %v", err)
		}
	*/

}
