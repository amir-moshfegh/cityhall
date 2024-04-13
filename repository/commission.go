package repository

import (
	"context"
	"github.com/amir-moshfegh/cityhall/domain"
	"gorm.io/gorm"
)

type commissionRepository struct {
	repo *gorm.DB
}

func NewCommissionRepository(repo *gorm.DB) domain.CommissionRepository {
	return &commissionRepository{repo: repo}

}

func (c commissionRepository) Create(ctx context.Context, commission *domain.Commission) (*domain.Commission, error) {
	if err := c.repo.WithContext(ctx).Create(&commission).Error; err != nil {
		return nil, err
	}

	return commission, nil
}

func (c commissionRepository) Update(ctx context.Context, commission *domain.Commission) error {
	return c.repo.WithContext(ctx).Save(&commission).Error
}

func (c commissionRepository) Delete(ctx context.Context, id uint) error {
	return c.repo.WithContext(ctx).Delete(&domain.Commission{}, id).Error
}

func (c commissionRepository) FindByID(ctx context.Context, id uint) (*domain.Commission, error) {
	var commission domain.Commission
	if err := c.repo.WithContext(ctx).First(&commission, id).Error; err != nil {
		return nil, err
	}

	return &commission, nil
}

func (c commissionRepository) FindAll(ctx context.Context) ([]*domain.Commission, error) {
	var commissions []*domain.Commission
	if err := c.repo.WithContext(ctx).Find(&commissions).Error; err != nil {
		return nil, err
	}

	return commissions, nil
}
