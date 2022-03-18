package main

import (
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"strconv"
)

func main() {
	names := lo.Uniq[string]([]string{"Samuel", "Marc", "Samuel"})
	log.Debug().Msgf("Uniq %v+", names)

	log.Debug().Msgf("Map %v+",
		lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
			return strconv.FormatInt(x, 10)
		}),
	)

	log.Debug().Msgf("GroupBy %v+",
		lop.GroupBy[int, int]([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
			return i % 3
		}),
	)
}
