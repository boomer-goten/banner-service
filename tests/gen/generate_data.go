package main

import (
	"banner-server/internal/api/model"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"
)

type value struct {
	Banner_id int `json:"banner_id"`
}

func SendRequestPost(id int32) {
	postData := model.BannerPostRequest{
		TagIds:    []int32{id, id + 1},
		FeatureId: id,
		Content: map[string]interface{}{
			"title": "hello, i'm title from 1 post test",
			"text":  "i have four tags",
		},
		IsActive: true,
	}
	postByte, err := json.Marshal(postData)
	value := value{}
	if err == nil {
		req, err := http.NewRequest("POST", "http://localhost:8080/banner", bytes.NewBuffer(postByte))
		if err != nil {
			panic(err)
		}
		req.Header.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI")
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			response, err := io.ReadAll(resp.Body)
			if err == nil {
				json.Unmarshal(response, &value)
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 4500; i > 3500; i-- {
		wg.Add(1)
		time.Sleep(1 * time.Millisecond)
		go func(id int32) {
			defer wg.Done()
			SendRequestPost(id)
		}(int32(i))
	}
	wg.Wait()
}
