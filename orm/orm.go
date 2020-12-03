package orm

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() {
	db := conn()
	err := db.AutoMigrate(Login{}, Role{})
	if err != nil {
		log.Fatal("create table false", err)
	}
	usr := Login{}
	usr.Usr = "lsy"
	usr.Role = 0
	usr.RoleValue = "管理员"
	usr.Pwd = "lsy0123"
	db.Create(&usr)
	role := Role{}
	role.Role = 0
	role.RoleValue = "管理员"
	db.Create(&role)
}
func conn() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "myview:lsy0123@tcp(127.0.0.1:3306)/myview?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql connect failed,", err)
	}
	return db
}

func Insert(i interface{}) {
	conn().Create(i)
}

func FindOne() {

}

func FindAll(tableName string,conditionField map[string]string,selectField ...string)(res interface{}) {
	if len(selectField) == 0{
		selectField = []string{"*"}
	}
	str := "SELECT "
	for _,v := range selectField{
		str += v + ", " 
	}
	str += " FROM "+ tableName
	if len(conditionField)>0{
		
		return
	}
	conn().Table(tableName).Select(selectField).Scan(res)
	return
}
func FindAllRaw(sql string,conditionField ...string)(res interface{}){
	conn().Raw(sql,)
	return
}
