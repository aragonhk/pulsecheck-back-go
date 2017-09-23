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
  
  if gin.SetMode(gin.ReleaseMode) ; testMode {
    gin.SetMode(gin.TestMode)
    r.Run(":3000")
  }
}


