type NumArray struct {
    prefix []int
}


func Constructor(nums []int) NumArray {
    numArray := new(NumArray)

	for i, v := range nums {
		if i == 0 {
			numArray.prefix = append(numArray.prefix, v)
			continue
		}
		
		numArray.prefix = append(numArray.prefix, numArray.prefix[i-1]+v)
	}

	return *numArray
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 {
		return this.prefix[right] - 0
	}
	return this.prefix[right] - this.prefix[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */