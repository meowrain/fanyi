package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"translate-cli/models"

	"github.com/fatih/color"
)

func ParseAndPrintDictInChinese(jsonData string) {
	// 创建 Dict 结构体实例
	var dict models.Dict

	// 解析 JSON 数据到结构体
	err := json.Unmarshal([]byte(jsonData), &dict)
	if err != nil {
		log.Fatalf("解析 JSON 出错: %v", err)
	}

	// 美化输出
	titleStyle := color.New(color.FgHiYellow).Add(color.Bold)
	subtitleStyle := color.New(color.FgCyan).Add(color.Bold)
	contentStyle := color.New(color.FgHiWhite)
	labelStyle := color.New(color.FgHiGreen).Add(color.Bold)
	separator := color.New(color.FgMagenta).Sprint(strings.Repeat("-", 40))

	// 输出解析结果
	titleStyle.Println("======== 词典数据 ========")
	fmt.Println(separator)
	labelStyle.Printf("单词: ")
	contentStyle.Printf("%s\n", dict.Word)

	labelStyle.Printf("词性: ")
	contentStyle.Printf("%v\n", dict.Pos)
	fmt.Println(separator)

	// 输出发音
	if len(dict.Pronunciations) > 0 {
		subtitleStyle.Println("-- 发音 --")
		for _, pron := range dict.Pronunciations {
			labelStyle.Printf("词性: ")
			contentStyle.Printf("%s, ", pron.Pos)

			labelStyle.Printf("语言: ")
			contentStyle.Printf("%s, ", pron.Lang)

			labelStyle.Printf("音标: ")
			contentStyle.Printf("%s\n", pron.Pron)
		}
		fmt.Println(separator)
	}

	// 输出释义
	if len(dict.Definitions) > 0 {
		subtitleStyle.Println("-- 释义 --")
		for _, def := range dict.Definitions {
			labelStyle.Printf("中文翻译: ")
			contentStyle.Printf("%s\n", def.Translation)

			labelStyle.Printf("词性: ")
			contentStyle.Printf("%s\n", def.Pos)

			labelStyle.Printf("英文定义: ")
			contentStyle.Printf("%s\n", def.Text)

			if len(def.Examples) > 0 {
				subtitleStyle.Println("例句:")
				for _, ex := range def.Examples {
					labelStyle.Printf("  原文: ")
					contentStyle.Printf("%s\n", ex.Text)
					labelStyle.Printf("  翻译: ")
					contentStyle.Printf("%s\n", ex.Translation)
				}
			}
			fmt.Println(separator)
		}
	}

	// 输出动词变形
	if dict.Verbs != nil && len(dict.Verbs) > 0 {
		subtitleStyle.Println("-- 动词变形 --")
		for _, verb := range dict.Verbs {
			labelStyle.Printf("类型: ")
			contentStyle.Printf("%s, ", verb.Type)

			labelStyle.Printf("文字: ")
			contentStyle.Printf("%s\n", verb.Text)
		}
		fmt.Println(separator)
	}
}
