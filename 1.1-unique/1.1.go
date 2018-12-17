/*
-----------------------
  @Time :             2018/12/16 9:19 AM 
  @Author :           pinglin
  @File :             1.1
  @Software:          GoLand
-----------------------
  Change Activity:    
                      2018/12/16 9:19 AM
                
*/
package main

import (
	"fmt"
)

var sMap map[string]bool

func main() {

	words:=[]string{"abcde", "hello", "apple", "kite", "padle"};
	for _,value:=range words {
		fmt.Println("Input : ",value," Is Unique? ",isUnique(value))
	}
}

func isUnique(str string) bool {

	sMap = make(map[string]bool)
	var ch string
	for i := 0; i < len(str); i++ {
		ch = string([]rune(str)[i])
		if sMap[ch] {
			return false
		} else {
			sMap[ch] = true
		}
	}
	return true
}