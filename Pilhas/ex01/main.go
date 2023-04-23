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
		if str == "{" || str == "}" {
			v.stack.Insert(str)
			continue
		}
		if str == "[" || str == "]" {
			v.stack.Insert(str)
			continue
		}
		if str == "(" || str == ")" {
			v.stack.Insert(str)
		}
	}

	vector := v.stack.Vector()

	for i := 0; i < v.stack.lastPos; i++ {
		str := vector[i]
		if str == "{" && v.stack.Tail() == "}" {
			v.stack.Remove()
			continue
		}

		if str == "[" && v.stack.Tail() == "]" {
			v.stack.Remove()
			continue
		}

		if str == "(" && v.stack.Tail() == ")" {
			v.stack.Remove()
			continue
		}
	}

	for _, strRune := range expression {
		str := strings.Trim(string(strRune), " ")
		if str == "{" || str == "[" || str == "(" {
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
	fmt.Println(v.Validate("a{b[c]d}(e)f"))

}
