/*
-----------------------
  @Time :             2018/12/16 10:41 AM 
  @Author :           pinglin
  @File :             1.7
  @Software:          GoLand
-----------------------
  Change Activity:    
                      2018/12/16 10:41 AM
                
*/
package main

import (
	"fmt"
)

//Main function
func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Input : ", matrix)
	fmt.Println("Output : ",rotateMatrix(matrix, len(matrix)))
}

//Function to rotate the matrix
func rotateMatrix(matrix [][]int, n int) [][]int {
	for layer := 0; layer < n/2;layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last;i++ {
			offset := i - first
			top := matrix[first][i]
			matrix[first][i] = matrix[last-offset][first]
			matrix[last-offset][first] = matrix[last][last-offset]
			matrix[last][last-offset] = matrix[i][last]
			matrix[i][last] = top
		}
	}

	return matrix
}


