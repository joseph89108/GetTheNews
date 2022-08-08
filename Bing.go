package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sync"
)

type BingNewsValue struct{
	DatePublished string
	Name string
	Url string
	Description string
	Image BingImage
}

type BingNews struct{
	Value []BingNewsValue
}

type BingImage struct{
	Thumbnail BingThumbnail
	IsLicensed bool
}

type BingThumbnail struct{
	ContentUrl string
	Width,Height int
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
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go news.ToTxt(wg)
	go news.ToHtml(wg)
	wg.Wait()
	c <- true
}