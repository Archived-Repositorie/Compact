package main

import "strconv"

type Error struct{}

func (e *Error) globalError(content string) {
	errors = append(errors, "Error: "+content)
	isError = true
}

func (e *Error) lineError(line int, content string) {
	e.globalError("line (" + strconv.Itoa(line) + "): " + content)
}

func (e *Error) tokenError(line int, char byte, content string) {
	e.lineError(line, "at '"+string(char)+"': "+content)
}