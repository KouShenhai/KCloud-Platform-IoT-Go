package main

import (
	"KCloud-Platform-Go/core"
	"encoding/json"
	"fmt"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var header = map[string]string{
		"Content-Type": "application/json",
	}
	body, err := core.SendGetRequest("http://jsonplaceholder.typicode.com/posts", &Post{}, nil, header)
	if err != nil {
		panic(err)
	}
	var b []Post
	_ = json.Unmarshal(body, &b)
	fmt.Println(b[0].Id)
}
