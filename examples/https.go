package main

import (
	curl "github.com/yanadhiwiranata/go-curl"
)

func main() {
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		easy.Setopt(curl.OPT_URL, "https://baidu.com/")
		easy.Perform()
	}
}
