package domain

import (
	"context"
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
	Parent     *uint
	Category   uint
	Name       string
	Parameters *string
}

type BaseRepository interface {
	Create(ctx context.Context, base *Base) (*Base, error)
	Update(ctx context.Context, base *Base) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*Base, error)
	FindByCategory(ctx context.Context, id uint) ([]Base, error)
	FindByName(ctx context.Context, name string, category uint) (bool, error)
}
