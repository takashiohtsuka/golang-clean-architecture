package input

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
