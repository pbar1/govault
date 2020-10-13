package govault

import (
	"os"
	"reflect"
	"testing"
)

func init() {
	os.Setenv("VAULT_ADDR", "http://127.0.0.1:8200")
	os.Setenv("VAULT_TOKEN", "test")
}

func TestKVv2ClientImpl_Configure(t *testing.T) {
	type fields struct {
		client    *Client
		MountPath string
	}
	type args struct {
		config *KVv2Config
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test",
			fields: fields{
				client:    NewDefaultClient(),
				MountPath: DefaultKVv2MountPath,
			},
			args: args{
				config: &KVv2Config{
					MaxVersions: 5,
					CASRequired: true,
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &kvv2ClientImpl{
				client:    tt.fields.client,
				MountPath: tt.fields.MountPath,
			}
			if err := k.Configure(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Configure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKVv2ClientImpl_ReadConfiguration(t *testing.T) {
	type fields struct {
		client    *Client
		MountPath string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *KVv2Config
		wantErr bool
	}{
		{
			name: "Test",
			fields: fields{
				client:    NewDefaultClient(),
				MountPath: DefaultKVv2MountPath,
			},
			want: &KVv2Config{
				MaxVersions: 0,
				CASRequired: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &kvv2ClientImpl{
				client:    tt.fields.client,
				MountPath: tt.fields.MountPath,
			}
			got, err := k.ReadConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
