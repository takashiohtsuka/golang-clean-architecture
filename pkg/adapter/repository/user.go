package repository

import (
	"golang-clean-architecture/pkg/adapter/mapper/user"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/infrastructure/model"
	"golang-clean-architecture/pkg/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"

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

func (ur *userRepository) WithTx(tx any) outputport.UserRepository {
	return &userRepository{db: tx.(*gorm.DB)}
}

func (ur *userRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.UserEntity], error) {
	var ormUsers []*model.User
	err := buildQuery(ur.db, conditions).Find(&ormUsers).Error
	if err != nil {
		return collection.NewCollection[entity.UserEntity](nil), err
	}

	if len(ormUsers) == 0 {
		return collection.NewCollection[entity.UserEntity](nil), nil
	}

	items := make([]entity.UserEntity, len(ormUsers))
	for i, ormUser := range ormUsers {
		items[i], _ = mapper.ToEntity(ormUser)
	}

	return collection.NewCollection(items), nil
}

func (ur *userRepository) Create(entityUser *entity.User) (*entity.User, error) {
	ormUser, _ := mapper.ToOrmModel(entityUser)
	if err := ur.db.Create(ormUser).Error; err != nil {
		return nil, err
	}

	return mapper.ToEntity(ormUser)
}
