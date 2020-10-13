package govault

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type (
	Client struct {
		lock       *sync.Mutex
		httpClient *http.Client
		Address    string
		Token      string
	}
)

// New constructs a Vault client.
func NewClient(httpClient *http.Client, address, token string) *Client {
	return &Client{&sync.Mutex{}, httpClient, address, token}
}

// NewDefault constructs a Vault client using environment variables VAULT_ADDR
// and VAULT_TOKEN for address and token respectively.
func NewDefaultClient() *Client {
	address := os.Getenv("VAULT_ADDR")
	if address == "" {
		address = "http://127.0.0.1:8200"
	}
	return &Client{&sync.Mutex{}, &http.Client{}, address, os.Getenv("VAULT_TOKEN")}
}

func (c *Client) do(method, endpoint string, reqBody io.Reader) ([]byte, error) {
	reqURL := fmt.Sprintf("%s/v1/%s", c.Address, endpoint)
	req, err := http.NewRequest(http.MethodGet, reqURL, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Vault-Token", c.Token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
