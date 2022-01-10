package usecase

import (
	"context"

	"github.com/jonascheng/gokv-demo/domain"
)

type fooUsecase struct {
	repoFooV1 domain.FooV1Repository
	repoFooV2 domain.FooV2Repository
}

func NewFooUsecase(repoFooV1 domain.FooV1Repository, repoFooV2 domain.FooV2Repository) domain.FooUseCase {
	return &fooUsecase{
		repoFooV1: repoFooV1,
		repoFooV2: repoFooV2,
	}
}

func (this *fooUsecase) StoreFooV1(ctx context.Context, foo *domain.FooV1) (*domain.FooV1, error) {
	return this.repoFooV1.Save(ctx, foo)
}

func (this *fooUsecase) GetFooV1ByID(ctx context.Context, id string) (*domain.FooV1, error) {
	return this.repoFooV1.GetByID(ctx, id)
}

func (this *fooUsecase) StoreFooV2(ctx context.Context, foo *domain.FooV2) (*domain.FooV2, error) {
	return this.repoFooV2.Save(ctx, foo)
}

func (this *fooUsecase) GetFooV2ByID(ctx context.Context, id string) (*domain.FooV2, error) {
	return this.repoFooV2.GetByID(ctx, id)
}

func (this *fooUsecase) GetFooV2fromV1(ctx context.Context, id string) (*domain.FooV2, error) {
	return this.repoFooV2.GetByID(ctx, id)
}

func (this *fooUsecase) GetFooV1fromV2(ctx context.Context, id string) (*domain.FooV1, error) {
	return this.repoFooV1.GetByID(ctx, id)
}

func (this *fooUsecase) MigrateFooV1toV2(ctx context.Context, id string) (*domain.FooV2, error) {
	// load from v1
	fooV2, _ := this.GetFooV2fromV1(ctx, id)
	return this.StoreFooV2(ctx, fooV2)
}
