package main

func validParentheses(s string) bool {
	stack := Stack{items: make([]int32, 0)}
	for _, char := range s {
		switch char {
		case '{', '[', '(':
			stack.Push(char)
		case '}', ']', ')':
			openParenthesis, ok := stack.Pop()
			if !ok {
				return false
			}

			closeParenthesis, ok := parenthesisPairs[openParenthesis]
			if !ok {
				return false
			}
			if closeParenthesis != char {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

var parenthesisPairs = map[int32]int32{
	'{': '}',
	'[': ']',
	'(': ')',
}

type Stack struct {
	items []int32
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(item int32) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int32, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	indexToPop := len(s.items) - 1
	item := s.items[indexToPop]
	if indexToPop == 0 {
		s.items = make([]int32, 0)
	} else {
		s.items = s.items[:indexToPop]
	}
	return item, true
}
