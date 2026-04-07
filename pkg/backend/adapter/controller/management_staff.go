package controller

import (
	"net/http"

	requestManagementStaffs "golang-clean-architecture/pkg/backend/adapter/request/managementstaffs"
	"golang-clean-architecture/pkg/backend/usecase/inputport"
)

type managementStaffController struct {
	managementStaffUsecase inputport.ManagementStaffUsecase
}

type ManagementStaff interface {
	CreateManagementStaff(c Context) error
}

func NewManagementStaffController(u inputport.ManagementStaffUsecase) ManagementStaff {
	return &managementStaffController{u}
}

func (mc *managementStaffController) CreateManagementStaff(ctx Context) error {
	var req requestManagementStaffs.Post
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := mc.managementStaffUsecase.Create(ctx.Request().Context(), req.ToInput()); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, nil)
}
