package main

func main() {
	arr := []int {1,0}
	dominantIndex(arr)
}

func dominantIndex(nums []int) int {

	if len(nums) == 1 {
		return 0
	}

	max,smax := nums[0],0
	maxIndex := 0

	for i := 1; i<len(nums);i++  {
		if nums[i] > max{
			max,smax,maxIndex= nums[i],max,i
		}else if smax < nums[i]{
			smax = nums[i]
		}
	}
	if max >= smax * 2 {
		return maxIndex
	}
	return -1
}