package main

import (
	"fmt"
	"github.com/AlexanderHOtt/lsmtree/pkg/lsmtree"
)

func main() {
	tree := lsmtree.New(10)
	tree.Put(1, 1)
	tree.Put(2, 2)
	tree.Put(3, 3)
	tree.Put(4, 4)
	println("buf ", tree)
	for i, e := range tree.GetRange(1, 3) {
		fmt.Println(i, " ", e)
	}

}
