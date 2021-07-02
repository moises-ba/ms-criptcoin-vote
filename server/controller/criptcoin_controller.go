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

func (c *criptCoinController) List(context.Context, *pb.EmptyParameter) (*pb.CriptCoinList, error) {

	coins, err := c.criptCoinService.List()
	if err != nil {
		return nil, err
	}

	return generateListReturn(coins), nil
}

func (c *criptCoinController) ListWithTotalVotes(context.Context, *pb.EmptyParameter) (*pb.CriptCoinList, error) {
	coins, err := c.criptCoinService.ListWithTotalVotes()
	if err != nil {
		return nil, err
	}

	return generateListReturn(coins), nil
}
func (c *criptCoinController) Find(context.Context, *pb.CriptCoinFilter) (*pb.CriptCoin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (c *criptCoinController) Insert(context.Context, *pb.CriptCoin) (*pb.CriptCoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (c *criptCoinController) Update(context.Context, *pb.CriptCoin) (*pb.CriptCoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (c *criptCoinController) Delete(context.Context, *pb.CriptCoin) (*pb.CriptCoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

//convert coin ou coinTotalizr em um pb.CriptCoin do grpc
func convert(coin model.Coin) *pb.CriptCoin {

	criptCoin := &pb.CriptCoin{Id: coin.Id,
		Name:                  coin.Name,
		Description:           coin.Description,
		TotalApprovedVotes:    uint32(coin.TotalApprovedVotes),
		TotalDisapprovedVotes: uint32(coin.TotalDisapprovedVotes),
	}

	return criptCoin
}

//gera lista de retorno da listagem de moedas
func generateListReturn(coins []model.Coin) *pb.CriptCoinList {
	criptCoinList := new(pb.CriptCoinList)
	criptCoinList.Items = make([]*pb.CriptCoin, 0, len(coins))
	for i, coin := range coins {
		criptCoinList.Items[i] = convert(coin)
	}

	return criptCoinList
}
