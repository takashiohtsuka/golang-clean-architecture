package controller

import (
	"net/http"

	"golang-clean-architecture/pkg/domain/model"
)

// 使う側（controller）がusecaseに必要なインターフェースを定義
type UserUsecase interface {
	List(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}

type userController struct {
	userUsecase UserUsecase
}

type User interface {
	GetUsers(c Context) error
	CreateUser(c Context) error
}

func NewUserController(us UserUsecase) User {
	return &userController{us}
}

func (uc *userController) GetUsers(ctx Context) error {
	var u []*model.User

	u, err := uc.userUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc *userController) CreateUser(ctx Context) error {
	var params model.User

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	u, err := uc.userUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}
