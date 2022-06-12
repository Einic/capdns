/**
 * @Author: Einic <einicyeo AT gmail.com>
 * @Description:
 * @File: capdns
 * @Version: 0.1.0
 * @Date: 2022/2/8 16:40
 * @BLOG:  https://www.infvie.com
 * @Project home page:
 *     @https://github.com/Einic/EnvoyinStack
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	version bool
	domain  string
	lary    int
	larray  []string
	nlag    int
)

const (
	VERSION = "v0.1.0"
)

func init() {
	flag.BoolVar(&version, "v", false, "`version`: show version")
	flag.StringVar(&domain, "d", "www.infvie.com", "`domain`: tcpdump filter custom grabs the specified domain name")
	flag.Usage = Usage
}

func Usage() {
	nargs := cap(os.Args)
	if nlag != 0 {
		Domain()
	} else {
		if domain != "" {
			if nargs != 1 {
				fmt.Fprintf(os.Stderr, `Usage: capdns [ OPTIONS ] OBJECT { COMMAND | help }
Options:
`)
				flag.PrintDefaults()
			} else {
				Domain()
			}
		} else {
			fmt.Fprintf(os.Stderr, `Usage: capdns [ OPTIONS ] OBJECT { COMMAND | help }
Options:
`)
			flag.PrintDefaults()
		}
	}
}

func Dascii() []int {
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

func Domain() {
	fmt.Println("domain：", domain)

	fmt.Printf("domain hex: 0x")
	for _, vrx := range Dascii() {
		fmt.Printf("%02x", vrx)
	}
	fmt.Println("")

	fmt.Println("domain length:", len(Dascii()))

	fmt.Println("tcpdump filter:")

	fmt.Printf("tcpdump -i any -nnX \"(udp and port 53 and (")
	fmt.Printf("ip[%d]=0x%02x", 40, Dascii()[0])
	for x := 1; x < len(Dascii()); x++ {
		fmt.Printf(" and ip[%d]=0x%02x", 40+x, Dascii()[x])
	}
	fmt.Printf("))")
	fmt.Printf(" or (ip[9]=0xfd and (")
	fmt.Printf("ip[%d]=0x%02x", 48, Dascii()[0])
	for y := 1; y < len(Dascii()); y++ {
		fmt.Printf(" and ip[%d]=0x%02x", 48+y, Dascii()[y])
	}
	fmt.Printf("))\"")
	fmt.Println("")
}

func main() {
	flag.Parse()

	nlag = flag.NFlag()

	if version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	flag.Usage()

}
