package repository

import (
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/outputport"

	"gorm.io/gorm"
)

/*
repositoryの具象classとなるファイル
*/
type userRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewUserRepository(db *gorm.DB) outputport.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) WithTx(tx interface{}) outputport.UserRepository {
	return &userRepository{db: tx.(*gorm.DB)}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) Create(u *model.User) (*model.User, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
