package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// CPU Profile
	f, err := os.Create("./profile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Memory Profile
	fm, err := os.Create("./heap")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fm.Close()
	pprof.WriteHeapProfile(fm)

	// goroutine Profile
	gm, err := os.Create("./goroutine")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fm.Close()
	ap := pprof.Profiles()
	gp := new(pprof.Profile)
	for _, profile := range ap {
		if profile.Name() != "goroutine" {
			continue
		}
		gp = profile
	}

	for i := 0; i < 100; i++ {
		fmt.Println("测试程序:", i)
	}

	loop()
	err = gp.WriteTo(gm, 0)
	fmt.Println("err == ", err)
	time.Sleep(time.Second)
}

func loop() {
	max := 10
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
