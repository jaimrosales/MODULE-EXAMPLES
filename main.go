package main

import (
	"fmt"
	"modules/slices"
	"strings"

	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	list := []string{"RSTech", "gophers", "golang", quote.Hello()}

	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "h")
	})

	fmt.Println(quoteV3.Concurrency())
}
