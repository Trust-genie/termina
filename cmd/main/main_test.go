package main_test

import (
	"bufio"
	"bytes"
	"io"
	"testing"
)

func TestCounter(t *testing.T) {
	b := bytes.NewBufferString("Newphrase Newphrase Newphrase Newphrase")
	exp := 4

	res := count(b)

	if res != exp {
		t.Errorf("Expected %d got %d", exp, res)
	}
}

func BenchmarkCounter(b *testing.B){
	for i:= 0; i< b.N ; i++{
		b := bytes.NewBufferString("Newphrase Newphrase Newphrase Newphrase")
		count(b)
	}
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