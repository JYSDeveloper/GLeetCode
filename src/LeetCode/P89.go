package main

import (
	"fmt"
	"math"
)

func main() {
	grayCode(2)
}

func grayCode(n int) []int {
	v := int(math.Pow(2,float64(n)))
	result := make([]int,v)
	help(result,n)
	fmt.Println(result)
	return result
}
func help(childV []int , level int) []int {
	if level == 0{
		childV[0] = 0
		return childV
	}else {
		childV = help(childV,level-1)
		t := int(math.Pow(2,float64(level-1)))
		for i:=0;i<t;i++ {
			childV[t + i] = childV[t-i-1] + t
		}
	}
	return childV
}



