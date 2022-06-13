package usecase

import (
	"context"

	"github.com/Neko2h/shortener/internal/entity"
)

type LinkUsecase interface {
	Get(context.Context, string) (*entity.Link, error)
	New(context.Context, *entity.Link) (*entity.Link, error)
}
