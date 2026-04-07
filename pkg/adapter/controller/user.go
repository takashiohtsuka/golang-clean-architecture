package controller

import (
	"net/http"

	requestUser "golang-clean-architecture/pkg/adapter/request/users"
	"golang-clean-architecture/pkg/usecase/inputport"
)

type userController struct {
	userUsecase inputport.UserUsecase
}

type User interface {
	GetUsers(c Context) error
	CreateUser(c Context) error
}

func NewUserController(us inputport.UserUsecase) User {
	return &userController{us}
}

func (uc *userController) GetUsers(ctx Context) error {
	var req requestUser.Get
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	input, err := req.ToInput()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	users, err := uc.userUsecase.List(input)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, users)
}

func (uc *userController) CreateUser(ctx Context) error {
	var req requestUser.Post

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	createdUser, err := uc.userUsecase.Create(ctx.Request().Context(), req.ToInput())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, createdUser)
}
