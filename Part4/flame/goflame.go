package flame

import (
	"net/http"
)

//HandlerFunc是路由映射的处理函数
type HandlerFunc func(*Context)

//分组
type RouterGroup struct {
	prefix string //路径前缀
	middlewares []HandlerFunc //当前路由组支持的中间件
	parent *RouterGroup //当前路由组的父路由组
	engine *Engine //所有的路由组共享一个engine实例
}

//Engine是所有请求的统一处理器，里面包含路由表
type Engine struct {
	*RouterGroup //engine本身也当成一个路由组
	router *router
	groups []*RouterGroup //记录全部路由组
}

//初始化处理器
func New()*Engine{
	//首先创建带有router的engine
	engine:=&Engine{router: newRouter()}
	//所有路由组共享同一个engine
	engine.RouterGroup=&RouterGroup{engine: engine}
	//engine是根路由组
	engine.groups=[]*RouterGroup{engine.RouterGroup}
	return engine
}

func(group *RouterGroup)Group(prefix string)*RouterGroup{
	engine:=group.engine
	newGroup:=&RouterGroup{
		prefix:      group.prefix+prefix,
		parent:      group,
		engine:      engine,
	}
	engine.groups=append(engine.groups,newGroup)
	return newGroup
}

//添加路由
func(group *RouterGroup)addRoute(method string,comp string,handler HandlerFunc){
	pattern:=group.prefix+comp
	group.engine.router.addRoute(method,pattern,handler)
}

//GET方法
func(group *RouterGroup)GET(pattern string,handler HandlerFunc){
	group.addRoute("GET",pattern,handler)
}

//POST方法
func(group *RouterGroup)POST(pattern string,handler HandlerFunc){
	group.addRoute("POST",pattern,handler)
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


