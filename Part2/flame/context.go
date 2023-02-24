package flame

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//简化返回的JSON数据
type H map[string]interface{}

//封装context
type Context struct {
	//起始对象
	Writer http.ResponseWriter
	Req *http.Request

	//Request信息
	Path string
	Method string

	//Response信息
	StatusCode int
}

//构造context
func newContext(w http.ResponseWriter,req *http.Request)*Context{
	return &Context{
		Writer:     w,
		Req:        req,
		Path:       req.URL.Path,
		Method:     req.Method,
	}
}

//PostForm方法
func(c *Context)PostForm(key string)string{
	return c.Req.FormValue(key)
}

//Query方法，返回请求url中第一个匹配的key
func(c *Context)Query(key string)string{
	return c.Req.URL.Query().Get(key)
}

//填充ResponseWriter头部
func(c *Context)SetHeader(key string,value string){
	c.Writer.Header().Set(key,value)
}

//填写状态码
func(c *Context)Status(code int){
	c.StatusCode=code
	c.Writer.WriteHeader(code)
}


/*
	常见响应类型
	text/html 文本方式的html
	text/plain 纯文本
	text/xml  文本方式的xml
	application/json 数据以json形式编码
	application/xml 数据xml形式编码
	multipart/form-data 表单上传图片等附件必须使用该类型
*/


//构造text/plain响应方法
func(c *Context)TEXT(code int,format string,value ...interface{}){
	c.SetHeader("Content-Type","text/plain")
	c.Status(code)
	fmt.Println(value)
	c.Writer.Write([]byte(fmt.Sprintf(format,value...)))
}

//构造application/json响应方法
func(c *Context)JSON(code int,obj interface{}){
	c.SetHeader("Content-Type","application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj);err!=nil{
		http.Error(c.Writer,err.Error(),500)
	}
}

//构造纯数据响应方法
func(c *Context)Data(code int, data []byte){
	c.Status(code)
	c.Writer.Write(data)
}

//构造Html响应方法
func(c *Context)HTML(code int,html string){
	c.SetHeader("Content-Type","text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
