func removeElement(nums []int, val int) int {
    var j int

	for _, v := range nums {
		if v == val {
			continue
		}
		nums[j] = v
		j++
	}

	return j
}
