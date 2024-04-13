package validator

import (
	"github.com/amir-moshfegh/cityhall/dto"
	validation "github.com/go-ozzo/ozzo-validation"
)

func BaseCreateReq(req dto.BaseCreateReq) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required),
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Parameters, validation.NilOrNotEmpty),
		validation.Field(&req.Parent, validation.NilOrNotEmpty),
	)
}

func BaseUpdateReq(req dto.BaseUpdateReq) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.ID, validation.Required),
		validation.Field(&req.Category, validation.Required),
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Parameters, validation.NilOrNotEmpty),
		validation.Field(&req.Parent, validation.NilOrNotEmpty),
	)
}
