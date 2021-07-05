package controller

import (
	"context"
	pb "moises-ba/ms-criptcoin-vote/criptcoinvote"
	"moises-ba/ms-criptcoin-vote/model"
	"moises-ba/ms-criptcoin-vote/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type criptCoinController struct {
	pb.UnimplementedCriptCoinApiServer
	criptCoinService service.CriptCoinService
}

func NewCriptCoinController(pCriptcoinService service.CriptCoinService) *criptCoinController {
	return &criptCoinController{criptCoinService: pCriptcoinService}
}

func (c *criptCoinController) List(context context.Context, emptyParameter *pb.EmptyParameter) (*pb.CriptCoinList, error) {

	coins, err := c.criptCoinService.List()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return generateListReturn(coins), nil
}

func (c *criptCoinController) ListWithTotalVotes(context context.Context, emptyParameter *pb.EmptyParameter) (*pb.CriptCoinList, error) {
	coins, err := c.criptCoinService.ListWithTotalVotes()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return generateListReturn(coins), nil
}
func (c *criptCoinController) Find(context context.Context, filter *pb.CriptCoinFilter) (*pb.CriptCoin, error) {

	coin, err := c.criptCoinService.Find(filter.CoinId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if coin == nil {
		return nil, status.Errorf(codes.NotFound, "Moeda nao encontrada")
	}

	return convert(coin), nil
}

func (c *criptCoinController) Insert(context context.Context, criptCoin *pb.CriptCoin) (*pb.CriptCoinReply, error) {

	err := c.criptCoinService.Insert(convertToCoin(criptCoin))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.CriptCoinReply{Message: "Moeda criada com sucesso"}, nil
}

func (c *criptCoinController) Update(context context.Context, criptCoin *pb.CriptCoin) (*pb.CriptCoinReply, error) {

	err := c.criptCoinService.Update(convertToCoin(criptCoin))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.CriptCoinReply{Message: "Moeda atualizada com sucesso"}, nil
}

func (c *criptCoinController) Delete(context context.Context, criptCoin *pb.CriptCoin) (*pb.CriptCoinReply, error) {
	err := c.criptCoinService.Delete(convertToCoin(criptCoin))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.CriptCoinReply{Message: "Moeda excluida com sucesso"}, nil
}

func convertToCoin(criptCoin *pb.CriptCoin) model.Coin {

	return model.Coin{
		Id:          criptCoin.Id,
		Name:        criptCoin.Name,
		Description: criptCoin.Description,
	}

}

//convert coin ou coinTotalizr em um pb.CriptCoin do grpc
func convert(coin *model.Coin) *pb.CriptCoin {

	criptCoin := &pb.CriptCoin{Id: coin.Id,
		Name:                  coin.Name,
		Description:           coin.Description,
		TotalApprovedVotes:    uint32(coin.TotalApprovedVotes()),
		TotalDisapprovedVotes: uint32(coin.TotalDisapprovedVotes()),
	}

	return criptCoin
}

//gera lista de retorno da listagem de moedas
func generateListReturn(coins []*model.Coin) *pb.CriptCoinList {
	criptCoinList := new(pb.CriptCoinList)
	criptCoinList.Items = make([]*pb.CriptCoin, 0, len(coins))
	for i, coin := range coins {
		criptCoinList.Items[i] = convert(coin)
	}

	return criptCoinList
}
