package main

import (
	"bytes"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/latiif/lail/cmd/repl"
)

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("Lail")

	registerCallbacks()

	<-c
}

func interpret(this js.Value, i []js.Value) interface{} {
	in := strings.NewReader(i[0].String())
	out := &bytes.Buffer{}
	repl.Start(in, out)
	js.Global().Set("output", out.String())
	return out.String()
}

func registerCallbacks() {
	js.Global().Set("interpret", js.FuncOf(interpret))
}
