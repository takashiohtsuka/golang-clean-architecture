package usecase

import (
	"errors"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"
)

type staffUsecase struct {
	staffRepository repository.StaffRepository
	dBRepository    repository.DBRepository
}

type Staff interface {
	List(u []*model.Staff) ([]*entity.Staff, error)
	Create(u *entity.Staff) (*entity.Staff, error)
}

// コンストラクタ
func NewStaffUsecase(r repository.StaffRepository, d repository.DBRepository) Staff {
	return &staffUsecase{r, d}
}

func (uu *staffUsecase) List(modelStaff []*model.Staff) ([]*entity.Staff, error) {
	entityStaffs, err := uu.staffRepository.FindAll(modelStaff)
	if err != nil {
		return nil, err
	}

	return entityStaffs, nil
}

func (uu *staffUsecase) Create(u *entity.Staff) (*entity.Staff, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		s, err := uu.staffRepository.Create(u)

		// do mailing
		// do logging
		// do another process
		return s, err
	})
	staff, ok := data.(*entity.Staff)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return staff, nil
}
