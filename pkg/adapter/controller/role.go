package controller

import (
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/usecase/usecase"
	"net/http"
)

type roleController struct {
	roleUsecase usecase.Role
}

type Role interface {
	CreateRole(c Context) error
}

func NewRoleController(r usecase.Role) Role {
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
