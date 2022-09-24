package GRIP

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type OutOut struct {
	Urls []string
}

var OO OutOut

func STDOUT(data map[string][]string) {
	for k, v := range data {
		for _, t := range v {
			if k != "" {
				if t != "" {
					fmt.Printf("\033[38;5;88m[TAG] > \033[38;5;55m%s -> \033[38;5;55m\t|+| -> \033[38;5;50m%s\n", k, t)
				}
			}
		}
	}
}

func Results() {
	OO.Urls = append(OO.Urls, Links_Crawled...)
	file, _ := json.MarshalIndent(OO, "", " ")
	_ = ioutil.WriteFile("Modules/Access/Links.json", file, 0644)
	Decoder()
}
