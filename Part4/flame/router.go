package flame

import (
	"log"
	"net/http"
	"strings"
)

//router里包含路由映射
type router struct {
	//不同方法对应的前缀树
	roots map[string]*node
	//存储路径对应的方法
	handlers map[string]HandlerFunc
}


//构造路由
func newRouter()*router{
	return &router{
		roots: make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

//解析pattern
func parsePattern(pattern string)[]string{
	vs := strings.Split(pattern,"/")
	parts:=make([]string,0)
	for _,item:=range vs{
		if item!=""{
			parts=append(parts,item)
			if item[0]=='*'{
				break
			}
		}
	}
	return parts
}


//添加路由映射
func(r *router)addRoute(method string,pattern string,handler HandlerFunc){
	parts:=parsePattern(pattern)

	key:=method+"-"+pattern

	_,ok:=r.roots[method]

	if !ok{
		r.roots[method]=&node{}
	}
	r.roots[method].insert(pattern,parts,0)
	r.handlers[key]=handler
	log.Printf("Route %4s - %s",method,pattern)
}

//获取路由方法
func(r *router)getRoute(method string,path string)(*node,map[string]string){
	searchParts := parsePattern(path)
	//：对应的参数
	params := make(map[string]string)
	root,ok:=r.roots[method]

	if !ok{
		return nil,nil
	}

	n:=root.search(searchParts,0)

	if n!=nil{
		parts:=parsePattern(n.pattern)
		for i,part := range parts {
			if part[0]==':'{
				params[part[1:]]=searchParts[i]
			}
			if part[0]=='*'&&len(part)>1{
				params[part[1:]]=strings.Join(searchParts[i:],"/")
				break
			}
		}
		return n,params
	}
	return nil,nil
}



//添加路由处理
func(r *router)handle(c *Context){
	n,params:=r.getRoute(c.Method,c.Path)
	if n!=nil{
		c.Params = params
		key:=c.Method+"-"+n.pattern
		r.handlers[key](c)
	}else{
		c.TEXT(http.StatusNotFound,"404 NOT FOUND",c.Path)
	}
}



