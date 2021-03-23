/*
Copyright 2019 BGBiao Ltd. All rights reserved.
@File   : main.go
@Desc   : None
*/
package main

import (
	"gin-jwt/controller"
	md "gin-jwt/middleware"
	"gin-jwt/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化db
	dbErr := model.InitMySQLCon()
	if dbErr != nil {
		panic(dbErr)
	}

	model.InitModel()
	defer model.DB.Close()

	// 初始化Gin实例
	router := gin.Default()
	v1 := router.Group("/apis/v1/")
	{
		v1.POST("/register", controller.RegisterUser)
		v1.POST("/login", controller.Login)
	}

	// secure v1
	sv1 := router.Group("/apis/v1/auth/")
	sv1.Use(md.JWTAuth())
	{
		sv1.GET("/time", controller.GetDataByTime)

	}
	router.Run(":8081")
}
