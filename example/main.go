package main

import (
	"context"
	"log"

	"github.com/xpzouying/xueqiu"
)

func main() {

	xq, err := xueqiu.NewWithEnvToken()
	if err != nil {
		log.Fatal(err)
	}

	news, err := xq.GetLiveNews(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range news.Items[:3] {
		log.Printf("%v", n)
	}
}
