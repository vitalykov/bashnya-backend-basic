package datastructs

import (
	"testing"
)

func checkFront[T comparable](t *testing.T, dq Deque[T], val T, e error) {
	gotVal, err := dq.Front()
	if err != e {
		t.Errorf("Front() returns error %v. Expected: %v", err, e)
	}
	if gotVal != val {
		t.Errorf("Front() = %v. Expected: %v", gotVal, val)
	}
}

func checkBack[T comparable](t *testing.T, dq Deque[T], val T, e error) {
	gotVal, err := dq.Back()
	if err != e {
		t.Errorf("Back() returns error %v. Expected: %v", err, e)
	}
	if gotVal != val {
		t.Errorf("Back() = %v. Expected: %v", gotVal, val)
	}
}

func checkDequeParams[T any](t *testing.T, dq Deque[T], sz int, empty bool) {
	if gotEmpty := dq.IsEmpty(); empty != gotEmpty {
		t.Errorf("IsEmpty() = %v. Expected: %v", gotEmpty, empty)
	}
	if gotSz := dq.Size(); sz != gotSz {
		t.Errorf("Size() = %d. Expected: %d", gotSz, sz)
		return
	}
}

func checkPopFront[T comparable](t *testing.T, dq *Deque[T], val T, e error) {
	gotVal, err := dq.PopFront()
	if err != e {
		t.Errorf("PopFront() returns error %v. Expected: %v", err, e)
	}
	if gotVal != val {
		t.Errorf("PopFront() = %v. Expected: %v", gotVal, val)
	}
}

func checkPopBack[T comparable](t *testing.T, dq *Deque[T], val T, e error) {
	gotVal, err := dq.PopBack()
	if err != e {
		t.Errorf("PopBack() returns error %v. Expected: %v", err, e)
	}
	if gotVal != val {
		t.Errorf("PopBack() = %v. Expected: %v", gotVal, val)
	}
}

func TestIntDeque(t *testing.T) {
	var dq Deque[int]
	nums := [...]int{5, 4, 3, 2, 1}
	checkDequeParams(t, dq, 0, true)
	checkBack(t, dq, 0, errEmptyDeque)
	checkFront(t, dq, 0, errEmptyDeque)

	t.Run("PushBack and PopBack", func(t *testing.T) {
		for i, num := range nums {
			dq.PushBack(num)
			checkDequeParams(t, dq, i+1, false)
			checkBack(t, dq, num, nil)
			checkFront(t, dq, nums[0], nil)
		}
		for i := len(nums) - 1; i > 0; i-- {
			checkPopBack(t, &dq, nums[i], nil)
			checkBack(t, dq, nums[i-1], nil)
			checkFront(t, dq, nums[0], nil)
			checkDequeParams(t, dq, i, false)
		}
		checkPopBack(t, &dq, nums[0], nil)
		checkDequeParams(t, dq, 0, true)
		checkBack(t, dq, 0, errEmptyDeque)
		checkFront(t, dq, 0, errEmptyDeque)
		checkPopBack(t, &dq, 0, errEmptyDeque)
	})

	t.Run("PushFront and PopFront", func(t *testing.T) {
		for i, num := range nums {
			dq.PushFront(num)
			checkDequeParams(t, dq, i+1, false)
			checkFront(t, dq, num, nil)
			checkBack(t, dq, nums[0], nil)
		}
		for i := len(nums) - 1; i > 0; i-- {
			checkPopFront(t, &dq, nums[i], nil)
			checkFront(t, dq, nums[i-1], nil)
			checkBack(t, dq, nums[0], nil)
			checkDequeParams(t, dq, i, false)
		}
		checkPopFront(t, &dq, nums[0], nil)
		checkDequeParams(t, dq, 0, true)
		checkFront(t, dq, 0, errEmptyDeque)
		checkBack(t, dq, 0, errEmptyDeque)
		checkPopFront(t, &dq, 0, errEmptyDeque)
	})

	t.Run("PushBack and PopFront", func(t *testing.T) {
		for i, num := range nums {
			dq.PushBack(num)
			checkDequeParams(t, dq, i+1, false)
			checkBack(t, dq, num, nil)
			checkFront(t, dq, nums[0], nil)
		}
		for i, num := range nums[:len(nums)-1] {
			checkPopFront(t, &dq, num, nil)
			checkFront(t, dq, nums[i+1], nil)
			checkBack(t, dq, nums[len(nums)-1], nil)
			checkDequeParams(t, dq, len(nums)-i-1, false)
		}
		checkPopFront(t, &dq, nums[len(nums)-1], nil)
		checkDequeParams(t, dq, 0, true)
		checkBack(t, dq, 0, errEmptyDeque)
		checkFront(t, dq, 0, errEmptyDeque)
		checkPopBack(t, &dq, 0, errEmptyDeque)
	})
}

func TestStringDeque(t *testing.T) {
	var dq Deque[string]
	strs := [...]string{"hello", "", "bashnya", "homework", "3"}
	checkDequeParams(t, dq, 0, true)
	checkBack(t, dq, "", errEmptyDeque)
	checkFront(t, dq, "", errEmptyDeque)

	t.Run("PushBack and PopBack", func(t *testing.T) {
		for i, s := range strs {
			dq.PushBack(s)
			checkDequeParams(t, dq, i+1, false)
			checkBack(t, dq, s, nil)
			checkFront(t, dq, strs[0], nil)
		}
		for i := len(strs) - 1; i > 0; i-- {
			checkPopBack(t, &dq, strs[i], nil)
			checkBack(t, dq, strs[i-1], nil)
			checkFront(t, dq, strs[0], nil)
			checkDequeParams(t, dq, i, false)
		}
		checkPopBack(t, &dq, strs[0], nil)
		checkDequeParams(t, dq, 0, true)
		checkBack(t, dq, "", errEmptyDeque)
		checkFront(t, dq, "", errEmptyDeque)
		checkPopBack(t, &dq, "", errEmptyDeque)
	})

	t.Run("PushFront and PopFront", func(t *testing.T) {
		for i, s := range strs {
			dq.PushFront(s)
			checkDequeParams(t, dq, i+1, false)
			checkFront(t, dq, s, nil)
			checkBack(t, dq, strs[0], nil)
		}
		for i := len(strs) - 1; i > 0; i-- {
			checkPopFront(t, &dq, strs[i], nil)
			checkFront(t, dq, strs[i-1], nil)
			checkBack(t, dq, strs[0], nil)
			checkDequeParams(t, dq, i, false)
		}
		checkPopFront(t, &dq, strs[0], nil)
		checkDequeParams(t, dq, 0, true)
		checkFront(t, dq, "", errEmptyDeque)
		checkBack(t, dq, "", errEmptyDeque)
		checkPopFront(t, &dq, "", errEmptyDeque)
	})

	t.Run("PushBack and PopFront", func(t *testing.T) {
		for i, num := range strs {
			dq.PushBack(num)
			checkDequeParams(t, dq, i+1, false)
			checkBack(t, dq, num, nil)
			checkFront(t, dq, strs[0], nil)
		}
		for i, num := range strs[:len(strs)-1] {
			checkPopFront(t, &dq, num, nil)
			checkFront(t, dq, strs[i+1], nil)
			checkBack(t, dq, strs[len(strs)-1], nil)
			checkDequeParams(t, dq, len(strs)-i-1, false)
		}
		checkPopFront(t, &dq, strs[len(strs)-1], nil)
		checkDequeParams(t, dq, 0, true)
		checkBack(t, dq, "", errEmptyDeque)
		checkFront(t, dq, "", errEmptyDeque)
		checkPopBack(t, &dq, "", errEmptyDeque)
	})
}
