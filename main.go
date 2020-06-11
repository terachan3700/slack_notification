package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("MY_SLACK_URL")
	param := "payload={\"text\": \"ハゲハゲ\"}"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}
