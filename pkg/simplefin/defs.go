package simplefin

import "context"

type SimplefinClient interface {
	GetAccounts(ctx context.Context) (*Accounts, error)
}

const (
	accountsEndpoint = "/accounts"
)

type Accounts struct {
	Errors   []any `json:"errors"`
	Accounts []struct {
		Org struct {
			Domain  string `json:"domain"`
			SfinURL string `json:"sfin-url"`
		} `json:"org"`
		ID               string `json:"id"`
		Name             string `json:"name"`
		Currency         string `json:"currency"`
		Balance          string `json:"balance"`
		AvailableBalance string `json:"available-balance"`
		BalanceDate      int    `json:"balance-date"`
		Transactions     []any  `json:"transactions"`
	} `json:"accounts"`
}
