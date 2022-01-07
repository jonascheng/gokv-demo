package domain

import (
	"context"

	"github.com/gofrs/uuid"
)

type FooV2 struct {
	Version string `json:"version"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     uint8  `json:"age"`
	Address string `json:"address"`
}

func newFooV2(name string, age uint, address string) *FooV2 {
	return &FooV2{
		Version: "2.0",
		ID:      uuid.Must(uuid.NewV4()).String(),

		Name:    name,
		Age:     uint8(age),
		Address: address,
	}
}

type FooV2Repository interface {
	Save(ctx context.Context, foo *FooV2) error
	GetByID(ctx context.Context, id string) (*FooV2, error)
	Update(ctx context.Context, foo *FooV2) error
}
