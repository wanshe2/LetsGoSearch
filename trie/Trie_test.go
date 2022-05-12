package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := newTrie()
	str := "你好我是你爷爷"
	trie.insertString(str)
	runes := []rune(str)
	node := trie.findByRunes(runes)
	fmt.Println(node.string())
}
