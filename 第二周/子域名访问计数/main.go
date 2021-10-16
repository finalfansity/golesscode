package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	var sum = make(map[string]int, 0)     //统计的map
	var res = make([]string, 0, len(sum)) //结果数组
	for _, v := range cpdomains {
		sp := strings.Split(v, " ")
		s, _ := strconv.Atoi(sp[0])
		for {
			sum[sp[1]] = sum[sp[1]] + s
			dot := strings.Index(sp[1], ".")
			if dot < 0 {
				break
			}
			sp[1] = sp[1][dot+1:]
		}
	}
	for k, v := range sum {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}

func main() {
	v := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	fmt.Println(subdomainVisits(v))
}
