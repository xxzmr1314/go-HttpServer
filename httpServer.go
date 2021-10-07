package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

func main()  {

	server := http.Server{
		Addr: "localhost:8081",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/healthz", healthz)
	err := server.ListenAndServe()
	if err != nil{
		fmt.Println("监听失败:",err.Error())
	}

}

func hello(res http.ResponseWriter, req *http.Request)  {
	//1、获取header并写入response
	header := req.Header
	for key,value := range header{
		res.Header().Set(key,strings.Join(value,","))
	}
	//2、获取本地VERSION
	version := os.Getenv("VERSION")
	res.Header().Set("VERSION",version)

	//3、记录客户端的ip,port,http返回码并输出
	ip,port,err :=net.SplitHostPort(req.RemoteAddr)
	code := httptest.NewRecorder().Code
	if err != nil{
		fmt.Println("获取ip失败:",err.Error())
	}
	fmt.Printf("IP:%v; Port: %v; StatusCode: %v",ip,port,code)//这里可以写入日志文件
	//页面打印返回
	fmt.Fprintln(res,"req中Header全部数据:",header)
	fmt.Fprintln(res,"本地VERSION:",version)
}

func healthz(res http.ResponseWriter, req *http.Request)  {
	//4、健康检测返回200
	res.WriteHeader(200)
}