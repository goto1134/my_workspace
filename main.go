package main

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
)

func main() {
	cases.Lower()
	m := map[string]bool{
		"work": false,
		"fun":  true,
	}
	log.Print("Keys", maps.Keys(m))

	list := []string{"foo", "bar", "zee"}
	log.Print(slices.Contains(list, "foo"))
	slices.SortFunc(list, func(a, b string) bool {
		return a < b
	})
}
