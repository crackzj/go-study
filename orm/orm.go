package orm

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	db := Conn()
	err := db.AutoMigrate(Login{}, Role{})
	if err != nil {
		log.Fatal("create table false", err)
	}
	usr := Login{}
	if res := db.Where("usr=?", "lsy").First(&usr); res.RowsAffected == 0 {
		usr.Usr = "lsy"
		usr.Role = 0
		usr.RoleValue = "管理员"
		usr.Pwd = "lsy0123"
		db.Create(&usr)
	}
	role := Role{}
	if res := db.Where("role=?", 0).First(&role); res.RowsAffected == 0 {
		role.Role = 0
		role.RoleValue = "管理员"
		db.Create(&role)
	}
}
func Conn() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "myview:lsy0123@tcp(127.0.0.1:3306)/myview?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql Connect failed,", err)
	}
	return db
}

//InsertOne 插入单条数据.
func InsertOne(i interface{}) {
	Conn().Create(i)
}

//FindOne 查询单条数据.
func FindOne() {

}

//FindAll use raw sql.
// res 结构体.
// tableName 查询表名.
// selectField 查询字段 为空则查询所有.
// conditionField 条件字段.
func FindAll(res interface{}, tableName string, conditionField string, selectField ...string) interface{} {
	if len(selectField) == 0 {
		selectField = []string{"*"}
	}
	str := "SELECT "
	for _, v := range selectField {
		str += v + ", "
	}
	str += " FROM " + tableName
	if len(conditionField) > 0 {
		Conn().Table(tableName).Select(selectField).Where(conditionField).Scan(res)
		return res
	}
	Conn().Table(tableName).Select(selectField).Scan(res)
	return res
}
