package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/pkorobeinikov/seeder/seeder"

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

	r := seeder.DefaultRegistry()

	seederYaml := flag.String("c", "seeder.yaml", "config file")
	showSeederHelp := flag.String("seeder-help", "", "show seeder help (ex: -seeder-help kafka)")
	knownSeeders := flag.Bool("known", false, "show known seeders list")
	flag.Parse()

	if *showSeederHelp != "" {
		known := r.ListKnownTypes()

		var isExists bool = false
		for _, v := range known {
			if v == *showSeederHelp {
				isExists = true
				break
			}
		}

		if !isExists {
			printErr(errors.Errorf("unknown seeder: %s", *showSeederHelp))
			return
		}

		r.ShowSeederHelp(*showSeederHelp, os.Stdout)
		return
	}

	if *knownSeeders {
		known := r.ListKnownTypes()
		fmt.Printf("Known seeders: %s\n", strings.Join(known, ", "))
		return
	}

	//
	// main
	//

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
