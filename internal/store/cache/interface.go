package cache

import (
	"context"

	"github.com/Neko2h/shortener/internal/entity"
)

type LinkCache interface {
	Get(context.Context, string) (*entity.LinkDB, error)
	Set(context.Context, *entity.LinkDB) error
}
