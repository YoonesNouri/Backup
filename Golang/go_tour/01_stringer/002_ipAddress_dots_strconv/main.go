// راه حل های زیادی در گیت هاب بود که این از همه قشنگتر بود بنظرم
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	s := make([]string, len(ip))
	for i, val := range ip {
		s[i] = strconv.Itoa(int(val)) //عدد داخل اسلایس اِس رو تبدیل به استرینگ میکند
	}
	return strings.Join(s, ".") //ورودی اسلایس استرینگ و سپریتور و خروجی استرینگ به هم چسبیده
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
