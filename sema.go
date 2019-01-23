package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int, 4) // 並列度を4に制限

// 第二引数は送信専用チャネル
func zhengTask(i int, tch chan <- string) {
	ch <- 1                 // zhengくんは4つ超えたら辻もっちゃんを未読、空きがあれば見る
	defer func() { <-ch }() // zhengくん新しい辻もっちゃんを受け付けれるようになる
	time.Sleep(time.Second)
	s := fmt.Sprintf("pull request #%d", i)
	fmt.Println(s)
	tch <- s
}

func tsujimotoReview(pr string) {
	fmt.Printf("review: %s\n", pr)
	time.Sleep(2 * time.Second)
}

func main() {
	var tch = make(chan string, 20)
	tsujimoto := sync.WaitGroup{}
	tsujimoto.Add(20)
	for i := 0; i < 20; i++ {
		// zhengくんは五月雨でタスクを振られている
		go func(i int) {
			zhengTask(i, tch)
			tsujimoto.Done() // Doneするまで下のWaitで待つ
		}(i)
	}

	go func() {
		tsujimoto.Wait()
		fmt.Println("zheng '仕事終わったよ'")
		close(tch)
		fmt.Println("zheng '....'")
	}()

	for pr := range tch {
		tsujimotoReview(pr)
	}
}
