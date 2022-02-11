/**
 * @Author: Einic <einicyeo AT gmail.com>
 * @Description:
 * @File: capdns
 * @Version: 1.0.0
 * @Date: 2022/2/8 16:40
 * @BLOG:  https://www.infvie.com
 * @Project home page:
 *     @https://github.com/Einic/EnvoyinStack
 */

package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	domain string
	lary   int
	larray []string
)

func init() {
	const (
		defaultdomain string = "www.infvie.com"
		usage         string = "tcpdump custom grabs the specified domain name"
	)
	flag.StringVar(&domain, "domain", defaultdomain, usage)
	flag.StringVar(&domain, "d", defaultdomain, usage+" (shorthand)")
	flag.Parse()
}

func dascii() []int {
	larray = strings.Split(domain, ".")
	var ret []int
	for lary = range larray {
		ret = append(ret, len(larray[lary]))
		for ary := range larray[lary] {
			ret = append(ret, int(larray[lary][ary]))

		}
	}
	ret = append(ret, 0)
	return ret
}

func main() {
	fmt.Println("domain：", domain)
	fmt.Printf("domain hex: 0x")
	for _, vrx := range dascii() {
		fmt.Printf("%02x", vrx)
	}
	fmt.Println("")
	fmt.Println("domain length:", len(dascii()))
	fmt.Println("tcpdump filter:")
	fmt.Printf("tcpdump -i any -nnX \"(udp and port 53 and (")
	fmt.Printf("ip[%d]=0x%02x", 40, dascii()[0])
	for x := 1; x < len(dascii()); x++ {
		fmt.Printf(" and ip[%d]=0x%02x", 40+x, dascii()[x])
	}
	fmt.Printf("))")
	fmt.Printf(" or (ip[9]=0xfd and (")
	fmt.Printf("ip[%d]=0x%02x", 48, dascii()[0])
	for y := 1; y < len(dascii()); y++ {
		fmt.Printf(" and ip[%d]=0x%02x", 48+y, dascii()[y])
	}
	fmt.Printf("))\"")
	fmt.Println("")
}
