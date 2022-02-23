package main

/*
 * gosafe
 * author: 0xlilith
 */

import (
	"database/sql"
	"flag"
	"fmt"
	"gosafe/gosafe"
	"math/rand"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*
	flag variables
*/
var (
	help *bool
	show *bool
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
	show = flag.Bool("show", false, "show passwords")
	add = flag.String("add", "", "add name to which you want to generate password for")
}

func main() {

	/* connecting database */
	db, err := sql.Open("sqlite3", "./gosafe.db")
	if err != nil {
		panic("failed to connect sql server: " + err.Error())
	}
	/* creating new table */
	locker := gosafe.NewPass(db)

	flag.Parse()

	/* check for help flag */
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	/* check for show flag */
	if *show {
		fmt.Println(*show)
		os.Exit(0)
	}

	pass := gen(length)
	fmt.Println("Name:", *add, "\tPassword:", pass)

	/* adding to the database */
	locker.Add(gosafe.Item{
		Name:     *add,
		Password: pass,
	})

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
