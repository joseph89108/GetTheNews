all:
	mkdir -p txt html
	go build main.go Bing.go Google.go NewsAPI.go

clean:
	rm main