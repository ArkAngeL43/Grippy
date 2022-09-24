package GRIP

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

var parser string
var FP Filepaths

type Filepaths struct {
	Saves []string
}

var Counter_requests int

// function will make the request and save the output to the file
func (Options *Opts) Request(uri string) {
	ct := time.Now()
	// actual extension name this will check vs the deired filetype
	u, x1 := url.Parse(uri)
	if x1 != nil {
		fmt.Println(x1)
	}
	p := strings.LastIndex(u.Path, ".")
	if p == -1 {
		msg := fmt.Sprintf("%v[WARN] :::: %v<%v %v:%v:%v %v> %vWARNING: COULD NOT FIND END URL SUFFIX LOOKING FOR PREFIX (.) NOT %v | SKIPPING", DNUJCFXC, UCDMBFDR, BLU, ct.Hour(), ct.Minute(), ct.Second(), UCDMBFDR, SEVDXFFI, u.Path[p+1:len(u.Path)])
		fmt.Println(msg)
	}
	if Options.Verbose {
		msg := fmt.Sprintf("%s[INFO] :::: %s<%v %v:%v:%v %v> Found end URL SUFFIX  %v < %v %s %v > ", DNUJCFXC, UCDMBFDR, BLU, ct.Hour(), ct.Minute(), ct.Second(), UCDMBFDR, DNUJCFXC, UCDMBFDR, u.Path[p+1:len(u.Path)], DNUJCFXC)
		fmt.Println(msg)
		msg1 := fmt.Sprintf("%s[INFO] :::: %s<%v %v:%v:%v %v> WANTED URL SUFFIX  %v < %v %s %v > ", DNUJCFXC, UCDMBFDR, BLU, ct.Hour(), ct.Minute(), ct.Second(), UCDMBFDR, DNUJCFXC, UCDMBFDR, strings.Trim(Options.FileFormat, "*."), DNUJCFXC)
		fmt.Println(msg1)
	}
	if u.Path[p+1:len(u.Path)] == strings.Trim(Options.FileFormat, "*.") {
		end, x := url.Parse(uri)
		if x != nil {
			fmt.Println(x)
		}
		parser = Options.Output + path.Base(end.Path)
		TextFilena.Filename = parser
		response, x := http.Get(uri)
		if x != nil {
			fmt.Println(x)
		}
		defer response.Body.Close()
		if response.StatusCode != 200 {
			fmt.Println("Invalid status code")
		} else {
			f, x := os.Create(parser)
			if x != nil {
				log.Fatal(x)
			}
			defer f.Close()
			_, x = io.Copy(f, response.Body)
			if x != nil {
				log.Fatal(x)
			}
			Counter_requests++
			FP.Saves = append(FP.Saves, parser)
		}
	} else {
		fmt.Println("The URL's end filepath did not match the extension you were looking for")
	}
}

func Directories() {
	file, _ := json.MarshalIndent(FP, "", " ")
	_ = ioutil.WriteFile("Modules/Access/Files.json", file, 0644)
	DecodeF()
}

/*
func (Options *Opts) Erase() {
	/*
		for _, l := range DL.DLF {
			fmt.Println(RED, "[INFO] :::: DELETING FILE - ", l, " ( REASON: FILE NOT WANTED BY USER [ specified by flag ] ) ")
			if Options.Verbose {
				FINF(l)
			}
		}

}
*/
