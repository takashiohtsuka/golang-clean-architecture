package repository

import "golang-clean-architecture/pkg/domain/model"
import "golang-clean-architecture/pkg/domain/entity"

type StaffRepository interface {
	//変数名 引数 戻り値の型
	//スライス(可変調配列)で構造体のentityが引数
	FindAll(s []*model.Staff) ([]*entity.Staff, error)
	Create(u *entity.Staff) (*entity.Staff, error)
}
