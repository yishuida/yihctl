package main

import (
	"fmt"
	"net/http"
)

var repsTpl = `%s
%v
%v
%v
%v
`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s\n", r.URL.RequestURI())
	fmt.Printf("%v\n", r.Proto)
	fmt.Printf("%v\n", r.Host)
	fmt.Printf("%v\n", r.UserAgent())
	fmt.Printf("%v\n", r.URL.Query().Get("mz_id"))

	reps := fmt.Sprintf(repsTpl, r.URL.RequestURI(), r.Proto, r.Host, r.UserAgent(), r.URL.Query().Get("mz_id"))

	_, err := w.Write([]byte(reps))
	if err != nil {
		fmt.Errorf("%s", err)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
