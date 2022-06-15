package app

import (
	"fmt"
	"os"

	"github.com/Neko2h/shortener/internal/controller"
	"github.com/Neko2h/shortener/internal/store/cache/redis"
	"github.com/Neko2h/shortener/internal/store/db/postgres"
	"github.com/Neko2h/shortener/internal/usecase"
	"github.com/Neko2h/shortener/pkg/migrations"
	"github.com/labstack/echo/v4"
)

func NewApp() {

	store, err := postgres.NewPgDb(os.Getenv("PG_URL"))
	if err != nil {
		panic(err)
	}

	err = migrations.RunMigrations(os.Getenv("PG_MIGRATIONS_PATH"), os.Getenv("PG_URL"))
	if err != nil {
		panic(err)
	}

	cacheImp, err := redis.NewRedisClient(os.Getenv("REDIS_HOST"))
	if err != nil {
		panic(err)
	}

	LinkUsecase := usecase.NewLinkUsecase(postgres.NewLinkRepository(store.DB), redis.NewLinkCache(cacheImp.Client))

	//postgres.NewLinkRepository(store.DB)
	LinkController := controller.NewLinkController(LinkUsecase)

	e := echo.New()
	e.GET("/:hash", LinkController.GetLink)
	e.POST("/", LinkController.Create)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))

}
