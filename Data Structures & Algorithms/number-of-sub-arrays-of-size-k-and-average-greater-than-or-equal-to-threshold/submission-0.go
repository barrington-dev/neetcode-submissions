func numOfSubarrays(arr []int, k int, threshold int) int {
	if len(arr) == 0 {
		return 0
	}

	sum := 0
	ans := 0

	for i := 0; i < k; i++ {
		sum+=arr[i]
	}

	if (sum/k) >= threshold {
		ans++
	}

	for i := k; i < len(arr); i++ {
		sum+=arr[i] - arr[i-k]

		if (sum/k) >= threshold {
			ans++
		}
	}

	return ans
}
