package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	minLeft,minRight := getMinIndex(len(nums1)+ len(nums2))
	i,j,total:=0,0,0
	for i < len(nums1) && j < len(nums2){
		current := 0
		if nums1[i] < nums2[j]{
			current =nums1[i]
			i ++
			total ++
		}else {
				current = nums2[j]
				j ++
				total++
		}
		if total == minLeft{
			if minRight == 0{
				return float64(current)
			}else {
				return float64(current)/2.0 + math.Min(float64(nums1[i]),float64(nums2[j]))/2.0
			}
		}
	}
	return 0.0
	if total < minLeft{
		if i == len(nums1){

		}
	}

}

func getMinIndex(len int) (left,right int)  {
	if len & 1 != 0{
		left ,right = len / 2+1,0
	}else {
		left,right = len/2 ,len/2+1
	}
	return
}