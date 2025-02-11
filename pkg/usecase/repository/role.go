package repository

import "golang-clean-architecture/pkg/domain/entity"

type RoleRepository interface {
	//変数名 引数 戻り値の型
	//スライス(可変調配列)で構造体のentityが引数
	Create(u *entity.Role) (*entity.Role, error)
	FindOne(map[string]interface{}) (*entity.Role, error)
}
