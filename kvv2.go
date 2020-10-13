package govault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"time"
)

const DefaultKVv2MountPath = "secret"

type (
	KVv2Client interface {
		SetMountPath(path string) KVv2Client
		Configure(options *KVv2Config) error
		ReadConfig() (*KVv2Config, error)
		ReadSecretVersion(path string, version int) (*KVv2Secret, error)
		CreateOrUpdateSecret(data map[string]string, options *KVv2CreateOrUpdateSecretOptions) error
		DeleteLatestSecretVersion(path string) error
		DeleteSecretVersions(path string, versions []int) error
		UndeleteSecretVersions(path string, versions []int) error
		DestroySecretVersions(path string, versions []int) error
		ListSecrets(path string) error
		ReadSecretMetadata(path string) error
		UpdateMetadata(path string, maxVersions int, casRequired bool, deleteVersionAfter time.Duration) error
		DeleteMetadataAndAllVersions(path string) error
	}

	kvv2ClientImpl struct {
		client    *Client
		MountPath string
	}

	KVv2Config struct {
		MaxVersions        int           `json:"max_versions"`
		CASRequired        bool          `json:"cas_required"`
		DeleteVersionAfter time.Duration `json:"delete_version_after,omitempty"`
	}

	KVv2Secret struct {
		Data     map[string]string `json:"data"`
		Metadata struct {
			CreatedTime  time.Time `json:"created_time"`
			DeletionTime string    `json:"deletion_time"`
			Destroyed    bool      `json:"destroyed"`
			Version      int       `json:"version"`
		} `json:"metadata"`
	}

	KVv2CreateOrUpdateSecretOptions struct {
		CAS int `json:"cas"`
	}
)

func (c *Client) KVv2() KVv2Client {
	return &kvv2ClientImpl{
		client:    c,
		MountPath: DefaultKVv2MountPath,
	}
}

func (k *kvv2ClientImpl) SetMountPath(path string) KVv2Client {
	kCopy := *k
	kCopy.MountPath = path
	return &kCopy
}

func (k *kvv2ClientImpl) do(method, endpoint string, reqBody io.Reader) (*responseContainer, error) {
	return k.client.doV1(method, path.Join(k.MountPath, endpoint), reqBody)
}

func (k *kvv2ClientImpl) Configure(config *KVv2Config) error {
	b, err := json.Marshal(*config)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(b)
	_, err = k.do(http.MethodPost, "config", buf)
	return err
}

func (k *kvv2ClientImpl) ReadConfig() (*KVv2Config, error) {
	r, err := k.do(http.MethodGet, "config", nil)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(r.Data)
	if err != nil {
		return nil, err
	}
	v := new(KVv2Config)
	if err := json.Unmarshal(b, v); err != nil {
		return nil, err
	}
	return v, nil
}

func (k *kvv2ClientImpl) ReadSecretVersion(path string, version int) (*KVv2Secret, error) {
	r, err := k.do(http.MethodGet, fmt.Sprintf("data/%s?version=%d", path, version), nil)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(r.Data)
	if err != nil {
		return nil, err
	}
	v := new(KVv2Secret)
	if err := json.Unmarshal(b, v); err != nil {
		return nil, err
	}
	return v, nil
}

func (k *kvv2ClientImpl) CreateOrUpdateSecret(data map[string]string, options *KVv2CreateOrUpdateSecretOptions) error {
	panic("implement me")
}

func (k *kvv2ClientImpl) DeleteLatestSecretVersion(path string) error {
	panic("implement me")
}

func (k *kvv2ClientImpl) DeleteSecretVersions(path string, versions []int) error {
	panic("implement me")
}

func (k *kvv2ClientImpl) UndeleteSecretVersions(path string, versions []int) error {
	panic("implement me")
}

func (k *kvv2ClientImpl) DestroySecretVersions(path string, versions []int) error {
	panic("implement me")
}

func (k *kvv2ClientImpl) ListSecrets(path string) error {
	panic("implement me")
}

func (k *kvv2ClientImpl) ReadSecretMetadata(path string) error {
	panic("implement me")
}

func (k *kvv2ClientImpl) UpdateMetadata(path string, maxVersions int, casRequired bool, deleteVersionAfter time.Duration) error {
	panic("implement me")
}

func (k *kvv2ClientImpl) DeleteMetadataAndAllVersions(path string) error {
	panic("implement me")
}
