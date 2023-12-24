package mer

import (
	"context"
	"errors"

	"github.com/shopspring/decimal"
	"github.com/winebarrel/mer/internal/yahoo"
)

func ExchangeContext(ctx context.Context, from string, to string, src decimal.Decimal) (decimal.Decimal, error) {
	resp, err := yahoo.Request(context.TODO(), from, to)

	if err != nil {
		return decimal.Zero, err
	}

	crs := resp.Chart.Result

	if len(crs) == 0 {
		return decimal.Zero, errors.New("Chart results not found")
	}

	frmp := crs[0].Meta.RegularMarketPrice
	rmp := decimal.NewFromFloat(frmp)

	return src.Mul(rmp), nil
}

func Exchange(from string, to string, src decimal.Decimal) (decimal.Decimal, error) {
	return ExchangeContext(context.Background(), from, to, src)
}
