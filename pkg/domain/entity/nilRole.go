package entity

type NilRole struct{}

func (n *NilRole) IsNil() bool    { return true }
func (n *NilRole) GetID() uint     { return 0 }
func (n *NilRole) GetName() string { return "" }
