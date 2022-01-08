package main

import (
	"context"
	"log"

	"github.com/philippgille/gokv/encoding"
	"github.com/philippgille/gokv/redis"

	"github.com/jonascheng/gokv-demo/domain"
	"github.com/jonascheng/gokv-demo/repository"
	"github.com/jonascheng/gokv-demo/usecase"
)

func runUsecaseWithRedis() {
	// run use cases with redis back storage
	options := redis.DefaultOptions // Address: "localhost:6379", Password: "", DB: 0
	options.Password = "supersecret"
	options.Codec = encoding.Gob

	// Create client
	client, err := redis.NewClient(options)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// repo
	repoFooV1 := repository.NewFooV1Repository(client)
	repoFooV2 := repository.NewFooV2Repository(client)

	// usecase
	usecaseFoo := usecase.NewFooUsecase(repoFooV1, repoFooV2)

	// Store, retrieve, print and delete a value
	usecaseRunner(usecaseFoo)
}

func main() {
	runUsecaseWithRedis()
}

func usecaseRunner(usecase domain.FooUseCase) {
	ctx := context.TODO()

	fooV1 := domain.NewFooV1("Jonas", 18)
	usecase.StoreFooV1(ctx, fooV1)
	log.Printf("Store fooV1: %v\n", fooV1)

	fooV2 := domain.NewFooV2("Jonas", 18, "Taipei City")
	usecase.StoreFooV2(ctx, fooV2)
	log.Printf("Store fooV2: %v\n", fooV2)

	fooV2FromV1, _ := usecase.GetFooV2fromV1(ctx, fooV1.ID)
	log.Printf("Get fooV2 from fooV1 to test adding new field: %v\n", fooV2FromV1)
}
