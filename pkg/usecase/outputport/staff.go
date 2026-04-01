package outputport

import "golang-clean-architecture/pkg/domain/model"
import "golang-clean-architecture/pkg/domain/entity"

type StaffRepository interface {
	Where(column string, value interface{}) StaffRepository
	FindAll(s []*model.Staff) ([]*entity.Staff, error)
	Create(u *entity.Staff) (*entity.Staff, error)
	Update(s *entity.Staff) (*entity.Staff, error)
	FindOne(conditions map[string]interface{}) (*entity.Staff, error)
}
