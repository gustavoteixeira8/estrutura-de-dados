package main

import (
	"errors"
	"fmt"
	"strings"
)

type Stack struct {
	lastPos int
	vector  []string
	maxLen  int
}

func (s *Stack) Insert(val string) error {
	if s.lastPos == s.maxLen-1 {
		return errors.New("vector is full")
	}

	s.lastPos += 1
	s.vector[s.lastPos] = val

	return nil
}

func (s *Stack) Remove() error {
	if s.lastPos == -1 {
		return errors.New("vector is empty")
	}

	s.lastPos -= 1

	return nil
}

func (s *Stack) Tail() string {
	if s.lastPos == -1 {
		return ""
	}
	return s.vector[s.lastPos]
}

func (s *Stack) Vector() []string {
	return s.vector
}

func (s *Stack) IsEmpty() bool {
	return s.lastPos == -1
}

func NewStack(maxLen int) *Stack {
	return &Stack{vector: make([]string, maxLen), maxLen: maxLen, lastPos: -1}
}

type ExpressionValidator struct {
	stack *Stack
}

func (v *ExpressionValidator) Validate(expression string) bool {
	for _, strRune := range expression {
		str := strings.Trim(string(strRune), " ")

		if str == "{" || str == "[" || str == "(" {
			v.stack.Insert(str)
			continue
		}

		tail := v.stack.Tail()
		brackets := str == "}" && tail == "{"
		squareBrackets := str == "]" && tail == "["
		parentheses := str == ")" && tail == "("

		if brackets || squareBrackets || parentheses {
			v.stack.Remove()
		}
	}

	return v.stack.IsEmpty()
}

func NewExpressionValidator() *ExpressionValidator {
	stack := NewStack(10)
	return &ExpressionValidator{stack: stack}
}

func main() {
	v := NewExpressionValidator()

	// c[d]
	// a{b[c]d}e
	// a{b(c]d}e
	// a[b{c}d]e}
	// a{b(c)

	fmt.Println(v.Validate("a{b(c)"))
}
