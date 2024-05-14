package main

import "encoding/json"

type Node struct {
	Prefix   string         `json:"prefix" rlp:""`
	Children map[byte]*Node `json:"children"`
}

func BuildTrie(words []string) *Node {
	root := &Node{
		Prefix:   "",
		Children: make(map[byte]*Node),
	}

	for _, word := range words {
		root.Insert(word)
	}

	return root
}

func (node *Node) Insert(word string) {
	if len(word) == 0 {
		return
	}

	letter := word[0]
	child := node.Children[letter]
	if child == nil {
		prefix := node.Prefix + string(letter)
		child = &Node{
			Prefix:   prefix,
			Children: make(map[byte]*Node),
		}
		node.Children[letter] = child
	}
	child.Insert(word[1:])
}

func (node Node) Exist(word string) bool {
	// TODO
	return false
}

func (node *Node) Delete(word string) {
	// TODO
}

func (node Node) FindAll(prefix string) []string {
	// TODO
	return nil
}

func (node Node) String() string {
	d, _ := json.Marshal(node)
	return string(d)
}
