package problem1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAddUp(t *testing.T) {
	var tests = []struct {
		list   []int
		k      int
		output bool
	}{
		{[]int{}, 17, false},
		{[]int{1}, 17, false},
		{[]int{1, 2, 3}, 5, true},
		{[]int{3, 2, 1}, 5, true},
		{[]int{10, 15, 3, 7}, 17, true},
		{[]int{10, 15, 3, 7}, 19, false},
		{[]int{10, 15, 3, 7, 20, 6, 8, 99, 12, 5, 100, 101, 3}, 23, true},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			out := CanAddUp(tt.list, tt.k)
			assert.Equal(t, tt.output, out)
		})
	}
}
