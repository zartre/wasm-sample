package main

import (
	"errors"
	"fmt"
	"net/mail"
	"syscall/js"
)

func main() {
	fmt.Println("Loaded Wasm")
	js.Global().Set("isEmail", isEmailWrapper)
	<-make(chan bool)
}

func isEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

var isEmailWrapper = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return errors.New("requires 1 argument")
	}
	s := args[0].String()
	return isEmail(s)
})
