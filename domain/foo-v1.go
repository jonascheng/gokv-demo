package domain

import (
	"context"

	"github.com/gofrs/uuid"
)

type FooV1 struct {
	Version string `json:"version"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     uint8  `json:"age"`
}

func NewFooV1(name string, age uint) *FooV1 {
	return &FooV1{
		Version: "1.0",
		ID:      uuid.Must(uuid.NewV4()).String(),
		Name:    name,
		Age:     uint8(age),
	}
}

type FooV1Repository interface {
	Save(ctx context.Context, foo *FooV1) error
	GetByID(ctx context.Context, id string) (*FooV1, error)
	Update(ctx context.Context, foo *FooV1) error
	RemoveByID(ctx context.Context, id string) error
}
