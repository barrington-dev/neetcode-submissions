func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] == nums[j] && abs(i-j) <= k {
				return true
			}
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
