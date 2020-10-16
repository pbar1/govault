package govault

import (
	"errors"
	"net/http"
	"path"
	"time"
)

const DefaultKVv2MountPath = "secret"

type (
	KVv2 interface {
		WithMountPath(path string) KVv2
		Configure(options *KVv2Config) error
		ReadConfig() (*KVv2Config, error)
		ReadSecretVersion(path string, version int) (*KVv2Secret, error)
		CreateOrUpdateSecret(path string, data map[string]string, options *KVv2CreateOrUpdateSecretOptions) error
		DeleteLatestSecretVersion(path string) error
		DeleteSecretVersions(path string, versions []int) error
		UndeleteSecretVersions(path string, versions []int) error
		DestroySecretVersions(path string, versions []int) error
		ListSecrets(path string) ([]string, error)
		ReadSecretMetadata(path string) error
		UpdateMetadata(path string, maxVersions int, casRequired bool, deleteVersionAfter time.Duration) error
		DeleteMetadataAndAllVersions(path string) error
	}

	kvv2Impl struct {
		client    *Client
		MountPath string
	}

	KVv2Config struct {
		MaxVersions        int    `json:"max_versions"`
		CASRequired        bool   `json:"cas_required"`
		DeleteVersionAfter string `json:"delete_version_after,omitempty"`
	}

	KVv2Secret struct {
		Data     map[string]string `json:"data"`
		Metadata struct {
			CreatedTime  string `json:"created_time"`
			DeletionTime string `json:"deletion_time"`
			Destroyed    bool   `json:"destroyed"`
			Version      int    `json:"version"`
		} `json:"metadata"`
	}

	kvv2CreateOrUpdateSecretRequest struct {
		Data    map[string]string               `json:"data"`
		Options KVv2CreateOrUpdateSecretOptions `json:"options"`
	}

	KVv2CreateOrUpdateSecretOptions struct {
		CAS int `json:"cas,omitempty"`
	}
)

func (c *Client) KVv2() KVv2 {
	return &kvv2Impl{
		client:    c,
		MountPath: DefaultKVv2MountPath,
	}
}

func (k *kvv2Impl) do(method, endpoint string, params map[string]interface{}, body interface{}) (*vaultResponse, error) {
	return k.client.doV1(method, path.Join(k.MountPath, endpoint), params, body)
}

func (k *kvv2Impl) WithMountPath(path string) KVv2 {
	kCopy := *k
	kCopy.MountPath = path
	k.client.Logger.Debug("using mount path: " + path)
	return &kCopy
}

// curl command: `curl -X POST -H "X-Vault-Request: true" -H "X-Vault-Token: $(vault print token)" -d '{"max_versions":5,"cas_required":false,"delete_version_after":"3h25m19s"}' http://127.0.0.1:8200/v1/secret/config`
func (k *kvv2Impl) Configure(config *KVv2Config) error {
	r, err := k.do(http.MethodPost, "config", nil, config)
	if err != nil && !errors.Is(err, &ErrSuccessNoData{}) {
		return err
	}
	k.client.Logger.Debug(r)
	return nil
}

func (k *kvv2Impl) ReadConfig() (*KVv2Config, error) {
	r, err := k.do(http.MethodGet, "config", nil, nil)
	if err != nil {
		return nil, err
	}
	k.client.Logger.Debug(r)
	v := new(KVv2Config)
	if err := typeConvert(r.Data, v); err != nil {
		return nil, err
	}
	return v, nil
}

// vault command: `vault kv get -version={version} secret/{path}`
// curl command: `curl -H "X-Vault-Request: true" -H "X-Vault-Token: $(vault print token)" http://127.0.0.1:8200/v1/secret/data/{path}?version={version}`
func (k *kvv2Impl) ReadSecretVersion(path string, version int) (*KVv2Secret, error) {
	q := map[string]interface{}{"version": version}
	r, err := k.do(http.MethodGet, "data/"+path, q, nil)
	if err != nil {
		return nil, err
	}
	k.client.Logger.Debug(r)
	v := new(KVv2Secret)
	if err := typeConvert(r.Data, v); err != nil {
		return nil, err
	}
	return v, nil
}

// vault command: `vault kv put -cas=1 secret/mysecret mykey=myval
// `curl -X PUT -H "X-Vault-Request: true" -H "X-Vault-Token: $(vault print token)" -d '{"data":{"mykey":"myval"},"options":{"cas":1}}' http://127.0.0.1:8200/v1/secret/data/mysecret`
func (k *kvv2Impl) CreateOrUpdateSecret(path string, data map[string]string, options *KVv2CreateOrUpdateSecretOptions) error {
	body := kvv2CreateOrUpdateSecretRequest{Data: data, Options: *options}
	r, err := k.do(http.MethodPut, "data/"+path, nil, body)
	if err != nil && !errors.Is(err, &ErrSuccessNoData{}) {
		return err
	}
	k.client.Logger.Trace(r)
	return err
}

func (k *kvv2Impl) DeleteLatestSecretVersion(path string) error {
	r, err := k.do(http.MethodDelete, "data/"+path, nil, nil)
	if err != nil && !errors.Is(err, &ErrSuccessNoData{}) {
		return err
	}
	k.client.Logger.Trace(r)
	return nil
}

// curl -X POST -H "X-Vault-Request: true" -H "X-Vault-Token: $(vault print token)" -d '{"versions":[1,2]}' https://127.0.0.1:8200/v1/secret/data/my-secret
func (k *kvv2Impl) DeleteSecretVersions(path string, versions []int) error {
	body := map[string]interface{}{"versions": versions}
	r, err := k.do(http.MethodPost, "delete/"+path, nil, body)
	if err != nil && !errors.Is(err, &ErrSuccessNoData{}) {
		return err
	}
	k.client.Logger.Trace(r)
	return nil
}

// curl -X POST -H "X-Vault-Request: true" -H "X-Vault-Token: $(vault print token)" -d '{"versions":[1,2]}' https://127.0.0.1:8200/v1/secret/undelete/my-secret
func (k *kvv2Impl) UndeleteSecretVersions(path string, versions []int) error {
	body := map[string]interface{}{"versions": versions}
	r, err := k.do(http.MethodPost, "undelete/"+path, nil, body)
	if err != nil && !errors.Is(err, &ErrSuccessNoData{}) {
		return err
	}
	k.client.Logger.Trace(r)
	return nil
}

// curl -X POST -H "X-Vault-Request: true" -H "X-Vault-Token: $(vault print token)" -d '{"versions":[1,2]}' https://127.0.0.1:8200/v1/secret/destroy/my-secret
func (k *kvv2Impl) DestroySecretVersions(path string, versions []int) error {
	body := map[string]interface{}{"versions": versions}
	r, err := k.do(http.MethodPost, "destroy/"+path, nil, body)
	if err != nil && !errors.Is(err, &ErrSuccessNoData{}) {
		return err
	}
	k.client.Logger.Trace(r)
	return nil
}

func (k *kvv2Impl) ListSecrets(path string) ([]string, error) {
	r, err := k.do(http.MethodGet, "metadata/"+path, nil, nil)
	if err != nil && !errors.Is(err, &ErrSuccessNoData{}) {
		return nil, err
	}
	k.client.Logger.Trace(r)
	data := new(struct {
		Keys []string `json:"keys"`
	})
	if err := typeConvert(r.Data, data); err != nil {
		return nil, err
	}
	return data.Keys, nil
}

func (k *kvv2Impl) ReadSecretMetadata(path string) error {
	panic("implement me")
}

func (k *kvv2Impl) UpdateMetadata(path string, maxVersions int, casRequired bool, deleteVersionAfter time.Duration) error {
	panic("implement me")
}

func (k *kvv2Impl) DeleteMetadataAndAllVersions(path string) error {
	panic("implement me")
}
