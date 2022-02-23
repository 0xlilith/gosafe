package main

/*
 * gosafe
 * author: 0xlilith
 */

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
	flag variables
*/
var (
	help *bool
	show *string
	add  *string
)

/*
 * lc = lower case characters
 * uc = upper case characters
 * sc = special characters
 * n  = numbers
 * a  = all
 * length = length of the password
 */
var (
	lc     = "abcdedfghijklmnopqrst"
	uc     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	n      = "0123456789"
	sc     = "~=+%^*/()[]{}/!@#$?|"
	all    = lc + uc + n + sc
	length = 25
)

func init() {
	help = flag.Bool("help", false, "Show help")
	show = flag.String("show", "all", "show passwords")
	add = flag.String("add", "", "add name to which you want to generate password for")
}

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage[To add] : gosafe [-add] twitter \nUsage[To show]: gosafe [-show] twitter (default all)")
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}
	pass := gen(length)
	fmt.Println(pass)
	fmt.Println("show:", *show)
	fmt.Println("add:", *add)
}

/*
	function gen: generates random values from the var list
	RULE: there has to be one digit, one sprcial char and one Upper case char
*/
func gen(length int) string {
	rand.Seed(time.Now().UnixNano())

	buf := make([]byte, length)
	buf[0] = n[rand.Intn(len(n))]   /* generating one random digit */
	buf[1] = sc[rand.Intn(len(sc))] /* generate one random special char */
	buf[2] = uc[rand.Intn(len(uc))] /* geneerate one random upper case char */

	/* generate other value from <all> */
	for i := 3; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}

	/* shuffle all the generated value */
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)
}
