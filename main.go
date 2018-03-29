package main

import (
	"fmt"
	"io/ioutil"
	golog "log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var tsRegexp = regexp.MustCompile(`\b(15\d{8})(\.\d{1,3})?\b`)
var log = golog.New(os.Stderr, "tsdin", golog.LstdFlags)

func getTimeString(ts string) string {
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	tm := time.Unix(i, 0)
	return tm.Format(`"2006-01-02 15:04:05 MST"`)
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Reading stdin failed", err)
	}
	stdin := string(bytes)

	result := tsRegexp.FindAllStringSubmatch(stdin, -1)
	for _, item := range result {
		ts := item[1]
		timeString := getTimeString(ts)
		stdin = strings.Replace(string(stdin), item[0], timeString, 1)
	}

	fmt.Print(stdin)
}
