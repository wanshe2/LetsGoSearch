package trie

import (
	"LetsGoSearch/constants"
	"bufio"
	"io"
	"os"
	"runtime"
	"time"
)

var strings []string
var counts []int32

func (n *Node) foreach(runes []rune, deep int) {
	if n.count != 0 {
		strings = append(strings, string(runes))
		counts = append(counts, n.count)
	}
	for _, v := range n.child {
		if len(runes) <= deep {
			runes = append(runes, '我')
		}
		runes[deep] = v.data
		v.foreach(runes, deep+1)
	}
}

func Serialize(node *Node) ([]string, []int32) {
	strings, counts = make([]string, 0), make([]int32, 0)
	if node == nil {
		return strings, counts
	}
	// todo encoding ....
	// 1. find all Node which count greater than zero, then use getPrefix(). Complexity: 10N
	// 2. bring prefix in function. Complexity: ?N
	runes := make([]rune, 0)
	node.foreach(runes, 0)
	// dfs (substring + count) > txt.
	return strings, counts
}

func Write(trie *Trie) {
	// todo write ....
	filepath := constants.TrieFileName
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic("TrieData open file!")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	str, counts := Serialize(trie.root)
	for pos, val := range str {
		_, err := writer.WriteString(val + string(counts[pos]))
		if err != nil {
			return
		}
	}
}

func Load() *Trie {
	filepath := constants.TrieFileName
	file, err := os.Open(filepath)
	if err != nil {
		panic("Trie load error: can find file!")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	trie := NewTrie()
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		val, _, err := reader.ReadRune() // _ 是val作为rune的实际字节长度
		//fmt.Println(str, int32(val))
		trie.insertRunesWithCount([]rune(str), int32(val))
	}

	return trie
}

// 自动保存索引，10秒钟检测一次
func (t *Trie) automaticFlush() {
	ticker := time.NewTicker(time.Second * 10)
	size := 0

	for {
		<-ticker.C
		//检查数据是否有变动
		if size != t.size {
			size = t.size
			t.FlushIndex()
		} else {
			//e.FlushIndex()
		}
		//定时GC
		runtime.GC()
	}

}

// FlushIndex 刷新缓存到磁盘
func (t *Trie) FlushIndex() {
	t.Lock()
	defer t.Unlock()

	Write(t)
	//for i, index := range e.Indexes {
	//	//fmt.Println(e.getFilePath(fmt.Sprintf("%s.%d", e.Option.KeyIndexName, i)))
	//	dump.Write(index.Root, e.getFilePath(fmt.Sprintf("%s.%d", e.Option.KeyIndexName, i)))
	//}
}
