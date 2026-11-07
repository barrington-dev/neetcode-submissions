func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}

	window := make(map[int]struct{})

	for i, v := range nums {
		if i > k {
			delete(window, nums[i-k-1])
		}

		if _, ok := window[v]; ok {
			return true
		}

		window[v] = struct{}{}
	}

	return false
}