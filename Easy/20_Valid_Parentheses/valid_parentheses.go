package validparentheses

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
