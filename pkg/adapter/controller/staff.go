package controller

import (
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/usecase"
	"net/http"
)

type staffController struct {
	staffUsecase usecase.Staff
}

type Staff interface {
	GetStaffs(c Context) error
	CreateStaff(c Context) error
}

func NewStaffController(st usecase.Staff) Staff {
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

// curl -X POST -H "Content-Type: application/json" -d '{"name":"James8", "age":"30", "is_active":"false", "created_at": "2024-12-22T12:00:00Z", "updated_at": "2024-12-22T12:00:00Z"}' http://localhost:8080/staffs
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
