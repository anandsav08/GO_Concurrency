package main

import "fmt"

type Node struct{
	isWord bool
	children [26]*Node
}

type Trie struct{
	root *Node
}

func NewTrie() Trie{
	trie := Trie{}
	trie.root = &Node{}
	return trie
}

func (trie *Trie) insert(word string){
	cur := trie.root
	for i:= range word{
		index := word[i] - 'a'
		if cur.children[index] == nil{
			cur.children[index] = &Node{isWord:false}
		}
		cur = cur.children[index]
	}
	cur.isWord = true
}

func (trie *Trie) search(word string) bool {
	cur := trie.root
	for i:=range word{
		index := word[i] - 'a'
		if cur.children[index] == nil{
			return false
		}
		cur = cur.children[index]
	}
	return cur.isWord
}

func main(){
	trie := NewTrie()
	trie.insert("bear")
	trie.insert("apple")
	trie.insert("beer")
	trie.insert("appspace")
	fmt.Println("*** [ New Trie created ] ***")
	fmt.Printf("TRIE : {'bear','apple','beer','appspace'} \n")
	fmt.Printf("Searching word %s - %v\n","apple",trie.search("apple"))
	fmt.Printf("Searching word %s - %v\n","bear",trie.search("bear"))
	fmt.Printf("Searching word %s - %v\n","goku",trie.search("goku"))
	fmt.Printf("Searching word %s - %v\n","nyu",trie.search("nyu"))
}