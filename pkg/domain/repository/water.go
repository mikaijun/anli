package repository

import (
	"context"

	"github.com/mikaijun/aquagent/pkg/domain/model"
)

type WaterRepository interface {
	CreateWater(ctx context.Context, water *model.Water) (*model.Water, error)
	CreateRandomWaters(ctx context.Context) ([]*model.Water, error)
	GetWaters(ctx context.Context, userId int64, filter map[string]interface{}) ([]*model.Water, error)
	GetWater(ctx context.Context, waterId int64) (*model.Water, error)
	DeleteWater(ctx context.Context, waterId int64) error
}
