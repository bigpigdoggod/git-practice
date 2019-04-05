// Test project main.go
package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	a := string("2*(2+4)")
	c := calculate(a)
	fmt.Printf("%d", c)

}

type Node struct {
	data string
	next *Node
}

type Stack struct {
	topeNode *Node
}

func NewStack() *Stack {
	return &Stack{
		nil,
	}
}

func (this *Stack) Push(value string) {
	this.topeNode = &Node{data: value, next: this.topeNode}
}

func (this *Stack) IsEmpty() bool {
	if this.topeNode == nil {
		return true
	}

	return false
}

func (this *Stack) Pop() string {
	value := this.topeNode.data
	this.topeNode = this.topeNode.next
	return value
}

func (this *Stack) Top() string {
	return this.topeNode.data
}

var stack Stack

func calculate(test string) int {
	postfix := infix2ToPostfix(test)
	fixLen := len(postfix)
	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		// 数字：直接压栈
		if unicode.IsDigit(rune(postfix[i])) {
			stack.Push(nextChar)
		} else {
			// 操作符：取出两个数字计算值，再将结果压栈
			num1, _ := strconv.Atoi(stack.Pop())
			num2, _ := strconv.Atoi(stack.Pop())
			switch nextChar {
			case "+":
				stack.Push(strconv.Itoa(num1 + num2))
			case "-":
				stack.Push(strconv.Itoa(num1 - num2))
			case "*":
				stack.Push(strconv.Itoa(num1 * num2))
			case "/":
				stack.Push(strconv.Itoa(num1 / num2))
			}
		}
	}
	result, _ := strconv.Atoi(stack.Top())
	return result
}
func infix2ToPostfix(exp string) string {
	postfix := ""
	expLen := len(exp)

	// 遍历整个表达式
	for i := 0; i < expLen; i++ {

		char := string(exp[i])

		switch char {
		case " ":
			continue
		case "(":
			// 左括号直接入栈
			stack.Push("(")
		case ")":
			// 右括号则弹出元素直到遇到左括号
			for !stack.IsEmpty() {
				preChar := stack.Top()
				if preChar == "(" {
					stack.Pop() // 弹出 "("
					break
				}
				postfix += preChar
				stack.Pop()
			}

			// 数字则直接输出
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			postfix += digit
			i = j - 1 // i 向前跨越一个整数，由于执行了一步多余的 j++，需要减 1

		default:
			// 操作符：遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符
			for !stack.IsEmpty() {
				top := stack.Top()
				if top == "(" || isLower(top, char) {
					break
				}
				postfix += top
				stack.Pop()
			}
			// 低优先级的运算符入栈
			stack.Push(char)
		}
	}

	// 栈不空则全部输出
	for !stack.IsEmpty() {
		postfix += stack.Pop()
	}

	return postfix
}

// 比较运算符栈栈顶 top 和新运算符 newTop 的优先级高低
func isLower(top string, newTop string) bool {
	// 注意 a + b + c 的后缀表达式是 ab + c +，不是 abc + +
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}
