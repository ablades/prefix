package prefix

import (
	"testing"
)

//test creation of a node object
func TestNewNode(t *testing.T) {

	var tests = []struct {
		char rune // input
	}{
		{'a'},
		{'b'},
		{'c'},
		{'d'},
		{'e'},
		{'A'},
		{'B'},
		{'@'},
	}

	for _, test := range tests {

		node := NewNode(test.char)

		//Check Character
		if node.GetChar() != test.char {
			t.Errorf("Expected Character: %c, got %c instead.", test.char, node.GetChar())
		}

		//Check Children Node length
		if len(node.GetChildren()) != 0 {
			t.Errorf("Expected children nodes of length %v, instead got length %v.", 0, len(node.GetChildren()))
		}

		//Check count of users
		if len(node.GetUsers()) != 0 {
			t.Errorf("Expected users number of users %v, found %v users instead.", 0, node.GetUsers())
		}
	}
}

func TestAddChild(t *testing.T) {
	node := NewNode('a')
	//Insert child
	child := NewNode('b')
	result := node.AddChild(&child)
	if result != true {
		t.Errorf("Expected insertion of %v and return value of true.", child)
	}

	//Test duplicate insertion
	child1 := NewNode('b')
	result1 := node.AddChild(&child1)
	if result1 != false {
		t.Errorf("Expected insertion to be invalid for %v", child1)
	}
}

func TestGetChild(t *testing.T) {
	node := NewNode('a')
	//Insert child
	child := NewNode('b')
	node.AddChild(&child)

	//
	result := node.GetChild('b')
	if result != &child {
		t.Errorf("Expected return of prefix node %v. Instead got value, %v", child, result)
	}

	//
	result1 := node.GetChild('c')
	if result1 != nil {
		t.Errorf("Expected nil value, got %v instead", result1)
	}

	node = child
	child1 := NewNode('c')
	node.AddChild(&child1)

	result2 := node.GetChild('c')
	if result2 != &child1 {
		t.Errorf("Expected node, got %v instead", result2)
	}
}

func TestAddUser(t *testing.T) {
	node := NewNode('a')

	node.AddUser("Test")
	t.Log(node.GetUsers())

	if len(node.GetUsers()) != 1 {
		t.Errorf("Expected size of 1, got size %v instead", len(node.GetUsers()))
	}

	testSlice := []string{"Test"}
	for i := range node.GetUsers() {
		if node.GetUsers()[i] != testSlice[i] {
			t.Errorf("Value: %v From Users  and %v from testSlice do not match", node.GetUsers()[i], testSlice[i])
		}
	}
}
