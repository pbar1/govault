package main

import "github.com/pbar1/govault"

func main() {
	// the default client looks for VAULT_ADDR and VAULT_TOKEN
	vault := govault.NewDefaultClient()

	// read latest version of KV v2 secret "/secret/foo"
	vault.KVv2().ReadSecretVersion("foo", 0)

	// read specific version of a KV v2 secret on a non-default mount path, ie "/kv/bar"
	vault.KVv2().WithMountPath("kv").ReadSecretVersion("bar", 3)
}
