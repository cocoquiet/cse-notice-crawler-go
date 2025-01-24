package main

import (
	"fmt"
	"time"

	"main.go/crawler"
)

func main() {
	fmt.Println(time.Now().UTC(), "- Start crawling")
	noticeList := crawler.CrawlNoticeFromWeb("전체", 50)
	fmt.Println(time.Now().UTC(), "- Finish crawling")

	crawler.SendNoticeToAPI("https://uqywguq.request.dreamhack.games", noticeList)

	fmt.Println(time.Now().UTC(), "- Finish Sending")
}
