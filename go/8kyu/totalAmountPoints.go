package kata

import (
	"strconv"
	"strings"
)

func Points(games []string) int {
  //games[0]
  total := 0
  for _, game := range games{
	parts := strings.Split(game,":")
	x, _ := strconv.Atoi(parts[0])
	y, _ :=strconv.Atoi(parts[1])
	if x > y{
		total += 3
	} else if x == y {
		total += 1
	} else {
		total += 0
	}
  }
  return total
}