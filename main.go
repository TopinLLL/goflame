package main

import (
	"goflame/Part4/flame"
	"net/http"
)




func main(){
	r:=flame.New()
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *flame.Context) {
			c.HTML(http.StatusOK,"<h1>hello go flame </h1>")
		})
	}

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
	r.GET("/hello/:name", func(c *flame.Context) {
		// expect /hello/geektutu
		c.TEXT(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *flame.Context) {
		c.JSON(http.StatusOK, flame.H{"filepath": c.Param("filepath")})
	})
	r.Run(":9997")
}
