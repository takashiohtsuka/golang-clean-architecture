package interactor

import (
	"errors"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/outputport"
)

type StaffUsecase struct {
	staffRepository     outputport.StaffRepository
	dBRepository        outputport.DBRepository
	roleRepository      outputport.RoleRepository
	staffRoleRepository outputport.StaffRoleRepository
}

// コンストラクタ
func NewStaffUsecase(
	s outputport.StaffRepository,
	d outputport.DBRepository,
	r outputport.RoleRepository,
	sr outputport.StaffRoleRepository) *StaffUsecase {
	return &StaffUsecase{s, d, r, sr}
}

func (uu *StaffUsecase) List(modelStaff []*model.Staff) ([]*entity.Staff, error) {
	entityStaffs, err := uu.staffRepository.FindAll(modelStaff)
	if err != nil {
		return nil, err
	}

	return entityStaffs, nil
}

func (uu *StaffUsecase) Create(u *entity.Staff) (*entity.Staff, error) {
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

/**　staff のupdate時のみstaff_roleの紐付けができる という業務を想定 **/
func (uu *StaffUsecase) Update(staffId uint, roleId uint, updateStaffName string) (*entity.Staff, error) {

	staffConditions := make(map[string]interface{})
	staffConditions["ID"] = staffId

	entityStaff, err := uu.staffRepository.FindOne(staffConditions)

	if err != nil {
		//staffのrecord not foundのerrが返る
		return nil, err
	}

	//roleの存在チェック
	//存在しないroleだったら何かしらのエラーレスポンス
	roleConditions := make(map[string]interface{})
	roleConditions["ID"] = roleId
	_, err2 := uu.roleRepository.FindOne(roleConditions)

	if err2 != nil {
		//roleのrecord not foundのerrが返る
		return nil, err2
	}

	//TODO 次のタスク
	//TODO コード上で日付インスタンスをセットする処理の仕方を検討する
	//TODO CleanArchitectureでNullEntityの在り方を検討する

	staffRoleConditions := make(map[string]interface{})
	staffRoleConditions["staff_id"] = entityStaff.ID

	//upSert対象のstaffRoleEntity
	entityStaffRole, _ := uu.staffRoleRepository.FindOne(staffRoleConditions)

	var isUpdate = false

	if entityStaffRole == nil {
		entityStaffRole = &entity.StaffRole{
			StaffId: entityStaff.ID,
			RoleId:  roleId,
		}

	} else {
		isUpdate = true
		entityStaffRole.StaffId = entityStaff.ID
		entityStaffRole.RoleId = roleId
	}

	staff, err3 := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {

		//TODO ここのerrのハンドリングの仕方は問題ないのか調べる
		entityStaff.Name = updateStaffName
		s, err4 := uu.staffRepository.Update(entityStaff)

		if isUpdate {
			_, err4 = uu.staffRoleRepository.Update(entityStaffRole)

		} else {
			_, err4 = uu.staffRoleRepository.Create(entityStaffRole)
		}

		// do mailing
		// do logging
		// do another process
		return s, err4
	})
	//TODO callbackみたいな記述の仕方なので、要確認
	updatedStaff, ok := staff.(*entity.Staff)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err3 != nil {
		return nil, err3
	}

	return updatedStaff, nil
}
