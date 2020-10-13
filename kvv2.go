package govault

import (
	"encoding/json"
	"net/http"
	"time"
)

type KVv2GetResponse struct {
	RequestID     string `json:"request_id"`
	LeaseID       string `json:"lease_id"`
	Renewable     bool   `json:"renewable"`
	LeaseDuration int    `json:"lease_duration"`
	Data          struct {
		Data     map[string]string `json:"data"`
		Metadata struct {
			CreatedTime  time.Time `json:"created_time"`
			DeletionTime string    `json:"deletion_time"`
			Destroyed    bool      `json:"destroyed"`
			Version      int       `json:"version"`
		} `json:"metadata"`
	} `json:"data"`
	WrapInfo interface{} `json:"wrap_info"`
	Warnings interface{} `json:"warnings"`
	Auth     interface{} `json:"auth"`
}

func (c *Client) KVv2Get(secretPath string) (*KVv2GetResponse, error) {
	resp, err := c.do(http.MethodGet, secretPath, nil)
	if err != nil {
		return nil, err
	}
	var v KVv2GetResponse
	if err := json.Unmarshal(resp, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
