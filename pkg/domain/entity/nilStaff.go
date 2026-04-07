package entity

import (
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/valueobject"
)

type NilStaff struct{}

func (n *NilStaff) IsNil() bool                                        { return true }
func (n *NilStaff) GetID() uint                                         { return 0 }
func (n *NilStaff) GetName() string                                     { return "" }
func (n *NilStaff) GetAge() string                                      { return "" }
func (n *NilStaff) GetIsActive() string                                 { return "" }
func (n *NilStaff) GetRoles() collection.Collection[valueobject.Role]   { return collection.NewCollection[valueobject.Role](nil) }
