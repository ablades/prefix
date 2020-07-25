package prefix

import (
	"fmt"
)

//PrefixTree for nodes
type PrefixTree struct {
	Name string
	Root *Node
	Size int
}

//NewTree tree with a default root node
func NewTree(name string) PrefixTree {
	rootNode := NewNode('0')

	tree := PrefixTree{
		Name: name,
		Root: &rootNode,
		Size: 0,
	}

	return tree
}

//InsertKeyword TODO: Replace USER STRING WITH USER OBJECT
func (tree PrefixTree) InsertKeyword(word string, userName string) {
	node := tree.Root

	for _, char := range word {
		//Node does not have child with character
		if !node.HasChild(char) {
			child := NewNode(char)
			node.AddChild(&child)
			fmt.Printf("%v \n", node)
			tree.Size++
		}

		//Child already exists advance pointer
		node = node.GetChild(char)
	}
	fmt.Printf("%v \n", node)
	fmt.Printf("%c \n", node.GetChar())
	node.AddUser(userName)
	fmt.Println(node.GetUsers())

}

//Contains returns list of users associated with word
func (tree PrefixTree) Contains(word string) []string {
	node := tree.Root

	//Loop until end of word is hit
	for _, char := range word {
		if node.HasChild(char) {
			node = node.GetChild(char)
		}
	}
	fmt.Printf("%v", node.GetUsers())
	return node.GetUsers()
}

//Getters and setters

//GetName returns the name of the tree
func (tree PrefixTree) GetName() string {
	return tree.Name
}

//GetSize returns size of current tree
func (tree PrefixTree) GetSize() int {
	return tree.Size
}

//GetRoot returns rootnode in a tree
func (tree PrefixTree) GetRoot() *Node {
	return tree.Root
}

//PrintTrie all words
func (tree PrefixTree) PrintTrie() []string {
	var words []string

	//traverse through a tree
	tree.traverse(tree.Root, "", &words)

	return words

}

//traverse the tree
func (tree PrefixTree) traverse(node *Node, currentString string, list *[]string) {
	//if users then we've hit a terminal character
	if len(node.GetUsers()) > 0 {
		*list = append(*list, currentString)
	}

	//recursively traverse children
	for _, n := range node.GetChildren() {
		tree.traverse(n, currentString+string(n.GetChar()), list)
	}
}
