# [Valid Parentheses](https://leetcode.com/problems/valid-parentheses)

## Solution 1: Stack with Direct Slice
- **Time Complexity**: O(n) - n is the length of the string
- **Space Complexity**: O(n) - in worst case, the stack could contain all characters of the string
- **Approach**: Use a stack to track opening brackets and match them with closing brackets. Create a mapping of closing brackets to their corresponding opening brackets. For each character, if it's an opening bracket, push it onto the stack. If it's a closing bracket, check if the stack is empty or if the top element matches the expected opening bracket. If not, return false. Finally, ensure all brackets have been matched by checking if the stack is empty.


```go
func isValid(s string) bool {
	matchMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}

	for _, v := range s {
		if _, isRightBracket := matchMap[v]; !isRightBracket {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top == matchMap[v] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
```


## Solution 2: Stack with Custom Type
- **Time Complexity**: O(n)
- **Space Complexity**: O(n)
- **Approach**: Same algorithm but implemented with a custom stack type for improved readability and encapsulation.

```go
type RuneStack []rune

func (s *RuneStack) Push(r rune) {
	*s = append(*s, r)
}

func (s *RuneStack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, true
}

func (s *RuneStack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return (*s)[len(*s)-1], true
}

func (s *RuneStack) IsEmpty() bool {
	return len(*s) == 0
}

func isValid(s string) bool {
	matchMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack RuneStack

	for _, v := range s {
		if _, isRightBracket := matchMap[v]; !isRightBracket {
			stack.Push(v)
		} else {
			if stack.IsEmpty() {
				return false
			}

			top, _ := stack.Peek()
			if top == matchMap[v] {
				stack.Pop()
			} else {
				return false
			}
		}
	}

	return stack.IsEmpty()
}


```