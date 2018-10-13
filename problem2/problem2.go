package problem2

func ListAugment(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	left := make([]int, len(list))
	left[0] = 1

	right := make([]int, len(list))
	right[len(right)-1] = 1

	prod := make([]int, len(list))

	for i := 1; i < len(list); i++ {
		left[i] = list[i-1] * left[i-1]
	}

	for i := len(list) - 2; i >= 0; i-- {
		right[i] = list[i+1] * right[i+1]
	}

	for i := range list {
		prod[i] = left[i] * right[i]
	}

	return prod
}
