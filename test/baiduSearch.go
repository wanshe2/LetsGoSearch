package test

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DictResponse struct {
	Q string `json:"q"`
	P bool   `json:"p"`
	G []struct {
		Type string `json:"type"`
		Sa   string `json:"sa"`
		Q    string `json:"q"`
	} `json:"g"`
	Slid    string `json:"slid"`
	Queryid string `json:"queryid"`
}

func BaiduSearch() {
	client := &http.Client{}

	filePath := "output.txt"
	file, err := os.Open(filePath)
	if err != nil {
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	filePath = "TrieData.txt"
	TrieFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	defer TrieFile.Close()
	writer := bufio.NewWriter(TrieFile)

	for {
		word, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}

		// 忽略掉不支持(>122的单字符)的字符,带空格所以用 <= 2,对于小于等于 122 的阿拉伯字符还是要的
		if len(word) <= 2 && word[0] > 122 {
			continue
		}
		log.Println(word)
		word = word[:len(word)-1]

		//var build = strings.Builder{}
		//build.WriteString("https://www.baidu.com/sugrec?pre=1&p=3&ie=utf-8&json=1&prod=pc&from=pc_web&sugsid=36427,36367,36005,35915,36167,34584,35978,36055,36419,26350,36348&wd=")
		//build.WriteString(word)
		//build.WriteString("&req=2&bs=%E4%BD%A0%E5%A5%BD&pbs=%E4%BD%A0%E5%A5%BD&csor=2&pwd=%E4%BD%A0%E5%A5%BD&cb=jQuery11020027690115017184436_1652173182282&_=1652173182553")
		//fmt.Println(build.String())
		url := "https://www.baidu.com/sugrec?pre=1&p=3&ie=utf-8&json=1&prod=pc&from=pc_web&sugsid=36427,36367,36005,35915,36167,34584,35978,36055,36419,26350,36348&wd=" + word + "&req=2&bs=%E4%BD%A0%E5%A5%BD&pbs=%E4%BD%A0%E5%A5%BD&csor=2&pwd=%E4%BD%A0%E5%A5%BD&cb=jQuery11020027690115017184436_1652173182282&_=1652173182553"
		//url = delEmpty(url)
		//fmt.Println(url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Accept", "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Cookie", "PSTM=1637892756; __yjs_duid=1_8065c540407bcf9fdd44e659b967d4031639841012332; BD_UPN=12314753; BIDUPSID=3B318AC7B171B7867AE3B5FAE455C0CD; BAIDUID=2CFECD3DCBE4BBA2E8739BE35733D867:FG=1; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BDUSS=klldmlNM2Nqc2tWWmRmMHNFV1BEdGVnR1ZnflBFVH5ISkpFajM4ajdTelhoS0ZpSVFBQUFBJCQAAAAAAAAAAAEAAADBgoPpeWVhcrK71qrD-9Ch1-QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANf3eWLX93lic; BDUSS_BFESS=klldmlNM2Nqc2tWWmRmMHNFV1BEdGVnR1ZnflBFVH5ISkpFajM4ajdTelhoS0ZpSVFBQUFBJCQAAAAAAAAAAAEAAADBgoPpeWVhcrK71qrD-9Ch1-QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANf3eWLX93lic; BDSFRCVID=mR8OJeC627YZvznDfCJgee3brm5S8o6TH6aoB7-2Q5ja_ZQHU_weEG0POx8g0Ku-T_hSogKKKgOTHICF_2uxOjjg8UtVJeC6EG0Ptf8g0f5; H_BDCLCKID_SF=tRCj_K_5tK03fP36q4OH5-_3-fTMetJyaR3LbDnvWJ5TMC_6Wq6-24Cvb4bZK-Q-tKrHWC5RfnA-ShPC-tnIjIrX0HoB5lRzLmbPM-3H3l02Vh79e-t2ynQD3bn3Q-RMW23i0h7mWpTTsxA45J7cM4IseboJLfT-0bc4KKJxbnLWeIJIjj6jK4JKDGDeJT5P; BCLID_BFESS=7531357301394653236; BDSFRCVID_BFESS=mR8OJeC627YZvznDfCJgee3brm5S8o6TH6aoB7-2Q5ja_ZQHU_weEG0POx8g0Ku-T_hSogKKKgOTHICF_2uxOjjg8UtVJeC6EG0Ptf8g0f5; H_BDCLCKID_SF_BFESS=tRCj_K_5tK03fP36q4OH5-_3-fTMetJyaR3LbDnvWJ5TMC_6Wq6-24Cvb4bZK-Q-tKrHWC5RfnA-ShPC-tnIjIrX0HoB5lRzLmbPM-3H3l02Vh79e-t2ynQD3bn3Q-RMW23i0h7mWpTTsxA45J7cM4IseboJLfT-0bc4KKJxbnLWeIJIjj6jK4JKDGDeJT5P; H_PS_PSSID=36427_36367_36005_35915_36167_34584_35978_36055_36419_26350_36348; H_PS_645EC=d4f7D5tdu2sRMEdYDVZf5%2FVaJlSx7z0YdZCZDWXET40kCqpbB9ye8N%2B0cC0; delPer=0; BDSVRTM=232; BA_HECTOR=a5812k25a08h2g04s81h7kq850r")
		req.Header.Set("Referer", "https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu&wd=%E4%BD%A0%E5%A5%BD&oq=dfa&rsv_pq=cfc99ee90002b1f5&rsv_t=e898DTC0P3ou6INjt2zEFLjXjDSTuy2knWPUCWmMTnY0uJkaWe4u5e9mb68&rqlang=cn&rsv_enter=1&rsv_dl=tb&rsv_btype=t&inputT=1415&rsv_sug3=245&rsv_sug1=24&rsv_sug7=100&rsv_sug2=0&rsv_sug4=1690")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
		req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Microsoft Edge";v="101"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"Windows"`)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%s\n", bodyText)
		if len(bodyText) == 0 {
			continue
		}

		var a int
		for i := 0; i < len(bodyText); i++ {
			if bodyText[i] == '(' {
				a = i
				break
			}
		}
		//fmt.Printf(string(bodyText[b:a]))
		str := string(bodyText[a+1 : len(bodyText)-1])
		var dictResponse DictResponse
		var ff []byte = []byte(str)
		//if word == "威世" {
		//	log.Println(str)
		//}

		//time.Sleep(6000)

		err = json.Unmarshal(ff, &dictResponse)
		if err != nil {
			log.Fatal(err)
		}

		_, err = writer.WriteString(word + "\n")
		if err != nil {
			return
		}
		for _, item := range dictResponse.G {
			//fmt.Println(item.Q)
			_, err := writer.WriteString(item.Q + "\n")
			if err != nil {
				return
			}
		}

	}

	err = writer.Flush()
	if err != nil {
		return
	}
}

func delEmpty(str string) string {
	result := ""
	for _, val := range str {
		if val == ' ' {
			continue
		}
		result += string(val)
	}
	return result
}
