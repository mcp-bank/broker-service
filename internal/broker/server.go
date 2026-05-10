package broker

import (
	"context"

	"github.com/brianvoe/gofakeit"
	"github.com/mcp-bank/proto/gen/brokerv1"
)

type Service struct {
	brokerv1.UnimplementedBrokerServiceServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetPortfolio(ctx context.Context, request *brokerv1.GetPortfolioRequest) (*brokerv1.GetPortfolioResponse, error) {
	types := []string{"stock", "bond", "futures", "cash", "option"}
	numberOfPositions := gofakeit.Number(0, 20)
	positions := make([]*brokerv1.PortfolioPosition, numberOfPositions)
	for i := range numberOfPositions {
		positions[i] = &brokerv1.PortfolioPosition{
			Name:     gofakeit.Name(),
			Type:     types[gofakeit.Number(0, 4)],
			Quantity: gofakeit.Float64Range(1, 1000000),
			Price: &brokerv1.Money{
				Amount:   int64(gofakeit.Price(0.0001, 1000000)),
				Currency: gofakeit.CurrencyShort(),
			},
		}
	}
	return &brokerv1.GetPortfolioResponse{Positions: positions}, nil
}

func (s *Service) GetStockPrice(ctx context.Context, request *brokerv1.GetStockPriceRequest) (*brokerv1.GetStockPriceResponse, error) {
	return &brokerv1.GetStockPriceResponse{Price: &brokerv1.Money{
		Amount:   int64(gofakeit.Price(0.0001, 1000000)),
		Currency: gofakeit.CurrencyShort(),
	}}, nil
}

func (s *Service) GetAccountBalance(ctx context.Context, request *brokerv1.GetAccountBalanceRequest) (*brokerv1.GetAccountBalanceResponse, error) {
	return &brokerv1.GetAccountBalanceResponse{Balance: &brokerv1.Money{
		Amount:   int64(gofakeit.Price(0.0001, 1000000)),
		Currency: gofakeit.CurrencyShort(),
	}}, nil
}
