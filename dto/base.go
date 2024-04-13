package dto

import "context"

type BaseCreateReq struct {
	Parent     *uint   `json:"parent"`
	Category   uint    `json:"category"`
	Name       string  `json:"name"`
	Parameters *string `json:"parameters"`
}

type BaseCreateRes struct {
	ID         uint    `json:"id"`
	Parent     *uint   `json:"parent"`
	Category   uint    `json:"category"`
	Name       string  `json:"name"`
	Parameters *string `json:"parameters"`
}

type BaseUpdateReq struct {
	ID         uint    `json:"id"`
	Parent     *uint   `json:"parent"`
	Category   uint    `json:"category"`
	Name       string  `json:"name"`
	Parameters *string `json:"parameters"`
}

type BaseFindByIDRes struct {
	ID         uint    `json:"id"`
	Parent     *uint   `json:"parent"`
	Category   uint    `json:"category"`
	Name       string  `json:"name"`
	Parameters *string `json:"parameters"`
}

type BaseFindByCategoryRes struct {
	ID         uint    `json:"id"`
	Parent     *uint   `json:"parent"`
	Category   uint    `json:"category"`
	Name       string  `json:"name"`
	Parameters *string `json:"parameters"`
}

type BaseUseCase interface {
	Create(ctx context.Context, base *BaseCreateReq) (*BaseCreateRes, error)
	Update(ctx context.Context, base *BaseUpdateReq) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*BaseFindByIDRes, error)
	FindByCategory(ctx context.Context, id uint) ([]BaseFindByCategoryRes, error)
}
