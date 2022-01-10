package main

import (
	"context"
	"log"

	"github.com/jonascheng/gokv-demo/domain"
	"github.com/jonascheng/gokv-demo/repository"
	"github.com/jonascheng/gokv-demo/usecase"
	"github.com/philippgille/gokv/encoding"
	"github.com/philippgille/gokv/postgresql"
	"github.com/philippgille/gokv/redis"
)

func runUsecaseWithRedis() {
	// run use cases with redis back storage
	options := redis.DefaultOptions  // Address: "localhost:6379",
	options.Password = "supersecret" // Password: "", DB: 0,
	options.Codec = encoding.Gob     // Codec: encoding.JSON,

	// create client
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

	usecaseRunner(usecaseFoo)
}

func runUsecaseWithPGSQL() {
	// run use cases with pgsql back storage
	options := postgresql.DefaultOptions                                               // TableName: "Item", MaxOpenConnections: 100,
	options.ConnectionURL = "postgres://postgres:supersecret@/kv_demo?sslmode=disable" // ConnectionURL: "postgres://postgres@/" + defaultDBname + "?sslmode=disable",
	options.Codec = encoding.Gob                                                       // Codec: encoding.JSON,

	// create client
	client, err := postgresql.NewClient(options)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// repo
	repoFooV1 := repository.NewFooV1Repository(client)
	repoFooV2 := repository.NewFooV2Repository(client)

	// usecase
	usecaseFoo := usecase.NewFooUsecase(repoFooV1, repoFooV2)

	usecaseRunner(usecaseFoo)
}

func main() {
	runUsecaseWithRedis()
	runUsecaseWithPGSQL()
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
	log.Printf("Get fooV2 from fooV1 to test extra new field: %v\n", fooV2FromV1)

	fooV1FromV2, _ := usecase.GetFooV1fromV2(ctx, fooV2.ID)
	log.Printf("Get fooV1 from fooV2 to test missing field: %v\n", fooV1FromV2)

	fooV2FromV1, _ = usecase.MigrateFooV1toV2(ctx, fooV1.ID)
	log.Printf("Migrate fooV1 to fooV2: %v\n", fooV2FromV1)
}
