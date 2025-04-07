package yahoo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	URL_PREFIX = "https://query1.finance.yahoo.com/v8/finance/chart/"
)

type ErrorResponse struct {
	Chart struct {
		Error struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"error"`
	} `json:"chart"`
}

type Response struct {
	Chart struct {
		Result []struct {
			Meta struct {
				RegularMarketPrice float64 `json:"regularMarketPrice"`
			} `json:"meta"`
		} `json:"result"`
	} `json:"chart"`
}

func Request(ctx context.Context, from string, to string) (*Response, error) {
	from = strings.ToUpper(from)
	to = strings.ToUpper(to)
	url := URL_PREFIX + from + to + "=X?interval=1d"
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("request to Yahoo! Finance API failed: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		if err != nil {
			return nil, fmt.Errorf("Yahoo! Finance API response is not OK: %s", resp.Status) //nolint:staticcheck
		}

		er := ErrorResponse{}
		err = json.Unmarshal(body, &er)

		if err != nil {
			return nil, fmt.Errorf("Yahoo! Finance API response is not OK: %s", body) //nolint:staticcheck
		}

		return nil, errors.New(er.Chart.Error.Description)
	}

	if err != nil {
		return nil, fmt.Errorf("reading Yahoo! Finance API response failed: %w", err)
	}

	r := &Response{}
	err = json.Unmarshal(body, &r)

	if err != nil {
		return nil, fmt.Errorf("parsing of Yahoo! Finance API response failed: %w", err)
	}

	return r, nil
}
