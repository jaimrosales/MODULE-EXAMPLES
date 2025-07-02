package main

import (
	"modules/slices"
	"strings"

	"rsc.io/quote"
)

func main() {
	list := []string{"RSTech", "gophers", "golang", quote.Hello()}

	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "h")
	})
}
