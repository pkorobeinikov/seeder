package seeder

import (
	"context"
	"io"

	"github.com/pkg/errors"
)

var defaultRegistry *registry

type (
	// Not thread safe.
	registry struct {
		seeders       map[string]SeedFunc
		seederHelpers map[string]HelpFunc
	}
	SeedFunc  func(ctx context.Context, cfg Config) error
	HelpFunc  func(w io.Writer)
	IRegistry interface {
		ListKnownTypes() []string
		RegisterSeeder(s SeedFunc, type_ string)
		RegisterSeederHelp(f HelpFunc, type_ string)
		RunSeeder(ctx context.Context, type_ string, cfg Config) error
		ShowSeederHelp(type_ string, w io.Writer)
	}
)

func (s *registry) RegisterSeederHelp(f HelpFunc, type_ string) {
	s.seederHelpers[type_] = f
}

func (s *registry) RegisterSeeder(f SeedFunc, type_ string) {
	s.seeders[type_] = f
}

func (s *registry) ListKnownTypes() []string {
	r := make([]string, 0, len(s.seeders))

	for i := range s.seeders {
		r = append(r, i)
	}

	return r
}

func (s *registry) RunSeeder(ctx context.Context, type_ string, cfg Config) error {

	f, ok := s.seeders[type_]
	if !ok {
		return errors.Errorf("seeder not found: %s", type_)
	}

	return f(ctx, cfg)
}

func (s *registry) ShowSeederHelp(type_ string, w io.Writer) {

	f, ok := s.seederHelpers[type_]
	if !ok {
		return
	}

	f(w)
}

func DefaultRegistry() IRegistry {
	return defaultRegistry
}

func init() {
	defaultRegistry = &registry{
		seeders:       map[string]SeedFunc{},
		seederHelpers: map[string]HelpFunc{},
	}
}
