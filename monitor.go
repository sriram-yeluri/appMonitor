package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	colorReset  string = "\033[0m"
	colorRed    string = "\033[31m"
	colorGreen  string = "\033[32m"
	colorYellow string = "\033[33m"
	colorBlue   string = "\033[34m"
	colorPurple string = "\033[35m"
	colorCyan   string = "\033[36m"
	colorWhite  string = "\033[37m"
)

var tools = make(map[string]interface{})

func checkUrl(url string) {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: transCfg,
	}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(string(colorRed), url+" ---> "+"[Offline]", string(colorReset))
	} else if resp.StatusCode == 200 {
		fmt.Println(string(colorGreen), url+" ---> "+"[ACTIVE]", string(colorReset))
	}
}

func fetchValue(value interface{}) {
	switch value.(type) {
	case string:
		//fmt.Printf("%v is an interface \n ", value)
		checkUrl(value.(string))
	case map[string]interface{}:
		//fmt.Printf("%v is a map \n ", value)
		for _, v := range value.(map[string]interface{}) { // use type assertion to loop over map[string]interface{}
			fetchValue(v)
		}
	default:
		fmt.Println("Undefined interface type")
	}
}

func readJson(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(data), &tools)
	for k, v := range tools {
		fmt.Println(string(colorCyan), "----------------", k, "----------------", string(colorReset))
		fetchValue(v)
	}
}

func main() {
	fmt.Println(string(colorPurple), "BEGIN ........ Monitor ........... BEGIN", string(colorReset))
	readJson("urls.json")
	fmt.Println(string(colorPurple), "END ........ Monitor ........... END ", string(colorReset))
}
