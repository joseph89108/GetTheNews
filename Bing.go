package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sync"
	"os"
	_ "github.com/joho/godotenv/autoload"
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

	apiKey := os.Getenv("BING_API_KEY")

	req.Header.Add("Ocp-Apim-Subscription-Key", apiKey)

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