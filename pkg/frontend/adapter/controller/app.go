package controller

type AppController struct {
	Store         interface{ Store }
	Woman         interface{ Woman }
	WomanDistrict interface{ WomanDistrict }
}
