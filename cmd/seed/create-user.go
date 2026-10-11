package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type userReq struct {
	IDStatus     int    `json:"id_status"`
	UserName     string `json:"user_name"`
	BirthDate    string `json:"birth_date"`
	Email        string `json:"email"`
	PhoneNum     string `json:"phone_num"`
	Address      string `json:"address"`
	PasswordHash string `json:"password_hash"`
}

func buildUser(i int) userReq {
	return userReq{
		IDStatus:     1,
		UserName:     fmt.Sprintf("Nguyen Van %d", i),
		BirthDate:    "05-11-2025",
		Email:        fmt.Sprintf("nguyenvan%d@gmail.com", i),
		PhoneNum:     fmt.Sprintf("0%d", 800000000+i),
		Address:      fmt.Sprintf("Quan Go Vap %d", i),
		PasswordHash: "12345678",
	}
}

func main() {
	url := flag.String("url", "http://localhost:8080/v1/api/manager/user", "API tạo user")
	start := flag.Int("start", 1, "số bắt đầu")
	total := flag.Int("total", 1000000, "tổng số user cần tạo")
	conc := flag.Int("c", 20, "số worker chạy song song")
	flag.Parse()

	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        *conc,
			MaxIdleConnsPerHost: *conc,
		},
	}

	jobs := make(chan int, *conc*2)
	var ok, fail, done int64
	var wg sync.WaitGroup
	begin := time.Now()

	for w := 0; w < *conc; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				body, _ := json.Marshal(buildUser(i))
				resp, err := client.Post(*url, "application/json", bytes.NewReader(body))
				if err != nil {
					if n := atomic.AddInt64(&fail, 1); n <= 10 {
						fmt.Printf("Lỗi i=%d: %v\n", i, err)
					}
				} else {
					msg, _ := io.ReadAll(io.LimitReader(resp.Body, 200))
					resp.Body.Close()
					if resp.StatusCode >= 200 && resp.StatusCode < 300 {
						atomic.AddInt64(&ok, 1)
					} else if n := atomic.AddInt64(&fail, 1); n <= 10 {
						fmt.Printf("Lỗi i=%d: HTTP %d %s\n", i, resp.StatusCode, msg)
					}
				}
				if n := atomic.AddInt64(&done, 1); n%1000 == 0 {
					fmt.Printf("Đã xử lý %d | thành công: %d | lỗi: %d | vị trí: %d | %s\n",
						n, atomic.LoadInt64(&ok), atomic.LoadInt64(&fail), i,
						time.Since(begin).Round(time.Second))
				}
			}
		}()
	}

	for i := *start; i < *start+*total; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	fmt.Printf("Xong trong %s. Thành công: %d, lỗi: %d\n",
		time.Since(begin).Round(time.Second), ok, fail)
}