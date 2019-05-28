package main

func twoSum(nums []int, target int) []int {
	helpMap := map[int]int{}
	for k, v := range nums {
		if pair,ok := helpMap[v];ok{
			return []int{pair,k}
		}
		helpMap[target-v] = k
	}
	return nil
}

func main() {

}