package main

import "time"

func longCommute(duration string) {
	dur, err := time.ParseDuration(duration)

	if err != nil {
		panic(err)
	}

	time.Sleep(dur)

	println(duration)
}

func main() {
	go longCommute("2s")
	go longCommute("1s")
	go longCommute("4s")
	go longCommute("3s")
	go longCommute("2s")
	longCommute("4s")
}
