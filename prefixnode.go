package prefix

//Node in the trie
type Node struct {
	Char     rune
	Children map[rune]*Node
	Users    *[]string
}

//NewNode creates a n with character and terminal value
func NewNode(char rune) Node {

	n := Node{
		Char:     char,
		Children: make(map[rune]*Node),
		Users:    &[]string{},
	}

	return n
}

//GetChar associated with node
func (n Node) GetChar() rune {
	return n.Char
}

//HasChild if n has child with given rune value
func (n Node) HasChild(char rune) bool {

	_, found := n.Children[char]

	if found {
		return true
	}
	return false
}

//AddChild Node to parents children
func (n Node) AddChild(child *Node) bool {

	if n.HasChild(child.Char) {
		return false
	}
	//Add Node
	n.Children[child.Char] = child

	return true
}

//GetChild Returns a pointer to child object
func (n Node) GetChild(char rune) *Node {
	if n.HasChild(char) {
		return n.Children[char]
	}

	return nil
}

//GetChildren of a given n
func (n Node) GetChildren() map[rune]*Node {
	return n.Children
}

//AddUser to usernames
func (n Node) AddUser(name string) {
	*n.Users = append(*n.Users, name)
}

//GetUsers for a given node
func (n Node) GetUsers() []string {
	return *n.Users
}
