package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"gopkg.in/yaml.v2"
)

const SeederYaml = "seeder.yaml"

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

	connstr := "postgres://postgres:secret@localhost:5432/seeder?sslmode=disable"

	conn, err := pgx.Connect(ctx, connstr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close(ctx)

	for _, state := range s.Seeder.State {
		for _, cfg := range state.Config {
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
	}
)
