package main

/*
 * gosafe
 * author: 0xlilith
 */

import (
	"fmt"
	"math/rand"
	"time"
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

func main() {
	pass := gen(length)
	fmt.Println(pass)
}

/*
	function gen: generates random values from the var list
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
