package flame

import (
	"log"
	"net/http"
)

//router里包含路由映射
type router struct {
	handlers map[string]HandlerFunc
}


//构造路由
func newRouter()*router{
	return &router{handlers: map[string]HandlerFunc{}}
}

//添加路由映射
func(r *router)addRoute(method string,pattern string,handler HandlerFunc){
	log.Printf("Route %4s - %s",method,pattern)
	key:=method+"-"+pattern
	r.handlers[key]=handler
}

func(r *router)handle(c *Context){
	key:=c.Req.Method+"-"+c.Req.URL.Path
	if handler,ok:=r.handlers[key];ok{
		handler(c)
	}else{
		c.TEXT(http.StatusNotFound,"404 NOT FOUND",c.Path)
	}
}



