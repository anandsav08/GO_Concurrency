// Binary Search Tree
package main

import "fmt"
type Node struct{
	left *Node
	right *Node
	key int
}

type BST struct{
	root *Node
}

func (tree *BST) insert(k int,root *Node) *Node{
	
	if root == nil{
		root := &Node{left:nil,right:nil,key:k}
		return root
	}
	if root.key < k{
		root.right = tree.insert(k,root.right)
	} else {
		root.left = tree.insert(k,root.left)
	}
	return root
}

func (tree *BST) print(root *Node){
	if root==nil{
		return
	}
	tree.print(root.left)
	fmt.Printf("%d\t",root.key)
	tree.print(root.right)
}

func main(){
	bst := BST{}
	bst.root = bst.insert(10,bst.root)
	bst.root = bst.insert(20,bst.root)
	bst.root = bst.insert(5,bst.root)
	bst.root = bst.insert(1,bst.root)
	bst.root = bst.insert(25,bst.root)
	bst.print(bst.root)
}