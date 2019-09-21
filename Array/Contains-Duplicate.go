func containsDuplicate(nums []int) bool {
	numMap := make(map[int]int, len(nums))
	total := len(nums)
	isDuplicated := false
	for i := 0; i < total; i++ {
			_, ok := numMap[nums[i]]
			if ok {
					isDuplicated = true
					break
			}
			numMap[nums[i]] = 1
	}
	return isDuplicated;
}