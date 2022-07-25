package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	var count int
	//此处可以用值也可以用指针，但涉及到传递的话还是用指针，应为在传参时使用值拷贝，相当于不是一把锁了
	mux := sync.Mutex{}
	http.HandleFunc("/count", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("%p\n", &mux)
		mux.Lock()
		count++
		fmt.Fprintf(writer, "hello,%s,%d", request.URL.Path, count)
		mux.Unlock()
	})
	/*	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL.Path, request.URL.Path[1:])
		http.ServeFile(writer, request, request.URL.Path[1:])
	})*/
	//example,  http://localhost:8081/index.html
	//相对gopath
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":8081", nil))

	//restful api    https://tutorialedge.net/golang/creating-restful-api-with-golang/

}
