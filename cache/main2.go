package main

import "time"

type Value struct {
	A string
	B int
	C time.Time
	D []byte
	E float32
	F *string
	T T
}

type T struct {
	H int
	I int
	J int
	K int
	L int
	M int
	N int
}

func main() {
	m := make(map[int]int, 10000000)
	for i := 0; i < 10000000; i++ {
		m[i] = i
	}

	for i := 0; ; i++ {
		delete(m, i)
		m[10000000+i] = i
		time.Sleep(5 * time.Millisecond)
	}
}
