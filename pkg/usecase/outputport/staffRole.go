package outputport

import "golang-clean-architecture/pkg/domain/entity"

type StaffRoleRepository interface {
	//変数名 引数 戻り値の型
	//スライス(可変調配列)で構造体のentityが引数
	Create(u *entity.StaffRole) (*entity.StaffRole, error)
	FindOne(map[string]interface{}) (*entity.StaffRole, error)
	Update(sr *entity.StaffRole) (*entity.StaffRole, error)
}
