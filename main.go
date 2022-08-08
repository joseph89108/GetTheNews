package main

import(
	"fmt"
	"time"
	"os"
	"sync"
)

type New interface{
	ToTxt(wg *sync.WaitGroup)
	ToHtml(wg *sync.WaitGroup)
}

func main(){
	c := make(chan bool)
	go Google(c)
	go Bing(c)
	go NewsAPI(c)
	<- c
	<- c
	<- c
}

func (n *BingNews)ToTxt(wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.Now()
    hour,min,sec := t.Clock()
    year, month, day := t.Date()
    str := fmt.Sprintf("txt/%d%02d%02d_%02d%02d%02d_Bing.txt",year,int(month),day,hour,min,sec)
    foutput, err := os.Create(str)
    if err != nil {
        panic(err)
    }
	for _, v := range n.Value {
		foutput.WriteString(v.DatePublished+" "+v.Name+"\n")
		foutput.WriteString(v.Url+"\n\n")
	}
	foutput.Close()
}

func (n *News)ToTxt(wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.Now()
    hour,min,sec := t.Clock()
    year, month, day := t.Date()
    str := fmt.Sprintf("txt/%d%02d%02d_%02d%02d%02d_News.txt",year,int(month),day,hour,min,sec)
    foutput, err := os.Create(str)
    if err != nil {
        panic(err)
    }
	for _, v := range n.Articles {
		foutput.WriteString(v.PublishedAt+" "+v.Title+"\n")
		foutput.WriteString(v.Url+"\n\n")
	}
	foutput.Close()
}

func (n *GoogleNews)ToTxt(wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.Now()
    hour,min,sec := t.Clock()
    year, month, day := t.Date()
    str := fmt.Sprintf("txt/%d%02d%02d_%02d%02d%02d_Google.txt",year,int(month),day,hour,min,sec)
    foutput, err := os.Create(str)
    if err != nil {
        panic(err)
    }
	for _, v := range n.Articles {
		foutput.WriteString(v.Published+" "+v.Title+"\n")
		foutput.WriteString(v.Link+"\n\n")
	}
	foutput.Close()
}


func (n *BingNews)ToHtml(wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.Now()
    hour,min,sec := t.Clock()
    year, month, day := t.Date()
    str := fmt.Sprintf("html/%d%02d%02d_%02d%02d%02d_Bing.html",year,int(month),day,hour,min,sec)
    foutput, err := os.Create(str)
    if err != nil {
        panic(err)
    }
    foutput.WriteString("<!DOCTYPE html>\n<html>\n<head><title>BingNews</title></head>")
	for _, v := range n.Value {
		foutput.WriteString("<br>"+v.DatePublished+"<br>\n<a href=\""+v.Url+"\">"+v.Name+"</a><br>")
	}
	foutput.WriteString("</html>")
	foutput.Close()
}

func (n *News)ToHtml(wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.Now()
    hour,min,sec := t.Clock()
    year, month, day := t.Date()
    str := fmt.Sprintf("html/%d%02d%02d_%02d%02d%02d_News.html",year,int(month),day,hour,min,sec)
    foutput, err := os.Create(str)
    if err != nil {
        panic(err)
    }
    foutput.WriteString("<!DOCTYPE html>\n<html>\n<head><title>News</title></head>")
	for _, v := range n.Articles {
		foutput.WriteString("<br>"+v.PublishedAt+"<br>\n<a href=\""+v.Url+"\">"+v.Title+"</a><br>")
	}
	foutput.WriteString("</html>")
	foutput.Close()
}

func (n *GoogleNews)ToHtml(wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.Now()
    hour,min,sec := t.Clock()
    year, month, day := t.Date()
    str := fmt.Sprintf("html/%d%02d%02d_%02d%02d%02d_Google.html",year,int(month),day,hour,min,sec)
    foutput, err := os.Create(str)
    if err != nil {
        panic(err)
    }
    foutput.WriteString("<!DOCTYPE html>\n<html>\n<head><title>GoogleNews</title></head>")
	for _, v := range n.Articles {
		foutput.WriteString("<br>"+v.Published+"<br>\n<a href=\""+v.Link+"\">"+v.Title+"</a><br>")
	}
	foutput.WriteString("</html>")
	foutput.Close()
}