package main

import (
  "fmt"
  "log"
  "net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request)  {
  fmt.Println("Inside HelloServer handler")
  fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}
func HFunc(w http.ResponseWriter, req *http.Request) {
  ...
}

func main() {
  http.HandleFunc("/", HelloServer)
  /*
  如果你需要使用安全的https连接，使用http.ListenAndServeTLS()代替http.ListenAndServe()
  http.HandleFunc("/", Hfunc)中的HFunc是一个处理函数
  */
  err := http.ListenAndServe("localhost:5000", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err.Error())
  }
}
