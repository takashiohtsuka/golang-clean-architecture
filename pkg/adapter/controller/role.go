package controller

import (
	"golang-clean-architecture/pkg/domain/entity"
	"net/http"
)

// 使う側（controller）がusecaseに必要なインターフェースを定義
type RoleUsecase interface {
	Create(u *entity.Role) (*entity.Role, error)
}

type roleController struct {
	roleUsecase RoleUsecase
}

type Role interface {
	CreateRole(c Context) error
}

func NewRoleController(r RoleUsecase) Role {
	return &roleController{r}
}

// curl -X POST -H "Content-Type: application/json" -d '{"name":"role3", "created_at": "2024-12-22T12:00:00Z", "updated_at": "2024-12-22T12:00:00Z"}' http://localhost:8080/roles
func (sc *roleController) CreateRole(ctx Context) error {
	var params entity.Role

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	u, err := sc.roleUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}
