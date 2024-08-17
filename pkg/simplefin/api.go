package simplefin

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *client) GetAccounts(ctx context.Context) (*Accounts, error) {
	url := c.accessUrl + accountsEndpoint

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var accounts Accounts
	err = json.Unmarshal(body, &accounts)
	if err != nil {
		return nil, err
	}

	return &accounts, nil
}
