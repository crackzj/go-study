package web

import (
	"golang/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Register 注册用户.
func Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

//Init 返回初始化数据
func Init(ctx *gin.Context) {
	// roles := orm.Role{}
	// orm.FindAll()
	

}

// func Login() {

// }
