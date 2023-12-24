package yahoo_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/mer/internal/yahoo"
)

func TestRequestOK(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://query1.finance.yahoo.com/v8/finance/chart/EURJPY=X?interval=1d", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusOK, `
			{
				"chart": {
					"result": [
						{
							"meta": {
								"currency": "JPY",
								"symbol": "EURJPY=X",
								"exchangeName": "CCY",
								"instrumentType": "CURRENCY",
								"firstTradeDate": 1043280000,
								"regularMarketTime": 1703284103,
								"gmtoffset": 0,
								"timezone": "GMT",
								"exchangeTimezoneName": "Europe/London",
								"regularMarketPrice": 156.878,
								"chartPreviousClose": 156.481,
								"priceHint": 4,
								"currentTradingPeriod": {
									"pre": {
										"timezone": "GMT",
										"start": 1703203200,
										"end": 1703203200,
										"gmtoffset": 0
									},
									"regular": {
										"timezone": "GMT",
										"start": 1703203200,
										"end": 1703289540,
										"gmtoffset": 0
									},
									"post": {
										"timezone": "GMT",
										"start": 1703289540,
										"end": 1703289540,
										"gmtoffset": 0
									}
								},
								"dataGranularity": "1d",
								"range": "1d",
								"validRanges": [
									"1d",
									"5d",
									"1mo",
									"3mo",
									"6mo",
									"1y",
									"2y",
									"5y",
									"10y",
									"ytd",
									"max"
								]
							},
							"timestamp": [
								1703203200
							],
							"indicators": {
								"quote": [
									{
									"volume": [
											0
										],
										"open": [
											156.3070068359375
										],
										"close": [
											156.3070068359375
										],
										"high": [
											156.99899291992188
										],
										"low": [
											156.1269989013672
										]
									}
								],
								"adjclose": [
									{
										"adjclose": [
											156.3070068359375
										]
									}
								]
							}
						}
					],
					"error": null
				}
			}
		`), nil
	})

	resp, err := yahoo.Request(context.Background(), "eur", "jpy")
	require.NoError(err)
	assert.Equal(156.878, resp.Chart.Result[0].Meta.RegularMarketPrice)
}

func TestRequestError(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://query1.finance.yahoo.com/v8/finance/chart/EURXXX=X?interval=1d", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusNotFound, `
			{
				"chart": {
					"result": null,
					"error": {
						"code": "Not Found",
						"description": "No data found, symbol may be delisted"
					}
				}
			}
		`), nil
	})

	_, err := yahoo.Request(context.Background(), "eur", "xxx")
	assert.ErrorContains(err, "No data found, symbol may be delisted")
}
