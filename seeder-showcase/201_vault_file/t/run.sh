VAULT_ADDR=http://localhost:8200 VAULT_TOKEN=secret vault kv get -format=json secret/seed/file/json
VAULT_ADDR=http://localhost:8200 VAULT_TOKEN=secret vault kv get -format=json secret/seed/file/yaml
VAULT_ADDR=http://localhost:8200 VAULT_TOKEN=secret vault kv get -format=json secret/seed/file/yml
