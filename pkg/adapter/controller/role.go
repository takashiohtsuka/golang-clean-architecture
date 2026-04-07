package controller

import (
	"context"
	"net/http"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
)

// 使う側（controller）がusecaseに必要なインターフェースを定義
type RoleUsecase interface {
	List() (collection.Collection[entity.RoleEntity], error)
	Create(ctx context.Context, u *entity.Role) (*entity.Role, error)
}

type roleController struct {
	roleUsecase RoleUsecase
}

type Role interface {
	GetRoles(c Context) error
	CreateRole(c Context) error
}

func NewRoleController(r RoleUsecase) Role {
	return &roleController{r}
}

func (sc *roleController) GetRoles(ctx Context) error {
	roles, err := sc.roleUsecase.List()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, roles)
}

// curl -X POST -H "Content-Type: application/json" -d '{"name":"role3", "created_at": "2024-12-22T12:00:00Z", "updated_at": "2024-12-22T12:00:00Z"}' http://localhost:8080/roles
func (sc *roleController) CreateRole(ctx Context) error {
	var params entity.Role

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	u, err := sc.roleUsecase.Create(ctx.Request().Context(), &params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}
