package main

import(
  "net/http"
  "fmt"
  "html/template"
  "strings"
)

func index(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])
  for k, v := range r.Form{
    fmt.Println("key:", k)
    fmt.Println("val:", strings.Join(v, ""))
  }
  fmt.Fprintf(w, "hello world")
}

func login(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  if r.Method == "GET"{
    t, _ := template.ParseFiles("login.gtpl")
    t.Execute(w, nil)
  } else {
    fmt.Println("username:", r.Form["username"])
    fmt.Println("password:", r.Form["password"])
  }
}


func main(){
  mux := http.NewServeMux()

  // ルーティングの設定
  mux.HandleFunc("/", index)
  mux.HandleFunc("/login", login)

  // server情報
  server := &http.Server{
    Addr:     "127.0.0.1:8080",
    Handler:  mux,
  }
  server.ListenAndServe()
}
