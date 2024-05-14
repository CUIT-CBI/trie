package main

// import "encoding/json"

// type Node struct {
// 	Prefix   string         `json:"prefix"`
// 	Children map[byte]*Node `json:"children"`
// }

// func BuildTrie(words []string) *Node {
// 	root := &Node{
// 		Prefix:   "",
// 		Children: make(map[byte]*Node),
// 	}
// 	for _, s := range words {
// 		cur := root

// 		for i := 0; i < len(s); i++ {
// 			c := s[i]
// 			if cur.Children[c] == nil {
// 				cur.Children[c] = &Node{
// 					Prefix:   cur.Prefix + string(c),
// 					Children: make(map[byte]*Node),
// 				}
// 			}
// 			cur = cur.Children[c]
// 		}
// 	}
// 	return root
// }

// func (node Node) String() string {
// 	d, _ := json.Marshal(node)
// 	return string(d)
// }
