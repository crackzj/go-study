package web

import (
	"fmt"
	"golang/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Msg  string        `json:"msg"`
	Code int           `json:"code"`
	Data []interface{} `json:"data"`
}

//Register 注册用户.
func Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

//Init 返回初始化数据
func Init(ctx *gin.Context) {
	type role struct {
		Role      int    `json:"role"`
		RoleValue string `json:"role_value"`
	}
	var rol role
	var res Response
	db := orm.Conn()
	// tx := conn().Find(&role)
	rows, _ := db.Raw("SELECT role,role_value from roles").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&rol.Role, &rol.RoleValue)
		res.Data = append(res.Data, rol)
	}
	res.Code = http.StatusOK
	res.Msg = "success"
	fmt.Println("rows=", res)
	ctx.JSON(http.StatusOK, res)
}

// func Login() {

// }
