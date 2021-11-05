package main

import (
	"errors"
	"fmt"

	"github.com/xwz123/test/animal"
)

func main() {
	fmt.Printf("hello github actions %v", "s")
	a, _ := mut(1, 0)
	fmt.Println(a)
	var an animal.Animaler
	an = animal.Dog{}
	an.Run()
	an.Eate()
}

func mut(a, b int) (int) {
	fmt.Println(a,b)
	return a / b
}
