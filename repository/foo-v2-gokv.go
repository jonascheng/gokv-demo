package repository

import (
	"context"
	"log"

	"github.com/jonascheng/gokv-demo/domain"
	"github.com/philippgille/gokv"
)

type fooV2Repository struct {
	store gokv.Store
}

func NewFooV2Repository(db gokv.Store) domain.FooV2Repository {
	return &fooV2Repository{db}
}

func (this *fooV2Repository) Save(ctx context.Context, foo *domain.FooV2) (*domain.FooV2, error) {
	// force version to 2.0
	foo.Version = "2.0"
	err := this.store.Set(foo.ID, foo)
	if err != nil {
		log.Println(err)
	}
	return foo, err
}

func (this *fooV2Repository) GetByID(ctx context.Context, id string) (*domain.FooV2, error) {
	retrievedVal := new(domain.FooV2)
	found, err := this.store.Get(id, retrievedVal)
	if err != nil {
		log.Println(err)
	}
	if !found {
		log.Println("Value not found")
	}
	return retrievedVal, err
}

func (this *fooV2Repository) Update(ctx context.Context, foo *domain.FooV2) (*domain.FooV2, error) {
	retrievedVal, err := this.GetByID(ctx, foo.ID)
	log.Println(retrievedVal)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// force version to 2.0
	foo.Version = "2.0"
	err = this.store.Set(foo.ID, foo)
	if err != nil {
		log.Println(err)
	}
	return foo, err
}

func (this *fooV2Repository) DeleteByID(ctx context.Context, id string) error {
	err := this.store.Delete(id)
	if err != nil {
		log.Println(err)
	}
	return err
}
