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

func (s *Service) GetNews(ctx context.Context, request *brokerv1.GetNewsRequest) (*brokerv1.GetNewsResponse, error) {
	numberOfItems := gofakeit.Number(0, 20)
	items := make([]*brokerv1.NewsItem, numberOfItems)
	for i := range numberOfItems {
		items[i] = &brokerv1.NewsItem{
			Title:   gofakeit.Word(),
			Summary: gofakeit.HipsterParagraph(gofakeit.Number(1, 5), gofakeit.Number(5, 50), gofakeit.Number(50, 500), "\n"),
			Url:     gofakeit.URL(),
		}
	}
	return &brokerv1.GetNewsResponse{Items: items}, nil
}

func (s *Service) GetNews2(ctx context.Context, request *brokerv1.GetNews2Request) (*brokerv1.GetNews2Response, error) {
	numberOfItems := gofakeit.Number(0, 20)
	items := make([]*brokerv1.News2Item, numberOfItems)
	for i := range numberOfItems {
		items[i] = &brokerv1.News2Item{
			Title:   gofakeit.Word(),
			Summary: gofakeit.HipsterParagraph(gofakeit.Number(1, 5), gofakeit.Number(5, 50), gofakeit.Number(50, 500), "\n"),
			Url:     gofakeit.URL(),
		}
	}
	return &brokerv1.GetNews2Response{Items: items}, nil
}
