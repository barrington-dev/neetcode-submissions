func maxSubArray(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	currSum := nums[0]
	maxSum := nums[0]

	for _, v := range nums[1:] {
		currSum = max(currSum + v, v)
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}
