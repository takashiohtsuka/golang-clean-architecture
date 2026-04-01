package controller

import (
	requestStaff "golang-clean-architecture/pkg/adapter/request/staffs"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/model"
	"net/http"
)

// 使う側（controller）がusecaseに必要なインターフェースを定義
type StaffUsecase interface {
	List(u []*model.Staff) ([]*entity.Staff, error)
	Create(u *entity.Staff) (*entity.Staff, error)
	Update(staffId uint, roleId uint, updateStaffName string) (*entity.Staff, error)
}

type staffController struct {
	staffUsecase StaffUsecase
}

type Staff interface {
	GetStaffs(c Context) error
	CreateStaff(c Context) error
	UpdateStaff(c Context) error
}

func NewStaffController(st StaffUsecase) Staff {
	return &staffController{st}
}

func (sc *staffController) GetStaffs(ctx Context) error {
	var s []*model.Staff

	//構造体のポインタ変数配列に構造体を追加
	//sts := &model.Staff{ID: 1}
	//s = append(s, sts)

	entityStaffs, err := sc.staffUsecase.List(s)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, entityStaffs)
}

// curl -X POST -H "Content-Type: application/json" -d '{"name":"James10", "age":"30", "is_active":"false", "created_at": "2024-12-22T12:00:00Z", "updated_at": "2024-12-22T12:00:00Z"}' http://localhost:8080/staffs
func (sc *staffController) CreateStaff(ctx Context) error {
	var params entity.Staff

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	u, err := sc.staffUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}

// curl -X PUT -H "Content-Type: application/json" -d '{"staff_id":1, "role_id":1, "name":"update Bob"}' http://localhost:8080/staffs
func (sc *staffController) UpdateStaff(ctx Context) error {
	var putRequest requestStaff.Put

	if err := ctx.Bind(&putRequest); err != nil {
		return err
	}

	u, err := sc.staffUsecase.Update(putRequest.StaffId, putRequest.RoleId, putRequest.Name)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}
