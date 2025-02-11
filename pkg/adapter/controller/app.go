package controller

type AppController struct {
	User  interface{ User }
	Staff interface{ Staff }
	Role  interface{ Role }
}
