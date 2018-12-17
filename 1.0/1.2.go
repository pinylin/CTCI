/*
-----------------------
  @Time :             2018/12/16 9:30 AM 
  @Author :           pinglin
  @File :             1.2
  @Software:          GoLand
-----------------------
  Change Activity:    
                      2018/12/16 9:30 AM
                
*/
package main

import "fmt"

func main(){
	testString := "abc"
	//testString := ""


	ans := reverseString(testString)

	fmt.Println(ans)
}

func reverseString(s string) string{
	runes := []rune(s)

	for from, to := 0, len(runes)-1;from < to;from, to = from +1, to -1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}


