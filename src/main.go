package main

import (
	"log"
	"os"
)

var isError bool
var errors []string
var e = Error{}
var s Scanner

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		e.globalError("use: compact [script file]")
	} else if len(args) == 1 {
		runFile(args[0])
		print(args[0])
	}
	if(isError == true) {
		for _, err := range errors {
			log.Println(err)
		}
		os.Exit(65)
	}
}

func runFile(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("Error: "+err.Error())
	}
	s = Scanner{chars: file, char: 0}
	s.scanTokens()
}