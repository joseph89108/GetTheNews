package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
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

	url := "https://google-news.p.rapidapi.com/v1/top_headlines?lang=en&country=TW"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "3acfdcd9c5msh937236f595a405cp1c0cc2jsn816014f178ff")
	req.Header.Add("X-RapidAPI-Host", "google-news.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var goo GoogleNews
	err := json.Unmarshal(body,&goo)
	if err != nil {
		return 
	}
	fmt.Println(time.Now()," Google:")
	for _, v := range goo.Articles {
		fmt.Println(v)
	}
	fmt.Println()
	c <- true
}