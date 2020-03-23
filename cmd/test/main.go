package main

import (
	"regexp"
)

func main() {

}

func CheckMatch(value string, exclude string) bool {
	r := regexp.MustCompile(exclude)
	if r.MatchString(value) {
		//log.Println()
		return true
	}
	return false
}
