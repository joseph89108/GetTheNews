package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
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

	url := "https://newsapi.org/v2/top-headlines?country=us&apiKey=fd7312e58db040d58caeead30d52a764"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var news News
	err := json.Unmarshal(body,&news)
	if err != nil {
		return 
	}
	fmt.Println(time.Now()," News:")
	for _, v := range news.Articles {
		fmt.Println(v)
	}
	fmt.Println()
	c <- true
}
/*
func writeNews(a NewsArticle){
	d := time.Now()
    hour,min,sec := d.Clock()
    year, month, day := d.Date()
	foutput, err := os.Create("News_%d%02d%02d_%02d%02d%02d.txt",year,int(month),day,hour,min,sec)
	//fmt.Printf("%d%02d%02d_%02d%02d%02d\n",year,int(month),day,hour,min,sec)
    if err != nil {
        panic(err)
    }
}
*/