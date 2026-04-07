package controller

import (
	requestStaff "golang-clean-architecture/pkg/adapter/request/staffs"
	"golang-clean-architecture/pkg/usecase/inputport"
	"net/http"
)

type staffController struct {
	staffUsecase inputport.StaffUsecase
}

type Staff interface {
	GetStaffs(c Context) error
	CreateStaff(c Context) error
	UpdateStaff(c Context) error
}

func NewStaffController(st inputport.StaffUsecase) Staff {
	return &staffController{st}
}

func (sc *staffController) GetStaffs(ctx Context) error {
	var req requestStaff.Get
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	input, err := req.ToInput()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	entityStaffs, err := sc.staffUsecase.List(input)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, entityStaffs)
}

// curl -X POST -H "Content-Type: application/json" -d '{"name":"James10", "age":"30", "is_active":"false", "created_at": "2024-12-22T12:00:00Z", "updated_at": "2024-12-22T12:00:00Z"}' http://localhost:8080/staffs
func (sc *staffController) CreateStaff(ctx Context) error {
	var req requestStaff.Post

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	createdStaff, err := sc.staffUsecase.Create(ctx.Request().Context(), req.ToInput())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, createdStaff)
}

// curl -X PUT -H "Content-Type: application/json" -d '{"staff_id":1, "role_ids":[1,2], "name":"update Bob", "age":"30", "is_active":"true"}' http://localhost:8080/staffs
func (sc *staffController) UpdateStaff(ctx Context) error {
	var req requestStaff.Put

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	ok, err := sc.staffUsecase.Update(ctx.Request().Context(), req.ToInput())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, ok)
}
