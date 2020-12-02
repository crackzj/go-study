package orm

import "gorm.io/gorm"

type NewModel interface{
	New() interface{}
}

type Login struct {
	gorm.Model
	Usr       string
	Pwd       string
	Role      int
	RoleValue string
}
