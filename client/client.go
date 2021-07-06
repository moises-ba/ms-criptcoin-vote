package main

import (
	"context"
	pb "moises-ba/ms-criptcoin-vote/criptcoinvote"
	"moises-ba/ms-criptcoin-vote/log"
	"strconv"

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

func listarMoedas(conn *grpc.ClientConn) {
	clientCiptcoin := pb.NewCriptCoinApiClient(conn)

	reply, err := clientCiptcoin.List(context.Background(), &pb.EmptyParameter{})
	if err != nil {
		log.Logger().Fatalf("Falha ao Listar %v", err)
	}

	for _, coin := range reply.Items {
		log.Logger().Infoln("Moeda -> " + " " + coin.Id + " " + coin.Name + " " + coin.Description)
	}

}

func findMoeda(conn *grpc.ClientConn, coinId string) {
	clientCiptcoin := pb.NewCriptCoinApiClient(conn)

	reply, err := clientCiptcoin.Find(context.Background(), &pb.CriptCoinFilter{CoinId: coinId})
	if err != nil {
		log.Logger().Fatalf("Falha ao localizar moeda pelo ID %v", err)
	}

	log.Logger().Infoln("Moeda -> " + " " + reply.Id + " " + reply.Name + " " + reply.Description)

}

func listarMoedasTotalizandoVotos(conn *grpc.ClientConn) {
	clientCiptcoin := pb.NewCriptCoinApiClient(conn)

	reply, err := clientCiptcoin.ListWithTotalVotes(context.Background(), &pb.EmptyParameter{})
	if err != nil {
		log.Logger().Fatalf("Falha ao Listar moeda totalizando os votos %v", err)
	}

	for _, coin := range reply.Items {
		log.Logger().Infoln("Moeda -> " + " " + coin.Id + " " + coin.Name + " " + coin.Description +
			" likes: " + strconv.Itoa(int(coin.GetTotalApprovedVotes())) +
			" deslikes: " + strconv.Itoa(int(coin.GetTotalDisapprovedVotes())))
	}

}

func likeUnlikeCoin(conn *grpc.ClientConn, coinId string, like bool) {
	clientCiptcoin := pb.NewCriptCoinVoterApiClient(conn)

	reply, err := clientCiptcoin.Vote(context.Background(), &pb.VoteRequest{CoinId: coinId, Approved: like})
	if err != nil {
		log.Logger().Fatalf("Falha ao votar na moeda  %v", err)
	}

	log.Logger().Infoln("Voto computado: " + reply.Message)

}

func main() {

	// dail server
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Logger().Fatalf("Impossivel conectar no servidor %v", err)
	}

	//listarMoedasTotalizandoVotos(conn)

	findMoeda(conn, "eoss")

	//inserirMoeda(conn, "eos", "Eos", "Moeda eos")

	//updateMoeda(conn, "btc", "Bitcoin", "Moeda bitcoin Alterado")

	//deleteMoeda(conn, "btc", "Bitcoin", "Moeda bitcoin Alterado")

	//likeUnlikeCoin(conn, "eos", true)

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
