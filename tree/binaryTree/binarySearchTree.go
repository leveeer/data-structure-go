package binaryTree

import (
	"sync"
)

type Node struct {
	binaryTreeEntry
	times  int
	parent *Node
	left   *Node
	right  *Node
}

type BinarySearchTree struct {
	root *Node
	lock sync.Mutex
}

func (t *BinarySearchTree) Height(node interface{}) int {
	panic("implement me")
}

func (t *BinarySearchTree) LeftRightRotate(node interface{}) interface{} {
	panic("binarySearchTree haven't implement LeftRightRotate method")
}

func (t *BinarySearchTree) RightLeftRotate(node interface{}) interface{} {
	panic("binarySearchTree haven't implement RightLeftRotate method")
}

// IsNil 判断节点是否为nil
func (t *BinarySearchTree) IsNil(n interface{}) bool {
	return n == nil
}

// Search 根据key查找节点
func (t *BinarySearchTree) Search(key uint32) interface{} {
	for cur := t.root; cur != nil; {
		if cur.Key == key {
			return cur
		} else if key < cur.Key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return nil
}

func (t *BinarySearchTree) Insert(key uint32, value interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()
	node := new(Node)
	node.Key, node.Value = key, value

	if t.root == nil {
		t.root = node
	} else {
		target := t.root
		for cur := t.root; cur != nil; {
			target = cur
			if node.Key < cur.Key {
				cur = cur.left
			} else {
				cur = cur.right
			}
		}
		node.parent = target
		if node.Key < target.Key {
			target.left = node
		} else {
			target.right = node
		}
	}
}

// Delete TODO 由于go没有实现可重入锁，该方法还有待优化
/**
1、被删除节点是叶子结点
  只需要把被删除节点的父节点指向该节点的指针置为nil
2、被删除节点有且只有一个孩子
  将被删除节点父结点的指针指向要删除结点的孩子结点；
3、被删除节点有两个孩子
  - 找到被删除节点的后继节点
  - 复制后继节点到被删除节点的位置
  - 删除后继节点在原来的位置
*/
func (t *BinarySearchTree) Delete(key uint32) interface{} {
	//t.lock.Lock()
	//defer t.lock.Unlock()
	node := t.Search(key).(*Node)
	if node == nil {
		return nil
	}
	if node.left == nil && node.right == nil {
		if node.parent.right == node {
			node.parent.right = nil
		} else {
			node.parent.left = nil
		}
		node = nil
	} else if node.left == nil || node.right == nil {
		var reConnectedNode *Node
		if node.left != nil {
			reConnectedNode = node.left
		} else {
			reConnectedNode = node.right
		}
		if node.parent.right == node {
			node.parent.right = reConnectedNode
		} else {
			node.parent.left = reConnectedNode
		}
	} else {
		successor := t.Successor(node, t.root).(*Node)
		_key, _value := successor.Key, successor.Value
		t.Delete(successor.Key)
		node.Key, node.Value = _key, _value
	}
	return node
}

// Predecessor 获取给定节点的前驱节点
func (t *BinarySearchTree) Predecessor(node interface{}, root interface{}) interface{} {
	//（1）该节点有一个左子树，因此该节点的前驱节点是其左子树中最大的孩子.
	//（2）该节点没有左子树，则沿着该节点的父节点一直往上找，直至父节点的右节点等于该节点，那么这个父节点就是该节点的前驱节点
	n := node.(*Node)
	if n == nil {
		return nil
	}
	if n.left != nil {
		return t.Max(n.left)
	}
	cur := n
	for cur.parent != nil && cur.parent.right != cur {
		cur = cur.parent
	}
	return cur.parent
}

// Successor 获取给定节点的后继节点
func (t *BinarySearchTree) Successor(node interface{}, root interface{}) interface{} {
	//（1）该节点具有一个右子树，则该节点的后继节点是其右子树中最小的孩子.
	//（2）该节点没有右子树，则沿着该节点的父节点一直往上找，直至父节点的左节点等于该节点，那么这个父节点就是该节点的后继节点
	n := node.(*Node)
	if n == nil {
		return nil
	}
	if n.right != nil {
		return t.Min(n.right)
	}
	cur := n
	for cur.parent != nil && cur.parent.left != cur {
		cur = cur.parent
	}
	return cur.parent
}

func (t *BinarySearchTree) LeftRotate(node interface{}) interface{} {
	panic("binarySearchTree haven't implement LeftRotate method")
}

func (t *BinarySearchTree) RightRotate(node interface{}) interface{} {
	panic("binarySearchTree haven't implement RightRotate method")
}

func (t *BinarySearchTree) Min(node interface{}) interface{} {
	current := node.(*Node)
	for current.left != nil {
		current = current.left
	}
	return current
}

func (t *BinarySearchTree) Max(node interface{}) interface{} {
	current := node.(*Node)
	for current.right != nil {
		current = current.right
	}
	return current
}

func (t *BinarySearchTree) Root() interface{} {
	return t.root
}

// PreOrderTraverse 前序遍历  根节点 -> 左子树 -> 右子树
func (t *BinarySearchTree) PreOrderTraverse(node interface{}, nodeSlice []interface{}) []interface{} {
	n := node.(*Node)
	if n != nil {
		nodeSlice = append(nodeSlice, n.Value)
		if n.left != nil {
			nodeSlice = t.PreOrderTraverse(n.left, nodeSlice)
		}
		if n.right != nil {
			nodeSlice = t.PreOrderTraverse(n.right, nodeSlice)
		}
	}
	return nodeSlice
}

// PostOrderTraverse 后序遍历  左子树 —> 右子树 —> 根结点
func (t *BinarySearchTree) PostOrderTraverse(node interface{}, nodeSlice []interface{}) []interface{} {
	n := node.(*Node)
	if n != nil {
		if n.left != nil {
			nodeSlice = t.PostOrderTraverse(n.left, nodeSlice)
		}
		if n.right != nil {
			nodeSlice = t.PostOrderTraverse(n.right, nodeSlice)
		}
		nodeSlice = append(nodeSlice, n.Value)
	}
	return nodeSlice
}

// InOrderTraverse 中序遍历 左子树 —> 根结点 —> 右子树
func (t *BinarySearchTree) InOrderTraverse(node interface{}, nodeSlice []interface{}) []interface{} {
	n := node.(*Node)
	if n != nil {
		if n.left != nil {
			nodeSlice = t.InOrderTraverse(n.left, nodeSlice)
		}
		nodeSlice = append(nodeSlice, n.Value)
		if n.right != nil {
			nodeSlice = t.InOrderTraverse(n.right, nodeSlice)
		}
	}
	return nodeSlice
}

func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}
