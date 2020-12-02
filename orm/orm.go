package orm

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	db := conn()
	err := db.AutoMigrate(Login{})
	if err != nil {
		log.Fatal("create table false", err)
	}
}
func conn() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "crackzj:zjzuo0808@tcp(127.0.0.1:3306)/myview?charset=utf8mb4&parseTime=True&loc=Local",
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

func Insert() {
	usr := Login{}
	usr.Usr = "lsy"
	usr.Role = 0
	usr.RoleValue = "管理员"
	usr.Pwd = "lsy0123"
	ret := conn().Create(&usr)
	fmt.Println(ret)
}
