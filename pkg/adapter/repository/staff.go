package repository

import (
	"github.com/jinzhu/gorm"
	"golang-clean-architecture/pkg/adapter/mapper/staff"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"
)

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) repository.StaffRepository {
	return &staffRepository{db}
}

func (ur *staffRepository) FindAll(ormStaffs []*model.Staff) ([]*entity.Staff, error) {
	//構造体のポインタ配列の引数だが、配列の要素を指定しないと検索条件として機能していない
	//whereメソッドを実装してまたは条件追記も可能
	//err := ur.db.Where(model.Staff{Name: "ddd"}).Find(&u).Error
	err := ur.db.Find(&ormStaffs).Error

	if err != nil {
		return nil, err
	}

	//mapperを他モジュールにすることで何か依存性の解決はされているのか？
	//repository内でprivate関数でも問題ない処理でもある
	entityStaffs := make([]*entity.Staff, len(ormStaffs))
	for i := 0; i < len(ormStaffs); i++ {
		entityStaffs[i], _ = mapper.ToEntity(ormStaffs[i])
	}

	return entityStaffs, nil
}

func (ur *staffRepository) Create(s *entity.Staff) (*entity.Staff, error) {
	ormStaff, _ := mapper.ToOrmModel(s)
	if err := ur.db.Create(ormStaff).Error; err != nil {
		return nil, err
	}

	return s, nil
}

func (ur *staffRepository) Update(s *entity.Staff) (*entity.Staff, error) {
	ormStaff, _ := mapper.ToOrmModel(s)
	if err := ur.db.Model(&ormStaff).Update(ormStaff).Error; err != nil {
		return nil, err
	}

	return s, nil
}

func (ur *staffRepository) FindOne(conditions map[string]interface{}) (*entity.Staff, error) {
	modelStaff := &model.Staff{}

	if err := ur.db.Where(conditions).First(&modelStaff).Error; err != nil {
		return nil, err
	}

	entityStaff, _ := mapper.ToEntity(modelStaff)

	return entityStaff, nil
}

//TODO staffsとリレーションが組まれたテーブルを追加。staff_roles。 staff update()を実装し、update時には追加したテーブルにも更新をかけるようにする
//db.transactionが意図した動作をするか検証する
