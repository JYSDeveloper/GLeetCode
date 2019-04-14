package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1,3},2))
}

func searchInsert(nums []int, target int) int {
	if nums[0] >= target {
		return 0
	}
	if nums[len(nums)-1] < target{
		return len(nums)
	}

	begin,end := 0,len(nums)-1
	for begin <= end {
		mid := (begin+end)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			begin = mid-1
		} else {
			end = mid+1
		}
	}

	return begin
}
