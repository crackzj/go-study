package orm

import "gorm.io/gorm"

//Role model
type Role struct {
	gorm.Model
	Role      int `gorm:"primaryKey"`
	RoleValue string
}

//Login model
type Login struct {
	gorm.Model
	Usr       string `gorm:"primaryKey"`
	Pwd       string
	Role      int
	RoleValue string
}
