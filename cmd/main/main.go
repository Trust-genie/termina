package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	//do womething awesome
	fmt.Println(count(os.Stdin))
	time.Sleep(15 * time.Second)
}

func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanWords)

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
