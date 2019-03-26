package main

import (
	"strconv"
)

func main() {
	//s:= "3[a]2[bc]"
	s :="100[leetCode]"
	decodeString(s)
}
func decodeString(s string) string {
	left := 0
	for i := 0 ;i< len(s); i++ {
		v := s[i]
		if isLeftFlag(v){
			left = i
		}
		if isRightFlag(v){
			j := left-1
			bs := string(s[left+1:i])
			println(bs)
			print("J")
			println(j)
			for isNumber(s[j]){
				j--
				if j == -1{
					break
				}
			}

			count,_ := strconv.Atoi(string(s[j+1:left]))
			println(count)
			ts := toSt(bs,count)
			s = string(s[0:j+1]) + ts + string(s[i+1:len(s)])
			println(s)
			i = 0
		}
	}
	println(s)
	return s
}

func toSt(s string,count int) string {
	res := ""
	for i := 0; i < count; i++ {
		res += s
	}
	return res
}

func isNumber(s byte)bool{
	_ ,err:= strconv.Atoi(string(s))
	return err == nil
}

func isLeftFlag(s byte ) bool{
	return  s == '['
}

func isRightFlag(s byte ) bool{
	return  s == ']'
}




