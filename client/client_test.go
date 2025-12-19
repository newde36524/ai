package client

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func Test_Ai(t *testing.T) {
	data := `{"models":[{"modelId":120,"modelName":"Qwen3-235B-A22B","active":false},{"modelId":9,"modelName":"Qwen3-30B-A3B","active":false},{"modelId":8,"modelName":"MiniMax-Text-01-456B","active":false},{"modelId":10,"modelName":"MiniMax-M1-456B","active":true}],"username":"xxxxx","password":"xxxxx","cookie":"xxxxxxxx","lastLogin":"1122-11-11"}`
	config := new(Config)
	_ = json.Unmarshal([]byte(data), config)
	cli := NewClient()
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
	_, err := cli.Completion("", "hi", os.Stdout)
	if err != nil {
		panic(err)
	}
}
