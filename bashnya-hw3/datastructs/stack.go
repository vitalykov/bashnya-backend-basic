package datastructs

import "errors"

var errEmptyStack = errors.New("Stack is empty")

type Stack[T any] struct {
	data []T
}

func (st Stack[T]) Size() int {
	return len(st.data)
}

func (st Stack[T]) IsEmpty() bool {
	return st.Size() == 0
}

func (st *Stack[T]) Push(el T) {
	st.data = append(st.data, el)
}

func (st Stack[T]) Top() (T, error) {
	var el T
	if st.IsEmpty() {
		return el, errEmptyStack
	}
	el = st.data[len(st.data)-1]
	return el, nil
}

func (st *Stack[T]) Pop() (T, error) {
	el, err := st.Top()
	if err == nil {
		st.data = st.data[:len(st.data)-1]
	}
	return el, err
}

func (st *Stack[T]) Clear() {
	st.data = make([]T, 0)
}
