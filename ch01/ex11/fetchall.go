package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {

	data, _ := ioutil.ReadFile("./moz500Domains.csv")
	sites := strings.Split(string(data), "\n")[0:500]
	start := time.Now()
	ch := make(chan string)
	for _, url := range sites {
		go fetch(url, ch) // ゴルーチンを開始
	}
	for range sites {
		fmt.Println(<-ch) // chチャネルから受信
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // chチャネルへ送信
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
