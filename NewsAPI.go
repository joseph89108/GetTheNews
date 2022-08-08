package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sync"
)

type NewsArticle struct{
	PublishedAt string
	Title string
	Url string
	Description string
}

type News struct{
	Articles []NewsArticle
}

func NewsAPI(c chan bool) {

	url := "https://newsapi.org/v2/top-headlines?country=tw&apiKey=fd7312e58db040d58caeead30d52a764"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var news News
	err := json.Unmarshal(body,&news)
	if err != nil {
		return 
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go news.ToTxt(wg)
	go news.ToHtml(wg)
	wg.Wait()
	c <- true
}