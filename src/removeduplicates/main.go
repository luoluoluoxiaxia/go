package main

import "fmt"

func main(){
	nums := []int{1,2,1,1,1,2,2,3,3,4}
	nlen := removeDuplicates(nums);
	fmt.Println(nlen)
	fmt.Println(nums)
}
func removeDuplicates(nums []int) int {
	number := 0//记录不同数字的下标
	lens := len(nums)
	for i := 0; i < lens ; i++ {
		if nums[i] != nums[number] {
			//不相同 下标前移 赋值
			number++
			nums[number] = nums[i]

		}
	}
	return number+1
}