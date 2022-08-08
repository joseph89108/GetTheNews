package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
)

type BingNewsValue struct{
	DatePublished string
	Name string
	Url string
	Description string
}

type BingNews struct{
	Value []BingNewsValue
}

func Bing(c chan bool) {

	url := "https://api.bing.microsoft.com/v7.0/news"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Ocp-Apim-Subscription-Key", "50e15c5f752b48d3a0594be991e2e8a4")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var news BingNews
	err := json.Unmarshal(body,&news)
	if err != nil {
		return 
	}
	fmt.Println(time.Now()," Bing:")
	for _, v := range news.Value {
		fmt.Println(v)
	}
	fmt.Println()
	c <- true
}