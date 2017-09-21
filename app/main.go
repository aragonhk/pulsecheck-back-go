package main

import (
  "github.com/gin-gonic/gin"
  "github.com/aragonhk/pulsecheck-back-go/app/controllers"
  )
//var router *gin.Engine

const (
  testMode = false
)

func main() {

  router := initializeRoutes()
  
  if gin.SetMode(gin.ReleaseMode) ; testMode {
    gin.SetMode(gin.TestMode)
    router.Run(":3000")
  }
}

func initializeRoutes() *gin.Engine {
  r := gin.Default()

  r.GET("/", controllers.Index)

  r.GET("/login", controllers.Login)
  
  r.GET("/search", controllers.Search)
  
	r.Static("/public", "./public")

	return r
}