package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Params struct {
	ServiceName string `json:"serviceName"`
	NewTag      string `json:"newTag"`
	Token       string `json:"token"`
}

func mustEnv(key string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		panic("env not found " + key)
	}
	return
}

func main() {
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}
	url := mustEnv("DRONE_URL")
	params := Params{
		ServiceName: mustEnv("DRONE_SERVICE_NAME"),
		NewTag:      mustEnv("DRONE_NEW_TAG"),
		Token:       mustEnv("DRONE_TOKEN"),
	}
	b, _ := json.Marshal(params)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(bodyBytes)

	if resp.StatusCode != http.StatusOK {
		log.Fatal(body)
	}
}
