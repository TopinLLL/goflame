package flame

import (
	"net/http"
)

//HandlerFunc是路由映射的处理函数
type HandlerFunc func(*Context)

//Engine是所有请求的统一处理器，里面包含路由表
type Engine struct {
	router *router
}

//初始化处理器
func New()*Engine{
	return &Engine{router:newRouter()}
}

//添加路由
func(e *Engine)addRoute(method string,pattern string,handler HandlerFunc){
	e.router.addRoute(method,pattern,handler)
}

//GET方法
func(e *Engine)GET(pattern string,handler HandlerFunc){
	e.addRoute("GET",pattern,handler)
}

//POST方法
func(e *Engine)POST(pattern string,handler HandlerFunc){
	e.addRoute("POST",pattern,handler)
}

//封装http.ListenAndServe
func(e *Engine)Run(addr string)(err error){
	err = http.ListenAndServe(addr,e)
	return err
}

//实现http.Handler------>http.ServeHTTP(ResponseWriter,*Request)
//原始的http.HandleFunc只能处理一个路由映射，但是实现了ServeHTTP就拥有了统一的控制入口，可以添加一些处理逻辑，比如日志或者异常处理
func(e *Engine)ServeHTTP(w http.ResponseWriter,req *http.Request){
	c:=newContext(w,req)
	e.router.handle(c)
}


