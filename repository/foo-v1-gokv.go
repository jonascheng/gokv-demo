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

func (this *fooV1Repository) Save(ctx context.Context, foo *domain.FooV1) (*domain.FooV1, error) {
	// force version to 1.0
	foo.Version = "1.0"
	err := this.store.Set(foo.ID, foo)
	if err != nil {
		log.Println(err)
	}
	return foo, err
}

func (this *fooV1Repository) GetByID(ctx context.Context, id string) (*domain.FooV1, error) {
	retrievedVal := new(domain.FooV1)
	found, err := this.store.Get(id, retrievedVal)
	if err != nil {
		log.Println(err)
	}
	if !found {
		log.Println("Value not found")
	}
	return retrievedVal, err
}

func (this *fooV1Repository) Update(ctx context.Context, foo *domain.FooV1) (*domain.FooV1, error) {
	retrievedVal, err := this.GetByID(ctx, foo.ID)
	log.Println(retrievedVal)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// force version to 1.0
	foo.Version = "1.0"
	err = this.store.Set(foo.ID, foo)
	if err != nil {
		log.Println(err)
	}
	return foo, err
}

func (this *fooV1Repository) DeleteByID(ctx context.Context, id string) error {
	err := this.store.Delete(id)
	if err != nil {
		log.Println(err)
	}
	return err
}
