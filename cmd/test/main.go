package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
	"text/template"
)

func main() {
	testTpl := "Hello {{ test }}!"
	tpl, _ := template.New("test").Funcs(template.FuncMap{"test": test}).Parse(testTpl)

	tpl.Execute(os.Stdout, tpl)

	in, _ := ioutil.ReadFile("./cmd/test/test.yaml")
	var test Test
	yaml.Unmarshal(in, &test)
	fmt.Println(test)
}

type Test struct {
	Name string
	Port int
}

func CheckMatch(value string, exclude string) bool {
	r := regexp.MustCompile(exclude)
	if r.MatchString(value) {
		//log.Println()
		return true
	}

	return false
}
