package main

import (
	"fmt"

	"github.com/coralproject/go-atoll"
)

type Post struct {
	Body string `json:"body"`
	User string `json:"user"`
}

func main() {
	post := &Post{
		Body: "yoyoy",
		User: "hey",
	}

	client := atoll.Client{
		BaseURL: "http://localhost:5001",
	}

	atollResp, err := client.Post(post, "/pipelines/score_post")
	if err != nil {
		panic(err)
	}
	fmt.Println(atollResp.Results)
	fmt.Println(atollResp.Response)
}
