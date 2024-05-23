package main

import "fmt"

func main(){
	fmt.Println(twoSum([]int{1,2,5,8},9))
	fmt.Println()
	fmt.Println(polindrom(122))
}

func twoSum(nums []int, target int) []int {
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i]+nums[j]==target{
				return []int{i, j}
			}
		}

	}
	return []int{}
}

func polindrom(son int) bool{
	teskari:=0
	tugri:=son
	for son>0{
		qoldiq:=son%10
		teskari=teskari*10+qoldiq
		son/=10
	}

	if tugri==teskari{
		return true
	}
	return false
}