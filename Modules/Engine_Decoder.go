package GRIP

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type JOUT struct {
	Urls []string `json:"Urls"`
}

func Decoder() {
	j, e := os.Open("Modules/Access/Links.json")
	if e != nil {
		log.Fatal(e)
	}
	defer j.Close()
	vtb, _ := ioutil.ReadAll(j)
	var vals JOUT
	json.Unmarshal(vtb, &vals)
	fmt.Println(vals)
	for i := 0; i < len(vals.Urls); i++ {
		if O.Verbose {
			fmt.Println("[OK++++] MAKING REQUEST TO - ", vals.Urls[i])
		}
		O.Request(vals.Urls[i]) // make request and save file
	}
}
