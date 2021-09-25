/*hard

给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
示例 1：
输入：s = "1 + 1"
输出：2
示例 2：
输入：s = " 2-1 + 2 "
输出：3
示例 3：
输入：s = "(1+(4+5+2)-3)+(6+8)"
输出：23
提示：
1 <= s.length <= 3 * 105
s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
s 表示一个有效的表达式

链接：https://leetcode-cn.com/problems/basic-calculator
*/

package main

import "fmt"

func main() {
	//args1 := "1+2-(30-10+2-( 2+10 ))" //-7
	//fmt.Println(calculate(args1))
	args2 := "(1+(4+5+2)-3)+(6+8)" //23
	fmt.Println(calculate(args2))
}

/*
遇到前括号把前操作符入栈。遇到后括号吧操作符出栈
res 记录计算结果
num 记录当前操作数
ops 栈记录括号前的符号
*/
func calculate(s string) int {
	//numStack := make([]int, 0)
	ops := make([]rune, 0)
	op := '+'
	num := 0
	res := 0
	for _, r := range []rune(s) {
		if r == ' ' {
			continue
		}
		//fmt.Println(string(r))
		if r >= '0' && r <= '9' {
			num = num*10 + int(r-'0')
			continue
			//numStack = append(numStack, r)
		}
		if len(ops) > 0 {
			temp_op := op
			for _, v := range ops {
				if temp_op == v {
					temp_op = '+'
				} else {
					temp_op = '-'
				}
			}
			if temp_op == '-' {
				num = 0 - num
			}
		} else if op == '-' {
			num = 0 - num
		}
		res += num
		num = 0
		if r == '(' {
			ops = append(ops, op)
			op = '+'
		} else if r == ')' {
			ops = ops[0 : len(ops)-1]
		} else {
			op = r
		}
	}
	//把最后一个数字加进来
	if op == '-' {
		num = 0 - num
	}
	res += num
	return res
}
