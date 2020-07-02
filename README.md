# `vaultkv`

Dead simple HashiCorp Vault key/value API for Go.

### Example

```go
package main

import (
	"github.com/pbar1/vaultkv"
)

func main() {
	// Creates a client respecting VAULT_ADDRESS and VAULT_TOKEN environment
	// variables, and assumes a KV v2 engine mounted at `secret`
	vkv := vaultkv.NewDefault()

	data := map[string]string{"password": "hunter2"}

	// Puts a new secret version at raw path `secret/data/my-accounts/cyprus-national-bank`
	vkv.Put("my-accounts/cyprus-national-bank", data)

	// Gets latest secret version at raw path `secret/data/my-accounts/cyprus-national-bank`
	vkv.Get("my-accounts/cyprus-national-bank")

	// Deletes latest secret version at raw path `secret/data/my-accounts/cyprus-national-bank`
	vkv.Delete("my-accounts/cyprus-national-bank")

	// Destroys secret at raw path `secret/metadata/my-accounts/cyprus-national-bank`
	vkv.Destroy("my-accounts/cyprus-national-bank")
}
```
