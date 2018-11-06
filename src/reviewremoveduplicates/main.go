package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nlen := removeDuplicates(nums);
	fmt.Println(nlen)
	fmt.Println(nums)
}
func removeDuplicates(nums []int) int {
	//记录重复的次数
	repeatNum := 0
	//记录数组的长度
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		//第一次循环
		// i = 0 != repeatNum = 0
		//i++
		//进入第二次循环
		//i = 1   repeatNum = 0
		//此时
		// nums[i] = 0 != nums[repeatNum] = 0  无法进入判断体
		//i++
		//进入第三次
		//i = 2 repeatNum = 0
		//此时
		// nums[i] = 1 != nums[repeatNum] = 0
		//
		if nums[i] != nums[repeatNum]{
			repeatNum += 1
			nums[repeatNum] = nums[i]
		}
	}
	return repeatNum + 1
}
