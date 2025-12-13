package datastructs

import (
	"errors"
	"slices"
)

var errEmptyDeque = errors.New("Deque is empty")

type Deque[T any] struct {
	frontData []T
	backData  []T
}

func (dq Deque[T]) Size() int {
	return len(dq.backData) + len(dq.frontData)
}

func (dq Deque[T]) IsEmpty() bool {
	return dq.Size() == 0
}

func (dq *Deque[T]) PushFront(el T) {
	dq.frontData = append(dq.frontData, el)
}

func (dq *Deque[T]) PushBack(el T) {
	dq.backData = append(dq.backData, el)
}

func (dq Deque[T]) Back() (T, error) {
	var el T
	if dq.IsEmpty() {
		return el, errEmptyDeque
	}
	backLen := len(dq.backData)
	if backLen != 0 {
		el = dq.backData[backLen-1]
	} else {
		el = dq.frontData[0]
	}
	return el, nil
}

func (dq Deque[T]) Front() (T, error) {
	var el T
	if dq.IsEmpty() {
		return el, errEmptyDeque
	}
	frontLen := len(dq.frontData)
	if frontLen != 0 {
		el = dq.frontData[frontLen-1]
	} else {
		el = dq.backData[0]
	}
	return el, nil
}

func (dq *Deque[T]) PopBack() (T, error) {
	el, err := dq.Back()
	if err != nil {
		return el, err
	}
	backLen := len(dq.backData)
	if backLen != 0 {
		dq.backData = dq.backData[:backLen-1]
	} else {
		backLen = len(dq.frontData)/2 + 1
		dq.backData = dq.frontData[1:backLen]
		slices.Reverse(dq.backData)
		dq.frontData = dq.frontData[backLen:]
	}
	return el, nil
}

func (dq *Deque[T]) PopFront() (T, error) {
	el, err := dq.Front()
	if err != nil {
		return el, err
	}
	frontLen := len(dq.frontData)
	if frontLen != 0 {
		dq.frontData = dq.frontData[:frontLen-1]
	} else {
		frontLen = len(dq.backData)/2 + 1
		dq.frontData = dq.backData[1:frontLen]
		slices.Reverse(dq.frontData)
		dq.backData = dq.backData[frontLen:]
	}
	return el, nil
}

func (dq *Deque[T]) Clear() {
	dq.frontData = make([]T, 0)
	dq.backData = make([]T, 0)
}
