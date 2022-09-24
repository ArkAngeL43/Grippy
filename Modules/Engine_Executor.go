package GRIP

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
)

func Run_Engine() {
	Param["q"] = OPTIONS.Search_Query
	Param["num"] = strconv.Itoa(OPTIONS.Results_per_page)
	page := 1
	for {
		paramss := url.Values{}
		for k, v := range Param {
			paramss.Add(k, v)
		}
		uri := URL + "?" + paramss.Encode()
		content, resp, err := Req_with_Response(uri)
		if err != nil {
			log.Printf("\033[38;5;55m|\033[38;5;88m-\033[38;5;55m| \033[38;5;88mGot connection error -> %s\n", err.Error())
			Connection_Attempt++
			if Connection_Attempt_max == Connection_Attempt {
				break
			}
			continue
		}
		watcher := Response_Watcher[resp.StatusCode]
		if watcher == "Whoops triggered a CAPTCHA" {
			fmt.Printf("\033[38;5;55m|\033[38;5;178m*\033[38;5;55m| \033[38;5;178mWarning: \033[38;5;88m%s \n", watcher)
		}
		if watcher == "Redirect" {
			red := resp.Header["location"]
			fmt.Printf("\033[38;5;55m|\033[38;5;178m*\033[38;5;55m| \033[38;5;178mWarning: \033[38;5;88m%s \t Location [ %s ]\n", watcher, red)
		}
		if watcher != "OK" {
			fmt.Printf("\033[38;5;55m|\033[38;5;178m*\033[38;5;55m| \033[38;5;178mWarning: \033[38;5;88mGot error when making request {0x01} - %s\n", resp.Status)
			continue
		}
		Pages_Crawled += content
		page++
		if page-1 >= OPTIONS.Pages_to_Crawl {
			break
		}
		Param["start"] = strconv.Itoa(SET_PAGE(page, OPTIONS.Results_per_page))
	}
	Spider()
	Results()
}

func Caller(Res_per_page, Pages_to_Crawl int, search_query, website string) {
	if Res_per_page == 0 {
		OPTIONS.Results_per_page = 1
	} else {
		OPTIONS.Results_per_page = Res_per_page
	}
	if Pages_to_Crawl == 0 {
		OPTIONS.Pages_to_Crawl = 1
	} else {
		OPTIONS.Pages_to_Crawl = Pages_to_Crawl
	}
	if search_query == "" {
		fmt.Println("Parser: Could not parse search query, data missing")
	} else {
		OPTIONS.Search_Query = search_query
	}
	if OPTIONS.Results_per_page >= 1 && OPTIONS.Pages_to_Crawl >= 1 && len(OPTIONS.Search_Query) > 0 {
		Run_Engine()
	} else {
		fmt.Println("Command parser: Got error when parsing data and information to run the OSINT engine, please specify some data")
		fmt.Println("====================")
		fmt.Println("Results per page (def =1) -> ", OPTIONS.Results_per_page)
		fmt.Println("Pages to crawl   (def =1) -> ", OPTIONS.Pages_to_Crawl)
		fmt.Println("Search query              -> ", search_query)
		fmt.Println("website to check for      -> ", website)
	}
}
