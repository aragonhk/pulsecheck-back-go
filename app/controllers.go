package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func initRoutes() *gin.Engine {
  r := gin.Default()
  r.GET("/", index)

  r.GET("/login", Login)
  
  r.GET("/search", Search)
  
	r.Static("/public", "./public")
	return r
}
func index(c *gin.Context){
	c.Redirect(http.StatusMovedPermanently, "https://pulsecheck-back-js.herokuapp.com/")
}

func Login(c *gin.Context){
	c.HTML(404, "404 - Page Not Found", nil)
}

func Search(c *gin.Context){
	searchResults, err := GetPersonByName( c.Query("lastname"), c.Query("firstname"), 0, 0)
    if err != "" {
      fmt.Print(err)
    } 
      fmt.Print(searchResults)
      c.JSON(200, gin.H{"oig": searchResults})
}