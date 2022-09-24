package GRIP

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type File struct {
	Filename string
}

/*
// Identify file based on signature

this function is to be called AFTER the file is downloaded, once it is downloaded and saved check the contents if the contents

of the file match given a signature then check if the fileformat or signature is matched up to the users wanted to be searched

file then move on and call it, if else then go on and delete the file. You do not want any files being accidentaly downloaded

by the engine that are HTML, you want to download them first then check them and erase them if they are NOT what the user wants

As of V0.1 the DELETE functions are not used due to various bugs, maybe later.






--- SEGMENT 1

type DeleteList struct {
	DLF []string
}

var DL DeleteList


func (F *File) DeleteCheck(val FileFormat_Sigs, filename string) {
	if O.Delte_IF_Wrong {
		DL.DLF = append(DL.DLF, filename)
		for _, b := range DL.DLF {
			fmt.Println("-] ", b, " -> ", filename)
			if b != filename {
				DL.DLF = append(DL.DLF, filename)
			}
		}
	}
}
*/

func (F *File) Sig(Filename string) {
	f, err := ioutil.ReadFile(Filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file: %v\n", err)
	}

	for _, val := range Signatures {
		ct := time.Now()
		msg := fmt.Sprintf("%s[INFO] :::: %s<%v %v:%v:%v %v> Possible      data suffix  %v < %v %s %v > ", DNUJCFXC, UCDMBFDR, BLU, ct.Hour(), ct.Minute(), ct.Second(), UCDMBFDR, DNUJCFXC, UCDMBFDR, val.FileFormat, DNUJCFXC)
		if strings.HasSuffix(Filename, val.SuffixFile) || bytes.Contains(f, []byte(val.Sign)) {
			if strings.Trim(O.FileFormat, "*") == strings.Trim(val.SuffixFile, "*") {
				fmt.Println(msg)
				return
			}
		} else {
			if O.Verbose {
				msg := fmt.Sprintf("%s[INFO] :::: %s<%v %v:%v:%v %v> Searching for data suffix  %v < %v %s %v > ", RED, RED, BLU, ct.Hour(), ct.Minute(), ct.Second(), RED, RED, BLU, val.SuffixFile, RED)
				fmt.Println(msg)
			}
		}
	}
}
