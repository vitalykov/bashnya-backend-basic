package datastructs

import (
	"cmp"
	"testing"
)

func checkBSTParams[T cmp.Ordered](t *testing.T, tr BST[T], sz int, empty bool) {
	if gotEmpty := tr.IsEmpty(); empty != gotEmpty {
		t.Errorf("IsEmpty() = %v. Expected: %v", gotEmpty, empty)
	}
	if gotSz := tr.Size(); sz != gotSz {
		t.Errorf("Size() = %d. Expected: %d", gotSz, sz)
		return
	}
}

func checkFind[T cmp.Ordered](t *testing.T, tr BST[T], el T, found bool) {
	gotNode := tr.Find(el)
	if !found && gotNode != nil {
		t.Errorf("Find() = %v with val = %v. Expected: %v", gotNode, gotNode.val, nil)
		return
	}
	if gotNode == nil {
		return
	}
	if val := gotNode.val; val != el {
		t.Errorf("Find node with val = %v. Expected: %v", val, el)
	}
}

func checkDepth[T cmp.Ordered](t *testing.T, tr BST[T], d int) {
	if depth := tr.Depth(); depth != d {
		t.Errorf("Depth() = %d. Expected: %d", depth, d)
	}
}

func TestIntBST(t *testing.T) {
	var tr BST[int]
	nums := [...]int{3, 6, 4, 1, 5}
	checkBSTParams(t, tr, 0, true)
	checkFind(t, tr, 5, false)
	checkDepth(t, tr, 0)
	t.Run("Insert", func(t *testing.T) {
		for i, num := range nums {
			tr.Insert(num)
			checkBSTParams(t, tr, i+1, false)
			checkFind(t, tr, num, true)
		}
		checkDepth(t, tr, 4)
	})
	t.Run("ToSlice", func(t *testing.T) {
		expSl := []int{1, 3, 4, 5, 6}
		sl := tr.ToSlice()
		for i := range sl {
			if sl[i] != expSl[i] {
				t.Errorf("sl[%d] = %v. expSl[%d] = %v", i, sl[i], i, expSl[i])
			}
		}
	})
	t.Run("Remove", func(t *testing.T) {
		for i, num := range nums[:len(nums)-1] {
			tr.Remove(num)
			checkBSTParams(t, tr, len(nums)-1-i, false)
			checkFind(t, tr, num, false)
		}
		tr.Remove(nums[len(nums)-1])
		checkBSTParams(t, tr, 0, true)
		checkFind(t, tr, 5, false)
		checkDepth(t, tr, 0)
	})
}

func TestStringBST(t *testing.T) {
	var tr BST[string]
	strs := [...]string{"basic", "go", "c++", "assembler", "fortran"}
	checkBSTParams(t, tr, 0, true)
	checkFind(t, tr, "go", false)
	checkDepth(t, tr, 0)
	t.Run("Insert", func(t *testing.T) {
		for i, str := range strs {
			tr.Insert(str)
			checkBSTParams(t, tr, i+1, false)
			checkFind(t, tr, str, true)
		}
		checkDepth(t, tr, 4)
	})
	t.Run("ToSlice", func(t *testing.T) {
		expSl := []string{"assembler", "basic", "c++", "fortran", "go"}
		sl := tr.ToSlice()
		for i := range sl {
			if sl[i] != expSl[i] {
				t.Errorf("sl[%d] = %v. expSl[%d] = %v", i, sl[i], i, expSl[i])
			}
		}
	})
	t.Run("Remove", func(t *testing.T) {
		for i, str := range strs[:len(strs)-1] {
			tr.Remove(str)
			checkBSTParams(t, tr, len(strs)-1-i, false)
			checkFind(t, tr, str, false)
		}
		tr.Remove(strs[len(strs)-1])
		checkBSTParams(t, tr, 0, true)
		checkFind(t, tr, "go", false)
		checkDepth(t, tr, 0)
	})
}
