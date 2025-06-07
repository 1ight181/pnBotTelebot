package sliceutils

import "strconv"

func StringsToInts(strs []string) ([]int, error) {
	nums := make([]int, 0, len(strs))
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}
