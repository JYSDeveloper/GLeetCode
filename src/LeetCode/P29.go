package main

import "math"

func divide(dividend int, divisor int) int {
	absDividend := abs(dividend)
	absDivisor := abs(divisor)
	up :=absDivisor
	for i:=0;i<math.MaxInt32;i++{
		absDivisor += up
		if absDivisor > absDividend{
			if dividend < 0 || divisor < 0{
				return -(i+1)
			}
			return i+1
		}
	}
	return math.MaxInt32
}

func abs(int int) int  {
	if int < 0{
		return -int
	}
	return int
}