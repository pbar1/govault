package govault

import (
	"encoding/json"
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

	responseContainer struct {
		RequestID     string      `json:"request_id"`
		LeaseID       string      `json:"lease_id"`
		Renewable     bool        `json:"renewable"`
		LeaseDuration int         `json:"lease_duration"`
		Data          interface{} `json:"data"`
		WrapInfo      interface{} `json:"wrap_info"`
		Warnings      interface{} `json:"warnings"`
		Auth          interface{} `json:"auth"`
	}
)

// NewClient constructs a Vault client.
func NewClient(httpClient *http.Client, address, token string) *Client {
	return &Client{&sync.Mutex{}, httpClient, address, token}
}

// NewDefaultClient constructs a Vault client using environment variables VAULT_ADDR
// and VAULT_TOKEN for address and token respectively.
func NewDefaultClient() *Client {
	address := os.Getenv("VAULT_ADDR")
	if address == "" {
		address = "http://127.0.0.1:8200"
	}
	return &Client{&sync.Mutex{}, &http.Client{}, address, os.Getenv("VAULT_TOKEN")}
}

func (c *Client) doV1(method, endpoint string, reqBody io.Reader) (*responseContainer, error) {
	// build request
	reqURL := fmt.Sprintf("%s/v1/%s", c.Address, endpoint)
	req, err := http.NewRequest(method, reqURL, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Vault-Token", c.Token)

	// execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	v := new(responseContainer)
	if err := json.Unmarshal(respBody, v); err != nil {
		return nil, err
	}
	// TODO: properly handle expected empty response

	return v, nil
}
