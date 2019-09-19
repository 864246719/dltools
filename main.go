package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type novelIndex struct {
	link   string
	name   string
	author string
}

var (
	rootURL = "https://www.37zw.net"
)

func main() {

	novelIndexList := getNovelIndexList()[0:100]

	for i, novelIndex := range novelIndexList {
		go getNovelInfo(i, novelIndex)
	}

	time.Sleep(200 * time.Second)

}

func getNovelIndexList() []novelIndex {
	var novelIndexList []novelIndex
	var novelIndex novelIndex
	re := regexp.MustCompile(`/\d/\d+/[^\s]*`)
	reSplitFlag := regexp.MustCompile(`(">|</a>/)`)

	// gbk编码类型字符bytes的解析器
	chDecoder := simplifiedchinese.GB18030.NewDecoder()

	resp, err := http.Get("https://www.37zw.net/xiaoshuodaquan/")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	body, _ = chDecoder.Bytes(body)
	rowContents := re.FindAll(body, -1)

	for _, item := range rowContents {
		itemStr := strings.Replace(string(item), "</li>", "", -1)

		tempList := reSplitFlag.Split(itemStr, -1)

		novelIndex.link = rootURL + tempList[0]
		novelIndex.name = tempList[1]
		novelIndex.author = tempList[2]
		novelIndexList = append(novelIndexList, novelIndex)
	}

	return novelIndexList
}

func getNovelInfo(i int, novelIndex novelIndex) {

	re := regexp.MustCompile(`<meta property="og:description" content="([\w\W]*?)>`)
	// gbk编码类型字符bytes的解析器
	chDecoder := simplifiedchinese.GB18030.NewDecoder()
	resp, err := http.Get(novelIndex.link)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	body, _ = chDecoder.Bytes(body)

	err = ioutil.WriteFile(novelIndex.name+".txt", body, 0644)
	if err != nil {
		fmt.Println(err)
	}

	rowContents := re.Find(body)
	info := strings.Replace(string(rowContents), `<meta property="og:description" content="`, ``, -1)
	info = strings.Replace(info, `"/>`, ``, -1)

	fmt.Printf("===========================%v: %v ============================\n %q \n\n\n", i, novelIndex.name, info)

}
