package GRIP

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type DownloadedSuccess struct {
	Filename []string `json:"Saves"`
}

func DecodeF() {
	j, e := os.Open("Modules/Access/Files.json")
	if e != nil {
		log.Fatal(e)
	}
	defer j.Close()
	vtb, _ := ioutil.ReadAll(j)
	var vals DownloadedSuccess
	json.Unmarshal(vtb, &vals)
	for i := 0; i < len(vals.Filename); i++ {
		if O.Verbose {
			fmt.Println("[OK+++++] CHECKING FILENAME - ", vals.Filename[i])
		}
		TextFilena.Sig(fmt.Sprint(vals.Filename[i]))
	}
}
