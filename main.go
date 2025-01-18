package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Notice struct {
	num       int
	link      string
	title     string
	category  string
	content   string
	createdAt string
}

var URLs = map[string]string{
	"전체":       "https://computer.knu.ac.kr/bbs/board.php?bo_table=sub5_1",
	"일반공지":     "https://computer.knu.ac.kr/bbs/board.php?bo_table=sub5_1&sca=%EC%9D%BC%EB%B0%98%EA%B3%B5%EC%A7%80",
	"학사":       "https://computer.knu.ac.kr/bbs/board.php?bo_table=sub5_1&sca=%ED%95%99%EC%82%AC",
	"장학":       "https://computer.knu.ac.kr/bbs/board.php?bo_table=sub5_1&sca=%EC%9E%A5%ED%95%99",
	"심컴":       "https://computer.knu.ac.kr/bbs/board.php?bo_table=sub5_1&sca=%EC%8B%AC%EC%BB%B4",
	"글솝":       "https://computer.knu.ac.kr/bbs/board.php?bo_table=sub5_1&sca=%EA%B8%80%EC%86%9D",
	"대학원":      "https://computer.knu.ac.kr/bbs/board.php?bo_table=sub5_1&sca=%EB%8C%80%ED%95%99%EC%9B%90",
	"대학원 계약학과": "https://computer.knu.ac.kr/bbs/board.php?bo_table=sub5_1&sca=%EB%8C%80%ED%95%99%EC%9B%90+%EA%B3%84%EC%95%BD%ED%95%99%EA%B3%BC",
}

var CATEGORY_ALIAS = map[string]string{
	"전체":       "ALL",
	"일반공지":     "NORMAL",
	"학사":       "STUDENT",
	"장학":       "SCHOLARSHIP",
	"심컴":       "SIM_COM",
	"글솝":       "GL_SOP",
	"인컴":       "SIM_COM",
	"대학원":      "GRADUATE_SCHOOL",
	"대학원 계약학과": "GRADUATE_CONTRACT",
}

const MAX_NOTICE_SIZE = 15

func parseNoticeTotalCount() int {
	res, err := http.Get(URLs["전체"])
	checkErr(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	total, err := strconv.Atoi(strings.TrimSpace(doc.Find("tbody>tr>td.td_num2").First().Text()))
	checkErr(err)

	return total
}

func parseNoticeTable(searchCategory string, page int) []*goquery.Selection {
	res, err := http.Get(URLs[searchCategory] + "&page=" + strconv.Itoa(page))
	checkErr(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	table := make([]*goquery.Selection, 0)

	doc.Find("tbody>tr").Each(func(i int, s *goquery.Selection) {
		table = append(table, s)
	})

	return table
}

func getNoticeData(notice *goquery.Selection, c chan Notice) {
	link, _ := notice.Find("td.td_subject>div.bo_tit>a").Attr("href")
	num, _ := strconv.Atoi(strings.TrimSpace(notice.Find("td.td_num2").Text()))

	res, err := http.Get(link)
	checkErr(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	title := strings.TrimSpace(doc.Find(".bo_v_tit").Text())
	category := CATEGORY_ALIAS[doc.Find(".bo_v_cate").Text()]
	content := strings.TrimSpace(strings.ReplaceAll(doc.Find("#bo_v_con").Text(), "\xa0", ""))
	createdAt := "20" + strings.TrimLeft(strings.Replace(doc.Find(".if_date").Text(), "작성일", "", 1), " ") + ":00"

	c <- Notice{num: num, link: link, title: title, category: category, content: content, createdAt: createdAt}
}

func CrawlNoticeFromWeb(searchCategory string, amount int) (noticeList []Notice) {
	c := make(chan Notice)

	noticeTotalCount := parseNoticeTotalCount()
	if amount == -1 || amount > noticeTotalCount {
		amount = noticeTotalCount
	}

	pages := amount / MAX_NOTICE_SIZE

	noticeTable := make([]*goquery.Selection, 0)
	for page := 1; page <= pages; page++ {
		noticeTable = append(noticeTable, parseNoticeTable(searchCategory, page)...)
	}
	noticeTable = append(noticeTable, parseNoticeTable(searchCategory, pages+1)[:amount%MAX_NOTICE_SIZE]...)

	for _, notice := range noticeTable {
		go getNoticeData(notice, c)
	}

	for i := 0; i < amount; i++ {
		noticeList = append(noticeList, <-c)
	}

	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	noticeList := CrawlNoticeFromWeb("전체", 50)
	for _, notice := range noticeList {
		log.Println(notice)
	}
}
