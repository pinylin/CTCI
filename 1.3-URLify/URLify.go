/*
-----------------------
  @Time :             2018/12/16 9:55 AM 
  @Author :           pinglin
  @File :             URLify.go
  @Software:          GoLand
-----------------------
  Change Activity:    
                      2018/12/16 9:55 AM
                
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " Helo Worl d! "
	fmt.Println("Input String :", str)
	fmt.Println(strings.Join(replaceSpaceByString(str),""))
}

func replaceSpaceByString(str string) []string {
	length := len(str)
	var count int

	for i := 0; i < length; i++ {
		ch := string([]rune(str)[i])
		if ch == " " {
			count++
		}
	}
	newLength := length + (2 * count)
	var sArr = make([]string, newLength)
	j := 0
	for i := 0; i < length; i++ {
		ch := string([]rune(str)[i])

		if ch == " " {
			sArr[j] = "%"
			j = j + 1
			sArr[j] = "2"
			j = j + 1
			sArr[j] = "0"
			j = j + 1
		} else {
			sArr[j] = ch
			j = j + 1
		}
	}
	return sArr
}
