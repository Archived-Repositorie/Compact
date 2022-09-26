package main


type Error struct {}

func (e *Error) globalError(content string) {
	errors = append(errors, "Error: "+content)
}

func (e *Error) lineError(line int, content string) {
	e.globalError("line ("+string(line)+"): "+content)
}