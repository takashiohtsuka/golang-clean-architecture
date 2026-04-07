package entity

import (
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/valueobject"
)

type StaffEntity interface {
	IsNil() bool
	GetID() uint
	GetName() string
	GetAge() string
	GetIsActive() string
	GetRoles() collection.Collection[valueobject.Role]
}

type RoleEntity interface {
	IsNil() bool
	GetID() uint
	GetName() string
}

type StaffRoleEntity interface {
	IsNil() bool
	GetID() uint
	GetStaffId() uint
	GetRoleId() uint
}

type UserEntity interface {
	IsNil() bool
	GetID() uint
	GetName() string
	GetAge() string
}
