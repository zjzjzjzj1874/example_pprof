package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	fmt.Println("start listen on 8080...")
	go loop()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("listen err:", err.Error())
	}
}

func loop() {
	max := 100
	for {
		go func(max int) {
			fmt.Println("max == ", max)
			select {}
		}(max)
		max--
		time.Sleep(1 * time.Second)

		if max == 0 {
			break
		}
	}
}
