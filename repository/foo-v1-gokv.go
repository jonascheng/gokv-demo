package repository

import (
	"context"
	"log"

	"github.com/jonascheng/gokv-demo/domain"
	"github.com/philippgille/gokv"
)

type fooV1Repository struct {
	store gokv.Store
}

func NewFooV1Repository(db gokv.Store) domain.FooV1Repository {
	return &fooV1Repository{db}
}

func (this *fooV1Repository) Save(ctx context.Context, foo *domain.FooV1) error {
	err := this.store.Set(foo.ID, foo)
	if err != nil {
		log.Error(err)
	}
}

// GetByID(ctx context.Context, id string) (*fooV1, error)
// Update(ctx context.Context, foo *fooV1) error
