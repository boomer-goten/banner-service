package main

import (
	"banner-server/internal/api/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var timePost int64

type value struct {
	Banner_id int `json:"banner_id"`
}

func sendRequestPost(id int32, chanData chan<- int) {
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
		start := time.Now()
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
		end := time.Now()
		timePost += end.Sub(start).Milliseconds()
	}
	chanData <- value.Banner_id
}

func sendRequestDelete(id int) {
	adrrURL := fmt.Sprintf("http://localhost:8080/banner/%d", id)
	req, err := http.NewRequest("DELETE", adrrURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	client.Do(req)
}

func main() {
	var wg sync.WaitGroup
	slice := make([]int, 0, 500)
	chanData := make(chan int)
	for i := 4500; i > 2500; i-- {
		wg.Add(1)
		time.Sleep(500 * time.Nanosecond)
		go func(id int32) {
			defer wg.Done()
			sendRequestPost(id, chanData)
		}(int32(i))
	}
	go func() {
		wg.Wait()
		close(chanData)
	}()
	for square := range chanData {
		slice = append(slice, square)
	}
	fmt.Printf("Выполнено %d вставок \n", len(slice))
	fmt.Printf("Время выполнения %d вставок %vms\n", len(slice), timePost)
	fmt.Printf("Среднее время выполнения одного удалений %vms\n", float64(timePost)/1000)
	var timeDelete int64
	for _, value := range slice {
		wg.Add(1)
		time.Sleep(500 * time.Nanosecond)
		go func(id int) {
			defer wg.Done()
			start := time.Now()
			sendRequestDelete(value)
			end := time.Now()
			timeDelete += end.Sub(start).Milliseconds()
		}(value)
	}
	wg.Wait()
	fmt.Printf("Выполнено %d удалений \n", len(slice))
	fmt.Printf("Время выполнения %d удалений %vms\n", len(slice), timeDelete)
	fmt.Printf("Среднее время выполнения одного удалений %vms\n", float64(timeDelete)/1000)
}
