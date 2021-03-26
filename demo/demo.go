package main

import "fmt"

func main() {
	//var a [99]int
	//
	//for i:=0;i<len(a);i++{
	//	a[i]=i
	//}
	//fmt.Println(a)
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums,target))
}
func twoSum(nums []int, target int) []int {
	a:=make([]int,2)
	m:=make(map[int]int)

	fmt.Println(len(nums))
	for i:=0;i<len(nums);i++{
		another_num:=target-nums[i]
		val,b:=m[another_num]
		if b{
			a[0]=val
			a[1]=i
			return a
		}
		m[nums[i]]=i
	}
	return a
}