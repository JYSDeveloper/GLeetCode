package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	near,result:=math.MaxInt32,math.MaxInt32
	sort.Ints(nums)

	for i:=0;i<len(nums)-2;i++{
		left,right := i+1,len(nums)-1
		for left < right{
			sum := nums[i] + nums[left] + nums[right]
			subTarget := target - sum
			if subTarget == 0{
				return sum
			}
			if subTarget > 0{
				left++
			}else {
				right--
				subTarget = -subTarget
			}
			if near > subTarget{
				near = subTarget
				result = sum
			}

		}
	}
	return result
}
