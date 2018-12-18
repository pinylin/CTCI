/*
-----------------------
  @Time :             2018/12/18 8:24 AM 
  @Author :           pinglin
  @File :             findNumAppearedOnce
  @Software:          GoLand
-----------------------
  Change Activity:    
                      2018/12/18 8:24 AM
                
*/
package main

import (
	"fmt"
	"github.com/imroc/biu"
)

func main() {
    nums := []uint8{ 1, 2, 3, 4, 5, 1, 2, 3, 4, 8}
    a,b := findNums(nums)
    fmt.Println(a,b)
}

func findNums(nums []uint8) (a,b uint8){
	var temp uint8 = 0
	for _,i := range nums{
		temp ^= i
	}
	idx := findIdxOf1(temp)
	var num1,num2 []uint8
	for _, item := range nums{
		if isBit1(item, idx){
			num1 = append(num1, item)
		}else{
			num2 = append(num2, item)
		}
	}

	for _, i := range num1{
		a ^= i
	}
	for _, i := range num2{
		b ^= i
	}
	return
}

func findIdxOf1(num uint8) int {
	digits := biu.ToBinaryString(num)
	for idx, j := range []rune(digits){
		//fmt.Println(idx, j)
		if string(j) == "1" {
			//fmt.Println(idx)
			return idx
		}
	}
	return -1
}

func isBit1(num uint8, idx int) bool{
	var m = uint8(8 - idx - 1)
	num = num >> m
	digits := biu.ToBinaryString(num)
	temp := []rune(digits)
	if string(temp[7]) != "1"{
		return false
	}
	return true
}