package crawler

import (
	"log"
	"net/http"
)

type notice struct {
	num        int
	link       string
	title      string
	category   string
	content    string
	created_at string
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
}

func CrawlNoticeFromWeb(searchCategory string, amount int) []notice {

}

func SendNoticeToAPI(url string, noticeList []notice) int {

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
