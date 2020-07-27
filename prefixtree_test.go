package prefix

import (
	"sort"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := NewTree("Test")

	//Check Name
	if tree.GetName() != "Test" {
		t.Errorf("Expected Name: Test: %v instead.", tree.GetName())
	}

	//Check size
	if tree.GetSize() != 0 {
		t.Errorf("Expected Size: 0, instead had size %v", tree.GetSize())
	}

	//Get Root rune
	if tree.GetRoot().GetChar() != '0' {
		t.Errorf("Expected Rune 0 instead had rune %v", tree.GetRoot().GetChar())
	}

}

func TestInsertKeyword(t *testing.T) {
	tree := NewTree("Test")

	tree.InsertKeyword("apple", "becky")

	//afmt.Println(tree.Contains("apple"))
}

func TestPrintTrie(t *testing.T) {
	tree := NewTree("Test")
	tree.InsertKeyword("pear", "becky")
	tree.InsertKeyword("able", "karen")
	tree.InsertKeyword("bear", "susan")
	tree.InsertKeyword("mamale", "getrude")
	testSlice := []string{"able", "bear", "mamale", "pear"}
	trieElements := tree.PrintTrie()
	sort.Strings(trieElements)

	for i := range testSlice {
		if testSlice[i] != trieElements[i] {
			t.Errorf("Slices do not contain same elements %v != %v", trieElements[i], testSlice[i])
		}
	}
}
