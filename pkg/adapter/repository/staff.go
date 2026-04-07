package repository

import (
	"errors"

	staffMapper "golang-clean-architecture/pkg/adapter/mapper/staff"
	staffRoleMapper "golang-clean-architecture/pkg/adapter/mapper/staffRole"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/valueobject"
	"golang-clean-architecture/pkg/infrastructure/model"
	"golang-clean-architecture/pkg/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) outputport.StaffRepository {
	return &staffRepository{db: db}
}

// staffWithRoleRow はJOIN結果の1行を表すスキャン用の構造体。
// staffが複数のroleを持つ場合、staff_idが同じ行が複数返される。
type staffWithRoleRow struct {
	StaffID       uint    `gorm:"column:staff_id"`
	StaffName     string  `gorm:"column:staff_name"`
	StaffAge      string  `gorm:"column:staff_age"`
	StaffIsActive string  `gorm:"column:staff_is_active"`
	RoleID        *uint   `gorm:"column:role_id"`
	RoleName      *string `gorm:"column:role_name"`
}

func (sr *staffRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.StaffEntity], error) {
	q := sr.db.Table("staffs s").
		Select(`
			s.id        AS staff_id,
			s.name      AS staff_name,
			s.age       AS staff_age,
			s.is_active AS staff_is_active,
			r.id        AS role_id,
			r.name      AS role_name
		`).
		Joins("LEFT JOIN staff_roles srl ON s.id = srl.staff_id").
		Joins("LEFT JOIN roles r ON srl.role_id = r.id").
		Where("s.deleted_at IS NULL")

	// conditionsのカラムはstaffsテーブルのものなので s. プレフィックスを付ける
	for _, c := range conditions {
		switch c.Kind {
		case query.KindWhere:
			q = q.Where("s."+c.Column+" = ?", c.Value)
		case query.KindWhereIn:
			q = q.Where("s."+c.Column+" IN ?", c.Value)
		case query.KindWhereBetween:
			q = q.Where("s."+c.Column+" BETWEEN ? AND ?", c.From, c.To)
		case query.KindWhereNotIn:
			q = q.Where("s."+c.Column+" NOT IN ?", c.Value)
		}
	}

	var rows []staffWithRoleRow
	if err := q.Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.StaffEntity](nil), err
	}

	if len(rows) == 0 {
		return collection.NewCollection[entity.StaffEntity](nil), nil
	}

	// JOIN結果をstaff_idでグループ化してentityに変換
	type staffAccumulator struct {
		staff *entity.Staff
		roles []valueobject.Role
	}
	staffOrder := []uint{}
	staffMap := map[uint]*staffAccumulator{}

	for _, row := range rows {
		if _, exists := staffMap[row.StaffID]; !exists {
			staffOrder = append(staffOrder, row.StaffID)
			staffMap[row.StaffID] = &staffAccumulator{
				staff: &entity.Staff{
					ID:       row.StaffID,
					Name:     row.StaffName,
					Age:      row.StaffAge,
					IsActive: row.StaffIsActive,
				},
			}
		}
		if row.RoleID != nil && row.RoleName != nil {
			staffMap[row.StaffID].roles = append(staffMap[row.StaffID].roles, valueobject.NewRole(*row.RoleName))
		}
	}

	items := make([]entity.StaffEntity, len(staffOrder))
	for i, id := range staffOrder {
		acc := staffMap[id]
		acc.staff.Roles = collection.NewCollection(acc.roles)
		items[i] = acc.staff
	}

	return collection.NewCollection(items), nil
}

func (sr *staffRepository) FindOne(conditions []query.Condition) (entity.StaffEntity, error) {
	modelStaff := &model.Staff{}

	if err := buildQuery(sr.db, conditions).First(&modelStaff).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &entity.NilStaff{}, nil
		}
		return &entity.NilStaff{}, err
	}

	entityStaff, _ := staffMapper.ToEntity(modelStaff)

	return entityStaff, nil
}

func (sr *staffRepository) Create(s *entity.Staff) (*entity.Staff, error) {
	ormStaff, _ := staffMapper.ToOrmModel(s)
	if err := sr.db.Create(ormStaff).Error; err != nil {
		return nil, err
	}

	return s, nil
}

func (sr *staffRepository) Update(s *entity.Staff) (*entity.Staff, error) {
	ormStaff, _ := staffMapper.ToOrmModel(s)
	if err := sr.db.Model(&ormStaff).Updates(ormStaff).Error; err != nil {
		return nil, err
	}

	return s, nil
}

// ReplaceRoles はstaff_rolesを全削除して指定のroleIdで再insertする。
func (sr *staffRepository) ReplaceRoles(staffId uint, roleIds []uint) error {
	if err := sr.db.Where("staff_id = ?", staffId).Delete(&model.StaffRole{}).Error; err != nil {
		return err
	}

	for _, roleId := range roleIds {
		ormStaffRole, _ := staffRoleMapper.ToOrmModel(&entity.StaffRole{
			StaffId: staffId,
			RoleId:  roleId,
		})
		if err := sr.db.Create(ormStaffRole).Error; err != nil {
			return err
		}
	}

	return nil
}
