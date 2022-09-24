package GRIP

import (
	"fmt"
	"strings"
	"time"
)

func Spider() {
	for _, n := range TAB0.FindAllStringSubmatch(Pages_Crawled, -1) {
		ct := time.Now()
		cond1 := strings.Contains(strings.ToLower(n[1]), REGEX_URL0)
		if !cond1 {
			Links_Crawled = append(Links_Crawled, n[1])
			msg := fmt.Sprintf("%s[INFO] :::: %s<%v %v:%v:%v %v> Found URL %v< %v %v> ", DNUJCFXC, UCDMBFDR, BLU, ct.Hour(), ct.Minute(), ct.Second(), UCDMBFDR, DNUJCFXC, n[1], UCDMBFDR)
			fmt.Println(msg)
		}
	}
	if len(Links_Crawled) < 1 {
		for _, link := range TAB1.FindAllStringSubmatch(Pages_Crawled, -1) {
			ct := time.Now()
			cond1 := strings.Contains(strings.ToLower(link[1]), REGEX_URL0)
			if !cond1 {
				msg := fmt.Sprintf("%s[INFO] :::: %s<%v %v:%v:%v %v> Found URL %v< %v %v> ", DNUJCFXC, UCDMBFDR, BLU, ct.Hour(), ct.Minute(), ct.Second(), UCDMBFDR, DNUJCFXC, link[1], UCDMBFDR)
				fmt.Println(msg)
				Links_Crawled = append(Links_Crawled, link[1])
			}
		}
	}
}
