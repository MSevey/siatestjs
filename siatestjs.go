package siatestjs

//go:generate gopherjs build --minify

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

func init() {
	js.Module.Get("exports").Set("hello", hello)
	js.Module.Get("exports").Set("newtestgroup", newGroup)
}

func hello(s string) string {
	return fmt.Sprintf("Hello %v!", s)
}
