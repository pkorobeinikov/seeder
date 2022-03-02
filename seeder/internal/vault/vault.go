package vault

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	"github.com/pkorobeinikov/seeder/seeder"
	"gopkg.in/yaml.v3"
)

const (
	SeederVaultAddressEnv = "SEEDER_VAULT_ADDRESS"
	SeederVaultTokenEnv   = "SEEDER_VAULT_TOKEN"
)

func Seed(_ context.Context, cfg seeder.Config) error {

	b, err := os.ReadFile(cfg.File)
	if err != nil {
		return errors.Wrap(err, "read file")
	}

	var secret map[string]interface{}
	switch {
	case strings.HasSuffix(cfg.File, ".json"):
		err := json.Unmarshal(b, &secret)
		if err != nil {
			return errors.Wrap(err, "unmarshal json")
		}
	case strings.HasSuffix(cfg.File, ".yaml"):
		fallthrough
	case strings.HasSuffix(cfg.File, ".yml"):
		err := yaml.Unmarshal(b, &secret)
		if err != nil {
			return errors.Wrap(err, "unmarshal yaml")
		}
	default:
		return errors.Errorf("unsupported file type: %s", cfg.File)
	}

	vaultAddr, found := os.LookupEnv(SeederVaultAddressEnv)
	if !found {
		return errors.New("vault address not set")
	}

	vaultToken, found := os.LookupEnv(SeederVaultTokenEnv)
	if !found {
		return errors.New("vault token not set")
	}

	config := &api.Config{
		Address: vaultAddr,
	}

	client, err := api.NewClient(config)
	if err != nil {
		return errors.Wrap(err, "vault: new api client")
	}

	client.SetToken(vaultToken)

	_, err = client.Logical().Write(cfg.Key, secret)
	if err != nil {
		return errors.Wrap(err, "vault: logical write")
	}

	return nil
}

func init() {
	seeder.DefaultRegistry().RegisterSeeder(Seed, "vault")
}
