package domain

import (
	"context"
	"gorm.io/gorm"
)

type Commission struct {
	gorm.DB
	Nosazi            string
	RequestNumber     string
	MeliNumber        string
	FileType          uint
	Name              string
	Family            string
	PhoneNumber       string
	Address           string
	DocumentType      uint
	CertificateNumber string
	CertificateDate   string
	OppositeMeter     float32
	ExpertName        uint
	Description       string
}

type CommissionRepository interface {
	Create(ctx context.Context, commission *Commission) (*Commission, error)
	Update(ctx context.Context, commission *Commission) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*Commission, error)
	FindAll(ctx context.Context) ([]*Commission, error)
}
