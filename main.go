package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	golog "log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var log = golog.New(os.Stderr, "tsdin ", golog.LstdFlags)

var unixRegexp = regexp.MustCompile(`\b(15\d{8})(\.)?(\d{1,3})?\b`)

func unixToString(ts string) string {
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	tm := time.Unix(i, 0)
	return tm.Format(timeLayout)
}

var isoRegexp = regexp.MustCompile(`("\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}?Z")`)

func isoToString(ts string) string {
	t, err := time.Parse(timeLayout, ts)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(int(t.Unix()))
}

var timeLayout string
var fromISOtoUnix bool

func init() {
	flag.StringVar(&timeLayout, "layout", `"`+time.RFC3339+`"`, "time layout")
	flag.BoolVar(&fromISOtoUnix, "reverse", false, "turn ISO time into unix timestamp")
	flag.Parse()
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Reading stdin failed", err)
	}
	stdin := string(bytes)

	re := unixRegexp
	f := unixToString

	if fromISOtoUnix {
		re = isoRegexp
		f = isoToString
	}

	result := re.FindAllStringSubmatch(stdin, -1)
	for _, item := range result {
		token := item[1]
		out := f(token)
		stdin = strings.Replace(string(stdin), item[0], out, 1)
	}
	fmt.Print(stdin)

}
