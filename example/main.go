package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/xpzouying/xueqiu"
)

func main() {
	stockNingde := "SZ300750" // ning de

	reports, err := xueqiu.GetCompanyReport(context.Background(), stockNingde)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.Marshal(reports)
	log.Printf("%s", data)
}
