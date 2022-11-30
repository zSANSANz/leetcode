func twoSum(nums []int, target int) []int {
	i := 0
	j := 0
	resulti := 0
	resultj := 0
	for i = 0; i < len(nums); i++ {
		for j = 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				resulti = i
				resultj = j
			}
		}
	}
	return []int{resultj, resulti}
}