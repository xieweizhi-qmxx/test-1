package animal

import "fmt"

type Animaler interface {
	Run()
	Eate()
}

type Dog struct {
}

func (d Dog) Run() {
	fmt.Println("sdsdfsfsfsfdaf")
}

func (d Dog) Eate() {
	fmt.Println("eate dog")
}
