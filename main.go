package main


func main(){
	c := make(chan bool)
	go Google(c)
	go Bing(c)
	go NewsAPI(c)
	<- c
	<- c
	<- c
}