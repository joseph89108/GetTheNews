package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sync"
	"os"
	_ "github.com/joho/godotenv/autoload"
)

type NewsArticle struct{
	PublishedAt string
	Title string
	Url string
	Description string
	UrlToImage string
}

type News struct{
	Articles []NewsArticle
}

func NewsAPI(c chan bool) {

	apiKey := os.Getenv("NEWS_API_KEY")

	url := "https://newsapi.org/v2/top-headlines?country=tw&apiKey=" + apiKey

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