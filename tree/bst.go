package tree

// Node is the binary search tree node representation.
type Node struct {
	Value, Count int
	Left, Right  *Node
}

// BST is the binary search tree object
type BST struct {
	Root *Node
}

// NewBST returns a BST with null root
func NewBST(values []int) *BST {
	return &BST{}
}

// Insert inserts value into the binary search tree
func (t *BST) Insert(value int) {
	t.Root = insert(t.Root, value)
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value, Count: 1}
	}
	if root.Value > value {
		root = insert(root.Left, value)
	} else if root.Value < value {
		root = insert(root.Right, value)
	} else {
		root.Value = value
	}
	root.Count = size(root.Left) + size(root.Right) + 1
	return root
}

func size(root *Node) int {
	if root == nil {
		return 0
	}
	return root.Count
}

// Get returns binary search tree node by given value or null if not found
func (t *BST) Get(value int) *Node {
	return get(t.Root, value)
}

func get(root *Node, value int) *Node {
	if root == nil {
		return nil
	}
	if root.Value > value {
		return get(root.Left, value)
	} else if root.Value < value {
		return get(root.Right, value)
	}
	return root
}

// Delete deletes binary search tree node by given value
func (t *BST) Delete(value int) {
	t.Root = delete(t.Root, value)
}

func delete(root *Node, value int) *Node {
	if root == nil {
		return nil
	}
	if root.Value > value {
		return delete(root.Left, value)
	} else if root.Value < value {
		return delete(root.Right, value)
	}
	min := min(root.Right)
	min.Left = root.Left
	min.Right = deleteMin(root.Right)
	min.Count = size(min.Left) + size(min.Right) + 1
	return min
}

func min(root *Node) *Node {
	for root != nil {
		root = root.Left
	}
	return root
}

func deleteMin(root *Node) *Node {
	if root == nil || root.Left == nil {
		return nil
	}
	root.Left = deleteMin(root.Left)
	root.Count = size(root.Left) + size(root.Right) + 1
	return root
}
