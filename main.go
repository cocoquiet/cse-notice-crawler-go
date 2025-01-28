package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"main.go/crawler"
)

func main() {
	fmt.Println(time.Now().UTC(), "- Start crawling")
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("Usage: %s <notice amount> <url>\n", args[0])
		return
	}

	amount, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid notice amount:", err)
		return
	}
	noticeList := crawler.CrawlNoticeFromWeb("전체", amount)
	fmt.Println(time.Now().UTC(), "- Finish crawling")

	crawler.SendNoticeToAPI(args[2], noticeList)

	fmt.Println(time.Now().UTC(), "- Finish Sending")
}
