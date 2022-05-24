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

func (tree *BST) search(k int,root *Node) bool {
	if root==nil{
		return false
	}

	if root.key == k{
		return true
	} else if(root.key < k){
		return tree.search(k,root.right)
	} else{
		return tree.search(k,root.left)
	}
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

	k1 := 10
	k2 := 100
	fmt.Printf("\nSearching %d - %v\n",k1,bst.search(k1,bst.root))
	fmt.Printf("\nSearching %d - %v\n",k2,bst.search(k2,bst.root))
}