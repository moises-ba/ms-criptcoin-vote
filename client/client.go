package main

import (
	"context"
	"io"
	"moises-ba/ms-criptcoin-vote/config"
	pb "moises-ba/ms-criptcoin-vote/criptcoinvote"
	"moises-ba/ms-criptcoin-vote/log"
	"moises-ba/ms-criptcoin-vote/security"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

	//gerando token
	jwtManager := security.NewJWTManager(config.GetJWTPassword(), 5*time.Minute)
	user := security.User{Username: "moises", Role: "ADMIN"}
	token, err := jwtManager.Generate(&user)
	if err != nil {
		log.Logger().Fatalf("Falha ao gerar token %v", err)
	}

	//adicionando o token no contexto
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization_jwt_token", token)

	clientCiptcoin := pb.NewCriptCoinVoterApiClient(conn)
	reply, err := clientCiptcoin.Vote(ctx, &pb.VoteRequest{CoinId: coinId, Approved: like})
	if err != nil {
		log.Logger().Fatalf("Falha ao votar na moeda  %v", err)
	}

	log.Logger().Infoln("Voto computado: " + reply.Message)

}

func streamVotes(conn *grpc.ClientConn) {
	//gerando token
	jwtManager := security.NewJWTManager(config.GetJWTPassword(), 5*time.Minute)
	user := security.User{Username: "Moises Almeida", Role: "ADMIN"}
	token, err := jwtManager.Generate(&user)
	if err != nil {
		log.Logger().Fatalf("Falha ao gerar token %v", err)
	}

	//adicionando o token no contexto
	ctx := context.Background()
	metadata.AppendToOutgoingContext(ctx, "authorization_jwt_token", token)

	clientCiptcoin := pb.NewCriptCoinVoterApiClient(conn)
	stream, err := clientCiptcoin.FetchVoteStream(ctx, &pb.EmptyParameterVote{})
	if err != nil {
		log.Logger().Fatalf("Falha ao votar na moeda  %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Logger().Fatalf("não pode receber %v", err)
			}
			log.Logger().Printf("Mensagem recebida: coin=%s like=%d dislike=%d",
				resp.CoinId,
				resp.TotalApprovedVotes,
				resp.TotalDisapprovedVotes)
		}
	}()

	<-done //we will wait until all response is received
	log.Logger().Printf("finished")
}

func main() {

	// dail server
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Logger().Fatalf("Impossivel conectar no servidor %v", err)
	}

	//listarMoedasTotalizandoVotos(conn)

	//findMoeda(conn, "eos")

	//inserirMoeda(conn, "eos", "Eos", "Moeda eos")

	//updateMoeda(conn, "btc", "Bitcoin", "Moeda bitcoin Alterado")

	//deleteMoeda(conn, "btc", "Bitcoin", "Moeda bitcoin Alterado")

	go func() {
		like := true
		for {

			likeUnlikeCoin(conn, "eos", like)
			like = !like

			time.Sleep(5 * time.Second)
		}

	}()

	streamVotes(conn)

}
