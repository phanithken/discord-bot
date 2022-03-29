package main

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	/// declare from env file
	provider := os.Getenv("IP_PROVIDER")
	webhook := os.Getenv("DISCORD_WEBHOOK")

	res, err := http.Get(provider)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)
	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// notify
	body := map[string]string{"username": "PK Master", "content": "IP Changed: " + string(ip)}
	bodyJson, err := json.Marshal(body)
	_, err = http.Post(webhook, "application/json", bytes.NewBuffer(bodyJson))
	if err != nil {
		panic(err)
	}
}