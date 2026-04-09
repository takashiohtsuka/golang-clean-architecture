package women

import (
	"fmt"
	"mime/multipart"

	"golang-clean-architecture/pkg/backend/usecase/input"
)

type Put struct {
	Name       string  `form:"name"        validate:"required,max=100"`
	Age        *int    `form:"age"         validate:"omitempty,min=18,max=60"`
	Birthplace *string `form:"birthplace"  validate:"omitempty,max=100"`
	BloodType  *string `form:"blood_type"  validate:"omitempty,oneof=A B O AB"`
	Hobby      *string `form:"hobby"       validate:"omitempty,max=200"`
	IsActive   bool    `form:"is_active"`
}

func (req *Put) ToInput(id uint, file multipart.File, header *multipart.FileHeader) input.UpdateWomanInput {
	i := input.UpdateWomanInput{
		ID:         id,
		Name:       req.Name,
		Age:        req.Age,
		Birthplace: req.Birthplace,
		BloodType:  req.BloodType,
		Hobby:      req.Hobby,
		IsActive:   req.IsActive,
	}

	if file != nil && header != nil {
		i.ImageFile = file
		i.ImageKey = fmt.Sprintf("women/%d/%s", id, header.Filename)
		i.ContentType = header.Header.Get("Content-Type")
	}

	return i
}
