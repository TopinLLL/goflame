package main

import (
	"goflame/flame"
	"net/http"
)




func main(){
	r:=flame.New()
	r.GET("/", func(c *flame.Context) {
		c.HTML(http.StatusOK,"<h1>hello go flame </h1>")
	})
	r.GET("/hello", func(c *flame.Context) {
		c.TEXT(http.StatusOK,"hello %s here is %s",c.Query("name"),c.Path)
	})
	r.POST("/login", func(c *flame.Context) {
		c.JSON(http.StatusOK,flame.H{
			"username":c.PostForm("username"),
			"password":c.PostForm("password"),
		})
	})
	r.Run(":8989")
}
