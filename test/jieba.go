package test

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/wangbin/jiebago"
	"os"
	"strings"
)

func GetWord() {
	seg := jiebago.Segmenter{}
	err := seg.LoadDictionary("./data/dictionary.txt")
	if err != nil {
		return
	}

	wordMap := make(map[string]int)

	for i := 0; i < 10; i++ {
		filepath := fmt.Sprintf("F:\\悟空数据集\\wukong_release\\wukong_100m_%d.csv", i)
		f, err := os.Open(filepath)
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				return
			}
		}(f)
		if err != nil {
			return
		}
		reader := csv.NewReader(f)
		preData, err := reader.ReadAll()
		if err != nil {
			return
		}

		for i := 1; i < len(preData); i++ {
			text := preData[i][1]

			text = strings.ToLower(text)

			resultChan := seg.CutForSearch(text, true)

			for {
				w, ok := <-resultChan
				if !ok {
					break
				}
				_, found := wordMap[w]
				if !found {
					//去除重复的词
					wordMap[w] = 1
				}
			}

		}
	}

	var wordsSlice []string
	for k, _ := range wordMap {
		wordsSlice = append(wordsSlice, k)
	}

	//for _, val := range wordsSlice {
	//	fmt.Println(val)
	//}

	filePath := "output.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("file open err")
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("file close err")
			return
		}
	}(file)

	fmt.Println(len(wordsSlice))
	writer := bufio.NewWriter(file)

	for _, val := range wordsSlice {
		//log.Println(val, len(val))
		// 非正常字符忽略
		if len(val) <= 1 {
			continue
		}
		_, err = writer.WriteString(val + " ")
		if err != nil {
			return
		}
		err = writer.Flush()
		if err != nil {
			return
		}
	}
}
