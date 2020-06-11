package main

import (
	"fmt"
	"net/http"
)

func HandleGithubPost(w http.ResponseWriter, r *http.Request){
	fmt.Println("githubからきたよ")
	sendMessage()
}

func main() {
	http.HandleFunc("/mention", HandleGithubPost)
	http.ListenAndServe(":8080", nil)
}


