package main

import (
  "github.com/gin-gonic/gin"
  )
var r *gin.Engine

const (
  testMode = false
)

func main() {
  
  r := initRoutes()
  
  if testMode {
    gin.SetMode(gin.TestMode)
    r.Run(":3000")
  } else {
    gin.SetMode(gin.ReleaseMode)
    r.Run()
  }
}


