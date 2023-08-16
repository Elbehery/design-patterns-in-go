package main

import (
	"fmt"
	"log"
)

type Form interface {
	Clone() Form
}

type Concrete struct {
	Name string
}

func (c *Concrete) Clone() Form {
	return &Concrete{Name: c.Name}
}

func main() {
	c := Concrete{Name: "Mustafa"}

	clone := c.Clone()

	c1, ok := clone.(*Concrete)
	if !ok {
		log.Fatal("illegal casting")
	}

	fmt.Println(c1)
}
