package glm

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	apiKey string
	client *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey, client: http.DefaultClient}
}

// notlint:errcheck
func (c *Client) Completion(content string, w io.Writer) error {
	urlPath := "https://open.bigmodel.cn/api/paas/v4/chat/completions"
	reqData, _ := json.Marshal(GlmReq{
		Model: "glm-4.6",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "system",
				Content: "你是一个有用的AI助手。",
			},
			{
				Role:    "user",
				Content: content,
			},
		},
		Temperature: 1.0,
		Stream:      true,
	})
	req, err := http.NewRequest(http.MethodPost, urlPath, bytes.NewReader(reqData))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		text := strings.TrimPrefix(scanner.Text(), "data: ")
		text = strings.TrimSuffix(text, "\n")
		text = strings.TrimSpace(text)
		// fmt.Println(text)
		if strings.Contains(text, "[done]") {
			break
		}
		var data GlmResp
		_ = json.Unmarshal([]byte(text), &data)
		if len(data.Choices) > 0 {
			_, err = w.Write([]byte(data.Choices[0].Delta.ReasoningContent))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return resp.Body.Close()
}

type GlmReq struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Temperature float64 `json:"temperature"`
	Stream      bool    `json:"stream"`
}

//	type GlmResp struct {
//		ID      string `json:"id"`
//		Created int    `json:"created"`
//		Model   string `json:"model"`
//		Choices []struct {
//			Index int `json:"index"`
//			Delta struct {
//				Role    string `json:"role"`
//				Content string `json:"content"`
//			} `json:"delta"`
//		} `json:"choices"`
//	}
type GlmResp struct {
	ID      string `json:"id"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index int `json:"index"`
		Delta struct {
			Role             string `json:"role"`
			ReasoningContent string `json:"reasoning_content"`
		} `json:"delta"`
	} `json:"choices"`
}
