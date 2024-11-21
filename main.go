package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"os"
	"strings"
)

var (
	longQuest bool
)

func init() {
	flag.BoolVar(&longQuest, "l", false, "进入上下文对话模式")
	flag.Parse()
}

func Filter[S ~[]E, E any](data S, callback func(E) bool) []E {
	var result []E
	for _, v := range data {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	if longQuest {
		fmt.Println("进入上下文对话模式")
		args := Filter(os.Args[1:], func(item string) bool {
			return item != "-l"
		})
		_ = LongChat(args...)
	}
	if len(os.Args) > 1 {
		_ = chatGPT([]Message{{Role: "user", Content: strings.Join(os.Args[1:], " ")}}, os.Stdout)
		fmt.Println("")
	} else {
		fmt.Println("例子: ai hi 或者 ai -l 进入上下文对话模式")
	}
}

func LongChat(msg ...string) error {
	msgs := []Message{}
	for {
		fmt.Println("---------------------------")
		var (
			text string
			err  error
		)
		if len(msg) > 0 {
			text = strings.Join(msg, " ")
			fmt.Println(text)
			msg = nil
		} else {
			text, err = ReadText()
			if err != nil {
				return err
			}
		}
		msgs = append(msgs, Message{
			Role:    "user",
			Content: text,
		})
		buf := bytes.NewBuffer(nil)
		_ = chatGPT(msgs, &Writer{buf})
		msgs = append(msgs, Message{
			Role:    "assistant",
			Content: buf.String(),
		})
		fmt.Println("")
	}
}

func chatGPT(msg []Message, output io.Writer) error {
	type AutoGenerated struct {
		Messages         []Message `json:"messages"`
		Stream           bool      `json:"stream"`
		Model            string    `json:"model"`
		Temperature      float64   `json:"temperature"`
		PresencePenalty  int       `json:"presence_penalty"`
		FrequencyPenalty int       `json:"frequency_penalty"`
		TopP             float64   `json:"top_p"`
	}
	type AutoGenerated2 struct {
		ID      string `json:"id"`
		Object  string `json:"object"`
		Created int    `json:"created"`
		Model   string `json:"model"`
		Choices []struct {
			Index int `json:"index"`
			Delta struct {
				Content string `json:"content"`
			} `json:"delta"`
			Logprobs     any `json:"logprobs"`
			FinishReason any `json:"finish_reason"`
		} `json:"choices"`
		SystemFingerprint string `json:"system_fingerprint"`
	}
	data := AutoGenerated{
		Messages:         msg,
		Stream:           true,
		Model:            "智普GLM-4",
		Temperature:      0.5,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		TopP:             1,
	}
	bs, _ := json.Marshal(data)
	urlPath := `https://newapi.yjie.fun/v1/chat/completions`
	req, err := http.NewRequest(http.MethodPost, urlPath, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	req.Header.Add("Host", "newapi.yjie.fun")
	req.Header.Add("authorization", "Bearer sk-rmkNimoGp5tjFzm3PCR9QxwajCZza2wF1274xWMRks4U")
	req.Header.Add("content-type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.HasPrefix(line, []byte("data: ")) {
			line = bytes.TrimPrefix(line, []byte("data: "))
			line = bytes.TrimSpace(line)
			var mp AutoGenerated2
			err = json.Unmarshal(line, &mp)
			if err != nil {
				break
			}
			_, err := output.Write([]byte(mp.Choices[0].Delta.Content))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ReadText() (string, error) {
	rd := textproto.NewReader(bufio.NewReader(os.Stdin))
	return rd.ReadLine()
}

type Writer struct {
	w io.Writer
}

func (w *Writer) Write(p []byte) (n int, err error) {
	if w.w != nil {
		_, _ = w.w.Write(p)
	}
	return os.Stdout.Write(p)
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
