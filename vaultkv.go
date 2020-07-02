package vaultkv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type (
	// VaultKV is a client for the Vault key/value secrets engine API.
	VaultKV struct {
		Address    string
		Token      string
		Mount      string
		V2         bool
		httpClient *http.Client
	}

	kvSecretV2Wrapper struct {
		Data *kvSecret `json:"data"`
	}

	kvSecret struct {
		Data map[string]string `json:"data"`
	}
)

// New constructs a VaultKV client.
func New(address, token, mount string, v2 bool) *VaultKV {
	return &VaultKV{address, token, mount, v2, &http.Client{}}
}

// NewDefault constructs a VaultKV client using environment variables VAULT_ADDR
// and VAULT_TOKEN for address and token respectively, assumes the mount path
// `secret`, and assumes the KV v2 API.
func NewDefault() *VaultKV {
	address := os.Getenv("VAULT_ADDR")
	if address == "" {
		address = "http://127.0.0.1:8200"
	}
	return &VaultKV{address, os.Getenv("VAULT_TOKEN"), "secret", true, &http.Client{}}
}

// Get retrieves the latest version of the secret at the given path.
// TODO: v1 is broken,
func (v *VaultKV) Get(path string) (map[string]string, error) {
	var url string
	if v.V2 {
		url = fmt.Sprintf("%s/v1/%s/data/%s", v.Address, v.Mount, path)
	} else {
		url = fmt.Sprintf("%s/v1/%s/%s", v.Address, v.Mount, path)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Vault-Token", v.Token)
	resp, err := v.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if v.V2 {
		var sec kvSecretV2Wrapper
		if err := json.Unmarshal(body, &sec); err != nil {
			return nil, err
		}
		return sec.Data.Data, nil
	} else {
		var sec kvSecret
		if err := json.Unmarshal(body, &sec); err != nil {
			return nil, err
		}
		return sec.Data, nil
	}
}

// Put creates a new secret at the given path if it does not exist, else updates
// the secret at that path. If KV v2, updates create new secret versions. If KV
// v1, updates overwrite the existing secret contents.
func (v *VaultKV) Put(path string, data map[string]string) error {
	var url string
	var payload []byte
	var err error
	if v.V2 {
		url = fmt.Sprintf("%s/v1/%s/data/%s", v.Address, v.Mount, path)
		payload, err = json.Marshal(kvSecret{data})
	} else {
		url = fmt.Sprintf("%s/v1/%s/%s", v.Address, v.Mount, path)
		payload, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Add("X-Vault-Token", v.Token)
	resp, err := v.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Delete either removes the latest version of a secret (KV v2), or completely
// removes a secret (KV v1).
func (v *VaultKV) Delete(path string) error {
	var url string
	if v.V2 {
		url = fmt.Sprintf("%s/v1/%s/data/%s", v.Address, v.Mount, path)
	} else {
		url = fmt.Sprintf("%s/v1/%s/%s", v.Address, v.Mount, path)
	}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-Vault-Token", v.Token)
	resp, err := v.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Destroy completely removes a secret.
func (v *VaultKV) Destroy(path string) error {
	var url string
	if v.V2 {
		url = fmt.Sprintf("%s/v1/%s/metadata/%s", v.Address, v.Mount, path)
	} else {
		return v.Delete(path)
	}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-Vault-Token", v.Token)
	resp, err := v.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
