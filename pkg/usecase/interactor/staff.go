package interactor

import (
	"context"
	"errors"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/usecase/input"
	"golang-clean-architecture/pkg/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"
)

type StaffUsecase struct {
	staffRepository outputport.StaffRepository
	uow             outputport.UnitOfWork
	roleRepository  outputport.RoleRepository
}

// コンストラクタ
func NewStaffUsecase(
	staffRepository outputport.StaffRepository,
	uow outputport.UnitOfWork,
	roleRepository outputport.RoleRepository) *StaffUsecase {
	return &StaffUsecase{staffRepository, uow, roleRepository}
}

func (su *StaffUsecase) List(input input.ListStaffInput) (collection.Collection[entity.StaffEntity], error) {
	conditions := []query.Condition{}

	if input.Name != "" {
		conditions = append(conditions, query.Where("name", input.Name))
	}
	if input.Age != "" {
		conditions = append(conditions, query.Where("age", input.Age))
	}
	if input.IsActive != nil {
		conditions = append(conditions, query.Where("is_active", *input.IsActive))
	}

	return su.staffRepository.FindAll(conditions)
}

func (su *StaffUsecase) Create(ctx context.Context, input input.CreateStaffInput) (*entity.Staff, error) {
	staff := &entity.Staff{
		Name:     input.Name,
		Age:      input.Age,
		IsActive: input.IsActive,
	}

	var createdStaff *entity.Staff
	err := su.uow.Do(ctx, func() error {
		var e error
		createdStaff, e = su.staffRepository.Create(staff)

		// do mailing
		// do logging
		// do another process
		return e
	})
	if err != nil {
		return nil, err
	}

	return createdStaff, nil
}

/**　staff のupdate時のみstaff_roleの紐付けができる という業務を想定 **/
func (su *StaffUsecase) Update(ctx context.Context, input input.UpdateStaffInput) (bool, error) {

	staffEntity, err := su.staffRepository.FindOne([]query.Condition{
		query.Where("ID", input.StaffId),
	})

	if err != nil {
		return false, err
	}

	if staffEntity.IsNil() {
		return false, errors.New("staff not found")
	}

	entityStaff, ok := staffEntity.(*entity.Staff)
	if !ok {
		return false, errors.New("cast error")
	}

	// 指定されたroleIdが全て存在するか確認（IN句で一括取得）
	roleCollection, err := su.roleRepository.FindAll([]query.Condition{
		query.WhereIn("ID", input.RoleIds),
	})
	if err != nil {
		return false, err
	}
	if roleCollection.TotalCount() != len(input.RoleIds) {
		return false, errors.New("role not found")
	}

	err = su.uow.Do(ctx, func() error {
		entityStaff.Name = input.Name
		entityStaff.Age = input.Age
		entityStaff.IsActive = input.IsActive
		if _, err := su.staffRepository.Update(entityStaff); err != nil {
			return err
		}

		// staff_rolesを全削除して再insert
		if err := su.staffRepository.ReplaceRoles(entityStaff.ID, input.RoleIds); err != nil {
			return err
		}

		// do mailing
		// do logging
		// do another process
		return nil
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
