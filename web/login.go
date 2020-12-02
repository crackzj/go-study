package web

import (
	"golang/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Register 注册用户.
func Register(ctx *gin.Context) {
	orm.Insert()
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}


// func Login() {

// }
