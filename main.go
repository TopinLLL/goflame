package main

import (
	"fmt"
	"goflame/flame"
	"net/http"
)




func main(){
	r:=flame.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		//http.ResponseWriter会将内容写入到http请求的Response中
		//Fprintf会将内容写入到实现了writer接口的类型中
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":9999")
}
