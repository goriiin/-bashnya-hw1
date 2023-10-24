package binaryTree

type node struct {
	data   int
	pLeft  *node
	pRight *node
}

type BinarySearchTree struct {
	root *node
}

func newNode(data int) *node {
	return &node{
		data,
		nil,
		nil,
	}
}

func (bst *BinarySearchTree) Insert(data int) {
	if bst.root == nil {
		bst.root = newNode(data)
	} else {
		insertNode(bst.root, data)
	}
}

func insertNode(node *node, data int) {
	if data < node.data {
		if node.pLeft == nil {
			node.pLeft = newNode(data)
		} else {
			insertNode(node.pLeft, data)
		}
	} else {
		if node.pRight == nil {
			node.pRight = newNode(data)
		} else {
			insertNode(node.pRight, data)
		}
	}
}

func (bst *BinarySearchTree) Search(data int) bool {
	return searchNode(bst.root, data)
}

func searchNode(node *node, data int) bool {
	if node == nil {
		return false
	}
	if data == node.data {
		return true
	} else if data < node.data {
		return searchNode(node.pLeft, data)
	} else {
		return searchNode(node.pRight, data)
	}
}

func (bst *BinarySearchTree) Balance() {
	nodes := bst.getNodesInOrder()
	bst.root = bst.buildBalancedTree(nodes, 0, len(nodes)-1)
}

func (bst *BinarySearchTree) getNodesInOrder() []*node {
	var nodes []*node
	bst.traverseInOrder(bst.root, func(node *node) {
		nodes = append(nodes, node)
	})
	return nodes
}

func (bst *BinarySearchTree) traverseInOrder(node *node, callback func(*node)) {
	if node == nil {
		return
	}
	bst.traverseInOrder(node.pLeft, callback)
	callback(node)
	bst.traverseInOrder(node.pRight, callback)
}

func (bst *BinarySearchTree) buildBalancedTree(nodes []*node, start, end int) *node {
	if start > end {
		return nil
	}
	middle := (start + end) / 2
	node := nodes[middle]
	node.pLeft = bst.buildBalancedTree(nodes, start, middle-1)
	node.pRight = bst.buildBalancedTree(nodes, middle+1, end)
	return node
}
