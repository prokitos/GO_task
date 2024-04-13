package tasks

import "fmt"

func MainE() {

	// мини бинарное дерево
	var tree BTree
	tree.addNode(5)
	tree.addNode(7)
	tree.addNode(1)
	tree.addNode(8)

	tree.delNode(1)
	tree.showTree()
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

type BTree struct {
	root *TreeNode
}

// добавление новой ноды
func (tree *BTree) addNode(val int) {

	var curNode TreeNode
	curNode.val = val

	if tree.root == nil {
		tree.root = &curNode
		return
	}

	tree.recurAddNode(tree.root, &curNode)

}

// добавление ноды рекурсивно
func (tree *BTree) recurAddNode(curNode *TreeNode, myNode *TreeNode) {

	if myNode == nil {
		return
	}

	if myNode.val >= curNode.val {
		if curNode.right == nil {
			curNode.right = myNode
			return
		} else {
			tree.recurAddNode(curNode.right, myNode)
			return
		}
	} else {
		if curNode.left == nil {
			curNode.left = myNode
			return
		} else {
			tree.recurAddNode(curNode.left, myNode)
			return
		}
	}
}

// удаление ноды
func (tree *BTree) delNode(val int) {
	if tree.root.val == val {
		if tree.root.right != nil {
			temp := tree.root
			tree.root = tree.root.right
			tree.recurAddNode(tree.root, temp.left)
			return
		}
		tree.root = tree.root.left
		return
	}

	tree.delRecur(tree.root, val)
}

// удаление ноды рекурсивно
func (tree *BTree) delRecur(curNode *TreeNode, val int) {

	if curNode == nil {
		return
	}
	if curNode.right.val == val {
		if curNode.right.right == nil {
			curNode.right = curNode.right.left
			return
		}
		temp := curNode.right.left
		curNode.right = curNode.right.right
		tree.recurAddNode(tree.root, temp)
		return
	}

	if curNode.left.val == val {
		if curNode.left.right == nil {
			curNode.left = curNode.left.left
			return
		}
		temp := curNode.left.left
		curNode.left = curNode.left.right
		tree.recurAddNode(tree.root, temp)
		return

	}

	if val > curNode.val {
		tree.delRecur(curNode.right, val)
		return
	}
	if val < curNode.val {
		tree.delRecur(curNode.left, val)
		return
	}
}

// показать все дерево
func (tree *BTree) showTree() {
	tree.showRecur(tree.root)
}

// рекурсивный показ
func (tree *BTree) showRecur(curNode *TreeNode) {
	if curNode != nil {
		tree.showRecur(curNode.left)
		fmt.Print(curNode.val, " ")
		tree.showRecur(curNode.right)
	}
}
