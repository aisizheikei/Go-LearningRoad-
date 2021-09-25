/*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]
示例 2：
输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
示例 3：
输入：matrix = [[1]]
输出：[[1]]
示例 4：
输入：matrix = [[1,2],[3,4]]
输出：[[3,1],[4,2]]
提示：
matrix.length == n
matrix[i].length == n
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000

链接：https://leetcode-cn.com/problems/rotate-image
*/
package main

import "fmt"

func main() {
	args1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	args2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}

	rotate(args1)
	rotate(args2)
	fmt.Println("args1：", args1)
	fmt.Println("args2：", args2)
}

//找规律，索引i,j，边长h
//变换后的索引：j,h-i-1
func rotate(matrix [][]int) {
	h := len(matrix)
	i := 0 //第几排
	for i = 0; i < (h+1)/2; i++ {
		for j := i; j < h-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[h-j-1][i]
			matrix[h-j-1][i] = matrix[h-i-1][h-j-1]
			matrix[h-i-1][h-j-1] = matrix[j][h-i-1]
			matrix[j][h-i-1] = temp

		}
	}
}
