func replaceElements(arr []int) []int {
	ans := make([]int, len(arr))
	i := len(arr)-1

	rightMax := arr[i]
    ans[i] = -1
	i--

	for ;i >= 0; i-- {
		ans[i] = rightMax
		if arr[i] > rightMax {
			rightMax = arr[i]
		}
	}

	return ans
}
