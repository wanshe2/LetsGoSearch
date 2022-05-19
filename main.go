package main

import (
	"LetsGoSearch/test"
	"LetsGoSearch/trie"
	"bufio"
	"fmt"
	"io"
	"os"
)

func testTrie() {
	//trieTree := trie.Load()
	trieTree := trie.NewTrie()

	for i := 1; i < 9; i++ {
		filepath := fmt.Sprintf("TrieData%d.txt", i)
		file, err := os.Open(filepath)
		if err != nil {
			panic("TrieData open file!")
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			str, ok := reader.ReadString('\n')
			if ok == io.EOF {
				break
			}
			trieTree.InsertString(str)
		}
	}

	fmt.Println("Insert success, start Search:")

	str := "回复"
	result := trieTree.Search([]rune(str))
	fmt.Println(len(result))
	for _, val := range result {
		fmt.Print(val)
	}

	trieTree.FlushIndex()
}

func main() {
	//test.GetWord()

	test.BaiduSearch()

	//testTrie()
}
