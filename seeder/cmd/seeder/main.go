package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/pkorobeinikov/seeder/seeder"
	"gopkg.in/yaml.v3"

	_ "github.com/pkorobeinikov/seeder/seeder/internal/kafka"
	_ "github.com/pkorobeinikov/seeder/seeder/internal/postgres"
	_ "github.com/pkorobeinikov/seeder/seeder/internal/s3"
	_ "github.com/pkorobeinikov/seeder/seeder/internal/vault"
)

func main() {
	var (
		err error
		s   seeder.Spec
	)

	seederYaml := flag.String("c", "./seeder.yml", "config file")
	flag.Parse()

	wd := filepath.Dir(*seederYaml)
	if wd != "" {
		fmt.Printf("working dir: %s\n", wd)
	}

	ctx := context.Background()

	b, err := os.ReadFile(*seederYaml)
	if err != nil {
		printErr(errors.Wrap(err, "read config"))
		return
	}

	err = yaml.Unmarshal(b, &s)
	if err != nil {
		printErr(errors.Wrap(err, "unmarshal config"))
		return
	}

	r := seeder.DefaultRegistry()

	for _, state := range s.Seeder.State {
		for i := range state.Config {
			cfg := state.Config[i]
			cfg.File = filepath.Join(wd, cfg.File)

			err := r.RunSeeder(ctx, state.Type, cfg)
			if err != nil {
				fmt.Println("error:", err.Error())
			}
		}
	}
}

func printErr(e error) {
	fmt.Println("error:", e.Error())
	os.Exit(1)
}
