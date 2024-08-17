package simplefin

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

type client struct {
	accessUrl  string
	httpClient http.Client
}

func NewSimplefinClient(accessUrl string) (SimplefinClient, error) {
	return &client{
		accessUrl:  accessUrl,
		httpClient: *http.DefaultClient,
	}, nil
}

func NewSimplefinClientFromSetupToken(setupToken string) (SimplefinClient, error) {
	claimUrl, err := base64.StdEncoding.DecodeString(setupToken)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 setupToken: %v", err)
	}

	req, err := http.NewRequest("POST", string(claimUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.ContentLength = 0
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error requesting access url: %v", err)
	}
	defer resp.Body.Close()

	accessUrl, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error getting access url: %v", err)
	}

	return NewSimplefinClient(string(accessUrl))
}
