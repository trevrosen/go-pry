package main

import (
	"../pry"

	"log"
)

func X() bool {
	return true
}

type Banana struct {
	Name string
	Cake []int
}

func (b Banana) Ly() string {
	return b.Name + "ly"
}

func main() {
	a := 1
	b := Banana{"Jeoffry", []int{1, 2, 3}}
	m := []int{1234}
	if d := X(); d {
		log.Println(d)
		for i, j := range []int{1} {
			k := 1
			log.Println(i, j, k)
			// Example comment
			pry.Pry()
		}
	}
	log.Println("Test", a, b, main)
}
