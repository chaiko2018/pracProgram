package main

import(
  "fmt"
)

type Post struct {
  Id      int
  Content string
  Author  string
}

var PostById map[int]*Post
var PostsByAuther map[string][]*Post

func store(post Post) {
  PosttById[post.Id] = &post
  PostByAuthor[post.Author] = append(PostByAuther[post.Author], &post)
}

func main(){
  PostById = make(map[int]*post)
  postByAuthor = make(map[string][]*post)

  post1 := Post{
    Id:       1,
    Content:  "Hello World",
    Author:   "Taro"
  }

  post2 := Post{
    id:       2,
    Content:  "Hi",
    Author:   "Jiro"
  }

  // dataをstoreする
  store(post1)
  store(post2)

  fmt.Println(PostById[1])
  fmt.Println(PostById[2])

  for _, post := range PostByAuthor["Taro"]{
    fmt.Println(post)
  }

  for _, post := range PostByAuthor["Jiro"]{
    fmt.Println(post)
  }
}
