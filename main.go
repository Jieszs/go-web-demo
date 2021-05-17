package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"web-orm-demo/pkg/setting"
	"web-orm-demo/routers"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
