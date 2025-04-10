package repository

import (
	"context"

	"github.com/gohantabeta/server-recruit-challenge-sample/model"
)

type SingerRepository interface {
	GetAll(ctx context.Context) ([]*model.Singer, error)
	Get(ctx context.Context, id model.SingerID) (*model.Singer, error)
	Add(ctx context.Context, singer *model.Singer) error
	Delete(ctx context.Context, id model.SingerID) error
}
