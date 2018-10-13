package problem1

import (
	"log"
	"sort"
)

func CanAddUp(list []int, k int) bool {
	sort.Ints(list)
	var out bool
	if len(list) <= 1 {
		out = false
	} else if list[0]+list[len(list)-1] == k {
		out = true
	} else if list[0]+list[len(list)-1] > k {
		out = CanAddUp(list[:len(list)-1], k)
	} else if list[0]+list[len(list)-1] < k {
		out = CanAddUp(list[1:], k)
	} else {
		log.Fatal("Something went catastrophically wrong.")
	}
	return out
}
