package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	url := mustEnv("PLUGIN_URL")
	params := Params{
		ServiceName: mustEnv("PLUGIN_SERVICE_NAME"),
		NewTag:      mustEnv("PLUGIN_NEW_TAG"),
		Token:       mustEnv("PLUGIN_TOKEN"),
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
		fmt.Println("code not 200: ", resp.StatusCode)
		fmt.Println("url: ", url[1:])
		fmt.Println("toke: ", params.Token[1:])
		fmt.Println("params: ", params)
		log.Fatal(body)
	}
	fmt.Println(body)
}
