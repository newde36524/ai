package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/newde36524/ai/client"
	"github.com/newde36524/ai/tool"
)

var (
	longQuest bool
	reset     bool
)

func init() {
	flag.BoolVar(&longQuest, "l", false, "进入上下文对话模式")
	flag.BoolVar(&reset, "r", false, "重置配置")
	flag.Parse()
}

func main() {
	ex, _ := os.Executable()
	config := client.NewConfig(filepath.Join(filepath.Dir(ex), "ai_config.json"))
	cli := client.NewClient()
	if !reset {
		config.Load()
	}
	cli.SetConfig(config)
	if err := cli.CheckUserInfo(); err != nil {
		fmt.Println(err)
		return
	}
	if err := cli.Login(); err != nil {
		fmt.Println(err)
		return
	}
	defer cli.Clear()
	if err := cli.CheckConfig(); err != nil {
		fmt.Println(err)
		return
	}
	if err := config.Save(); err != nil {
		fmt.Println(err)
		return
	}
	args := tool.Filter(os.Args[1:], func(item string) bool {
		return item != "-l" && item != "-r"
	})
	if longQuest {
		fmt.Println("进入上下文对话模式")
		_ = cli.LongChat(args...)
	}
	if len(os.Args) > 1 {
		_, _ = cli.Completion("", strings.Join(args, " "), os.Stdout)
		fmt.Println("")
	} else {
		fmt.Println("例子: ai hi 或者 ai -l 进入上下文对话模式")
	}
}
