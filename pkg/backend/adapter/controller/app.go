package controller

type AppController struct {
	Company         interface{ Company }
	Store           interface{ Store }
	ManagementStaff interface{ ManagementStaff }
	Woman           interface{ Woman }
}
