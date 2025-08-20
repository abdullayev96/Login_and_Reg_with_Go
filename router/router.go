package router

import (
    "github.com/gin-gonic/gin"

    "Login_Reg_Go/controllers"
    "Login_Reg_Go/repositories/db"
)



func SetupRouter() *gin.Engine {
    database := db.ConnectDatabase()

    r := gin.Default()
    r.POST("/register", func(c *gin.Context) { controllers.Register(c, database) })
    r.POST("/login", func(c *gin.Context) { controllers.Login(c, database) })

    return r
}