package main

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"

	"translate-cli/utils"
)

func main() {
	var word string = os.Args[1]
	var mode string = "en-cn"
	if len(os.Args) > 2 {
		mode = os.Args[2]
	}
	client := resty.New()
	url := fmt.Sprintf("https://dict.meowrain.cn/api/dictionary/%s/%s", mode, word)
	resp, err := client.R().Get(url)
	if err != nil {
		fmt.Println(err)
	}

	// 输出解析结果
	utils.ParseAndPrintDictInChinese(string(resp.Body()))
}
