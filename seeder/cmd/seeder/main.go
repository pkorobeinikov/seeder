package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	var (
		err error
		s   Spec
	)

	b, err := os.ReadFile("seeder-showcase/101_postgres_file/seeder.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = yaml.Unmarshal(b, &s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", s)
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
