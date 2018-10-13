package problem2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAugment(t *testing.T) {
	var tests = []struct {
		in  []int
		out []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1, 2, 3}, []int{6, 3, 2}},
		{[]int{3, 2, 1}, []int{2, 3, 6}},
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			out := ListAugment(tt.in)
			assert.Equal(t, tt.out, out)
		})
	}
}
