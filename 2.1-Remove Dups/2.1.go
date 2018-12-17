/*
-----------------------
  @Time :             2018/12/16 10:51 AM 
  @Author :           pinglin
  @File :             2.1
  @Software:          GoLand
-----------------------
  Change Activity:    
                      2018/12/16 10:51 AM
                
*/
package main

import (
	"fmt"
	"container/list"
)

var sMap map[int]bool

func main() {
	l := list.New()
	l.PushFront(4)
	l.PushFront(5)
	l.PushFront(7)
	l.PushFront(6)
	l.PushFront(5)
	l.PushFront(4)
	l.PushFront(5)
	l.PushFront(7)
	l.PushBack(9)
	l = removeDuplicate(l)
	//print the result list
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
func removeDuplicate(l *list.List) *list.List {
	sMap = make(map[int]bool)
	var next *list.Element
	for e := l.Front(); e != nil; e = next {
		next = e.Next()
		m := e.Value.(int)
		//To verify whether the node value present in the map
		if sMap[m] == true {
			l.Remove(e)
		} else {
			sMap[m] = true
		}
	}
	return l
}

