// Copyright (c) 2022, wanshe li. All rights reserved

package trie

import "fmt"

// Node is a single element within the Trie
type Node struct {
	data  rune
	child map[rune]*Node
	flag  bool // 当前是否为词语的结尾
	count int  // 累计查找次数(用于排序)
	//countToday int // 当日查找次数
}

// get a new Node
func newNode() *Node {
	return &Node{child: make(map[rune]*Node)}
}

// string returns a string describe Node
func (n *Node) string() string {
	return fmt.Sprintf("Node: {data: %s, flag: %v, count: %d, child: %v}", string(n.data), n.flag, n.count, n.child)
}

// Trie holds elements of the Trie tree.
type Trie struct {
	size int
	root *Node
}

// get a new Trie
func newTrie() *Trie {
	return &Trie{root: newNode()}
}

// string returns a string representation of container
func (t *Trie) string() string {
	str := "Trie\n"
	if t.root != nil {
		output(t.root, &str)
	}
	return str
}

func output(rt *Node, str *string) {

}

// get new root Node and set size to zero
func (t *Trie) init() {
	t.size = 0
	t.root = newNode()
}

// change root Node to nil
func (t *Trie) clear() {
	t.size = 0
	t.root = nil // 自动 GC
}

// insert bytes into Trie
func (t *Trie) insertBytes(bytes []byte) {
	str := string(bytes)
	t.insertString(str)
}

// insert string into Trie
func (t *Trie) insertString(str string) {
	runes := []rune(str)
	t.insertRunes(runes)
}

// insert runes into Trie
func (t *Trie) insertRunes(runes []rune) {
	if len(runes) == 0 {
		return
	}
	now := t.root
	for _, val := range runes {
		nxt, ok := now.child[val]
		if !ok {
			nxt = newNode()
			nxt.data = val
			now.child[val] = nxt
			t.size += 1
		}
		now = nxt
		now.count += 1
	}
	now.flag = true
}

// find the last character in string from Trie.
// If not find, return nil
func (t *Trie) findByString(str string) *Node {
	runes := []rune(str)
	return t.findByRunes(runes)
}

// find the last character in bytes from Trie.
// If not find, return nil
func (t *Trie) findByBytes(bytes []byte) *Node {
	str := string(bytes)
	return t.findByString(str)
}

// find the last character in runes from Trie
// If not find, return nil
func (t *Trie) findByRunes(runes []rune) *Node {
	if len(runes) == 0 {
		return nil
	}
	now := t.root
	for _, val := range runes {
		nxt, ok := now.child[val]
		if !ok {
			return nil
		}
		now = nxt
		now.count += 1
	}
	return now
}
