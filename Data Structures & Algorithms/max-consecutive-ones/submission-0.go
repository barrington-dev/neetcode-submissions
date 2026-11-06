func findMaxConsecutiveOnes(nums []int) int {
    var ans int
    var count int

    for _, v := range nums {
        if v == 1 {
            count = count + 1
            if ans < count {
                ans = count
            }
            continue
        }
        count = 0
    }

    return ans
}
