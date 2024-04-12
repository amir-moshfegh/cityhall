package repository

import (
	"context"
	"github.com/amir-moshfegh/cityhall/domain"
	"gorm.io/gorm"
)

type baseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) domain.BaseRepository {
	return baseRepository{
		db: db,
	}
}

func (b baseRepository) Create(ctx context.Context, base *domain.Base) (*domain.Base, error) {
	if err := b.db.WithContext(ctx).Create(&base).Error; err != nil {
		return nil, err
	}

	return base, nil
}

func (b baseRepository) Update(ctx context.Context, base *domain.Base) error {
	//TODO implement me
	panic("implement me")
}

func (b baseRepository) Delete(ctx context.Context, id uint) error {
	if err := b.db.WithContext(ctx).Delete(&domain.Base{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (b baseRepository) FindByID(ctx context.Context, id uint) (*domain.Base, error) {
	var base domain.Base
	if err := b.db.WithContext(ctx).First(&base, id).Error; err != nil {
		return nil, err
	}

	return &base, nil
}

func (b baseRepository) FindByCategory(ctx context.Context, id uint) ([]domain.Base, error) {
	var bases []domain.Base
	if err := b.db.WithContext(ctx).Where("category = ?", id).Find(&bases).Error; err != nil {
		return nil, err
	}
	return bases, nil
}
