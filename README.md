# goflame
//第一版：实现了HTTP基本功能：实现了http.Handler接口，统一了路由入口，可以进行日志或者异常处理的操作
//第二版：构造http.ResponseWriter和*http.Request的粒度太细，每次都要set很多请求头，因此考虑将其封装起来。
//同时框架需要引入中间件和动态路由，需要有位置存储它们，因此引入context
