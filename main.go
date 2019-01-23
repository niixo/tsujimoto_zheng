package main

import (
	"log"
	"sync"
	"time"
)

func fetchURL(wg *sync.WaitGroup, q chan string) {
	defer wg.Done()
	for {
		url, ok := <-q
		if !ok {
			return
		}
		log.Printf("ダウンロード: %s\n", url)
		time.Sleep(3 * time.Second)
		log.Println("3 time second")
	}
}

func main() {
	var wg sync.WaitGroup
	//q := make(chan string, 5)
	// ワーカーを3つ作る
	for i := 0; i < 3; i++ {
		wg.Add(1)
		//go fetchURL(&wg, q)
	}
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			log.Println("hoge")
			wg.Done()
		}
	}()

	//q <- "http://www.example.com"
	//q <- "http://www.example.net"
	//q <- "http://www.example.net/foo"
	//q <- "http://www.example.net/bar"
	//q <- "http://www.example.net/baz"
	//close(q)

	wg.Wait()
	log.Println("end")
}