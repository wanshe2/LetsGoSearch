package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	str := "你好我是你爷爷"
	trie.InsertString(str)
	runes := []rune(str)
	node, _ := trie.findByRunes(runes)
	fmt.Println(node.string())

	// protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --{binary}_out=. trie/trie.proto
	//bs, _ := proto.Marshal(&trie)
	//var trie2 Trie
	//_ = proto.Unmarshal(bs, &trie2)

}
