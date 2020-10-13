package govault

import (
	"net/http"
	"reflect"
	"sync"
	"testing"
)

func TestClient_KVv2Get(t *testing.T) {
	type fields struct {
		lock       *sync.Mutex
		httpClient *http.Client
		Address    string
		Token      string
	}
	type args struct {
		secretPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KVv2GetResponse
		wantErr bool
	}{
		{
			"Test",
			fields{
				lock:       &sync.Mutex{},
				httpClient: &http.Client{},
				Address:    "http://127.0.0.1:8200",
				Token:      "test",
			},
			args{secretPath: "secret/data/data/foo"},
			&KVv2GetResponse{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				lock:       tt.fields.lock,
				httpClient: tt.fields.httpClient,
				Address:    tt.fields.Address,
				Token:      tt.fields.Token,
			}
			got, err := c.KVv2Get(tt.args.secretPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("KVv2Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KVv2Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
