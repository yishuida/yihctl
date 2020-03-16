package main

import "log"

func init() {
	log.SetPrefix("【UserCenter】")
	log.SetFlags(log.Ldate | log.Lshortfile)
}
func main() {
	log.Print("hahaha", "adsf")
}
