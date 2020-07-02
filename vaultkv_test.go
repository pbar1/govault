package vaultkv

import (
	"net/http"
	"reflect"
	"testing"
)

func TestVaultKV_Put(t *testing.T) {
	type fields struct {
		Address    string
		Token      string
		Mount      string
		V2         bool
		httpClient *http.Client
	}
	type args struct {
		path string
		data map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"CanPutV2",
			fields{"http://127.0.0.1:8200", "test", "secret", true, &http.Client{}},
			args{"test", map[string]string{"foo": "foo"}},
			false,
		},
		{
			"CanPutV1",
			fields{"http://127.0.0.1:8200", "test", "kv", false, &http.Client{}},
			args{"test", map[string]string{"foo": "foo"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VaultKV{
				Address:    tt.fields.Address,
				Token:      tt.fields.Token,
				Mount:      tt.fields.Mount,
				V2:         tt.fields.V2,
				httpClient: tt.fields.httpClient,
			}
			if err := v.Put(tt.args.path, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVaultKV_Get(t *testing.T) {
	type fields struct {
		Address    string
		Token      string
		Mount      string
		V2         bool
		httpClient *http.Client
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			"CanGetV2",
			fields{"http://127.0.0.1:8200", "test", "secret", true, &http.Client{}},
			args{"test"},
			map[string]string{"foo": "foo"},
			false,
		},
		{
			"CanGetV1",
			fields{"http://127.0.0.1:8200", "test", "kv", false, &http.Client{}},
			args{"test"},
			map[string]string{"foo": "foo"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VaultKV{
				Address:    tt.fields.Address,
				Token:      tt.fields.Token,
				Mount:      tt.fields.Mount,
				V2:         tt.fields.V2,
				httpClient: tt.fields.httpClient,
			}
			got, err := v.Get(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVaultKV_Delete(t *testing.T) {
	type fields struct {
		Address    string
		Token      string
		Mount      string
		V2         bool
		httpClient *http.Client
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"CanDeleteV2",
			fields{"http://127.0.0.1:8200", "test", "secret", true, &http.Client{}},
			args{"test"},
			false,
		},
		{
			"CanDeleteV1",
			fields{"http://127.0.0.1:8200", "test", "kv", false, &http.Client{}},
			args{"test"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VaultKV{
				Address:    tt.fields.Address,
				Token:      tt.fields.Token,
				Mount:      tt.fields.Mount,
				V2:         tt.fields.V2,
				httpClient: tt.fields.httpClient,
			}
			if err := v.Delete(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
