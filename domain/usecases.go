package domain

import "context"

type FooUseCase interface {
	StoreFooV1(ctx context.Context, foo *FooV1) error
	GetFooV1ByID(ctx context.Context, id string) (*FooV1, error)
	StoreFooV2(ctx context.Context, foo *FooV2) error
	GetFooV2ByID(ctx context.Context, id string) (*FooV2, error)
	GetFooV2fromV1(ctx context.Context, id string) (*FooV2, error)
	GetFooV1fromV2(ctx context.Context, id string) (*FooV1, error)
	MigrateFooV1toV2(ctx context.Context, id string) (*FooV2, error)
}
