package input

import "io"

type CreateWomanInput struct {
	CompanyID  uint
	Name       string
	Age        *int
	Birthplace *string
	BloodType  *string
	Hobby      *string
	IsActive   bool
	StoreIDs   []uint
}

type UpdateWomanInput struct {
	ID          uint
	Name        string
	Age         *int
	Birthplace  *string
	BloodType   *string
	Hobby       *string
	IsActive    bool
	ImageFile   io.Reader
	ImageKey    string
	ContentType string
}
