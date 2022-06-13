package db

import (
	"context"

	"github.com/Neko2h/shortener/internal/entity"
)

type LinkRepository interface {
	Get(context.Context, string) (*entity.LinkDB, error)
	New(context.Context, *entity.LinkDB) (*entity.LinkDB, error)
}
