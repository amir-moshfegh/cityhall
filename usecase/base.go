package usecase

import (
	"context"
	"github.com/amir-moshfegh/cityhall/domain"
	"github.com/amir-moshfegh/cityhall/dto"
	"gorm.io/gorm"
	"time"
)

type baseUseCase struct {
	repo    domain.BaseRepository
	timeout time.Duration
}

func NewBaseUseCase(repo domain.BaseRepository, timeoute time.Duration) dto.BaseUseCase {
	return &baseUseCase{
		repo:    repo,
		timeout: timeoute,
	}
}

// Create is a method to create a new base entity.
// It takes a context and a BaseCreateReq as input and returns a BaseCreateRes and an error.
func (b baseUseCase) Create(ctx context.Context, base *dto.BaseCreateReq) (*dto.BaseCreateRes, error) {
	// Create a new context with a timeout
	c, cancel := context.WithTimeout(ctx, b.timeout)
	defer cancel()

	// Call the repo's Create method to create a new base entity
	bs, err := b.repo.Create(c, &domain.Base{
		Parent:     base.Parent,
		Category:   base.Category,
		Name:       base.Name,
		Parameters: base.Parameters,
	})

	if err != nil {
		return nil, err
	}

	// Return the created base entity in the response
	return &dto.BaseCreateRes{
		ID:         bs.ID,
		Parent:     bs.Parent,
		Category:   bs.Category,
		Name:       bs.Name,
		Parameters: bs.Parameters,
	}, nil
}

func (b baseUseCase) Update(ctx context.Context, base *dto.BaseUpdateReq) error {
	c, cancel := context.WithTimeout(ctx, b.timeout)
	defer cancel()

	// Call the repo's Update method to update a base entity
	return b.repo.Update(c, &domain.Base{
		Model:      gorm.Model{ID: base.ID},
		Parent:     base.Parent,
		Category:   base.Category,
		Name:       base.Name,
		Parameters: base.Parameters,
	})
}

func (b baseUseCase) Delete(ctx context.Context, id uint) error {
	c, cancel := context.WithTimeout(ctx, b.timeout)
	defer cancel()

	// Call the repo's Delete method to delete a base entity
	return b.repo.Delete(c, id)
}

func (b baseUseCase) FindByID(ctx context.Context, id uint) (*dto.BaseFindByIDRes, error) {
	c, cancel := context.WithTimeout(ctx, b.timeout)
	defer cancel()

	// Call the repo's FindByID method to find a base entity
	base, err := b.repo.FindByID(c, id)
	if err != nil {
		return nil, err
	}
	return &dto.BaseFindByIDRes{
		ID:         base.ID,
		Parent:     base.Parent,
		Category:   base.Category,
		Name:       base.Name,
		Parameters: base.Parameters,
	}, nil
}

func (b baseUseCase) FindByCategory(ctx context.Context, id uint) ([]dto.BaseFindByCategoryRes, error) {
	c, cancel := context.WithTimeout(ctx, b.timeout)
	defer cancel()

	// Call the repo's FindByCategory method to find a base entity
	bases, err := b.repo.FindByCategory(c, id)
	if err != nil {
		return nil, err
	}

	var newBases []dto.BaseFindByCategoryRes
	for _, base := range bases {
		newBases = append(newBases, dto.BaseFindByCategoryRes{
			ID:         base.ID,
			Parent:     base.Parent,
			Category:   base.Category,
			Name:       base.Name,
			Parameters: base.Parameters,
		})
	}

	return newBases, nil
}
