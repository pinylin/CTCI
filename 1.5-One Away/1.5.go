/*
-----------------------
  @Time :             2018/12/16 10:11 AM 
  @Author :           pinglin
  @File :             1.5
  @Software:          GoLand
-----------------------
  Change Activity:    
                      2018/12/16 10:11 AM
                
*/
package main

import "fmt"

func main() {
	s1 := "dfw2"
	s2 := "dfw"
	fmt.Println(oneAway(s1, s2))

}

func oneAway(str1, str2 string) bool{
	count := 0
	sMap := make(map[string]int)

	for i := 0; i < len(str1); i++ {
		ch := string([]rune(str1)[i])
		sMap[ch] = sMap[ch] + 1
	}

	for i := 0; i < len(str2); i++ {
		ch := string([]rune(str2)[i])
		sMap[ch] = sMap[ch] - 1
	}

	for k, _ := range sMap {
		//fmt.Println(k,sMap[k])
		if sMap[k] != 0 {
			count = count + 1
		}
	}
	if count > 1 {
		return false
	}
	return true
}
