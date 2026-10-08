func twoSum(nums []int, target int) []int {
    
	var seen map[int]int = map[int]int{}
	
	for i, num := range nums {
		remainder := target - num
		
		s, ok := seen[remainder]
		if ok {
			return []int{s, i}
		}
		seen[num] = i
	}

	return []int{}
}
