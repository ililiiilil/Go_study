package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go sendMessage(ch)

	for i := 0; i < 10; i++ {
		select {
		case <-ch:
			fmt.Println("메시지가 발송되었습니다.")
		default:
			fmt.Println("초 안에 메세지를 입력하시오.")
		}
	}
}

func sendMessage(ch chan string) {
	var s string
	fmt.Scanf("%s", &s)
	ch <- s
}
