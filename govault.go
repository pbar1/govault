package govault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type (
	Client struct {
		httpClient *http.Client
		Address    string
		Token      string
		Logger     Logger
	}

	vaultResponse struct {
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
func NewClient(httpClient *http.Client, address, token string, logger Logger) *Client {
	return &Client{httpClient, address, token, logger}
}

// NewDefaultClient constructs a Vault client using environment variables VAULT_ADDR
// and VAULT_TOKEN for address and token respectively.
func NewDefaultClient() *Client {
	address := os.Getenv("VAULT_ADDR")
	if address == "" {
		address = "http://127.0.0.1:8200"
	}
	return &Client{&http.Client{}, address, os.Getenv("VAULT_TOKEN"), NewStdLogger()}
}

func (c *Client) doV1(method, endpoint string, params map[string]interface{}, body interface{}) (*vaultResponse, error) {
	// serialize request body
	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(b)
	}

	// build request, add query parameters and headers
	reqURL := c.Address + "/v1/" + endpoint
	req, err := http.NewRequest(method, reqURL, reqBody)
	if err != nil {
		return nil, err
	}
	if params != nil {
		for p, q := range params {
			req.URL.Query().Set(p, fmt.Sprint(q))
		}
	}
	req.Header.Add("X-Vault-Token", c.Token)
	req.Header.Add("X-Vault-Request", "true")

	// execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// check known status codes
	if err := checkStatus(resp.StatusCode); err != nil {
		return nil, err
	}

	// parse response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	v := new(vaultResponse)
	if err := json.Unmarshal(respBody, v); err != nil {
		return nil, err
	}

	return v, nil
}

// typeConvert takes an object "from" and, using JSON marshal/unmarshal, converts it into the given "to" object pointer.
// Note: "to" must be a pointer to an object, not the object itself.
func typeConvert(from, toPtr interface{}) error {
	b, err := json.Marshal(from)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, toPtr); err != nil {
		return err
	}
	return nil
}
