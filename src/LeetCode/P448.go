package main

/*
 * https://www.jianshu.com/p/c8cc75d2f9dc
 */
func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	findDisappearedNumbers(nums)
}

func findDisappearedNumbers(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{}
	}
	var res []int
	for i := 0; i < len(nums); i++ {
		value := nums[i]
		for nums[i] != i+1 && nums[value-1] != nums[i] {
			nums[i], nums[value-1] = nums[value-1], nums[i]
			value = nums[i]
		}
	}
	for i := range nums {
		if i+1 != nums[i] {
			res = append(res, i+1)
		}
	}
	return res
}
