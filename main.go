package main

import (
	"golang/orm"
	"golang/web"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	orm.Init()
	router := gin.Default()

	s := &http.Server{
		Addr:           ":9000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	router.GET("/init", web.Init)
	s.ListenAndServe()
}
