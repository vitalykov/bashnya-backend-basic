package datastructs

import (
	"testing"
)

func checkTop[T comparable](t *testing.T, st Stack[T], val T, e error) {
	gotVal, err := st.Top()
	if err != e {
		t.Errorf("Top() returns error %v. Expected: %v", err, e)
	}
	if gotVal != val {
		t.Errorf("Top() = %v. Expected: %v", gotVal, val)
	}
}

func checkStackParams[T any](t *testing.T, st Stack[T], sz int, empty bool) {
	if gotEmpty := st.IsEmpty(); empty != gotEmpty {
		t.Errorf("IsEmpty() = %v. Expected: %v", gotEmpty, empty)
	}
	if gotSz := st.Size(); sz != gotSz {
		t.Errorf("Size() = %d. Expected: %d", gotSz, sz)
		return
	}
}

func checkPop[T comparable](t *testing.T, st *Stack[T], val T, e error) {
	gotVal, err := st.Pop()
	if err != e {
		t.Errorf("Pop() returns error %v. Expected: %v", err, e)
	}
	if gotVal != val {
		t.Errorf("Pop() = %v. Expected: %v", gotVal, val)
	}
}

func TestIntStack(t *testing.T) {
	var st Stack[int]
	nums := [...]int{5, 4, 3, 2, 1}
	checkStackParams(t, st, 0, true)
	checkTop(t, st, 0, errEmptyStack)
	for i, num := range nums {
		st.Push(num)
		checkStackParams(t, st, i+1, false)
		checkTop(t, st, num, nil)
	}
	for i := len(nums) - 1; i > 0; i-- {
		checkPop(t, &st, nums[i], nil)
		checkTop(t, st, nums[i-1], nil)
		checkStackParams(t, st, i, false)
	}
	checkPop(t, &st, nums[0], nil)
	checkStackParams(t, st, 0, true)
	checkTop(t, st, 0, errEmptyStack)
	checkPop(t, &st, 0, errEmptyStack)
}

func TestStringStack(t *testing.T) {
	var st Stack[string]
	strs := [...]string{"hello", "", "bashnya", "homework", "3"}
	checkStackParams(t, st, 0, true)
	checkTop(t, st, "", errEmptyStack)
	for i, s := range strs {
		st.Push(s)
		checkStackParams(t, st, i+1, false)
		checkTop(t, st, s, nil)
	}
	for i := len(strs) - 1; i > 0; i-- {
		checkPop(t, &st, strs[i], nil)
		checkTop(t, st, strs[i-1], nil)
		checkStackParams(t, st, i, false)
	}
	checkPop(t, &st, strs[0], nil)
	checkStackParams(t, st, 0, true)
	checkTop(t, st, "", errEmptyStack)
	checkPop(t, &st, "", errEmptyStack)
}
