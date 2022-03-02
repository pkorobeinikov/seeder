package seeder

import (
	"context"

	"github.com/pkg/errors"
)

var defaultRegistry *registry

type (
	registry struct {
		seeders map[string]SeedFunc
	}
	SeedFunc  func(ctx context.Context, cfg Config) error
	IRegistry interface {
		ListKnownTypes() []string
		RegisterSeeder(s SeedFunc, type_ string)
		RunSeeder(ctx context.Context, type_ string, cfg Config) error
	}
)

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

func DefaultRegistry() IRegistry {
	return defaultRegistry
}

func init() {
	defaultRegistry = &registry{
		seeders: map[string]SeedFunc{},
	}
}
