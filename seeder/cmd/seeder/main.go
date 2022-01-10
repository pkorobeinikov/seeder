package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
	"github.com/jackc/pgx/v4"
	"gopkg.in/yaml.v2"
)

const (
	SeederYaml = "seeder.yaml"

	SeederPgConnStrEnv    = "SEEDER_PG_CONNSTR"
	SeederVaultAddressEnv = "SEEDER_VAULT_ADDRESS"
	SeederVaultTokenEnv   = "SEEDER_VAULT_TOKEN"
)

func main() {
	var (
		err error
		s   Spec
	)

	ctx := context.Background()

	b, err := os.ReadFile(SeederYaml)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = yaml.Unmarshal(b, &s)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, state := range s.Seeder.State {
		for _, cfg := range state.Config {
			switch state.Type {
			case "postgres":
				connstr, found := os.LookupEnv(SeederPgConnStrEnv)
				if !found {
					fmt.Println(errors.New("connection string is not set"))
					return
				}

				conn, err := pgx.Connect(ctx, connstr)
				if err != nil {
					fmt.Println(err)
					return
				}

				b, err := os.ReadFile(cfg.File)
				if err != nil {
					fmt.Println(err)
					return
				}

				_, err = conn.Exec(ctx, string(b))
				if err != nil {
					fmt.Println(err)
					return
				}

				err = conn.Close(ctx)
				if err != nil {
					fmt.Println(err)
					return
				}
			case "vault":
				b, err := os.ReadFile(cfg.File)
				if err != nil {
					fmt.Println(err)
					return
				}

				var secret map[string]interface{}
				switch {
				case strings.HasSuffix(cfg.File, ".json"):
					err := json.Unmarshal(b, &secret)
					if err != nil {
						fmt.Println(err)
						return
					}
				case strings.HasSuffix(cfg.File, ".yaml"):
					fallthrough
				case strings.HasSuffix(cfg.File, ".yml"):

				default:
					fmt.Println(errors.New("unsupported file type"))
					return
				}

				vaultAddr, found := os.LookupEnv(SeederVaultAddressEnv)
				if !found {
					fmt.Println(errors.New("vault address not set"))
					return
				}

				vaultToken, found := os.LookupEnv(SeederVaultTokenEnv)
				if !found {
					fmt.Println(errors.New("vault token not set"))
					return
				}

				config := &api.Config{
					Address: vaultAddr,
				}

				client, err := api.NewClient(config)
				if err != nil {
					fmt.Println(err)
					return
				}

				client.SetToken(vaultToken)

				_, err = client.Logical().Write(cfg.Key, secret)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

type (
	Spec struct {
		Seeder Seeder `yaml:"seeder"`
	}

	Seeder struct {
		State []State
	}

	State struct {
		Name   string `yaml:"name"`
		Type   string `yaml:"type"`
		Config []Config
	}

	Config struct {
		File string `yaml:"file"`
		Key  string `yaml:"key"`
	}
)
