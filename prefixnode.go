package prefix

//Node in the trie
type Node struct {
	char     rune
	children map[rune]*Node
	users    *[]string
}

//NewNode creates a n with character and terminal value
func NewNode(char rune) Node {

	n := Node{
		char:     char,
		children: make(map[rune]*Node),
		users:    &[]string{},
	}

	return n
}

//GetChar associated with node
func (n Node) GetChar() rune {
	return n.char
}

//HasChild if n has child with given rune value
func (n Node) HasChild(char rune) bool {

	_, found := n.children[char]

	if found {
		return true
	}
	return false
}

//AddChild Node to parents children
func (n Node) AddChild(child *Node) bool {

	if n.HasChild(child.char) {
		return false
	}
	//Add Node
	n.children[child.char] = child

	return true
}

//GetChild Returns a pointer to child object
func (n Node) GetChild(char rune) *Node {
	if n.HasChild(char) {
		return n.children[char]
	}

	return nil
}

//GetChildren of a given n
func (n Node) GetChildren() map[rune]*Node {
	return n.children
}

//AddUser to usernames
func (n Node) AddUser(name string) {
	*n.users = append(*n.users, name)
}

//GetUsers for a given node
func (n Node) GetUsers() []string {
	return *n.users
}
