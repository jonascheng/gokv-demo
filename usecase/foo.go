package usecase

import (
	"context"

	"github.com/jonascheng/gokv-demo/domain"
)

type fooUsecase struct {
	repoFooV1 domain.FooV2Repository
	repoFooV2 domain.FooV1Repository
}

func NewFooUsecase(repoFooV1 domain.FooV2Repository, repoFooV2 domain.FooV1Repository) domain.FooUseCase {
	return &fooUsecase{
		repoFooV1: repoFooV1,
		repoFooV2: repoFooV2,
	}
}

func (this *fooUsecase) StoreFooV1(ctx context.Context, foo *domain.FooV1) error {
	return nil
}

func (this *fooUsecase) GetFooV1ByID(ctx context.Context, id string) (*domain.FooV1, error) {
	return nil, nil
}

func (this *fooUsecase) StoreFooV2(ctx context.Context, foo *domain.FooV2) error {
	return nil
}

func (this *fooUsecase) GetFooV2ByID(ctx context.Context, id string) (*domain.FooV2, error) {
	return nil, nil
}

func (this *fooUsecase) GetFooV1toV2(ctx context.Context, id string) (*domain.FooV2, error) {
	return nil, nil
}

func (this *fooUsecase) GetFooV2toV1(ctx context.Context, id string) (*domain.FooV1, error) {
	return nil, nil
}

func (this *fooUsecase) MigrateFooV1toV2(ctx context.Context, id string) (*domain.FooV2, error) {
	return nil, nil
}
