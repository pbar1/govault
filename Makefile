export VAULT_ADDR := http://127.0.0.1:8200
export VAULT_TOKEN := test

test:
	@vault server -dev -dev-root-token-id=$(VAULT_TOKEN) &
	@vault secrets enable -version=1 || true
	@go test -v ./...
