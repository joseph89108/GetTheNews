package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sync"
	"os"
	_ "github.com/joho/godotenv/autoload"
)

type GoogleArticle struct{
	Published string
	Title string
	Link string
}

type GoogleNews struct{
	Articles []GoogleArticle
}

func Google(c chan bool) {

	url := "https://google-news.p.rapidapi.com/v1/top_headlines?lang=zh&country=TW"

	req, _ := http.NewRequest("GET", url, nil)

	apiKey := os.Getenv("GOOGLE_RapidAPI_KEY")

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "google-news.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var goo GoogleNews
	err := json.Unmarshal(body,&goo)
	if err != nil {
		return 
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go goo.ToTxt(wg)
	go goo.ToHtml(wg)
	wg.Wait()
	c <- true
}