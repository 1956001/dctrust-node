package service

import (
	"context"
	"fmt"
	pb "github.com/kvant-node/api/v2/api_pb"
	"github.com/kvant-node/core/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Addresses(_ context.Context, req *pb.AddressesRequest) (*pb.AddressesResponse, error) {
	cState, err := s.getStateForHeight(req.Height)
	if err != nil {
		return new(pb.AddressesResponse), status.Error(codes.NotFound, err.Error())
	}

	cState.Lock()
	defer cState.Unlock()

	response := &pb.AddressesResponse{
		Addresses: make([]*pb.AddressesResponse_Result, 0, len(req.Addresses)),
	}

	for _, address := range req.Addresses {
		addr := types.StringToAddress(address)
		data := &pb.AddressesResponse_Result{
			Address:           address,
			Balance:           make(map[string]string),
			CountTransactions: fmt.Sprintf("%d", cState.Accounts.GetNonce(addr)),
		}

		balances := cState.Accounts.GetBalances(addr)
		for k, v := range balances {
			data.Balance[k.String()] = v.String()
		}

		if _, exists := data.Balance[types.GetBaseCoin().String()]; !exists {
			data.Balance[types.GetBaseCoin().String()] = "0"
		}

		response.Addresses = append(response.Addresses, data)
	}

	return response, nil
}
