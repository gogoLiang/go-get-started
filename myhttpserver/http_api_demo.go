package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// http.ListenAndServe(":9999", nil)
// 第二个参数是Handler, 需要实现ServeHTTP方法，所有http请求都经过他处理
// 可用于统一请求分发器
func main() {
	log.SetOutput(os.Stdout)
	ds := New()
	ds.Register("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})
	ds.Open(":9999")
}

type DServlet struct {
	HandlerMap map[string]func(w http.ResponseWriter, req *http.Request)
}

func (ds *DServlet) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	fmt.Println("request", path)
	handler := ds.HandlerMap[path]
	if handler == nil {
		e := fmt.Errorf("path %s not found", path)
		//http.NotFound(w, req)
		http.Error(w, e.Error(), http.StatusNotFound)
		return
	}
	handler(w, req)
}
func (ds *DServlet) Register(path string, f func(w http.ResponseWriter, req *http.Request)) {
	ds.HandlerMap[path] = f
}

func (ds *DServlet) Open(address string) {
	fmt.Println("open http server", address)
	err := http.ListenAndServe(address, ds)
	if err != nil {
		return
	}
	fmt.Println("http server down")
}

func New() *DServlet {
	return &DServlet{
		HandlerMap: make(map[string]func(w http.ResponseWriter, req *http.Request)),
	}
}
