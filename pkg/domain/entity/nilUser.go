package entity

type NilUser struct{}

func (n *NilUser) IsNil() bool    { return true }
func (n *NilUser) GetID() uint     { return 0 }
func (n *NilUser) GetName() string { return "" }
func (n *NilUser) GetAge() string  { return "" }
