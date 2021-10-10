package main

import "fmt"

func main() {
	ch := make(chan bool, 20)

	go func() {
		for i := 1; i < 21; i++ {
			ch <- true
		}
		fmt.Println("송신 루틴 완료")
	}()

	for i := 1; i < 21; i++ {
		<-ch
		fmt.Println("수신한 데이터 :", i)
	}
}
