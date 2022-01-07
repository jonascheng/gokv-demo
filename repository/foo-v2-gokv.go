package repository

import (
	"github.com/jonascheng/gokv-demo/domain"
	"github.com/philippgille/gokv"
)

type fooV2Repository struct {
	store *gokv.Store
}

func NewFooV2Repository(db *gokv.Store) domain.FooV1Repository {
	return &fooV2Repository{db}
}
