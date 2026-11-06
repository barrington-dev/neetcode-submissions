func maxSubArray(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	currSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = max(currSum + nums[i], nums[i])
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}
