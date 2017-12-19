package main

import "fmt"

type orang struct {
	nama string
	umur int
}

func hello(o orang) {
	fmt.Printf("hello %s, umur kamu adalah %d\n", o.nama, o.umur)
}

func buatJadiMuda(o orang) {
	o.umur = 17
}

func main() {
	udin := orang{"udin", 25}

	hello(udin)

	buatJadiMuda(udin)

	hello(udin) //umur udin tetap 25, karena semua pass by value
}
