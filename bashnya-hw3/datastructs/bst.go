package datastructs

import (
	"cmp"
)

type node[T cmp.Ordered] struct {
	val   T
	left  *node[T]
	right *node[T]
}

type BST[T cmp.Ordered] struct {
	root *node[T]
	size int
}

func (t BST[T]) IsEmpty() bool {
	return t.root == nil
}

func (t BST[T]) Size() int {
	return t.size
}

func (t BST[T]) Find(el T) *node[T] {
	cur := t.root
	for cur != nil {
		if cur.val == el {
			return cur
		}
		if el < cur.val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return nil
}

func (t *BST[T]) Insert(el T) {
	t.size++
	if t.root == nil {
		t.root = &node[T]{val: el, left: nil, right: nil}
		return
	}
	cur := t.root
	for cur != nil {
		if el <= cur.val {
			if cur.left == nil {
				cur.left = &node[T]{val: el, left: nil, right: nil}
				return
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &node[T]{val: el, left: nil, right: nil}
				return
			}
			cur = cur.right
		}
	}
}

func (t BST[T]) findWithParent(el T) (*node[T], *node[T]) {
	cur := t.root
	var parent *node[T]
	for cur != nil {
		if cur.val == el {
			break
		}
		parent = cur
		if el < cur.val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return cur, parent
}

func (t *BST[T]) Remove(el T) {
	cur, parent := t.findWithParent(el)
	if cur == nil {
		return
	}
	var onLeft bool
	if parent != nil && cur == parent.left {
		onLeft = true
	}
	right := cur.right
	left := cur.left
	var subNode *node[T]
	if left != nil {
		subNode = left
		subNodeParent := cur
		for subNode.right != nil {
			subNodeParent = subNode
			subNode = subNode.right
		}
		if subNode != left {
			// Using leftCopy to avoid GC from deleting subNode.left
			leftCopy := subNode.left
			subNode.left = left
			subNodeParent.right = leftCopy
		}
		subNode.right = right
	} else {
		subNode = right
	}
	if parent == nil {
		t.root = subNode
	} else {
		if onLeft {
			parent.left = subNode
		} else {
			parent.right = subNode
		}
	}
	t.size--
}

func (t BST[T]) Depth() int {
	var depth int
	var findDepth func(nd *node[T], d int)
	findDepth = func(nd *node[T], d int) {
		if nd == nil {
			depth = max(depth, d)
			return
		}
		findDepth(nd.left, d+1)
		findDepth(nd.right, d+1)
	}
	findDepth(t.root, 0)
	return depth
}

func (t BST[T]) ToSlice() []T {
	var sl []T
	if t.IsEmpty() {
		return sl
	}
	var writeToSlice func(nd *node[T])
	writeToSlice = func(nd *node[T]) {
		if nd.left != nil {
			writeToSlice(nd.left)
		}
		sl = append(sl, nd.val)
		if nd.right != nil {
			writeToSlice(nd.right)
		}
	}
	writeToSlice(t.root)
	return sl
}
