/*
-----------------------
  @Time :             2018/12/16 9:36 AM 
  @Author :           pinglin
  @File :             1.3
  @Software:          GoLand
-----------------------
  Change Activity:    
                      2018/12/16 9:36 AM
                
*/
package main

import (
	"fmt"
)

var sMap map[string]int


func main() {
	var inputMap = map[string]string{"apple": "papel","carrot":"tarroc","hello":"llloh",}
	for str1,str2 := range inputMap{
		fmt.Println(str1,",",str2," Is Anagram? ",isPermutation(str1, str2))
	}
}

func isPermutation(str1, str2 string) bool {

	length1 := len(str1)
	length2 := len(str2)

	if length1 != length2 {
		return false
	}

	sMap = make(map[string]int)

	for i := 0; i < length1; i++ {
		ch := string([]rune(str1)[i])
		sMap[ch] = sMap[ch] + 1
	}

	for i := 0; i < length2; i++ {
		ch := string([]rune(str2)[i])
		sMap[ch] = sMap[ch] - 1
	}

	for k, _ := range sMap {
		//fmt.Println(k,sMap[k])
		if sMap[k] != 0 {
			return false
		}
	}
	return true
}