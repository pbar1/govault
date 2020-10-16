package govault

import (
	"os"
	"reflect"
	"testing"
)

var testClient *Client

func init() {
	os.Setenv("VAULT_ADDR", "http://127.0.0.1:8200")
	os.Setenv("VAULT_TOKEN", "test")

	testClient = NewDefaultClient()
	testClientLogger := NewStdLogger()
	testClientLogger.Level = LevelTrace
	testClient.Logger = testClientLogger
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
				client:    testClient,
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
			k := &kvv2Impl{
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
				client:    testClient,
				MountPath: DefaultKVv2MountPath,
			},
			want: &KVv2Config{
				MaxVersions: 5,
				CASRequired: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &kvv2Impl{
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

func Test_kvv2ClientImpl_ReadSecretVersion(t *testing.T) {
	type fields struct {
		client    *Client
		MountPath string
	}
	type args struct {
		path    string
		version int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KVv2Secret
		wantErr bool
	}{
		{
			name: "Test",
			fields: fields{
				client:    testClient,
				MountPath: DefaultKVv2MountPath,
			},
			args: args{
				path:    "foo",
				version: 1,
			},
			want: &KVv2Secret{
				Data: map[string]string{
					"foo": "foo",
				},
				Metadata: struct {
					CreatedTime  string `json:"created_time"`
					DeletionTime string `json:"deletion_time"`
					Destroyed    bool   `json:"destroyed"`
					Version      int    `json:"version"`
				}{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &kvv2Impl{
				client:    tt.fields.client,
				MountPath: tt.fields.MountPath,
			}
			got, err := k.WithMountPath("secret").ReadSecretVersion(tt.args.path, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadSecretVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadSecretVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kvv2Impl_CreateOrUpdateSecret(t *testing.T) {
	type fields struct {
		client    *Client
		MountPath string
	}
	type args struct {
		path    string
		data    map[string]string
		options *KVv2CreateOrUpdateSecretOptions
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
				client:    testClient,
				MountPath: DefaultKVv2MountPath,
			},
			args: args{
				path: "foo",
				data: map[string]string{
					"foo": "supersec",
				},
				options: &KVv2CreateOrUpdateSecretOptions{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &kvv2Impl{
				client:    tt.fields.client,
				MountPath: tt.fields.MountPath,
			}
			if err := k.CreateOrUpdateSecret(tt.args.path, tt.args.data, tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdateSecret() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_kvv2Impl_DeleteLatestSecretVersion(t *testing.T) {
	type fields struct {
		client    *Client
		MountPath string
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
			name: "Test",
			fields: fields{
				client:    testClient,
				MountPath: DefaultKVv2MountPath,
			},
			args: args{
				path: "foo",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &kvv2Impl{
				client:    tt.fields.client,
				MountPath: tt.fields.MountPath,
			}
			if err := k.DeleteLatestSecretVersion(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("DeleteLatestSecretVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_kvv2Impl_DeleteSecretVersions(t *testing.T) {
	type fields struct {
		client    *Client
		MountPath string
	}
	type args struct {
		path     string
		versions []int
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
				client:    testClient,
				MountPath: DefaultKVv2MountPath,
			},
			args: args{
				path:     "foo",
				versions: []int{1, 2, 3, 4},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &kvv2Impl{
				client:    tt.fields.client,
				MountPath: tt.fields.MountPath,
			}
			if err := k.DeleteSecretVersions(tt.args.path, tt.args.versions); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSecretVersions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_kvv2Impl_UndeleteSecretVersions(t *testing.T) {
	type fields struct {
		client    *Client
		MountPath string
	}
	type args struct {
		path     string
		versions []int
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
				client:    testClient,
				MountPath: DefaultKVv2MountPath,
			},
			args: args{
				path:     "foo",
				versions: []int{1, 2, 3, 4},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &kvv2Impl{
				client:    tt.fields.client,
				MountPath: tt.fields.MountPath,
			}
			if err := k.UndeleteSecretVersions(tt.args.path, tt.args.versions); (err != nil) != tt.wantErr {
				t.Errorf("UndeleteSecretVersions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_kvv2Impl_DestroySecretVersions(t *testing.T) {
	type fields struct {
		client    *Client
		MountPath string
	}
	type args struct {
		path     string
		versions []int
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
				client:    testClient,
				MountPath: DefaultKVv2MountPath,
			},
			args: args{
				path:     "foo",
				versions: []int{1, 2, 3, 4},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &kvv2Impl{
				client:    tt.fields.client,
				MountPath: tt.fields.MountPath,
			}
			if err := k.DestroySecretVersions(tt.args.path, tt.args.versions); (err != nil) != tt.wantErr {
				t.Errorf("DestroySecretVersions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_kvv2Impl_ListSecrets(t *testing.T) {
	type fields struct {
		client    *Client
		MountPath string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			"Test",
			fields{
				client:    testClient,
				MountPath: DefaultKVv2MountPath,
			},
			args{path: "/"},
			[]string{"foo"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &kvv2Impl{
				client:    tt.fields.client,
				MountPath: tt.fields.MountPath,
			}
			got, err := k.ListSecrets(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSecrets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListSecrets() got = %v, want %v", got, tt.want)
			}
		})
	}
}
